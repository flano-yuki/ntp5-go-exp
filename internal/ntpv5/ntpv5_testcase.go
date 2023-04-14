package ntpv5

import (
//	"fmt"
)

type Testcase struct {
	Description string

	// Send Data
        SendNtpv5Data Ntpv5Data

	// For test manipulate send data
        ForTestShrinkBufferSize int
        ForTestOverwriteBuffer []OverwriteData

	// received data check
        ResponseTimeout bool
        ResponseMatch []PropertyTest
        ResponseUnmatch []PropertyTest
        ResponseMayMatch []PropertyTest
        ResponseMayUnmatch []PropertyTest
}

type PropertyTest  struct {
        Property string
	Value uint // TODO: interface{} ?
}
type OverwriteData  struct {
        Index uint
	Value byte
}
