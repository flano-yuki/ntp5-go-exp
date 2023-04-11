# ntp5-go-exp

This is an experimental NTPv5 server and client implementation in Go language.

This tool is primarily focused on protocol implementation. Therefore, it does not provide filtering algorithms or time alignment functions on the nodes.

## Client usage
Example:

```
$ go run ./app/ntpv5-cli/main.go client 127.0.0.1 -p 10123
Send NTPv5 Data(127.0.0.1:10123): 
 LI: 0, VN: 5, Mode: 3, Stratum: 0, Poll: 1, Precision: 1, Timescale: 0, Era: 0, Flags: 0, RootDelay: 0, RootDispersion: 0, ServerCookie: 0, ClientCookie: 43981, ReceiveTimestamp: 0, TransmitTimestamp: 0
Received NTPv5 Data(127.0.0.1:10123):
 LI: 0, VN: 5, Mode: 4, Stratum: 1, Poll: 1, Precision: 236, Timescale: 0, Era: 0, Flags: 0, RootDelay: 0, RootDispersion: 0, ServerCookie: 0, ClientCookie: 43981, ReceiveTimestamp: 16708366057156395717, TransmitTimestamp: 16708366057156853140
 
```

Options:

```
$ go run ./app/ntpv5-cli/main.go client --help
ntpv5 client

Usage:
  ntpv5-cli client HOSTNAME [flags]

Flags:
  -a, --draft string    Draft Identification
  -f, --flags int       Flags
  -h, --help            help for client
  -i, --info            Server Information
  -d, --padding int     Padding length
  -p, --port int        Target Port number (default 123)
  -r, --refreq int      ReferenceIDsRequest length
  -t, --timescale int   Timescale type
  -v, --verbose         verbose

```

## Server usage 
Example: 

```
$ go run main.go server -p 10123  
Receive NTPv5 Data(127.0.0.1:46938): 
 LI: 0, VN: 5, Mode: 3, Stratum: 0, Poll: 1, Precision: 1, Timescale: 0, Era: 0, Flags: 0, RootDelay: 0, RootDispersion: 0, ServerCookie: 0, ClientCookie: 43981, ReceiveTimestamp: 0, TransmitTimestamp: 0
Send NTPv5 Data(127.0.0.1:46938): 
 LI: 0, VN: 5, Mode: 4, Stratum: 1, Poll: 1, Precision: 236, Timescale: 0, Era: 0, Flags: 0, RootDelay: 0, RootDispersion: 0, ServerCookie: 0, ClientCookie: 43981, ReceiveTimestamp: 16708366712497773237, TransmitTimestamp: 16708366712498263731

```

Options:

```
$ go run main.go server --help
ntpv5 server

Usage:
  ntpv5-cli server [flags]

Flags:
  -b, --bind string     Bind Adress (default "0.0.0.0")
  -a, --draft string    Draft Identification (default "draft-mlichvar-ntp-ntpv5-07 ")
  -f, --flags int       Flags
  -h, --help            help for server
  -i, --info int        Server Information (default 16)
  -p, --port int        Target Port number (default 10123)
  -t, --timescale int   Timescale type
  -v, --verbose         verbose

```
