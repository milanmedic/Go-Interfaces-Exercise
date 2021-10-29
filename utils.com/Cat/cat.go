package cat

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// checkout flag package

func Print(reader bufio.Reader, lineNumFlag bool) {
	lineNum := 1
	for {
		line, _, err := reader.ReadLine()

		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error while reading the file!")
			panic(err)
		}
		if !lineNumFlag {
			fmt.Printf("%s\n", line)
		} else {
			fmt.Printf("%d	%s\n", lineNum, line)
			lineNum++
		}
	}
}

func Cat() {

	filename := os.Args[1]
	lineNumFlag := false
	if len(os.Args) > 2 && os.Args[2] == "-n" {
		lineNumFlag = true
	}

	file, err := os.Open(filename)

	if err != nil {
		fmt.Println("Error while reading the file!")
		panic(err)
	}

	defer file.Close() // when enclosing scope finishes, this line triggers

	reader := bufio.NewReader(file)

	Print(*reader, lineNumFlag)

}
