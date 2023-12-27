package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type Invoice struct {
	ID         int
	CustomerID int
	Date       string
	Details    []InvoiceDetails
}

type InvoiceDetails struct {
	ID        int
	InvoiceID int
	ProductID int
	Quantity  int
}

func main() {
	connStr := "postgres://postgres:postgres@localhost/testdb?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// two queries
	start := time.Now()
	rows, err := db.Query("SELECT * FROM invoice")
	if err != nil {
		log.Fatal(err)
	}
	invoicesTwoQ := make([]Invoice, 0)
	for rows.Next() {
		var invoice Invoice
		err = rows.Scan(&invoice.ID, &invoice.CustomerID, &invoice.Date)
		if err != nil {
			log.Fatal(err)
		}

		details, err := db.Query("SELECT * FROM invoice_details WHERE invoice_id = $1", invoice.ID)
		if err != nil {
			log.Fatal(err)
		}
		invoiceDetails := make([]InvoiceDetails, 0)
		for details.Next() {
			var detail InvoiceDetails
			err = details.Scan(&detail.ID, &detail.InvoiceID, &detail.ProductID, &detail.Quantity)
			if err != nil {
				log.Fatal(err)
			}
			invoiceDetails = append(invoiceDetails, detail)
		}
		details.Close()

		invoice.Details = invoiceDetails
		invoicesTwoQ = append(invoicesTwoQ, invoice)
	}
	rows.Close()
	fmt.Printf("Time taken for two queries: %v.Got %d invoces.\n", time.Since(start), len(invoicesTwoQ))

	// join query
	start = time.Now()
	rows, err = db.Query("SELECT invoice.id, invoice.customer_id, invoice.date, invoice_details.id, invoice_details.invoice_id, invoice_details.product_id, invoice_details.quantity FROM invoice JOIN invoice_details ON invoice.id = invoice_details.invoice_id")
	if err != nil {
		log.Fatal(err)
	}
	invoicesJ := make(map[int]Invoice)
	for rows.Next() {
		var invoice Invoice
		var detail InvoiceDetails
		err = rows.Scan(&invoice.ID, &invoice.CustomerID, &invoice.Date, &detail.ID, &detail.InvoiceID, &detail.ProductID, &detail.Quantity)
		if err != nil {
			log.Fatal(err)
		}
		if existingInvoice, ok := invoicesJ[invoice.ID]; ok {
			existingInvoice.Details = append(existingInvoice.Details, detail)
			invoicesJ[invoice.ID] = existingInvoice
		} else {
			invoice.Details = append(invoice.Details, detail)
			invoicesJ[invoice.ID] = invoice
		}
	}
	rows.Close()
	fmt.Printf("Time taken for join query: %v. Got %d invoices\n", time.Since(start), len(invoicesJ))

	db.Close()
}
