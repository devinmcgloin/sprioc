package store

import (
	"log"
	"math/rand"
	"time"

	"github.com/sprioc/sprioc-core/pkg/env"

	"gopkg.in/mgo.v2"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

var dbname = env.Getenv("MONGODB_NAME", "morph")
var dbURI = env.Getenv("MONGODB_URI", "mongodb://localhost:27017/morph")

type MgoStore struct {
	session *mgo.Session
}

func ConnectStore() MgoStore {
	session, err := mgo.Dial(dbURI)
	if err != nil {
		log.Fatal(err)
	}

	err = session.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return MgoStore{session}
}

func (ds *MgoStore) getSession() *mgo.Session {
	return ds.session.Copy()
}