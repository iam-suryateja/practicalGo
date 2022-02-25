package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

type config struct { // The config structure is used for in-memory representation of data on which the application will rely on runtime behavior
	numTimes   int
	printUsage bool
}

var usageString = fmt.Sprintf(`Usage: %s <integer> [-h|--help]
A greeter application which prints the name you entered <integer> number of times.
`, os.Args[0])

func printUsage(w io.Writer) {
	fmt.Fprintf(w, usageString)
}

/*
The parseArgs() creates an object c of config type to store the data.
*/

func validateArgs(c config) error {
	if !(c.numTimes > 0) {
		fmt.Println(c.numTimes)
		return errors.New("Must specify a number greater than 0")
	}
	return nil
}
func parseArgs(args []string) (config, error) {
	var numTimes int
	var err error
	c := config{}

	if len(args) != 1 {
		/*Function first checks to see if the number of command line arguments is not equal to 1
		If so, it returns an empty config object and error
		*/
		return c, errors.New("Invalid number of arguments")
	}

	/* If only one argument is specified and it is -h or -help, the printUsage field is specified to true and th
	   the objectc and nil error are returned using the following snippet
	*/

	if args[0] == "-h" || args[0] == "--help" {
		c.printUsage = true
		c.numTimes = len(args)
		return c, nil

	}

	numTimes, err = strconv.Atoi(args[0]) //Atoi() function from strconv package is used to convert the arguments - a string to its integer equivalent
	if err != nil {
		return c, err
	}
	c.numTimes = numTimes
	return c, nil
}

/*
The getName() accepts two arguments.
1. The first argument, r is a variable whose value satisfies the Reader interface defined in the io package.
   Ex: stdin as defined in the os package - usually the terminal session in which you are executing the program.
2.The second argument,w is a variable whose value satisfies the writer interface
   Ex: stdout
*/
func getName(r io.Reader, w io.Writer) (string, error) {
	msg := "your name please? Press the Enter key when done. \n"
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

func greetUser(c config, name string, w io.Writer) {
	msg := fmt.Sprintf("Nice to meet you %s\n", name)
	for i := 0; i < c.numTimes; i++ {
		fmt.Fprintf(w, msg)
	}
}

func runCmd(r io.Reader, w io.Writer, c config) error {
	if c.printUsage {
		printUsage(w)
		return nil
	}
	name, err := getName(r, w)
	if err != nil {
		return err
	}
	greetUser(c, name, w)
	return nil
}

func main() {
	c, err := parseArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		printUsage(os.Stdout)
		os.Exit(1)
	}
	err = validateArgs(c)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		printUsage(os.Stdout)
		os.Exit(1)
	}
	err = runCmd(os.Stdin, os.Stdout, c)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
}
