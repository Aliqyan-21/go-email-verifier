package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Domain, hasMX, hasSPF, spfRecords, hasDMARC, demarcRecords\n")

	for scanner.Scan() {
		checkDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error: %v\n", err)
	}
}

func checkDomain(domain string) {

}
