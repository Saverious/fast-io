/*
	Golang program to show synchronized concurrent I/O operations
*/

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	src_file, err := os.Open("./bigSrcFile.txt")
	done_reading := make(chan int) // channel for synchronization

	if err != nil {
		fmt.Println("error opening file for reading")
		panic(err)
	}

	dest_file, err := os.Create("./bigDestFile.txt")

	if err != nil {
		fmt.Println("error opening file for writing")
		panic(err)
	}

	reader := io.Reader(src_file)
	buffer := make([]byte, 4096)

	go readFromSource(reader, buffer, dest_file, done_reading)
	<-done_reading // receive signal from readFromSource()
	defer src_file.Close()
	defer dest_file.Close()
	fmt.Println("finished reading and writing data")

}

func readFromSource(reader io.Reader, buffer []byte, dest_file *os.File, done_reading chan int) {
	for {
		n, err := reader.Read(buffer)
		done_writing := make(chan int)

		if err == io.EOF {
			done_reading <- 1 // send signal to main()
			fmt.Println("reached end of file")
			break
		}

		if err != nil {
			fmt.Println("error reading file")
			panic(err)
		}

		fmt.Printf("read %v bytes\n", n)
		go writeToDestination(dest_file, buffer, done_writing)
		<-done_writing // receive signal from writeToDestination()
	}
}

func writeToDestination(writer io.Writer, buffer []byte, done_writing chan int) {
	n, err := writer.Write(buffer)

	if err != nil {
		fmt.Println("error writing to file")
		panic(err)
	}

	fmt.Printf("wrote %v bytes\n", n)
	done_writing <- 1 // send signal to readFromSource()
}
