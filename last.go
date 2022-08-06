package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

var (

	//Flags for -n and -o (offset)
	off = flag.Int("o", 0, "Offset")
	n   = flag.Int("n", 10, "Number of lines")
	f   = flag.String("f", "", "File to read")
)

func main() {

	flag.Parse()
	if *f == "" {
		fmt.Println("No file specified")
		os.Exit(1)
	}
	tailOffset(*f, *n, *off)

}

//Tail n lines from the end of the file, with an offset of bytes
func tailOffset(filepath string, n int, k int) {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	//Seek to the end of the file
	file.Seek(seekBack(file, (k+n)), 2)
	r := bufio.NewReader(file)

	for i := 0; i < n; i++ {
		//Read a line
		line, _ := r.ReadBytes('\n')
		//Print the line
		fmt.Printf(string(line))
	}
}

//Seek n lines back from the end of the file without loading the entire file into memory
//It returns an int64 offset that can be used with the os.Seek function to seek back n lines from the end of the file by reading the file backwards
//and finding each newline character until n lines have been read
func seekBack(file *os.File, n int) int64 {
	//Store the bytes read
	loc := int64(-1)
	stat, _ := file.Stat()
	size := stat.Size()
	//Loop n times reading backwards through the file, breaking when we find a newline
	for i := 0; i < n; i++ {
		for {
			b := make([]byte, 1)
			//Move the cursor back one byte, and seek to it
			loc -= 1
			file.Seek(loc, io.SeekEnd)
			file.Read(b)
			if b[0] == 13 || b[0] == 10 {
				break
			}

			//if we're at the beginning of the file, break
			if loc == -size {
				return loc
			}
		}
	}
	//Return the offset +1 to account for the newline we read
	return loc + 1

}
