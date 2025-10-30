package database

import (
	"fmt"

	"github.com/google/uuid"
	"go.etcd.io/bbolt"
)

const (
	defaultDBName = "default"
)

type M map[string]string

type Database struct {
	db *bbolt.DB
}

type Collection struct {
	*bbolt.Bucket
}

func New() (*Database, error) {
	dbname := fmt.Sprintf("%s.database", defaultDBName)
	db, err := bbolt.Open(dbname, 0666, nil)
	if err != nil {
		return nil, err
	}

	return &Database{
		db: db,
	}, nil
}

func (d *Database) CreateCollection(name string) (*Collection, error) {

	tx, err := d.db.Begin(true)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	// check if bucket exists
	// b := tx.Bucket([]byte(name))
	// if b != nil {
	// 	return &Collection{Bucket: b}, nil
	// }
	// create new bucket if not exists
	b, err := tx.CreateBucketIfNotExists([]byte(name))
	if err != nil {
		return nil, err
	}
	return &Collection{Bucket: b}, nil
}

func (d *Database) Insert(collName string, data M) (uuid.UUID, error) {
	id := uuid.New()
	tx, err := d.db.Begin(true)
	if err != nil {
		return id, err
	}
	defer tx.Rollback()
	b, err := tx.CreateBucketIfNotExists([]byte(collName))
	if err != nil {
		return id, err
	}
	for k, v := range data {
		if err := b.Put([]byte(k), []byte(v)); err != nil {
			return id, err
		}
	}
	if err := b.Put([]byte("id"), []byte(id.String())); err != nil {
		return id, err
	}
	return id, nil
}

func (d *Database) Select(coll string, k string, query any) {

}

// db.Update(func(tx *bbolt.Tx ) error {
// 	id := uuid.New()
// 	for k, v := range user {
// 		if err := b.Put([]byte(k), []byte(v)); err != nil {
// 			return err
// 		}
// 	}
// 	if err := b.Put([]byte("id"), []byte(id.String())); err != nil {
// 		return err
// 	}
// 	return nil
// })
