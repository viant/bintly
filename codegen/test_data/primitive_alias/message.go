package primitive_alias

type Aint int
type SliceInt []int
type Aints []Aint

//type Cbyte time.Time

type Message struct {
	C1 SliceInt
	A1 Aint
	D1 Aints
//	E1 Cbyte
}
