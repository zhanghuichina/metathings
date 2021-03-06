package common

import (
	"math/rand"
	"strings"

	uuid "github.com/satori/go.uuid"
)

func NewId() string {
	return strings.Replace(uuid.NewV4().String(), "-", "", -1)
}

func NewUint64Id() uint64 {
	return rand.Uint64()
}
