/*
Copyright © 2023 flano_yuki

*/
package cmd

import (
	"fmt"
	"net"
	"log"
	"time"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/flano-yuki/ntp5-go-exp/internal/ntpv5"
)

var clientCmd = &cobra.Command{
	Use:   "client HOSTNAME",
	Short: "ntpv5 client",
	Args: cobra.MinimumNArgs(1),
	Run: execClient,
}

func execClient(cmd *cobra.Command, args []string){
	// handle, flag and args
	port, _ := cmd.Flags().GetInt("port")
	host := args[0]

	refreq, _ := cmd.Flags().GetInt("refreq")
	padding, _ := cmd.Flags().GetInt("padding")
	info, _ := cmd.Flags().GetBool("info")
	draft, _ := cmd.Flags().GetString("draft")
	timescale, _ := cmd.Flags().GetInt("timescale")
	flags, _ := cmd.Flags().GetInt("flags")

	// struct packet data
	ntpv5data := ntpv5.NewClientNtpv5Data()
	ntpv5data.Timescale = uint8(timescale)
	ntpv5data.Flags = uint16(flags)

	if (padding > 0){
		ntpv5data.PaddingEx = ntpv5.Padding{
			Length: uint16(padding),
		}
	}
	if (refreq > 0){
		ntpv5data.ReferenceIDsRequestEx = ntpv5.ReferenceIDsRequest{
			Length: uint16(refreq),
			Offset: 0,
		}
	}
	if (info){
		ntpv5data.ServerInformationEx = ntpv5.ServerInformation{
			Length: uint16(8),
		}
	}
	if (len(draft) > 0){
		ntpv5data.DraftIdentificationEx = ntpv5.DraftIdentification{
			Length: uint16(4 + len(draft)),
			Draft: draft,
		}
	}
	buffer := ntpv5.Encode(ntpv5data)

	// udp send
	conn, err := net.Dial("udp", host + ":" + strconv.Itoa(port))
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	conn.SetReadDeadline(time.Now().Add(3 * time.Second))
	defer conn.Close()

	fmt.Println("Send NTPv5 Data(" + conn.RemoteAddr().String() + "): \n", ntpv5data)
	_, err = conn.Write(buffer)
	if err != nil {
		panic(err)
	}
	readBuffer := make([]byte, 1500)
	readLength, err := conn.Read(readBuffer)
	if err != nil {
		panic(err)
	}
	receivedNtpv5data := ntpv5.Decode(readBuffer[:readLength])
	fmt.Println("Received NTPv5 Data(" + conn.RemoteAddr().String() + "):\n", receivedNtpv5data)
	fmt.Println("\n")

}

func init() {
	rootCmd.AddCommand(clientCmd)
	clientCmd.Flags().IntP("port", "p", 123, "Target Port number")
	clientCmd.Flags().BoolP("verbose", "v", false, "verbose")
	clientCmd.Flags().IntP("refreq", "r", 0, "ReferenceIDsRequest length")
	clientCmd.Flags().IntP("flags", "f", 0, "Flags")
	clientCmd.Flags().IntP("padding", "d", 0, "Padding length")
	clientCmd.Flags().BoolP("info", "i", false, "Server Information")
	clientCmd.Flags().StringP("draft", "a", "", "Draft Identification")
	clientCmd.Flags().IntP("timescale", "t", 0, "Timescale type")
}
