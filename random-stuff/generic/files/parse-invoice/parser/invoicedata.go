package parser

import (
	"io"
	"time"
)

const (
	// FileType is used by text formats
	FileType = "INVOICES"
	// MagicNumber is used by binary formats
	MagicNumber = 0x1250
	// FileVersion used by all formats
	FileVersion = 100
	// DateFormat date format used
	DateFormat = "2006-01-02"
	// NanosecondsToSeconds nano seconds to seconds conversion
	NanosecondsToSeconds = 1e9
)

// Invoice struct represents the fields from the invoice file
type Invoice struct {
	ID         int
	CustomerID int
	Raised     time.Time
	Due        time.Time
	Paid       bool
	Note       string
	Items      []*Item
}

// Item represents the child fields present in invoice
type Item struct {
	ID       string
	Price    float64
	Quantity int
	Note     string
}

// InvoiceMarshaler interface uses writer for marshalling the
// data into a generic format
type InvoiceMarshaler interface {
	MarshalInvoices(writer io.Writer, invoices []*Invoice) error
}

// InvoiceUnmarshaler interface uses reader for unmarshaling the
// data into a generic format
type InvoiceUnmarshaler interface {
	UnmarshalInvoices(reader io.Reader) ([]*Invoice, error)
}
