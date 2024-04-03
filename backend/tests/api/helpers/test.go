package helpers

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/huandu/go-assert"
)

type TestAppConfigurator func(*TestApp)

func WithScrapedData() TestAppConfigurator {
	return func(app *TestApp) {
		migrationSQL, err := os.ReadFile(filepath.Join("..", "..", "src", "migrations", "data.sql"))
		if err != nil {
			panic(fmt.Errorf("error reading migration file: %w", err))
		}

		if err := app.Conn.Exec(string(migrationSQL)).Error; err != nil {
			panic(fmt.Errorf("error executing migration: %w", err))
		}
	}
}

func InitTest(t *testing.T, configurators ...TestAppConfigurator) ExistingAppAssert {
	assert := assert.New(t)
	app, err := spawnApp()

	for _, configurator := range configurators {
		configurator(app)
	}

	assert.NilError(err)

	return ExistingAppAssert{
		App:    *app,
		Assert: assert,
	}
}
