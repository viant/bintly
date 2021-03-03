package basic_struct

import (
	"github.com/viant/bintly"
)


func (m *Message) EncodeBinary(coder *bintly.Writer) error {
	coder.Int(int(m.A1))
	return nil
}

func (m *Message) DecodeBinary(coder *bintly.Reader) error {
	t := int(0)
	coder.Int(&t)
	m.A1 = Aint(t)
	return nil
}
