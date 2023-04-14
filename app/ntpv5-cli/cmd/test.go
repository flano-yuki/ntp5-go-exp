/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"net"
	"os"
	"time"
	"log"
	"strconv"
	"strings"
	"encoding/json"
	"reflect"
	"io/ioutil"

	"github.com/spf13/cobra"
	"github.com/flano-yuki/ntp5-go-exp/internal/ntpv5"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test HOSTNAME",
	Short: "",
	Args: cobra.MinimumNArgs(1),
	Run: execTest,
}

func execTest(cmd *cobra.Command, args []string){
	port, _ := cmd.Flags().GetInt("port")
	timeoutSec, _ := cmd.Flags().GetInt("timeout")
	dir, _ := cmd.Flags().GetString("dir")
	host := args[0]

	fmt.Println("Connect: " + host + ":" + strconv.Itoa(port))

	files, _ := ioutil.ReadDir(dir)
	for _, file := range files {
		if (!strings.Contains(file.Name(), ".json")){
			continue
		}
		fmt.Println(dir + file.Name())

		f, _ := os.ReadFile(dir + file.Name())
		testcase := ntpv5.Testcase{}
		json.Unmarshal(f, &testcase)

		execTestcase(host + ":" + strconv.Itoa(port), timeoutSec, testcase)
	}

}

func execTestcase(target string, timeoutSec int, testcase ntpv5.Testcase){

	ntpv5data := testcase.SendNtpv5Data
	buffer := ntpv5.Encode(ntpv5data)

	// manipulate ntpv5data for test if testcase json specify it
	buffer = buffer[:len(buffer) - testcase.ForTestShrinkBufferSize]

	for _, overwrite := range testcase.ForTestOverwriteBuffer {
		buffer[overwrite.Index] = overwrite.Value
	}

	// udp send
	conn, err := net.Dial("udp", target)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	conn.SetReadDeadline(time.Now().Add( time.Duration(timeoutSec) * time.Second))
	defer conn.Close()

	_, err = conn.Write(buffer)
	if err != nil {
		panic(err)
	}
	readBuffer := make([]byte, 1500)
	isTimeout := false
	readLength, err := conn.Read(readBuffer)
	if err != nil {
		if netErr, isNetErr := err.(net.Error); isNetErr && netErr.Timeout() {
			isTimeout = true
		} else {
			panic(err)
		}
	}

	// Check Timeout test
	if testcase.ResponseTimeout && isTimeout {
		fmt.Println("	[OK] Response SHOULD timeout\n")
		return;
	} else if (!testcase.ResponseTimeout && isTimeout){
		fmt.Println("	[NG] Response SHOULD NOT timeout (but timeout)\n")
		return;
	} else if (testcase.ResponseTimeout && !isTimeout){
		fmt.Println("	[NG] Response SHOULD timeout (but received)")
	} else {
		fmt.Println("	[OK] Response SHOULD NOT timeout")
	}

	receivedNtpv5data := ntpv5.Decode(readBuffer[:readLength])
	reflectValue := reflect.ValueOf(receivedNtpv5data)

	// Check ResponseMatch tests
	for _, test := range testcase.ResponseMatch {
		receivedValue := getPropertyValue(test.Property, reflectValue)
		printCheck(test.Property, uint(receivedValue), uint(test.Value), true, false)
	}
	for _, test := range testcase.ResponseUnmatch {
		receivedValue := getPropertyValue(test.Property, reflectValue)
		printCheck(test.Property, uint(receivedValue), uint(test.Value), false, false)
	}
	for _, test := range testcase.ResponseMayMatch {
		receivedValue := getPropertyValue(test.Property, reflectValue)
		printCheck(test.Property, uint(receivedValue), uint(test.Value), true, true)
	}
	for _, test := range testcase.ResponseMayUnmatch {
		receivedValue := getPropertyValue(test.Property, reflectValue)
		printCheck(test.Property, uint(receivedValue), uint(test.Value), false, true)
	}

	fmt.Println("")
}

func printCheck(property string, receivedValue uint, testValue uint, match bool, optional bool){
	// Setup conditions
	matchStr := "	[OK] "
	unmatchStr := "	[NG] "
	auxiliaryVerb := "SHOULD"
	if (optional){
		matchStr = "	[OptionalOK] "
		unmatchStr = "	[OptionalNG] "
		auxiliaryVerb = "MAY"
	}
	if (!match){
		matchStr, unmatchStr = unmatchStr, matchStr
		auxiliaryVerb = auxiliaryVerb + " NOT"
	}

	// Output
	if (uint(receivedValue) == uint(testValue)){
		fmt.Print(matchStr)
	} else {
		fmt.Print(unmatchStr)
	}
	fmt.Println(property + " " + auxiliaryVerb +
	  " be " + strconv.FormatUint(uint64(testValue), 10) +
	  " (" + strconv.FormatUint(uint64(receivedValue), 10)  + ")")
}

func getPropertyValue(property string, v reflect.Value) uint {
	var returnValue uint
	if (strings.Contains(property, ".")){
		value := v
		parts := strings.Split(property, ".")
		for _, part := range parts {
			value = reflect.Indirect(value).FieldByName(part)
		}
		returnValue = uint(value.Uint())
	} else {
		returnValue = uint(v.FieldByName(property).Uint())
	}

	return returnValue
}

func init() {
	rootCmd.AddCommand(testCmd)
	testCmd.Flags().IntP("port", "p", 123, "Target Port number")
	testCmd.Flags().IntP("timeout", "t", 1, "Timeout Secound")
	testCmd.Flags().StringP("dir", "d", "./testcase/", "Testcase Directory")

}
