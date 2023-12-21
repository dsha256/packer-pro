package packet

import (
	"context"
)

//encore:service
type Packet struct {
}

func initPacket() (*Packet, error) {

	return &Packet{}, nil
}

func (f *Packet) Shutdown(force context.Context) {
}
