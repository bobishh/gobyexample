package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	dateCmd := exec.Command("date", "+'%H:%M'")
	dateOut, err := dateCmd.Output()
	check(err)
	puts := fmt.Println

	puts("> date +'%H:%M'")
	puts(string(dateOut))
	// Next we’ll look at a slightly more involved case where we pipe data to the external process on its stdin and collect the results from its stdout.

	grepCmd := exec.Command("grep", "hello")
	// Here we explicitly grab input/output pipes, start the process, write some input to it, read the resulting output, and finally wait for the process to exit.

	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()
	// We ommited error checks in the above example, but you could use the usual if err != nil pattern for all of them. We also only collect the StdoutPipe results, but you could collect the StderrPipe in exactly the same way.

	puts("> grep hello")
	puts(string(grepBytes))

	// Note that when spawning commands we need to provide an explicitly delineated command and argument array, vs. being able to just pass in one command-line string. If you want to spawn a full command with a string, you can use bash’s -c option:

	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	puts("> ls -a -l -h")
	puts(string(lsOut))
}
