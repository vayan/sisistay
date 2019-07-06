package apiutils

import (
	"encoding/json"
	"errors"
	"io"
)

func ParseTo(input io.Reader, dest interface{}) error {
	if input == nil {
		return errors.New("MissingBody")
	}

	decoder := json.NewDecoder(input)
	err := decoder.Decode(dest)

	return err
}

func Serialize(response interface{}) []byte {
	payload, err := json.Marshal(response)

	if err != nil {
		// TODO
	}

	return payload
}
