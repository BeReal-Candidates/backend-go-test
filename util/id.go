package util

import (
	"github.com/jaevor/go-nanoid"
)

func GetNewID() (string, error) {
	canonicID, err := nanoid.Standard(21)
	if err != nil {
		return "", err
	}
	return canonicID(), nil
}

func MustGetNewID() string {
	v, err := GetNewID()
	if err != nil {
		panic(err)
	}
	return v
}
