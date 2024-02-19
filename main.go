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
	var spfRecords, dmarcRecords string

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

	tmpdemarcRecords, err := net.LookupTXT("_dmarc." + domain)
	checkError(err)

	for _, record := range tmpdemarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecords = record
			break
		}
	}

	fmt.Printf("Domain: %v\n", domain)
	fmt.Printf("Hasmx: %v\n", hasMX)
	fmt.Printf("hasSPF: %v\n", hasSPF)
	fmt.Printf("spfRecord: %v\n", spfRecords)
	fmt.Printf("hasDMARC: %v\n", hasDMARC)
	fmt.Printf("dmarcRecords: %v\n", dmarcRecords)
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
