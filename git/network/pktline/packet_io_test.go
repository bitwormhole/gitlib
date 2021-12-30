package pktline

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/bitwormhole/starter/util"
)

func TestPacketIO(t *testing.T) {
	ter := packetIOTester{}
	err := ter.test()
	if err != nil {
		t.Error(err)
	}
}

type packetIOTester struct{}

func (inst *packetIOTester) test() error {

	buffer := bytes.Buffer{}
	reader := StreamReadCloser{}
	writer := StreamWriteCloser{}
	plist1 := inst.prepareTestingPackets()

	writer.InitWithWriter(&buffer)
	for _, p := range plist1 {
		writer.Write(p)
	}

	plist2 := make([]*Packet, 0)
	reader.InitWithReader(&buffer)
	for {
		p, err := reader.Read()
		if err == nil {
			plist2 = append(plist2, p)
		} else {
			fmt.Println("err: ", err.Error())
			break
		}
	}

	len1 := len(plist1)
	len2 := len(plist2)
	if len1 != len2 {
		return errors.New("len(packs1) != len(packs2)")
	}
	for i := 0; i < len1; i++ {
		err := inst.compareForEq(plist1[i], plist2[i])
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *packetIOTester) getBody(p *Packet) []byte {
	b := p.Body
	if b == nil {
		b = []byte{}
	}
	return b
}

func (inst *packetIOTester) compareForEq(p1, p2 *Packet) error {

	if p1.Length != p2.Length {
		if len(p1.Head) <= 4 {
			return errors.New("bad pktline.packet.length")
		}
	}

	if p1.Head != p2.Head {
		return errors.New("bad pktline.packet.head")
	}

	b1 := inst.getBody(p1)
	b2 := inst.getBody(p2)
	if len(b1) != len(b2) {
		return errors.New("bad pktline.packet.body")
	}
	n := bytes.Compare(b1, b2)
	if n != 0 {
		return errors.New("bad pktline.packet.body")
	}

	return nil
}

func (inst *packetIOTester) prepareTestingPackets() []*Packet {
	plist := make([]*Packet, 0)

	//	plist = plist[0:0]

	p := &Packet{}
	plist = append(plist, p)

	p = &Packet{}
	p.Length = 1
	plist = append(plist, p)

	p = &Packet{}
	p.Length = 4
	plist = append(plist, p)

	p = &Packet{}
	p.Length = 0
	p.Head = inst.makeLargeString(32)
	p.Body = inst.makeLargeBytes(32)
	plist = append(plist, p)

	p = &Packet{}
	p.Length = 0
	p.Head = inst.makeLargeString(32760) + "\n"
	p.Body = inst.makeLargeBytes(32760)
	plist = append(plist, p)

	return plist
}

func (inst *packetIOTester) makeLargeString(wantSize int) string {
	ha := sha256.New()
	sum := []byte{}
	buffer := strings.Builder{}
	for buffer.Len() < wantSize {
		sum = ha.Sum(sum)
		buffer.WriteString(util.StringifyBytes(sum))
	}
	res := buffer.String()
	return res[0:wantSize]
}

func (inst *packetIOTester) makeLargeBytes(wantSize int) []byte {
	ha := sha256.New()
	sum := []byte{}
	buffer := bytes.Buffer{}
	for buffer.Len() < wantSize {
		sum = ha.Sum(sum)
		buffer.Write(sum)
	}
	res := buffer.Bytes()
	return res[0:wantSize]
}
