package repositories

import (
	"context"
	"github.com/fastid/fastid/internal/config"
	"github.com/fastid/fastid/internal/db"
	"github.com/fastid/fastid/internal/logger"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSession(t *testing.T) {
	cfg, err := config.NewConfig("../../configs/fastid.yml")
	if err != nil {
		t.Fatalf("%s", err.Error())
	}
	log := logger.NewLogger(cfg)

	database, err := db.NewDB(cfg, "sqlite3")
	if err != nil {
		log.Fatalln(err.Error())
	}

	ctx := context.Background()

	sessionsRepository := NewSessionsRepository(cfg, log, database)

	t.Run("First test", func(t *testing.T) {
		id, err := sessionsRepository.GetByID(ctx)
		require.NoError(t, err)
		require.Equal(t, id, "hello word")
	})
}
