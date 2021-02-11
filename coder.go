package bintly

import (
	"fmt"
	"github.com/viant/bintly/conv"
	"reflect"
	"sync"
)

type structFields struct {
	//indexes of the exported fields
	indexes     []int
	convertible []convField
}

type convField struct {
	index  int
	origin reflect.Type
	native reflect.Type
}

//structCoder represents a struct coder
type structCoder struct {
	mapped map[reflect.Type]*structFields
	ptr    *reflect.Value
	isNil  bool
	v      reflect.Value
	t      reflect.Type
	fields *structFields
}

func (c *structCoder) Alloc() uint32 {
	if c.isNil {
		return 0
	}
	return 1
}

//SetAlloc set allocation, if zero the pointer to struct is nil
func (c *structCoder) SetAlloc(allocation uint32) {
	if allocation == 0 {
		c.ptr = nil
	}
}

func (c *structCoder) setFields(t reflect.Type) error {
	if fields, ok := c.mapped[t]; ok {
		c.fields = fields
		return nil
	}
	fields := &structFields{
		indexes:     make([]int, 0),
		convertible: make([]convField, 0),
	}

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.PkgPath != "" {
			continue
		}
		if f.Anonymous {
			return fmt.Errorf("anonymous field %v not yet supported", f.Name)
		}
		if !conv.IsNative(f.Type) {
			if native := conv.MatchNative(f.Type); native != nil {
				fields.convertible = append(fields.convertible, convField{
					index:  i,
					origin: f.Type,
					native: *native,
				})
				continue
			}
		}
		fields.indexes = append(fields.indexes, i)
	}
	c.mapped[t] = fields
	c.fields = c.mapped[t]
	return nil
}

func (c *structCoder) set(v reflect.Value, t reflect.Type) error {
	c.isNil = false
	c.t = t
	c.v = v
	if v.Kind() == reflect.Ptr {
		c.ptr = &v
		c.isNil = v.IsNil()
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
	for _, f := range c.fields.convertible {
		v := c.v.Field(f.index).Convert(f.native)
		if err := stream.Any(v.Interface()); err != nil {
			return err
		}
	}
	return nil
}

//DecodeBinary decodes struct from reader
func (c *structCoder) DecodeBinary(stream *Reader) error {
	if c.ptr == nil {
		return nil
	}
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
	for _, f := range c.fields.convertible {
		v := reflect.New(f.native)
		if err := stream.Any(v.Interface()); err != nil {
			return err
		}
		c.v.Field(f.index).Set(v.Elem().Convert(f.origin))
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
