package cache

import (
	"sync"
	"time"

	"gorm.io/gorm"
)

type Config struct {
	Easer  bool
	Cacher Cacher
	TTL    time.Duration
}

type Caches struct {
	callbacks map[queryType]func(db *gorm.DB)
	Conf      *Config

	queue *sync.Map
}

func (c *Caches) Name() string {
	return "gorm:caches"
}

func (c *Caches) Initialize(db *gorm.DB) error {
	if c.Conf == nil {
		c.Conf = &Config{
			Easer:  false,
			Cacher: nil,
		}
	}

	if c.Conf.Easer {
		c.queue = &sync.Map{}
	}

	callbacks := make(map[queryType]func(db *gorm.DB), 4)
	callbacks[uponQuery] = db.Callback().Query().Get("gorm:query")
	callbacks[uponCreate] = db.Callback().Create().Get("gorm:query")
	callbacks[uponUpdate] = db.Callback().Update().Get("gorm:query")
	callbacks[uponDelete] = db.Callback().Delete().Get("gorm:query")
	c.callbacks = callbacks

	if err := db.Callback().Query().Replace("gorm:query", c.query); err != nil {
		return err
	}

	if err := db.Callback().Create().Replace("gorm:query", c.getMutatorCb(uponCreate)); err != nil {
		return err
	}

	if err := db.Callback().Update().Replace("gorm:query", c.getMutatorCb(uponUpdate)); err != nil {
		return err
	}

	if err := db.Callback().Delete().Replace("gorm:query", c.getMutatorCb(uponDelete)); err != nil {
		return err
	}

	return nil
}

func (c *Caches) query(db *gorm.DB) {
	useCache, ok := db.Statement.Context.Value(useCacheKey).(bool)
	if !ok {
		useCache = false
	}

	cacheTTL, ok := db.Statement.Context.Value(cacheTTLKey).(time.Duration)
	if !ok {
		cacheTTL = c.Conf.TTL
	}

	if !useCache || (!c.Conf.Easer && c.Conf.Cacher == nil) {
		c.callbacks[uponQuery](db)
		return
	}

	if !c.Conf.Easer && c.Conf.Cacher == nil {
		c.callbacks[uponQuery](db)
		return
	}

	identifier := buildIdentifier(db)

	if c.checkCache(db, identifier) {
		return
	}

	c.ease(db, identifier)
	if db.Error != nil {
		return
	}

	c.storeInCache(db, identifier, cacheTTL)
	if db.Error != nil {
		return
	}
}

// getMutatorCb returns a decorator which calls the Cacher's Invalidate method
func (c *Caches) getMutatorCb(typ queryType) func(db *gorm.DB) {
	return func(db *gorm.DB) {
		if c.Conf.Cacher != nil {
			if err := c.Conf.Cacher.Invalidate(db.Statement.Context); err != nil {
				_ = db.AddError(err)
			}
		}
		if cb := c.callbacks[typ]; cb != nil { // By default, gorm has no callbacks associated with mutating behaviors
			cb(db)
		}
	}
}

func (c *Caches) ease(db *gorm.DB, identifier string) {
	if !c.Conf.Easer {
		c.callbacks[uponQuery](db)
		return
	}

	res := ease(&queryTask{
		id:      identifier,
		db:      db,
		queryCb: c.callbacks[uponQuery],
	}, c.queue).(*queryTask)

	if db.Error != nil {
		return
	}

	if res.db.Statement.Dest == db.Statement.Dest {
		return
	}

	detachedQuery := &Query[any]{
		Dest:         db.Statement.Dest,
		RowsAffected: db.Statement.RowsAffected,
	}

	easedQuery := &Query[any]{
		Dest:         res.db.Statement.Dest,
		RowsAffected: res.db.Statement.RowsAffected,
	}
	if err := easedQuery.copyTo(detachedQuery); err != nil {
		_ = db.AddError(err)
	}

	detachedQuery.replaceOn(db)
}

func (c *Caches) checkCache(db *gorm.DB, identifier string) bool {
	if c.Conf.Cacher != nil {
		res, err := c.Conf.Cacher.Get(db.Statement.Context, identifier, &Query[any]{
			Dest:         db.Statement.Dest,
			RowsAffected: db.Statement.RowsAffected,
		})
		if err != nil {
			_ = db.AddError(err)
		}

		if res != nil {
			res.replaceOn(db)
			return true
		}
	}
	return false
}

func (c *Caches) storeInCache(db *gorm.DB, identifier string, ttl time.Duration) {
	if c.Conf.Cacher != nil {
		err := c.Conf.Cacher.Store(db.Statement.Context, identifier, &Query[any]{
			Dest:         db.Statement.Dest,
			RowsAffected: db.Statement.RowsAffected,
		},
			ttl,
		)
		if err != nil {
			_ = db.AddError(err)
		}
	}
}

type key byte

const (
	useCacheKey key = iota
	cacheTTLKey
)

type queryType byte

const (
	uponQuery queryType = iota
	uponCreate
	uponUpdate
	uponDelete
)
