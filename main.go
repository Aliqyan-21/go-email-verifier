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
	fmt.Println("Type 0 if want to exit program")

	for scanner.Scan() {
		if scanner.Text() == "0" {
			os.Exit(0)
		}
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
	fmt.Printf("dmarcRecords: %v\n\n", dmarcRecords)
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
