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
	invoices := make([]Invoice, 0)
	for rows.Next() {
		var invoice Invoice
		err = rows.Scan(&invoice.ID, &invoice.CustomerID, &invoice.Date)
		if err != nil {
			log.Fatal(err)
		}
		invoices = append(invoices, invoice)

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
	}
	rows.Close()
	fmt.Printf("Time taken for two queries: %v.Got %d invoces.\n", time.Since(start), len(invoices))

	// join query
	start = time.Now()
	rows, err = db.Query("SELECT invoice.id, invoice.customer_id, invoice.date FROM invoice JOIN invoice_details ON invoice.id = invoice_details.invoice_id GROUP BY invoice.id")
	if err != nil {
		log.Fatal(err)
	}
	invoices = make([]Invoice, 0)
	for rows.Next() {
		var invoice Invoice
		err = rows.Scan(&invoice.ID, &invoice.CustomerID, &invoice.Date)
		if err != nil {
			log.Fatal(err)
		}
		invoices = append(invoices, invoice)
	}
	rows.Close()
	fmt.Printf("Time taken for join query: %v. Got %d invoices\n", time.Since(start), len(invoices))

	db.Close()
}
