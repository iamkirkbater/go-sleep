package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

const (
	USAGE = `
usage: sleep time-duration

time-duration defaults to seconds if no duration length is passed.

examples: 
  gosleep 1 - sleeps for 1 second
  gosleep 3m - sleeps for 3 minutes
  gosleep 1h3m45s - sleeps for 1 hour, 3 minutes, 45 seconds
`
)

var (
	ERR_TOO_FEW_ARGS  = fmt.Errorf("Too Few Arguments Provided")
	ERR_TOO_MANY_ARGS = fmt.Errorf("Too Many Arguments Provided")
)

func echo_usage() {
	os.Stderr.WriteString(USAGE)
}

func parseArgs(args []string) (time.Duration, error) {
	allOtherArgs := args[1:]

	if len(allOtherArgs) == 0 {
		return (0 * time.Second), ERR_TOO_FEW_ARGS
	}
	if len(allOtherArgs) > 1 {
		return (0 * time.Second), ERR_TOO_MANY_ARGS
	}

	arg := allOtherArgs[0]

	i, err := strconv.Atoi(arg)
	if err == nil {
		return (time.Duration(i) * time.Second), nil
	}

	return time.ParseDuration(arg)
}

func main() {
	sleepTime, err := parseArgs(os.Args)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		echo_usage()
		os.Exit(2)
	}

	time.Sleep(sleepTime)
}
