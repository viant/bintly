# Bintly (super fast binary serialization for go) 

[![GoReportCard](https://goreportcard.com/badge/github.com/viant/bintly)](https://goreportcard.com/report/github.com/viant/bintly)
[![GoDoc](https://godoc.org/github.com/viant/bintly?status.svg)](https://godoc.org/github.com/viant/bintly)

This library is compatible with Go 1.11+

Please refer to [`CHANGELOG.md`](CHANGELOG.md) if you encounter breaking changes.

- [Motivation](#motivation)
- [Introduction](#introduction)
- [Contribution](#contributing-to-bintly)
- [License](#license)

## Motivation

The goal of library to provide super fast binary oriented decoding and encoding capability for any go data type, critical
for low latency applications.


## Introduction

Typical streamlined binary serialization format store primitive types with their native size, and all collection type
got pre seeded with the repeated data size. Imagine the follow struct:

```go
type Employee struct {
	ID int
	Name string
	RolesIDs []int
	Titles []string
    DeptIDs []int
}

var emp := Employee{
    ID: 100,
    Name: "test",
    RolesIDs: []int{1000,1002,1003},
    Titles: []string{"Lead", "Principal"},
    DeptIDs: []int{10,13},
}
```
This maps to the following binary stream representation:
```
100,4,test,3,1000,1002,1003,2,4,Lead,9,Principal,2,10,13
```

In examples presented coma got preserved only for visualisation, also numeric/alphanumerics usage is for simplification.

When decoding this binary format each repeated type requires new memory allocation, in this case 6 allocations:
3 for slices, and 3 for string type. 

Since it's possible to copy any primitive slice to memory back and forth, we can go about binary serialization way faster than the originally presented approach.
Instead of allocation memory for each repeated type (string,slice), we could simply reduce number allocation to number of 
primitive data type used + 1 to track allocations.
In that case binary data stream for emp variable will look like the following. 

```yaml
alloc: [4,3,2,4,9,2] 
ints: [100,1000,1002,1003,10,13]
uint8s: [test,Lead,Principal]
```


## Usage

```go
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
```

#### Working with Map


#### Working with Objects


### Benchmark

Benchmark uses [BenchStruct](stress/bench.go)  where slices got populated with 80 random items.

```bash
BenchmarkUnmarshalBintly
BenchmarkUnmarshalBintly-16           	  886302	      1311 ns/op	    3762 B/op	       6 allocs/op
BenchmarkMarshalBintly
BenchmarkMarshalBintly-16             	  884871	      1246 ns/op	    2484 B/op	       3 allocs/op
BenchmarkUnmarshalBintlyReflect
BenchmarkUnmarshalBintlyReflect-16    	  659103	      1670 ns/op	    3796 B/op	       7 allocs/op
BenchmarkMarshalBintlyReflect
BenchmarkMarshalBintlyReflect-16      	  643982	      1574 ns/op	    2506 B/op	      10 allocs/op
BenchmarkUnmarshalBinary
BenchmarkUnmarshalBinary-16           	  324630	      3406 ns/op	    3152 B/op	      78 allocs/op
BenchmarkMarshalBinary
BenchmarkMarshalBinary-16             	  241238	      4949 ns/op	    4536 B/op	       7 allocs/op
BenchmarkMarshalGob
BenchmarkMarshalGob-16                	   99133	     10373 ns/op	    9805 B/op	      36 allocs/op
BenchmarkMarshalCbor
BenchmarkMarshalCbor-16               	  234380	      5057 ns/op	    2194 B/op	       2 allocs/op
BenchmarkUnmarshalCbor
BenchmarkUnmarshalCbor-16             	  109696	     10075 ns/op	    3472 B/op	      84 allocs/op
BenchmarkMarshalMsgPack
BenchmarkMarshalMsgPack-16            	   93470	     11834 ns/op	    4722 B/op	       8 allocs/op
BenchmarkUnmarshalMsgPack
BenchmarkUnmarshalMsgPack-16          	   63546	     17666 ns/op	    4899 B/op	      89 allocs/op
BenchmarkUnMarshalGob
BenchmarkUnMarshalGob-16              	   44133	     26423 ns/op	   13928 B/op	     322 allocs/op
BenchmarkJSONUnmarshal
BenchmarkJSONUnmarshal-16             	   18291	     62308 ns/op	   15296 B/op	     310 allocs/op
BenchmarkJSONMarshal
BenchmarkJSONMarshal-16               	   65025	     18538 ns/op	    3717 B/op	       3 allocs/op
```



<a name="License"></a>
## License

The source code is made available under the terms of the Apache License, Version 2, as stated in the file `LICENSE`.

Individual files may be made available under their own specific license,
all compatible with Apache License, Version 2. Please see individual files for details.

<a name="Credits-and-Acknowledgements"></a>

## Contributing to Bintly

Bintly is an open source project and contributors are welcome!

See [TODO](TODO.md) list

## Credits and Acknowledgements

**Library Author:** Adrian Witas

