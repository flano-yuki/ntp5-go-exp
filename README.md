# ntp5-go-exp

This is an experimental NTPv5 client, server and server-test tool implementation in Go language.

This tool is primarily focused on protocol implementation. Therefore, it does not provide filtering algorithms or time alignment functions on the nodes.

This is implemented with reference to `draft-ietf-ntp-ntpv5-00` .

## Server Test Tool usage
The server test tool sends a packet specified in a json format test case and checks the value of the received response.

The default test cases are in the [./testcase](https://github.com/flano-yuki/ntp5-go-exp/tree/main/testcase) directory. and, The explanation of test case json is [here](https://github.com/flano-yuki/ntp5-go-exp/tree/main/testcase#testcase-json).

```
$ go run ./app/ntpv5-cli/main.go test localhost -p 10123 
Connect: localhost:10123

./testcase/000-simple.json
	[OK] Response SHOULD NOT timeout
	[OK] Request size and Response size SHOULD be same (88, 88)
	[OK] VN SHOULD be 5 (5)
	[OK] Mode SHOULD be 4 (4)
	[OK] ClientCookie SHOULD be 1234 (1234)
	[OK] ServerCookie SHOULD be 0 (0)
	[OK] Stratum SHOULD NOT be 0 (2)
	[OK] ReceiveTimestamp SHOULD NOT be 0 (16710952955318547277)
	[OK] TransmitTimestamp SHOULD NOT be 0 (16710952955319174996)

./testcase/001-refid-request.json
	[OK] Response SHOULD NOT timeout
	[OK] Request size and Response size SHOULD be same (148, 148)
	[OK] ReferenceIDsResponseEx.Length SHOULD NOT be 0 (100)

...
```

Options:

```
$ go run ./app/ntpv5-cli/main.go test --help
Usage:
  ntpv5-cli test HOSTNAME [flags]

Flags:
  -d, --dir string    Testcase Directory (default "./testcase/")
  -h, --help          help for test
  -p, --port int      Target Port number (default 123)
  -t, --timeout int   Timeout Secound (default 1)

```

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
  -a, --draft string    Draft Identification (default "draft-ietf-ntp-ntpv5-00")
  -f, --flags int       Flags
  -h, --help            help for server
  -i, --info int        Server Information (default 16)
  -p, --port int        Target Port number (default 10123)
  -t, --timescale int   Timescale type
  -v, --verbose         verbose

```
