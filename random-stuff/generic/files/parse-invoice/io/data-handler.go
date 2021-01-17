package io

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/fpdevil/goprog/random-stuff/generic/files/parse-invoice/parser"
)

// OpenInvoiceFile reads input input invoice file and returns handlers
// for reading and writing the data
func OpenInvoiceFile(filename string) (io.ReadCloser, func(), error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}

	closer := func() {
		file.Close()
	}

	var reader io.ReadCloser = file
	var decompressor *gzip.Reader
	if strings.HasSuffix(filename, ".gz") {
		if decompressor, err = gzip.NewReader(file); err != nil {
			return file, closer, err
		}
		closer = func() {
			decompressor.Close()
			file.Close()
		}
		reader = decompressor
	}
	return reader, closer, nil
}

// ReadInvoices function is useful for reading and parsing
// the invoices from an open file
func ReadInvoices(reader io.Reader, suffix string) ([]*parser.Invoice, error) {
	var unmarshaler parser.InvoiceUnmarshaler

	switch suffix {
	case "*.gob":
		unmarshaler = parser.GobMarshaler{}
	}

	if unmarshaler != nil {
		return unmarshaler.UnmarshalInvoices(reader)
	}
	return nil, fmt.Errorf("unrecognized input suffix: %s", suffix)
}
