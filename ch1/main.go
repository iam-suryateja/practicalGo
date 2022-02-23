package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
)

/*
The getName() accepts two arguments.
1. The first argument, r is a variable whose value satisfies the Reader interface defined in the io package.
   Ex: stdin as defined in the os package - usually the terminal session in which you are executing the program.
2.The second argument,w is a variable whose value satisfies the writer interface
   Ex: stdout
*/
func getName(r io.Reader, w io.Writer) (string, error) {
	msg := "your name pleasae? Press the Enter key when done. \n"
	fmt.Fprint(w, msg) // To write a prompt to the specified writer w

	scanner := bufio.NewScanner(r)
	scanner.Scan() //lets you scan the reader for any input data using the scan function
	if err := scanner.Err(); err != nil {
		return "", err
	}
	name := scanner.Text() // returns the read data as string
	if len(name) == 0 {
		return "", errors.New("you didn't enter your name")
	}
	return name, nil
}
