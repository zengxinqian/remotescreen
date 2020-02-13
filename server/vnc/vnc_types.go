package vnc

import (
	"fmt"
	"io"

	"github.com/pkg/errors"

	"github.com/zengxinqian/remotescreen/object/pool"
)

const (
	ProtocolVersionMajor = 3
	ProtocolVersionMinor = 8
)

type RFBMessage interface {
	Encode(writer io.Writer) error
	Decode(reader io.Reader) error
}

type ProtocolVersion struct {
	Major int
	Minor int
}

var DefaultProtocolVersion = ProtocolVersion{Minor: ProtocolVersionMajor, Major: ProtocolVersionMinor}

func (p *ProtocolVersion) Encode(w io.Writer) error {

	_, err := fmt.Fprintf(w, "RFB %03d:%03d\n", p.Major, p.Minor)
	return err

}

func (p *ProtocolVersion) Decode(r io.Reader) error {

	buffer := pool.GetBytes(12)
	_, err := io.ReadFull(r, buffer)
	if err != nil {
		return errors.Wrap(err, "error reading ProtocolVersion Handshake")
	}

	l, err := fmt.Sscanf(string(buffer), "RFB %d.%d\n", &p.Minor, &p.Minor)
	if l != 2 {
		return errors.New("error parsing ProtocolVersion.")
	}
	if err != nil {
		return errors.Wrap(err, "error parsing ProtocolVersion.")
	}

	return nil

}

type PixelFormat struct {
	BPP        uint8
	Depth      uint8
	BigEndian  bool
	TrueColor  bool
	RedMax     uint16
	GreenMax   uint16
	BlueMax    uint16
	RedShift   uint8
	GreenShift uint8
	BlueShift  uint8
}
