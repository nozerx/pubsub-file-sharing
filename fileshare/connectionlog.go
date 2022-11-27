package fileshare

import (
	"fmt"
	"os"
)

var filename_const string = "log/connectionlog"

func OpenConnectionStatusLog() *os.File {
	i := 0
	for {
		_, err := os.Stat(filename_const + ".txt")
		if err != nil {
			break
		} else {
			filename_const = fmt.Sprintf("%s_%d", filename_const[:len(filename_const)-2], i)
		}
		i++
	}
	file, err := os.Create(filename_const + ".txt")
	if err != nil {
		fmt.Println("Error while opening the file", filename_const)
	}
	fmt.Println("Using [", filename_const, "] for connection status log")
	return file
}
