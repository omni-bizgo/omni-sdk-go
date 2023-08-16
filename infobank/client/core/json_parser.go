package core

import (
	"bytes"
	"encoding/json"
)

func MarshalAndConvertStr(vo interface{}) (*string, *error) {
	b, err := marshal(vo)
	if err != nil {
		return nil, &err
	} else {
		s := string(b)
		return &s, nil
	}
}

func marshal(vo interface{}) ([]byte, error) {
	return json.Marshal(vo)
}

func Unmarshal(data []byte, vo interface{}) error {
	des := json.NewDecoder(bytes.NewReader(data))
	des.DisallowUnknownFields()
	if err := des.Decode(vo); err != nil {
		return err
	} else {
		return nil
	}
}

func Valid(data []byte) bool {
	return json.Valid(data)
}
