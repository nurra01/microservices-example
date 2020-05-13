package utils

import (
	"encoding/json"
	"io"
)

// FromJSON deserializes the object from JSON string
// in an io.Reader to the given interface
func FromJSON(i interface{}, r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(i)
}

// ToJSON serializes the given interface into a string based JSON format
func ToJSON(i interface{}, w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(i)
}

// FromByteToObject converts byte stream to object structure
func FromByteToObject(data []byte, obj interface{}) error {
	err := json.Unmarshal(data, obj)
	if err != nil {
		return err
	}
	return nil
}

// FromObjectToByte converts object structure to byte stream
func FromObjectToByte(obj interface{}) ([]byte, error) {
	data, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return data, nil
}
