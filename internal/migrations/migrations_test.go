package migrations

import (
	"github.com/fastid/fastid/internal/config"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMigrations(t *testing.T) {
	cfg, err := config.NewConfig("../../configs/fastid.yml")
	require.NoError(t, err)

	migration, err := NewMigration(cfg, "sqlite3")
	require.NoError(t, err)

	err = migration.Upgrade()
	require.NoError(t, err)

	err = migration.Downgrade()
	require.NoError(t, err)
}
