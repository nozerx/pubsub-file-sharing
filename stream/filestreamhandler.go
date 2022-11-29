package stream

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/libp2p/go-libp2p/core/network"
)

func sendToStream(str network.Stream) {
	fmt.Println("Sending file to ", str.Conn().RemotePeer())
	file, err := os.Open("go.sum")
	sendBytes := make([]byte, 100)
	if err != nil {
		fmt.Println("Error while opening the sending file")
	} else {
		bufstr := bufio.NewWriter(str)
		for {
			_, err = file.Read(sendBytes)
			if err == io.EOF {
				fmt.Println("Send the file completely")
				count := 0
				for {
					bufstr.Write([]byte("jfkda"))
					count++
					if count > 200 {
						break
					}
				}
				break
			}
			if err != nil {
				fmt.Println("Error while reading from the file")
			}
			_, err = bufstr.Write(sendBytes)
			if err != nil {
				fmt.Println("Error while sending to the stream")
			}
		}
		fmt.Println("Outside the loop")
		time.Sleep(60 * time.Second)
		fmt.Println("Closing the stream")
		str.Close()
	}

}

func ReceivedFromStream(str network.Stream, logfile *os.File) {
	file, err := os.Create("Recieved.txt")
	fmt.Fprintln(logfile, "Recieving file from stream to :", str.Conn().RemotePeer())
	readBytes := make([]byte, 100)
	if err != nil {
		fmt.Fprintln(logfile, "Error while creating the recieving file")
	} else {
		bufstr := bufio.NewReader(str)
		for {
			_, err := bufstr.Read(readBytes)
			if err == io.EOF {
				fmt.Fprintln(logfile, "End of file reached")
				break
			}
			if err != nil {
				fmt.Fprintln(logfile, "Error while reading from stream")
				break
			}
			fmt.Print("\nline\n" + string(readBytes))
			_, err = file.Write(readBytes)
			if err != nil {
				fmt.Fprintln(logfile, "Error while writing to the stream")
			} else {
				fmt.Println("Writing to the file")
			}

		}
		fmt.Fprintln(logfile, "Completed reading from the stream")
		fmt.Println("Completed reading from the stream")
	}
	file.Close()
	time.Sleep(1 * time.Minute)

}
