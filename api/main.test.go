package api

import (
	"os"
	"testing"
	"time"

	db "github.com/JILSE7/simplebank/db/sqlc"
	"github.com/JILSE7/simplebank/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := utils.Config{
		TokenSymmetrictKey:  utils.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store)

	require.NoError(t, err)

	return server

}
func main(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
