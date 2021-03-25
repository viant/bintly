package maps

type SubMessage struct {
	Id   int
	Name string
}

type SubMessages []SubMessage

type Message struct {
	//M1 map[string]SubMessage
	//M1 map[string]*SubMessage
	M1 map[string][]SubMessage
	//M2 map[string]*[]SubMessage


//	M2 map[string][]SubMessage
	//M0  map[string]int
	//M1  map[int]*float32
	//M2  map[string][]int

	//	M3  map[byte][]byte

	//M2 map[string]SubMessage
	//M3 map[string]*SubMessage
	//Map
	//M4 *Map
}
