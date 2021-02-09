package bintly_test

import (
	"github.com/viant/bintly"
	"log"
)

//Employee represents a test struct
type Employee struct {
	ID       int
	Name     string
	RolesIDs []int
	Titles   []string
	DeptIDs  []int
}

func Example_Marshal() {
	emp := Employee{
		ID:       100,
		Name:     "test",
		RolesIDs: []int{1000, 1002, 1003},
		Titles:   []string{"Lead", "Principal"},
		DeptIDs:  []int{10, 13},
	}
	data, err := bintly.Marshal(emp)
	if err != nil {
		log.Fatal(err)
	}
	clone := Employee{}
	err = bintly.Unmarshal(data, &clone)
	if err != nil {
		log.Fatal(err)
	}
}

//DecodeBinary decodes data to binary stream
func (e *Employee) DecodeBinary(stream *bintly.Reader) error {
	stream.Int(&e.ID)
	stream.String(&e.Name)
	stream.Ints(&e.RolesIDs)
	stream.Strings(&e.Titles)
	stream.Ints(&e.DeptIDs)
	return nil
}

//EncodeBinary encodes data from binary stream
func (e *Employee) EncodeBinary(stream *bintly.Writer) error {
	stream.Int(e.ID)
	stream.String(e.Name)
	stream.Ints(e.RolesIDs)
	stream.Strings(e.Titles)
	stream.Ints(e.DeptIDs)
	return nil
}
