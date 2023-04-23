package ntpv5

import (
	"fmt"
	"time"
)

// https://www.ietf.org/archive/id/draft-mlichvar-ntp-ntpv5-07.html#section-4
type Ntpv5Data struct {
	LI uint8
	VN uint8
	Mode uint8
	Stratum uint8
	Poll uint8
	Precision uint8
	Timescale uint8
	Era uint8
	Flags uint16
	RootDelay uint32
	RootDispersion uint32
	ServerCookie uint64
	ClientCookie uint64
	ReceiveTimestamp uint64
	TransmitTimestamp uint64

	//TODO: Extension
	PaddingEx Padding
	ReferenceIDsRequestEx ReferenceIDsRequest
	ReferenceIDsResponseEx ReferenceIDsResponse
	ServerInformationEx ServerInformation
	ReferenceTimestampEx ReferenceTimestamp
	SecondaryReceiveTimestampExs []SecondaryReceiveTimestamp
	DraftIdentificationEx DraftIdentification
	UnknownExs []Unknown
}

func (d Ntpv5Data) String() string{
	return fmt.Sprintf(
		"LI: %d, VN: %d, Mode: %d, Stratum: %d, Poll: %d, " +
		"Precision: %d, Timescale: %d, Era: %d, Flags: %d, " +
		"RootDelay: %d, RootDispersion: %d, ServerCookie: %d, " +
		"ClientCookie: %d, ReceiveTimestamp: %d, TransmitTimestamp: %d",
		d.LI, d.VN, d.Mode, d.Stratum, d.Poll,
		d.Precision, d.Timescale, d.Era, d.Flags,
		d.RootDelay, d.RootDispersion, d.ServerCookie,
		d.ClientCookie, d.ReceiveTimestamp, d.TransmitTimestamp)
}

func NewClientNtpv5Data() Ntpv5Data{
	return Ntpv5Data{
		LI: 0,
		VN: 5,
		Mode: 3,
		Stratum: 0,
		Poll: 1,
		Precision: 1,
		Timescale: 0,
		Era: 0,
		Flags: 0,
		RootDelay: 0,
		RootDispersion: 0,
		ServerCookie: 0,
		ClientCookie: 0xabcd,
		ReceiveTimestamp: 0,
		TransmitTimestamp: 0,
	}
}

func NewServerNtpv5Data() Ntpv5Data{
	return Ntpv5Data{
		LI: 0,
		VN: 5,
		Mode: 4,
		Stratum: 1,
		Poll: 1,
		Precision: 236,
		Timescale: 0,
		Era: 0,
		Flags: 0,
		RootDelay: 0,
		RootDispersion: 0,
		ServerCookie: 0,
		ClientCookie: 0,
		ReceiveTimestamp: 0,
		TransmitTimestamp: 0,

	}
}

type Padding struct {
	Length uint16
}
func (d Padding) String() string{
	return fmt.Sprintf(
		"Type: 0xF501, Name: Padding, Length: %d", d.Length)
}

type ReferenceIDsRequest struct {
	Length uint16
	Offset uint16
}
func (d ReferenceIDsRequest) String() string{
	return fmt.Sprintf(
		"Type: 0xF503, Name: ReferenceIDsRequest, Length: %d, Offset: %d",
		d.Length, d.Offset )
}

type ReferenceIDsResponse struct {
	Length uint16
	BloomFilterchunk []byte
}
func (d ReferenceIDsResponse) String() string{
	return fmt.Sprintf(
		"Type: 0xF504, Name: ReferenceIDsRequest, Length: %d, BloomFilterchunk: <ommited>",
		d.Length,)
}

type ServerInformation struct {
	Length uint16
	SupportedNtpVersions uint16
}
func (d ServerInformation) String() string{
	return fmt.Sprintf(
		"Type: 0xF505, Name: ServerInformation, Length: %d, SupportedNtpVersions: %d",
		d.Length, d.SupportedNtpVersions)
}

type ReferenceTimestamp struct {
	Length uint16
	ReferenceTimestamp uint64
}
func (d ReferenceTimestamp) String() string{
	return fmt.Sprintf(
		"Type: 0xF507, Name: ReferenceTimestamp, Length: %d, ReferenceTimestamp: %d",
		d.Length, d.ReferenceTimestamp )
}

type SecondaryReceiveTimestamp struct {
	Length uint16
	Timescale uint8
	Era uint8
	SecondaryReceiveTimestamp uint64
}
func (d SecondaryReceiveTimestamp) String() string{
	return fmt.Sprintf(
		"Type: 0xF509, Name: SecondaryReceiveTimestamp, Length: %d, Timescale: %d"+
		"Era: %d, SecondaryReceiveTimestamp:%d",
		d.Length, d.Timescale, d.Era, d.SecondaryReceiveTimestamp )
}

type DraftIdentification struct {
	Length uint16
	Draft string
}
func (d DraftIdentification) String() string{
	return fmt.Sprintf(
		"Type: 0xF5FF, Name: DraftIdentification, Length: %d, Draft: %s",
		d.Length, d.Draft )
}

type Unknown struct {
	Type uint16
	Length uint16
	Payload []byte
}
func (d Unknown) String() string{
	return fmt.Sprintf(
		"Type: %s, Name: Unknown Extension, Length: %d, Draft: %s",
		d.Type, d.Length, d.Payload )
}

// Utils
func GetTimestampNow() uint64 {
	now := time.Now()
	sec := uint64(now.Unix()) + 2208988800
	nanosec := uint64( (float64(now.Nanosecond())/1000000000) * (1<<32))
	timestamp := (sec << 32) +  nanosec

	return timestamp
}

