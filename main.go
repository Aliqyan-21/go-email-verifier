package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
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

	txtRecords, err := net.LookupTXT(domain) // spf records
	checkError(err)

	for _, records := range txtRecords {
		if strings.HasPrefix(records, "v=spf1") {
			hasSPF = true
			spfRecords = records
			break
		}
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
