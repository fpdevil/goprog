/*
Parsing of the bank transaction files

We will be  ingesting a csv formatted transaction  file from some
bank. It includes budget categories for transactions in the file.
It is required to create  a command-line program that will accept
two flags: the location of  csv transaction file and the location
of a log file.

We will  check that the log  and bank file location(s)  are valid
before parsing of CSV file starts. The program will parse the CSV
file  and log  any errors  it encounters  to the  log. Upon  each
restart  of the  program, it  will also  delete the  previous log
file.
*/
package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// create budget category types for fuel, food, mortgage, repairs,
// insurance, utilities, and retirement.
type category string

const (
	fuel       category = "fuel"
	food       category = "food"
	mortgage   category = "mortgage"
	repairs    category = "repairs"
	insurance  category = "insurance"
	utilities  category = "utilities"
	retirement category = "retirement"
)

var (
	// ErrInvalidcategory handles any errors for budget categories
	ErrInvalidcategory = errors.New("unable to find listed budget category")
)

// bank transactions entities
type transaction struct {
	id       int
	payee    string
	spent    float64
	category category
}

func main() {
	fmt.Println("==== Bank-CSV-Transactions ====")
	fmt.Println()

	// check for the existence of bank data and log file initially
	bankFile := flag.String("b", "", "path to the bank transaction's csv")
	logFile := flag.String("l", "", "path to log file for error logging")
	flag.Parse()

	if *bankFile == "" {
		fmt.Println("csv transaction file required.")
		flag.PrintDefaults()
		return
	}

	if *logFile == "" {
		fmt.Println("log file path is required.")
		flag.PrintDefaults()
		return
	}

	// get file info of the named file
	_, err := os.Stat(*bankFile)
	if os.IsNotExist(err) {
		fmt.Printf("Bank transaction data %s unavailable", *bankFile)
		return
	}

	// if a previous log file exists remove the same
	_, err = os.Stat(*logFile)
	if !os.IsNotExist(err) {
		fmt.Println(">>> cleaning up the log file")
		fmt.Println()
		_ = os.Remove(*logFile)
	}

	// open the CSV file for parsing
	csvFile, err := os.Open(*bankFile)
	if err != nil {
		log.Fatal(err)
	}

	txns := parseCSV(csvFile, *logFile)
	for _, txn := range txns {
		fmt.Printf("%+v\n", txn)
	}

}

func parseCSV(bankTxns io.Reader, logFile string) []transaction {
	var (
		txns   []transaction
		header bool
		cline  int
	)

	rr := csv.NewReader(bufio.NewReader(bankTxns))

	for {
		txn := transaction{}
		record, err := rr.Read()
		cline++

		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error on line %d\n", cline)
			log.Fatal(err)
		}

		// header contains the top heading words which may be discarded
		if header {
			// check each record
			for i, val := range record {
				switch i {
				case 0:
					val = strings.TrimSpace(val)
					txn.id, err = strconv.Atoi(val)
					if err != nil {
						fmt.Printf("error on record %d\n", cline)
						log.Fatal(err)
					}
				case 1:
					val := strings.TrimSpace(val)
					txn.payee = val
				case 2:
					val := strings.TrimSpace(val)
					txn.spent, err = strconv.ParseFloat(val, 64)
					if err != nil {
						fmt.Printf("error on record %d\n", cline)
						log.Fatal(err)
					}
				case 3:
					txn.category, err = mapBudgetToCategory(val)
					if err != nil {
						s := strings.Join(record, ", ")
						_ = errorLog("error parsing csv catgeory column - ", err, s, logFile)
					}
				}
			}
			txns = append(txns, txn)
		}
		header = true
	}
	return txns
}

func errorLog(msg string, err error, data string, logFile string) error {
	msg += "\n" + err.Error() + "\nData: " + data + "\n\n"
	f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer f.Close()
	if _, err = f.WriteString(msg); err != nil {
		return err
	}
	return nil
}

func mapBudgetToCategory(value string) (category, error) {
	value = strings.TrimSpace(strings.ToLower(value))
	switch value {
	case "fuel", "gas":
		return fuel, nil
	case "food":
		return food, nil
	case "mortgage":
		return mortgage, nil
	case "repairs":
		return repairs, nil
	case "life insurance", "car insurance":
		return insurance, nil
	case "utilities":
		return utilities, nil
	default:
		return "", ErrInvalidcategory
	}
}
