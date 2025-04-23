package gapi

import (
	"testing"
	"time"

	db "github.com/adified/Go-Bank/db/sqlc"
	"github.com/adified/Go-Bank/util"
	"github.com/adified/Go-Bank/worker"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store db.Store, taskDistributor worker.TaskDistributor) *Server {
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store, taskDistributor)
	require.NoError(t, err)

	return server
}
