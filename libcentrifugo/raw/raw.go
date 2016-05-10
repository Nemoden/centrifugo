// package raw contains Raw type (alias to slice of bytes). This type used by Centrifugo
// as type for fields in structs which value we want to left untouched. For example custom
// application specific JSON payload data in published message. This is very similar to
// json.RawMessage type but have some extra methods to fit gogoprotobuf custom type interface.
package raw

import (
	"bytes"
	"errors"
)

type Raw []byte

func (r Raw) Marshal() ([]byte, error) {
	if len(r) == 0 {
		return nil, nil
	}
	return []byte(r), nil
}

func (r Raw) MarshalTo(data []byte) (n int, err error) {
	if len(r) == 0 {
		return 0, nil
	}
	copy(data, r)
	return len(r), nil
}

func (r *Raw) Unmarshal(data []byte) error {
	if len(data) == 0 {
		r = nil
		return nil
	}
	id := Raw(make([]byte, len(data)))
	copy(id, data)
	*r = id
	return nil
}

func (r *Raw) Size() int {
	if r == nil {
		return 0
	}
	return len(*r)
}

// MarshalJSON returns *m as the JSON encoding of m.
func (r *Raw) MarshalJSON() ([]byte, error) {
	return *r, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (r *Raw) UnmarshalJSON(data []byte) error {
	if r == nil {
		return errors.New("Raw: UnmarshalJSON on nil pointer")
	}
	*r = append((*r)[0:0], data...)
	return nil
}

func (r Raw) Equal(other Raw) bool {
	return bytes.Equal(r[0:], other[0:])
}

func (r Raw) Compare(other Raw) int {
	return bytes.Compare(r[0:], other[0:])
}

type intn interface {
	Intn(n int) int
}

func NewPopulatedRaw(r intn) *Raw {
	v1 := r.Intn(100)
	data := make([]byte, v1)
	for i := 0; i < v1; i++ {
		data[i] = byte(r.Intn(256))
	}
	raw := Raw(data)
	return &raw
}
