package ntpv5

import (
	"encoding/binary"
	//"fmt"
)

func Encode(d Ntpv5Data) []byte {
	b := make([]byte, 48)
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

	if (d.PaddingEx.Length > 0){
		ex := make([]byte, d.PaddingEx.Length)
		ex[0], ex[1] = 0xF5, 0x01
		ex[2] = byte(d.PaddingEx.Length >> 8)
		ex[3] = byte(d.PaddingEx.Length & 255)

		b = append(b, ex...)
		length += int(d.PaddingEx.Length)
	}
	if (d.ReferenceIDsRequestEx.Length > 0){
		ex := make([]byte, d.ReferenceIDsRequestEx.Length)
		ex[0], ex[1] = 0xF5, 0x03
		ex[2] = byte(d.ReferenceIDsRequestEx.Length >> 8)
		ex[3] = byte(d.ReferenceIDsRequestEx.Length & 255)

		ex[4] = byte(d.ReferenceIDsRequestEx.Offset >> 8)
		ex[5] = byte(d.ReferenceIDsRequestEx.Offset & 255)

		b = append(b, ex...)

		length += int(d.ReferenceIDsRequestEx.Length)
	}
	if (d.ReferenceIDsResponseEx.Length > 0){
		ex := make([]byte, d.ReferenceIDsResponseEx.Length)
		ex[0], ex[1] = 0xF5, 0x04
		ex[2] = byte(d.ReferenceIDsResponseEx.Length >> 8)
		ex[3] = byte(d.ReferenceIDsResponseEx.Length & 255)

		b = append(b, ex...)

		length += int(d.ReferenceIDsResponseEx.Length)
	}
	if (d.ServerInformationEx.Length > 0){
		ex := make([]byte, d.ServerInformationEx.Length)
		ex[0], ex[1] = 0xF5, 0x05
		ex[2] = byte(d.ServerInformationEx.Length >> 8)
		ex[3] = byte(d.ServerInformationEx.Length & 255)

		ex[4] = byte(d.ServerInformationEx.SupportedNtpVersions >> 8)
		ex[5] = byte(d.ServerInformationEx.SupportedNtpVersions & 255)


		b = append(b, ex...)

		length += int(d.ServerInformationEx.Length)
	}
	if (d.ReferenceTimestampEx.Length > 0){
		ex := make([]byte, 12)
		ex[0], ex[1] = 0xF5, 0x07
		ex[2] = byte(d.ReferenceTimestampEx.Length >> 8)
		ex[3] = byte(d.ReferenceTimestampEx.Length & 255)

		binary.BigEndian.PutUint64(tmp, uint64(d.ReferenceTimestampEx.ReferenceTimestamp))
		copy(ex[4:12], tmp[0:8])

		b = append(b, ex...)

		length += int(d.ReferenceTimestampEx.Length)
	}
	if (d.DraftIdentificationEx.Length > 0){
		ex := make([]byte, 4)
		ex[0], ex[1] = 0xF5, 0xFF
		ex[2] = byte(d.DraftIdentificationEx.Length >> 8)
		ex[3] = byte(d.DraftIdentificationEx.Length & 255)

		ex = append(ex, []byte(d.DraftIdentificationEx.Draft)...)

		b = append(b, ex...)

		length += int(d.DraftIdentificationEx.Length)
	}
	for _, unknown := range d.UnknownExs {

		ex := make([]byte, unknown.Length)
		ex[0], ex[1] = byte(unknown.Type >> 8), byte(unknown.Type & 255)
		ex[2] = byte(unknown.Length >> 8)
		ex[3] = byte(unknown.Length & 255)

		b = append(b, ex...)
		length += int(unknown.Length)
	}

	return b[:length]
}
