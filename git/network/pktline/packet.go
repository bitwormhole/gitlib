package pktline

// Packet ...
type Packet struct {

	// * '0000' Flush Packet (flush-pkt) - indicates the end of a message
	// * '0001' Delimiter Packet (delim-pkt) - separates sections of a message
	// * '0002' Response End Packet (response-end-pkt) - indicates the end of a response for stateless connections
	Special bool

	Length int
	Head   string
	Body   []byte
}
