package migrations

import (
	"github.com/fastid/fastid/internal/config"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestMigrations(t *testing.T) {
	err := os.Chdir("../../")
	require.NoError(t, err)

	cfg, err := config.New("configs/fastid.yml")
	require.NoError(t, err)

	migration, err := New(cfg, "sqlite3")
	require.NoError(t, err)

	err = migration.Upgrade()
	require.NoError(t, err)

	err = migration.Downgrade()
	require.NoError(t, err)
}

func TestMigrationsPostgres(t *testing.T) {
	err := os.Chdir("../../")
	require.NoError(t, err)

	cfg, err := config.New("configs/fastid.yml")
	require.NoError(t, err)

	if os.Getenv("CI") == "" {
		t.Skip("Skipping testing in not CI environment")
	}

	_, err = New(cfg, "postgres")
	require.NoError(t, err)

}
