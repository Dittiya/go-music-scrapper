package db

import "errors"

type Storage interface {
	Save(string, string, int) (string, error)
	Load(string) (string, error)
}

var ErrNotFound = errors.New("not found")
