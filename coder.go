package bintly

import (
	"fmt"
	"reflect"
	"sync"
)

type structFields struct {
	//indexes of the exported fields
	indexes []int
}

//structCoder represents a struct coder
type structCoder struct {
	mapped map[reflect.Type]*structFields
	ptr    *reflect.Value
	v      reflect.Value
	t      reflect.Type
	fields *structFields
}

func (c *structCoder) setFields(t reflect.Type) error {
	if fields, ok := c.mapped[t]; ok {
		c.fields = fields
		return nil
	}
	fields := &structFields{indexes: make([]int, 0)}
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.PkgPath != "" {
			continue
		}
		if f.Anonymous {
			return fmt.Errorf("anonymous field %v not yet supported", f.Name)
		}
		fields.indexes = append(fields.indexes, i)
	}
	c.mapped[t] = fields
	c.fields = c.mapped[t]
	return nil
}

func (c *structCoder) set(v reflect.Value, t reflect.Type) error {
	c.t = t
	c.v = v
	if v.Kind() == reflect.Ptr {
		c.ptr = &v
		if !c.ptr.IsNil() {
			c.v = v.Elem()
		}
	}
	return c.setFields(t)
}

//EncodeBinary writes struct to stream
func (c *structCoder) EncodeBinary(stream *Writer) error {
	for _, i := range c.fields.indexes {
		v := c.v.Field(i).Interface()
		if err := stream.Any(v); err != nil {
			return err
		}
	}
	return nil
}

//DecodeBinary decodes struct from reader
func (c *structCoder) DecodeBinary(stream *Reader) error {
	if c.ptr.IsNil() {
		c.v = reflect.New(c.t).Elem()
		c.ptr.Elem().Set(c.v)
	}
	for _, i := range c.fields.indexes {
		v := c.v.Field(i).Addr().Interface()
		if err := stream.Any(v); err != nil {
			return err
		}
	}
	return nil
}

type structCoderPool struct {
	sync.Pool
}

func (s *structCoderPool) Get() *structCoder {
	return s.Pool.Get().(*structCoder)
}

func newStructCoderPool() *structCoderPool {
	return &structCoderPool{

		Pool: sync.Pool{
			New: func() interface{} {

				return &structCoder{
					mapped: make(map[reflect.Type]*structFields),
				}
			},
		},
	}
}

var structCoders = newStructCoderPool()
