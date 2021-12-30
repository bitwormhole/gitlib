package pktline

// Packet 表示一个pktline包
type Packet struct {
	Length int
	Head   string
	Body   []byte
}
