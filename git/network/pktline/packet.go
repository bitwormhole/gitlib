package pktline

// Packet ...
type Packet struct {
	Length int
	Flush  bool
	Head   string
	Body   []byte
}
