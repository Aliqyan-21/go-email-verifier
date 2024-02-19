package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Domain, hasMX, hasSPF, spfRecords, hasDMARC, demarcRecords\n")

	for scanner.Scan() {
		checkDomain(scanner.Text())
	}

	checkError(scanner.Err())
}

func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecords, demarcRecords string

	mxRecords, err := net.LookupMX(domain)
	checkError(err)

	if len(mxRecords) > 0 {
		hasMX = true
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
