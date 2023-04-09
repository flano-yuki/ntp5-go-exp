package ntpv5

import (
	"encoding/binary"
)

func Encode(b []byte,p *Ntpv5Data)  {
	//var b []byte
	b[0] = byte( ((p.LI & 3) << 6) + (p.VN & 7) << 3 + (p.Mode & 7) )
	b[1] = byte(p.Stratum)
	b[2] = byte(p.Poll)
	b[3] = byte(p.Precision)
	b[4] = byte(p.Timescale)
	b[5] = byte(p.Era)
	b[6] = byte((p.Flags>>8) & 255)
	b[7] = byte(p.Flags & 255)

	tmp := make([]byte, 4)
	binary.BigEndian.PutUint32(tmp, uint32(p.RootDelay))
	copy(b[8:12], tmp[0:4])
	binary.BigEndian.PutUint32(tmp, uint32(p.RootDispersion))
	copy(b[12:16], tmp[0:4])

	tmp = make([]byte, 8)
        binary.BigEndian.PutUint64(tmp, uint64(p.ServerCookie))
	copy(b[16:24], tmp[0:8])
        binary.BigEndian.PutUint64(tmp, uint64(p.ClientCookie))
	copy(b[24:32], tmp[0:8])


        binary.BigEndian.PutUint64(tmp, uint64(p.ReceiveTimestamp))
	copy(b[32:40], tmp[0:8])
        binary.BigEndian.PutUint64(tmp, uint64(p.TransmitTimestamp))
	copy(b[40:48], tmp[0:8])
//	return b
}
