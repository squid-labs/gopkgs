package database_test

import (
	"gopkgs/database"
	"os"
	"testing"

	"github.com/ilyakaznacheev/cleanenv"
)

// TODO
// Test connection to invalid config, wrong adapter, and more
func Test_Connect(t *testing.T) {
	t.Run("Connection to sqlite", func(t *testing.T) {
		_, err := database.Connect(&database.Config{
			Driver: database.SqliteDriver,
			Pool:   5,
		})
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("Wrong driver", func(t *testing.T) {
		_, err := database.Connect(&database.Config{
			Driver: "foo1",
		})

		if err == nil {
			t.Error("Has to fail with unsupported driver")
		}
	})

	t.Run("Empty configuration", func(t *testing.T) {
		_, err := database.Connect(&database.Config{})
		if err == nil {
			t.Error(err)
		}
	})

	t.Run("Proper configuration", func(t *testing.T) {
		const configFile = `
port: 6009
database:
  driver: sqlite
  host: localhost
  port: 3006
  name: opendax
  user: user
  pass: changeme
`

		type config struct {
			ServerPort string          `yaml:"port"`
			DbConfig   database.Config `yaml:"database"`
		}

		// In your application you can just create the file ( "config/config.yml" for example )
		// And pass the path to the file in ika.ReadConfig
		tmpFile, err := os.CreateTemp(os.TempDir(), "*.yml")
		if err != nil {
			t.Error(err)
		}
		defer os.Remove(tmpFile.Name())

		input := []byte(configFile)
		if _, err = tmpFile.Write(input); err != nil {
			t.Error(err)
		}

		cfg := config{}
		cleanenv.ReadConfig(tmpFile.Name(), &cfg)

		_, err = database.Connect(&cfg.DbConfig)
		if err != nil {
			t.Error(err)
		}

		t.Logf("Selected database driver is %s", cfg.DbConfig.Driver)
	})
}

func Test_Create(t *testing.T) {
	t.Run("Create database with sqlite driver", func(t *testing.T) {
		err := database.Create(&database.Config{
			Driver: database.SqliteDriver,
			Pool:   5,
		})
		if err != nil {
			t.Error(err)
		}
	})
}

func Test_Drop(t *testing.T) {
	t.Run("Drop database with sqlite driver", func(t *testing.T) {
		err := database.Drop(&database.Config{
			Driver: database.SqliteDriver,
			Pool:   5,
		})
		if err != nil {
			t.Error(err)
		}
	})
}
