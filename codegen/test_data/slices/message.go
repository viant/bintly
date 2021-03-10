package slices

type SubMessage struct {
	Id   int
	Name string
}

type SubMessages []SubMessage
type SubMessages2 []*SubMessage

type Message struct {
	M1 []SubMessage
	M2 []*SubMessage
	*SubMessages
	SubMessages2
}
