/*
This package is required to provide functionality to process a csv file and return a sorted (by email domain) data
structure of your choice containing the email domains along with the number of customers for each domain. The customer_data.csv
file provides an example csv file to work with. Any errors should be logged (or handled) or returned to the consumer of
this package. Performance matters, the sample file may only contain 1K lines but the package may be expected to be used on
files with 10 million lines or run on a small machine.

Write this package as you normally would for any production grade code that would be deployed to a live system.

Please stick to using the standard library.
*/

package emaildomainstats

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"sort"
	"strconv"
	"strings"
)

// domainStats holds the running total of domains and a unique list of customers
// per domain
type domainStats map[string]map[string]bool

// ProcessData holds the main functionality of the emaildomainstats package
// it takes an io.Reader providing CSV data and an io.Writer to writeh the output CSV data too
func ProcessData(dataReader io.Reader, dataWriter io.Writer) error {
	csvReader := csv.NewReader(dataReader)
	domainRegister := make(domainStats)

	// ignore the header row
	if _, err := csvReader.Read(); err != nil {
		return fmt.Errorf("error reading header row of file: %s", err)
	}

	// Loop through the file until the end of the file is reached
	for i := 0; ; i++ {
		row, err := csvReader.Read()
		if err == io.EOF {
			log.Printf("Finished processing file: %d lines processed", i)
			break
		}
		if err != nil {
			log.Printf("error reading csv line %d: %s\n", i, err)
			continue
		}

		customer, domain, err := splitEmail(row[2])
		if err != nil {
			log.Printf("error processing line %d: %s\n", i, err)
			continue
		}

		// initalise the domain if not seen before
		if domainRegister[domain] == nil {
			domainRegister[domain] = make(map[string]bool)
		}
		domainRegister[domain][customer] = true
	}

	return processDomainStatsOutput(domainRegister, dataWriter)
}

// splitEmail takes an email in the form of a string and returns the domain
// and customer as separate strings
func splitEmail(email string) (string, string, error) {
	splitEmail := strings.Split(email, "@")
	if len(splitEmail) != 2 {
		return "", "", fmt.Errorf("email is malformed")
	}

	if splitEmail[0] == "" {
		return "", "", fmt.Errorf("customer is missing from email")
	}

	if splitEmail[1] == "" {
		return "", "", fmt.Errorf("domain is missing from email")
	}

	return splitEmail[0], splitEmail[1], nil
}

// processDomainStatsOutput takes a domainStats dataset, formats it, sort it and
// writes the domain and customer count to the provided io.Writer
func processDomainStatsOutput(stats domainStats, outputWriter io.Writer) error {
	writer := csv.NewWriter(outputWriter)
	outputDomainStats := [][]string{}
	for domain, customerList := range stats {
		outputDomainStats = append(outputDomainStats, []string{domain, strconv.Itoa(len(customerList))})
	}
	// Sort the output alphabetically, grouping capital and lower case letters together
	sort.Slice(outputDomainStats, func(i, j int) bool {
		if strings.EqualFold(outputDomainStats[i][0], outputDomainStats[j][0]) {
			return outputDomainStats[i][0] < outputDomainStats[j][0]
		}
		return strings.ToLower(outputDomainStats[i][0]) < strings.ToLower(outputDomainStats[j][0])
	})
	if err := writer.WriteAll(outputDomainStats); err != nil {
		return fmt.Errorf("error while writing CSV to output: %s", err)
	}
	return nil
}
