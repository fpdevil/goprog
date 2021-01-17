package parser

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"time"
)

// JSONInvoice represents the underlying data packet of the
// invoice for parsing the json data
type JSONInvoice struct {
	ID         int
	CustomerID int
	Raised     string // seconds since unix epoch time.Time in invoice struct
	Due        string // seconds since unix epoch time.Time in invoice struct
	Paid       bool
	Note       string
	Items      []*Item
}

// JSONMarshaler represents an empty template for holding the json data
type JSONMarshaler struct{}

// MarshalJSON function aids in marshaling the json formatted invoice
func (invoice Invoice) MarshalJSON() ([]byte, error) {
	jsonInvoice := JSONInvoice{
		invoice.ID,
		invoice.CustomerID,
		invoice.Raised.Format(DateFormat),
		invoice.Due.Format(DateFormat),
		invoice.Paid,
		invoice.Note,
		invoice.Items,
	}
	return json.Marshal(jsonInvoice)
}

// UnmarshalJSON function deserializes the data supplied as a byte slice
func (invoice *Invoice) UnmarshalJSON(data []byte) (err error) {
	var jsonInvoice JSONInvoice
	if err = json.Unmarshal(data, &jsonInvoice); err != nil {
		return err
	}

	var raised, due time.Time
	if raised, err = time.Parse(DateFormat, jsonInvoice.Raised); err != nil {
		return err
	}
	if due, err = time.Parse(DateFormat, jsonInvoice.Due); err != nil {
		return err
	}

	*invoice = Invoice{
		jsonInvoice.ID,
		jsonInvoice.CustomerID,
		raised,
		due,
		jsonInvoice.Paid,
		jsonInvoice.Note,
		jsonInvoice.Items,
	}
	return nil
}

// MarshalInvoices function satisfies the InvoiceMarshaler interface and
// implements its declared method
func (JSONMarshaler) MarshalInvoices(writer io.Writer, invoices []*Invoice) error {
	encoder := json.NewEncoder(writer)
	if err := encoder.Encode(FileType); err != nil {
		return err
	}
	if err := encoder.Encode(FileVersion); err != nil {
		return err
	}
	return encoder.Encode(invoices)
}

// UnmarshalInvoices satisfies the interface InvoiceUnmarshaler and
// implements its declated method
func (JSONMarshaler) UnmarshalInvoices(reader io.Reader) ([]*Invoice, error) {
	decoder := json.NewDecoder(reader)
	var kind string
	if err := decoder.Decode(&kind); err != nil {
		return nil, err
	}
	if kind != FileType {
		return nil, errors.New("cannot read non-invoices json data files")
	}
	var version int
	if err := decoder.Decode(&version); err != nil {
		return nil, err
	}
	if version > FileVersion {
		return nil, fmt.Errorf("version %d is newer to read", version)
	}
	var invoices []*Invoice
	err := decoder.Decode(&invoices)
	return invoices, err
}
