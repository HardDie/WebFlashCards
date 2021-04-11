package bolt_wrapper

import (
	"errors"
	"os"
	"time"

	"github.com/boltdb/bolt"
)

type (
	BoltWrapper struct {
		db *bolt.DB
	}
)

var (
	ErrorDBNotExists = errors.New("bolt_wrapper: DB file is not exists")
	ErrorDBClosed    = errors.New("bolt_wrapper: DB is closed")
)

func NewBoltWrapperOpen(dbpath string) (bw *BoltWrapper, err error) {
	_, err = os.Stat(dbpath)
	if os.IsNotExist(err) {
		err = ErrorDBNotExists
		return
	}

	db, err := bolt.Open(dbpath, 0644, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return
	}
	bw = &BoltWrapper{
		db: db,
	}
	return
}

func NewBoltWrapperCreate(dbpath string) (bw *BoltWrapper, err error) {
	db, err := bolt.Open(dbpath, 0644, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return
	}
	bw = &BoltWrapper{
		db: db,
	}
	return
}
