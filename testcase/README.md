# Testcase json
The test case specifies the request data to be sent, and checks that the received response values are as specified.

## Testcase filename Class
Each test case file has a 3-digit prefix.

- **0xx**: Sends valid data
- **1xx**: Sends data containing unknown values 
- **2xx**: Sends data containing invalid values that do not affect the server's response generation
- **3xx**: Sends data containing extensions (Possibly, The server does not support it.)
- **4xx**: Sends invalid data

## How to write testcase
See [000-simple.json](https://github.com/flano-yuki/ntp5-go-exp/blob/maub/testcase/000-simple.json).

example: 
```
{
	"Description": "Simple valid data set",
	"SendNtpv5Data": {
		"VN": 5,
		"Mode": 3,
		"ClientCookie": 1234,
		"TransmitTimestamp": 0,
		"PaddingEx": {
			"Length": 40
		}
	},
	"ResponseTimeout": false,
	"ResponseMatch": [
		{"Property": "VN", "Value": 5},
		{"Property": "Mode", "Value": 4}
	],
	"ResponseUnmatch": [
		{"Property": "Stratum", "Value": 0}
	],
	"ResponseMayUnmatch": [
		{"Property": "RootDelay", "Value": 0}
	]
}

```

- **Description**: Test case description.
- **SendNtpv5Data**: NTP request to be sent. The structure is defined in [ntpv5_structs.go](https://github.com/flano-yuki/ntp5-go-exp/blob/featuer/test/internal/ntpv5/ntpv5_structs.go)
- **ResponseTimeout**: Check response timeout
- **ResponseMatch**: Checks that the response value matches the specified value.
- **ResponseUnmatch**: Checks that the response value does not matches the specified value.
- **ResponseMayMatch**: Checks that the response value matches the specified value. However, even if they do not match, it is not a violation of the specification.
- **ResponseMayUnmatch**: Checks that the response value does not matches the specified value. However, even if they do not match, it is not a violation of the specification.
