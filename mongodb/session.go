package mongodb

import (
	"context"
	"fmt"
	"github.com/kohirens/stdlib/web/session"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type SessionStore struct {
	collection *mongo.Collection
}

func NewSessionStore(c *mongo.Client, database, collection string) session.Storage {
	return &SessionStore{
		collection: c.Database(database).Collection(collection),
	}
}

func (ss *SessionStore) Save(id string, value session.Store, exp time.Time) error {
	query := map[string]string{"session_id": id}

	_, e1 := UpsertOne(
		query,
		session.OfflineStore{Id: id, Expiration: exp, Data: value},
		ss.collection,
	)
	if e1 != nil {
		return fmt.Errorf(stderr.CannotSaveSession, e1.Error())
	}

	return nil
}

func (ss *SessionStore) Load(id string) (*session.OfflineStore, error) {
	sd := &session.OfflineStore{}

	query := bson.M{"session_id": id}
	result := ss.collection.FindOne(context.TODO(), query)

	if e := result.Decode(sd); e != nil {
		return nil, fmt.Errorf(stderr.CannotLoadSession, e.Error())
	}

	return sd, nil
}
