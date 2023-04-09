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
}

func (p Ntpv5Data) String() string{
	return fmt.Sprintf(
		"LI: %d, VN: %d, Mode: %d, Stratum: %d, Poll: %d, " +
		"Precision: %d, Timescale: %d, Era: %d, Flags: %d, " +
		"RootDelay: %d, RootDispersion: %d, ServerCookie: %d, " +
		"ClientCookie: %d, ReceiveTimestamp: %d, TransmitTimestamp: %d",
		p.LI, p.VN, p.Mode, p.Stratum, p.Poll,
		p.Precision, p.Timescale, p.Era, p.Flags,
		p.RootDelay, p.RootDispersion, p.ServerCookie,
		p.ClientCookie, p.ReceiveTimestamp, p.TransmitTimestamp)
}

func NewClientNtpv5Data() *Ntpv5Data{
	return &Ntpv5Data{
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

func NewServerNtpv5Data() *Ntpv5Data{
	return &Ntpv5Data{
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

func GetTimestampNow() uint64 {
	now := time.Now()
	sec := uint64(now.Unix()) + 2208988800
	fmt.Println(now.Nanosecond())
	nanosec := uint64( (float64(now.Nanosecond())/1000000000) * (1<<32))
	timestamp := (sec << 32) +  nanosec

	return timestamp
}
