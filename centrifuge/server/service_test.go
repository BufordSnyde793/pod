package server

import (
	"testing"
	"context"
	"bytes"
	"fmt"
	"github.com/CentrifugeInc/go-centrifuge/centrifuge/invoice"
)

func TestCoreDocumentService(t *testing.T) {
	identifier := []byte("identifier")
	identifierIncorrect := []byte("incorrectIdentifier")
	s := newInvoiceDocumentService()
	doc := invoice.InvoiceDocument{
		DocumentIdentifier: identifier,
	}

	sentDoc, err := s.SendInvoiceDocument(context.Background(), &invoice.SendInvoiceEnvelope{[][]byte{}, &doc})
	if err != nil {
		t.Fatal("Error in RPC Call", err)
	}
	if !bytes.Equal(sentDoc.DocumentIdentifier, identifier) {
		t.Fatal("DocumentIdentifier doesn't match")
	}

	receivedDoc, err := s.GetInvoiceDocument(context.Background(),
		&invoice.GetInvoiceDocumentEnvelope{DocumentIdentifier: identifier})
	if err != nil {
		t.Fatal("Error in RPC Call", err)
	}
	if !bytes.Equal(receivedDoc.DocumentIdentifier, identifier) {
		t.Fatal("DocumentIdentifier doesn't match")
	}

	docIncorrect, err := s.GetInvoiceDocument(context.Background(),
		&invoice.GetInvoiceDocumentEnvelope{DocumentIdentifier: identifierIncorrect})
	fmt.Println(docIncorrect)
	if err == nil {
		t.Fatal("RPC call should have raised exception")
	}
}