package account

import (
	"context"

	_ "github.com/go-kivik/couchdb" // The CouchDB driver
	"github.com/go-kivik/kivik"
)

var db *kivik.DB

//InitDB establishes database connection
func InitDB(dbName string, dbType string) (*kivik.DB, error) {
	client, err := kivik.New(dbType, "http://admin:admin@127.0.0.1:5984/")
	return client.DB(context.Background(), dbName), err
}
