package cache

type task interface {
	GetId() string
	Run()
}
