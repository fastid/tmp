package db

import (
	"github.com/fastid/fastid/internal/config"
	"github.com/stretchr/testify/require"
	"testing"
)

type Version struct {
	Name string `db:"name"`
}

func TestDBSqlite3(t *testing.T) {
	cfg, _ := config.New("../../configs/fastid.yml")
	db, _ := NewDB(cfg, "sqlite")
	sql := db.GetConnect()

	var result Version
	err := sql.Get(&result, "SELECT 'hello word' AS name")
	require.NoError(t, err)
	require.Equal(t, result.Name, "hello word")

	err = db.Close()
	require.NoError(t, err)
}

func TestInvalid(t *testing.T) {
	cfg, _ := config.New("../../configs/fastid.yml")

	t.Run("FakeDriver", func(t *testing.T) {
		_, err := NewDB(cfg, "FakeDriver")
		require.Error(t, err)
		require.Equal(t, err.Error(), "Unable to load the driver FakeDriver")
	})

	t.Run("ErrorConnection", func(t *testing.T) {
		cfg.DATABASE.Host = "127.0.0.1"
		cfg.DATABASE.Port = "9999"

		_, err := NewDB(cfg, "postgres")
		require.Error(t, err)
		require.Equal(t, err.Error(), "dial tcp 127.0.0.1:9999: connect: connection refused")
	})
}
