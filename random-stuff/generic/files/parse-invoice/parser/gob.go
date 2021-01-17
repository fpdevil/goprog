package parser

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"time"
)

// GobInvoice is a custom type for satisfying the gob.Encoder and
// gob.Decoder interfaces
type GobInvoice struct {
	ID         int
	CustomerID int
	Raised     int64 // seconds since unix epoch
	Due        int64 // seconds since unix epoch
	Paid       bool
	Note       string
	Items      []*Item
}

// GobMarshaler struct is for handling the marshalled data
type GobMarshaler struct{}

// GobEncode function takes an invoice as receiver type and returns the
// byte slice representation of the invoice and an error
func (invoice *Invoice) GobEncode() ([]byte, error) {
	gobInvoice := GobInvoice{
		invoice.ID,
		invoice.CustomerID,
		invoice.Raised.Unix(),
		invoice.Due.Unix(),
		invoice.Paid,
		invoice.Note,
		invoice.Items,
	}
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(gobInvoice)
	return buffer.Bytes(), err
}

// GobDecode function takes an invoice as receiver type and decodes the
// the input data represented as a byte slice
func (invoice *Invoice) GobDecode(data []byte) error {
	var gobInvoice GobInvoice
	buffer := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buffer)
	if err := decoder.Decode(&gobInvoice); err != nil {
		return err
	}
	raised := time.Unix(gobInvoice.Raised, 0)
	due := time.Unix(gobInvoice.Due, 0)
	*invoice = Invoice{
		gobInvoice.ID,
		gobInvoice.CustomerID,
		raised,
		due,
		gobInvoice.Paid,
		gobInvoice.Note,
		gobInvoice.Items,
	}
	return nil
}

// MarshalInvoices satisfies the interface InvoiceMarshaler and implements
// the defined method for gob marshalling
func (GobMarshaler) MarshalInvoices(writer io.Writer, invoices []*Invoice) error {
	encoder := gob.NewEncoder(writer)
	if err := encoder.Encode(MagicNumber); err != nil {
		return err
	}

	if err := encoder.Encode(FileVersion); err != nil {
		return err
	}
	return encoder.Encode(invoices)
}

// UnmarshalInvoices satisfies the InvoiceUnmarshaler interface and implements
// the method defined in it for unmarshaling gob data
func (GobMarshaler) UnmarshalInvoices(reader io.Reader) ([]*Invoice, error) {
	decoder := gob.NewDecoder(reader)
	var magic int
	if err := decoder.Decode(&magic); err != nil {
		return nil, err
	}
	var version int
	if err := decoder.Decode(&version); err != nil {
		return nil, err
	}

	if version > FileVersion {
		return nil, fmt.Errorf("version: %d is newer to read", version)
	}
	var invoices []*Invoice
	err := decoder.Decode(&invoices)
	return invoices, err
}
