package store

import (
	"fmt"
	"github.com/siddontang/ledisdb/config"
	"github.com/siddontang/ledisdb/store/driver"
	"os"
	"path"

	_ "github.com/siddontang/ledisdb/store/boltdb"
	_ "github.com/siddontang/ledisdb/store/goleveldb"
	_ "github.com/siddontang/ledisdb/store/hyperleveldb"
	_ "github.com/siddontang/ledisdb/store/leveldb"
	_ "github.com/siddontang/ledisdb/store/mdb"
	_ "github.com/siddontang/ledisdb/store/rocksdb"
)

func getStorePath(cfg *config.Config) string {
	return path.Join(cfg.DataDir, fmt.Sprintf("%s_data", cfg.DBName))
}

func Open(cfg *config.Config) (*DB, error) {
	s, err := driver.GetStore(cfg)
	if err != nil {
		return nil, err
	}

	path := getStorePath(cfg)

	if err := os.MkdirAll(path, 0755); err != nil {
		return nil, err
	}

	idb, err := s.Open(path, cfg)
	if err != nil {
		return nil, err
	}

	db := &DB{idb, s.String()}

	return db, nil
}

func Repair(cfg *config.Config) error {
	s, err := driver.GetStore(cfg)
	if err != nil {
		return err
	}

	path := getStorePath(cfg)

	return s.Repair(path, cfg)
}

func init() {
}
