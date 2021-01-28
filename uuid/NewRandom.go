package uuid

import (
	google_uuid "github.com/google/uuid"
)

func NewRandom() (string, error) {
	uuidObj, err := google_uuid.NewRandom()
	return uuidObj.String(), err
}
