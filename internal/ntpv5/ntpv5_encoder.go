package ntpv5

import (
	"encoding/binary"
	//"fmt"
)

func Encode(b []byte,d *Ntpv5Data) int {
	length := 48

	//var b []byte
	b[0] = byte( ((d.LI & 3) << 6) + (d.VN & 7) << 3 + (d.Mode & 7) )
	b[1] = byte(d.Stratum)
	b[2] = byte(d.Poll)
	b[3] = byte(d.Precision)
	b[4] = byte(d.Timescale)
	b[5] = byte(d.Era)
	b[6] = byte((d.Flags>>8) & 255)
	b[7] = byte(d.Flags & 255)

	tmp := make([]byte, 4)
	binary.BigEndian.PutUint32(tmp, uint32(d.RootDelay))
	copy(b[8:12], tmp[0:4])
	binary.BigEndian.PutUint32(tmp, uint32(d.RootDispersion))
	copy(b[12:16], tmp[0:4])

	tmp = make([]byte, 8)
        binary.BigEndian.PutUint64(tmp, uint64(d.ServerCookie))
	copy(b[16:24], tmp[0:8])
        binary.BigEndian.PutUint64(tmp, uint64(d.ClientCookie))
	copy(b[24:32], tmp[0:8])


        binary.BigEndian.PutUint64(tmp, uint64(d.ReceiveTimestamp))
	copy(b[32:40], tmp[0:8])
        binary.BigEndian.PutUint64(tmp, uint64(d.TransmitTimestamp))
	copy(b[40:48], tmp[0:8])

	if (d.ReferenceIDsResponseEx.Length != 0){
		ex := make([]byte, d.ReferenceIDsResponseEx.Length/8-1)
		ex[0], ex[1] = 0xF5, 0x04
		ex[2] = byte(d.ReferenceIDsResponseEx.Length >> 8)
		ex[3] = byte(d.ReferenceIDsResponseEx.Length & 255)
		b = append(b, ex...)

		length += int(d.ReferenceIDsResponseEx.Length)
	}

	return length
}
