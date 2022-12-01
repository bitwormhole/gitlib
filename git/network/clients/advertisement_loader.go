package clients

import (
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/bitwormhole/gitlib/git/network/pktline"
	"github.com/bitwormhole/starter/vlog"
)

// AdvertisementLoader ...
type AdvertisementLoader struct {
	context *Context

	wantService string
	HEAD        string
	refs        map[string]string
}

// Load ...
func (inst *AdvertisementLoader) Load() error {

	// init
	inst.refs = make(map[string]string)

	// open new connection
	conn1 := inst.context.Connection
	conn2, err := conn1.NewConnection(&pktline.ConnParams{
		Service: pktline.ServiceGitUploadPack,
	})
	if err != nil {
		return err
	}

	// open reader
	reader, ctype, err := conn2.OpenReader()
	if err != nil {
		return err
	}
	defer func() { reader.Close() }()

	if ctype != pktline.TypeGitUploadPackAdvertisement {
		return fmt.Errorf("bad Advertisement content-type:%v", ctype)
	}

	for {
		p, err := reader.Read()
		inst.handlePacket(p)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
	}

	return nil
}

func (inst *AdvertisementLoader) handlePacket(p *pktline.Packet) error {

	if p == nil {
		return nil
	}

	if p.Special {
		return nil
	}

	src := strings.Split(p.Head, string(' '))
	dst := make([]string, 0)
	for _, item := range src {
		item = strings.TrimSpace(item)
		if item != "" {
			dst = append(dst, item)
		}
	}
	count := len(dst)

	if count == 2 {
		str0 := dst[0]
		str1 := dst[1]
		if str0 == "#" {
			// service=xxx
			return inst.handlePacketAsService(str0, str1, p)
		} else if inst.isHashID(str0) {
			if str1 == "HEAD" {
				return inst.handlePacketAsHEAD(str0, str1, p)
			}
			// normal ref
			return inst.handlePacketAsRef(str0, str1, p)
		}
	}

	vlog.Warn("unsupported pktline packet head: ", p.Head)
	return nil
}

func (inst *AdvertisementLoader) handlePacketAsService(s0, s1 string, p *pktline.Packet) error {
	inst.wantService = s1
	return nil
}

func (inst *AdvertisementLoader) handlePacketAsHEAD(s0, s1 string, p *pktline.Packet) error {

	body := string(p.Body)
	options := strings.Split(body, string(' '))
	sort.Strings(options)

	inst.HEAD = s0
	return nil
}

func (inst *AdvertisementLoader) handlePacketAsRef(s0, s1 string, p *pktline.Packet) error {
	inst.refs[s1] = s0
	return nil
}

func (inst *AdvertisementLoader) isHashID(str string) bool {
	const minSize = 32
	size := len(str)
	if size < minSize {
		return false
	}
	if (size & 0x01) != 0 {
		return false
	}
	chs := []rune(str)
	for _, ch := range chs {
		if '0' <= ch && ch <= '9' {
			// NOP
		} else if 'a' <= ch && ch <= 'f' {
			// NOP
		} else {
			return false
		}
	}
	return true
}
