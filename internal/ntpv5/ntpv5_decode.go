package ntpv5

import (
	"encoding/binary"
	"bytes"
)

func Decode(b []byte) Ntpv5Data {
	ntpv5data := Ntpv5Data{}
	ntpv5data.LI = (b[0] >> 6) & 3
	ntpv5data.VN = (b[0] >> 3) & 7
	ntpv5data.Mode = b[0] & 7
	ntpv5data.Stratum = b[1]
	ntpv5data.Poll = b[2]
	ntpv5data.Precision = b[3]

	ntpv5data.Timescale = b[4]
	ntpv5data.Era = b[5]
	ntpv5data.Flags = (uint16(b[6]) << 8) + uint16(b[7])

	binary.Read(bytes.NewReader(b[8:12]), binary.BigEndian, &ntpv5data.RootDelay)
	binary.Read(bytes.NewReader(b[12:16]), binary.BigEndian, &ntpv5data.RootDispersion)
	binary.Read(bytes.NewReader(b[16:24]), binary.BigEndian, &ntpv5data.ServerCookie)
	binary.Read(bytes.NewReader(b[24:32]), binary.BigEndian, &ntpv5data.ClientCookie)
	binary.Read(bytes.NewReader(b[32:40]), binary.BigEndian, &ntpv5data.ReceiveTimestamp)
	binary.Read(bytes.NewReader(b[40:48]), binary.BigEndian, &ntpv5data.TransmitTimestamp)

	return ntpv5data
}
