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
        info, _ := cmd.Flags().GetInt("info")
        draft, _ := cmd.Flags().GetString("draft")
        timescale, _ := cmd.Flags().GetInt("timescale")
        flags, _ := cmd.Flags().GetInt("flags")

	// Listen server
	udpAddr := &net.UDPAddr{
		IP:   net.ParseIP(bind),
		Port: port,
	}
	conn, err := net.ListenUDP("udp", udpAddr)
	fmt.Println("Listen on:", bind, port)

	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	readBuffer := make([]byte, 1024)
	for {
		// Receive
		readLength, addr, err := conn.ReadFromUDP(readBuffer)
		if err != nil {
			log.Fatalln(err)
		}
		if (readLength < 48 || readLength % 4 != 0){
			continue
		}

		go func() {
			// Handling Receive Data
			receiveTimestamp := ntpv5.GetTimestampNow(0)

			receivedNtpv5data, err := ntpv5.Decode(readBuffer[:readLength])

			// failure decode
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println("Receive NTPv5 Data(" + addr.String() + "): \n", receivedNtpv5data)

			if (!verifyNtpv5Data(receivedNtpv5data)){
				return
			}

			// Struct Response
		        ntpv5data := ntpv5.NewServerNtpv5Data()
			ntpv5data.ClientCookie = receivedNtpv5data.ClientCookie
			ntpv5data.ReceiveTimestamp = receiveTimestamp
			ntpv5data.Timescale = uint8(timescale)
			ntpv5data.Flags = uint16(flags)

			if (receivedNtpv5data.ReferenceIDsRequestEx.Length > 0){
				ntpv5data.ReferenceIDsResponseEx = ntpv5.ReferenceIDsResponse{
					Length: receivedNtpv5data.ReferenceIDsRequestEx.Length,
				}
			}
			if (receivedNtpv5data.ServerInformationEx.Length > 0){
				ntpv5data.ServerInformationEx = ntpv5.ServerInformation{
					Length: 8,
					SupportedNtpVersions: uint16(info),
				}
			}
			if (receivedNtpv5data.ReferenceTimestampEx.Length > 0){
				ntpv5data.ReferenceTimestampEx = ntpv5.ReferenceTimestamp{
					Length: 12,
					ReferenceTimestamp: ntpv5.GetTimestampNow(0),
				}
			}
			if (len(receivedNtpv5data.SecondaryReceiveTimestampExs) > 0){
				//TODO:impl
			}
			if (receivedNtpv5data.DraftIdentificationEx.Length > 0){
				l := receivedNtpv5data.DraftIdentificationEx.Length - 4
				if(l > uint16(len(draft)) ){
					l = uint16(len(draft))
				}


				ntpv5data.DraftIdentificationEx = ntpv5.DraftIdentification{
					Length: uint16(l) + 4,
					Draft: draft[:l],
				}
			}

			// Padding rest
			restLength := readLength - len(ntpv5.Encode(ntpv5data))
			if (restLength > 0){
				ntpv5data.PaddingEx = ntpv5.Padding{
					Length: uint16(restLength),
				}
			} else {
				//Todo: ?
				fmt.Println("Todo?")
			}


			transmitTimestamp := ntpv5.GetTimestampNow(0)
			ntpv5data.TransmitTimestamp = transmitTimestamp

			fmt.Println("Send NTPv5 Data(" + addr.String() + "): \n", ntpv5data)

			buffer := ntpv5.Encode(ntpv5data)

			// Send
			_, _ = conn.WriteTo(buffer, addr)
			fmt.Println()
		}()
	}
}

func verifyNtpv5Data (d ntpv5.Ntpv5Data) bool {
	if( d.VN != 5 || d.Mode != 3){
		return false
	}
	return true
}

func init() {
	rootCmd.AddCommand(serverCmd)
        serverCmd.Flags().IntP("port", "p", 10123, "Target Port number")
        serverCmd.Flags().StringP("bind", "b", "0.0.0.0", "Bind Adress")
	serverCmd.Flags().IntP("timescale", "t", 0, "Timescale type")
	serverCmd.Flags().IntP("flags", "f", 0, "Flags")
	serverCmd.Flags().StringP("draft", "a", "draft-ietf-ntp-ntpv5-00", "Draft Identification")
	serverCmd.Flags().IntP("info", "i", 16, "Server Information")
        serverCmd.Flags().BoolP("verbose", "v", false, "verbose")
}
