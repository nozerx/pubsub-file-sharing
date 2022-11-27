package fileshare

import (
	"fmt"
	"os"
)

var filename string = "log/peerconnlog"

func OpenPeerConnectionLog() *os.File {
	i := 0
	for {
		_, err := os.Stat(filename + ".txt")
		if err != nil {
			break
		} else {
			filename = fmt.Sprintf("%s_%d", filename[:len(filename)-2], i)
		}
		i++
	}
	file, err := os.Create(filename + ".txt")
	if err != nil {
		fmt.Println("Error while opening the file", filename)
	}
	fmt.Println("Using [", filename, "] for peer connection log")
	return file
}
