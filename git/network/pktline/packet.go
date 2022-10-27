package pktline

// Packet ...
type Packet struct {
	Length int
	Head   string
	Body   []byte
}
