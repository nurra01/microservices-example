package utils

import "encoding/json"

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
