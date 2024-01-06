package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func getNumberOfNetworkInterfaces() float64 {
	cmd := "ls /sys/class/net | wc -l"
	commandOutput, err := exec.Command("bash", "-c", cmd).Output()

	if err != nil {
		fmt.Printf("%s", err)
	}

	converted, err := strconv.ParseFloat(strings.TrimSuffix(string(commandOutput), "\n"), 64)
	if err != nil {
		fmt.Printf("%s", err)
	}

	return converted
}

func main() {
	fmt.Printf("\n\n%f\n\n", getNumberOfNetworkInterfaces())
}