package bintly

var writers = NewWriters()

//Marshal converts e into []byte, or error
func Marshal(v interface{}) ([]byte, error) {
	stream := writers.Get()
	defer writers.Put(stream)
	return MarshalStream(stream, v)
}

//MarshalStream converts e into []byte, or error
func MarshalStream(stream *Writer, v interface{}) ([]byte, error) {
	err := stream.Any(v)
	if err != nil {
		return nil, err
	}
	bs := stream.Bytes()
	return bs, nil
}
