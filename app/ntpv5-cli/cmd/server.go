/*
Copyright © 2023 flano_yuki

*/
package cmd

import (
	"fmt"
	"net"
	"log"
	"os"

	"github.com/spf13/cobra"
        "github.com/flano-yuki/ntp5-go-exp/internal/ntpv5"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "ntpv5 server",
	Run: execServer,
}

func execServer(cmd *cobra.Command, args []string){
        // handle, flag and args
        port, _ := cmd.Flags().GetInt("port")
        bind, _ := cmd.Flags().GetString("bind")

	fmt.Println("port, bind:", port, bind)

	// Listen server
	udpAddr := &net.UDPAddr{
		IP:   net.ParseIP(bind),
		Port: port,
	}
	conn, err := net.ListenUDP("udp", udpAddr)

	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	readBuffer := make([]byte, 256)
	for {
		readLength, addr, err := conn.ReadFromUDP(readBuffer)
		if err != nil {
			log.Fatalln(err)
		}

		go func() {
			receiveTimestamp := ntpv5.GetTimestampNow()

			receivedNtpv5data := ntpv5.Decode(readBuffer[:readLength])
			fmt.Println("Receive NTPv5 Data: \n", receivedNtpv5data)

			verifyNtpv5Data(receivedNtpv5data)

		        ntpv5data := ntpv5.NewServerNtpv5Data()
			ntpv5data.ClientCookie = receivedNtpv5data.ClientCookie
			ntpv5data.ReceiveTimestamp = receiveTimestamp

			transmitTimestamp := ntpv5.GetTimestampNow()
			ntpv5data.TransmitTimestamp = transmitTimestamp
			fmt.Println("Send NTPv5 Data: \n", ntpv5data)
			buffer := ntpv5.Encode(ntpv5data)

			_, _ = conn.WriteTo(buffer, addr)
			fmt.Println()
		}()
	}
}

func verifyNtpv5Data (d ntpv5.Ntpv5Data) bool {
	return true
}

func init() {
	rootCmd.AddCommand(serverCmd)
        serverCmd.Flags().IntP("port", "p", 10123, "Target Aort number")
        serverCmd.Flags().StringP("bind", "b", "0.0.0.0", "Bind Adress")
        serverCmd.Flags().BoolP("verbose", "v", false, "verbose")
}
