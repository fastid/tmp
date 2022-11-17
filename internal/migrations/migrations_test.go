package migrations

import (
	"github.com/fastid/fastid/internal/config"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestMigrations(t *testing.T) {
	cfg, err := config.New("../../configs/fastid.yml")
	require.NoError(t, err)

	migration, err := New(cfg, "sqlite")
	require.NoError(t, err)

	err = migration.Upgrade()
	require.NoError(t, err)

	err = migration.Downgrade()
	require.NoError(t, err)
}

func TestMigrationsPostgres(t *testing.T) {
	cfg, err := config.New("../../configs/fastid.yml")
	require.NoError(t, err)

	if os.Getenv("CI") == "" {
		t.Skip("Skipping testing in not CI environment")
	}

	migration, err := New(cfg, "postgres")
	require.NoError(t, err)

	err = migration.Upgrade()
	require.NoError(t, err)

	err = migration.Downgrade()
	require.NoError(t, err)
}

func TestMigrationFakeDriver(t *testing.T) {
	cfg, err := config.New("../../configs/fastid.yml")
	require.NoError(t, err)

	_, err = New(cfg, "fake")
	require.Error(t, err)
	require.Equal(t, err.Error(), "Unable to load the driver fake")
}
