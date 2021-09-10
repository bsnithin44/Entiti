package database

import (
	"os"
	"sync"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
)

var dbonce sync.Once

var DBSession gocqlx.Session

func GetDbSession() gocqlx.Session {

	dbonce.Do(
		func() {
			env := os.Getenv("ENVIRONMENT")

			// Connect to cluster
			cluster := gocql.NewCluster(os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT"))

			if env == "production" || env == "dev" {
				// only auth while aws keyspaces
				cluster.Authenticator = gocql.PasswordAuthenticator{
					Username: os.Getenv("DB_USERNAME"),
					Password: os.Getenv("DB_PASSWORD")}
				// provide the path to the sf-class2-root.crt
				cluster.SslOpts = &gocql.SslOptions{
					CaPath: os.Getenv("CERT_PATH")}
			}

			cluster.Consistency = gocql.LocalQuorum
			// DBSession, _ = cluster.CreateSession()

			DBSession, _ = gocqlx.WrapSession(cluster.CreateSession())

		})

	return DBSession
}
