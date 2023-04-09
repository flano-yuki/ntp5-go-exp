/*
Copyright © 2023 flano_yuki

*/
package cmd

import (
	"fmt"
	"net"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/flano-yuki/ntp5-go-exp/internal/ntpv5"
)

var clientCmd = &cobra.Command{
	Use:   "client HOSTNAME",
	Short: "ntp",
	Args: cobra.MinimumNArgs(1),
	Run: execClient,
}

func execClient(cmd *cobra.Command, args []string){
	// handle, flag and args
	port, _ := cmd.Flags().GetInt("port")
	host := args[0]

	// struct packet data
	ntpv5data := ntpv5.NewClientNtpv5Data()
	fmt.Println("Send NTPv5 Data")
	fmt.Printf("", ntpv5data)
	fmt.Println("\n")

	buffer := make([]byte, 48)
	ntpv5.Encode(buffer, ntpv5data)


	// udp send
	conn, err := net.Dial("udp", host + ":" + strconv.Itoa(port))
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	defer conn.Close()

	_, err = conn.Write(buffer)
	if err != nil {
		panic(err)
	}
	readBuffer := make([]byte, 1500)
	length, err := conn.Read(readBuffer)
	if err != nil {
		panic(err)
	}
	receivedNtpv5data := ntpv5.Decode(readBuffer[:length])
	fmt.Println("Received NTPv5 Data")
	fmt.Printf("", receivedNtpv5data)
	fmt.Println("\n")

}

func init() {
	rootCmd.AddCommand(clientCmd)
	clientCmd.Flags().IntP("port", "p", 123, "Target Aort number")
	clientCmd.Flags().BoolP("verbose", "v", false, "verbose")
}
