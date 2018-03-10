package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/coreos/bbolt"
)

// Clipboard database store settings
const (
	DBDir  = ".clipstore"
	DBName = "store.db"
	BUCKET = "clipboard"
)

func init() {
	mustHaveDBFile()
	mustHaveBucket()
}

func getAll() (map[string]string, error) {
	db := openDB()
	defer db.Close()

	clips := make(map[string]string)
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BUCKET))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			clips[string(k)] = string(v)
		}
		return nil
	})
	if err != nil {
		return nil, errors.New("Error occurs while retrieving all keys")
	}
	if clips == nil {
		return nil, errors.New("There is no clip data")
	}
	return clips, err
}

func get(key string) (string, error) {
	db := openDB()
	defer db.Close()

	v := make([]byte, 0)
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BUCKET))
		v = b.Get([]byte(key))
		return nil
	})
	if err != nil {
		return "", fmt.Errorf("Error occurs while retrieving key %s", key)
	}
	if v == nil {
		return "", fmt.Errorf("Key %s does not exist", key)
	}
	return string(v), nil
}

func set(key, value string) error {
	db := openDB()
	defer db.Close()

	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BUCKET))
		e := b.Put([]byte(key), []byte(value))
		return e
	})
	if err != nil {
		return fmt.Errorf("Error occurs while setting %s:%s", key, value)
	}
	return nil
}

func del(key string) error {
	db := openDB()
	defer db.Close()

	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BUCKET))
		e := b.Delete([]byte(key))
		return e
	})
	if err != nil {
		return fmt.Errorf("Error occurs while deleting key %s", key)
	}
	return nil
}

func openDB() *bolt.DB {
	home := homeDir()
	dbFilePath := filepath.Join(home, DBDir, DBName)
	db, err := bolt.Open(dbFilePath, os.FileMode(0600), nil)
	checkError(err)
	return db
}

func mustHaveDBFile() {
	home := homeDir()
	dbFilePath := filepath.Join(home, DBDir, DBName)
	if _, err := os.Stat(dbFilePath); err != nil {
		dir := filepath.Dir(dbFilePath)
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			os.MkdirAll(dir, 0766)
		}
		if _, err := os.Create(dbFilePath); err != nil {
			panic(err)
		}
	}
}

func mustHaveBucket() {
	db := openDB()
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte(BUCKET))
		if err != nil {
			return fmt.Errorf("Create bucket: %s", err)
		}
		return nil
	})
}
