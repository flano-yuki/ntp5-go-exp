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

	// decode extensions
	extensions := b[48:]
	for {
		if len(extensions) < 4 {
			// TODO error Handling
			break
		}
		extensionType := (uint16(extensions[0]) << 8) + uint16(extensions[1])
		extensionLenght := (uint16(extensions[2]) << 8) + uint16(extensions[3])

		switch extensionType {
		    case 0xF503:
			ex := ReferenceIDsRequest{}
			ex.Length = extensionLenght
			ex.Offset = (uint16(extensions[4]) << 8) + uint16(extensions[5])
			ntpv5data.ReferenceIDsRequestEx = ex
		    case 0xF505:
			ex := ServerInformation{}
			ex.Length = extensionLenght
			ex.SupportedNtpVersions = (uint16(extensions[4]) << 8) + uint16(extensions[5])
			ntpv5data.ServerInformationEx = ex
		    case 0xF507:
			ex := ReferenceTimestamp{}
			ex.Length = extensionLenght
			binary.Read(bytes.NewReader(b[4:12]), binary.BigEndian, &ex.ReferenceTimestamp)
			ntpv5data.ReferenceTimestampEx = ex
		    case 0xF509:
			ex := SecondaryReceiveTimestamp{}
			ex.Length = extensionLenght
			ex.Timescale = extensions[4]
			ex.Era = extensions[5]
			binary.Read(bytes.NewReader(b[8:16]), binary.BigEndian, &ex.SecondaryReceiveTimestamp)
			ntpv5data.SecondaryReceiveTimestampEx = ex
		    case 0xF5FF:
			ex := DraftIdentification{}
			ex.Length = extensionLenght
			ex.draftVersion = string(extensions[4:(ex.Length)])
			ntpv5data.DraftIdentificationEx = ex
		}

		extensions = extensions[extensionLenght:]

	}

	return ntpv5data
}
