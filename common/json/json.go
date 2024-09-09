package json

import (
	"fmt"
	"reflect"

	jsoniter "github.com/json-iterator/go"
)

// Decoder represent decoder for client es
type Decoder struct{}

// IsJSONEqual : check jsons are equal
func IsJSONEqual(s1, s2 string) (bool, error) {
	var o1 interface{}
	var o2 interface{}
	var err error
	err = jsoniter.ConfigFastest.Unmarshal([]byte(s1), &o1)
	if err != nil {
		return false, fmt.Errorf("fail mashalling string 1 :: %s", err.Error())
	}
	err = jsoniter.ConfigFastest.Unmarshal([]byte(s2), &o2)
	if err != nil {
		return false, fmt.Errorf("fail mashalling string 2 :: %s", err.Error())
	}
	return reflect.DeepEqual(o1, o2), nil
}

// Decode mock of Unmarshal using jsoniter
func (u *Decoder) Decode(data []byte, v interface{}) error {
	return jsoniter.ConfigFastest.Unmarshal(data, v)
}

// Unmarshal mock of Unmarshal using jsoniter
func (u *Decoder) Unmarshal(data []byte, v interface{}) error {
	return jsoniter.ConfigFastest.Unmarshal(data, v)
}

// Marshal mock of Marshal using jsoniter
func (u *Decoder) Marshal(v interface{}) ([]byte, error) {
	return jsoniter.ConfigFastest.Marshal(v)
}
