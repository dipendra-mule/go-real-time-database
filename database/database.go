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
	coll := Collection{}
	err := d.db.Update(func(tx *bbolt.Tx) error {
		var (
			err error
			b   *bbolt.Bucket
		)
		b = tx.Bucket([]byte(name))
		if b == nil {
			b, err = tx.CreateBucket([]byte("users"))
			if err != nil {
				return err
			}
		}
		coll.Bucket = b
		return nil
	})

	if err != nil {
		return nil, err
	}
	return &coll, nil
}

func (d *Database) Insert(collName string, data M) (uuid.UUID, error) {
	id := uuid.New()
	coll, err := d.CreateCollection(collName)
	if err != nil {
		return id, err
	}
	d.db.Update(func(tx *bbolt.Tx) error {
		for k, v := range data {
			if err := coll.Put([]byte(k), []byte(v)); err != nil {
				return err
			}
		}
		if err := coll.Put([]byte("id"), []byte(id.String())); err != nil {
			return err
		}
		return nil
	})
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
