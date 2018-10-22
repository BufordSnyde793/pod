package purchaseorder

import (
	"bytes"
	"context"
	"fmt"

	"github.com/centrifuge/go-centrifuge/centrifuge/centerrors"
	"github.com/centrifuge/go-centrifuge/centrifuge/code"

	"github.com/centrifuge/centrifuge-protobufs/gen/go/coredocument"
	"github.com/centrifuge/centrifuge-protobufs/gen/go/p2p"
	"github.com/centrifuge/go-centrifuge/centrifuge/coredocument/processor"
	"github.com/centrifuge/go-centrifuge/centrifuge/documents"
	"github.com/centrifuge/go-centrifuge/centrifuge/notification"
	clientpopb "github.com/centrifuge/go-centrifuge/centrifuge/protobufs/gen/go/purchaseorder"
)

// Service defines specific functions for purchase order
type Service interface {
	documents.Service

	// DeriverFromPayload derives purchase order from clientPayload
	DeriveFromCreatePayload(payload *clientpopb.PurchaseOrderCreatePayload) (documents.Model, error)

	// DeriveFromUpdatePayload derives purchase order from update payload
	DeriveFromUpdatePayload(payload *clientpopb.PurchaseOrderUpdatePayload) (documents.Model, error)

	// Create validates and persists purchase order and returns a Updated model
	Create(ctx context.Context, po documents.Model) (documents.Model, error)

	// Update validates and updates the purchase order and return the updated model
	Update(ctx context.Context, po documents.Model) (documents.Model, error)

	// DerivePurchaseOrderData returns the purchase order data as client data
	DerivePurchaseOrderData(po documents.Model) (*clientpopb.PurchaseOrderData, error)

	// DerivePurchaseOrderResponse returns the purchase order in our standard client format
	DerivePurchaseOrderResponse(po documents.Model) (*clientpopb.PurchaseOrderResponse, error)
}

// service implements Service and handles all purchase order related persistence and validations
// service always returns errors of type `centerrors` with proper error code
type service struct {
	repo             documents.Repository
	coreDocProcessor coredocumentprocessor.Processor
	notifier         notification.Sender
}

// DefaultService returns the default implementation of the service
func DefaultService(repo documents.Repository, processor coredocumentprocessor.Processor) Service {
	return service{repo: repo, coreDocProcessor: processor, notifier: &notification.WebhookSender{}}
}

// DeriveFromCoreDocument takes a core document and returns a purchase order
func (s service) DeriveFromCoreDocument(cd *coredocumentpb.CoreDocument) (documents.Model, error) {
	return nil, fmt.Errorf("implement me")
}

// Create validates, persists, and anchors a purchase order
func (s service) Create(ctx context.Context, po documents.Model) (documents.Model, error) {
	return nil, fmt.Errorf("implement me")
}

// Update validates, persists, and anchors a new version of purchase order
func (s service) Update(ctx context.Context, po documents.Model) (documents.Model, error) {
	return nil, fmt.Errorf("implement me")
}

// DeriveFromCreatePayload derives purchase order from create payload
func (s service) DeriveFromCreatePayload(payload *clientpopb.PurchaseOrderCreatePayload) (documents.Model, error) {
	return nil, fmt.Errorf("implement me")
}

// DeriveFromUpdatePayload derives purchase order from update payload
func (s service) DeriveFromUpdatePayload(payload *clientpopb.PurchaseOrderUpdatePayload) (documents.Model, error) {
	return nil, fmt.Errorf("implement me")
}

// DerivePurchaseOrderData returns po data from the model
func (s service) DerivePurchaseOrderData(po documents.Model) (*clientpopb.PurchaseOrderData, error) {
	return nil, fmt.Errorf("implement me")
}

// DerivePurchaseOrderResponse returns po response from the model
func (s service) DerivePurchaseOrderResponse(po documents.Model) (*clientpopb.PurchaseOrderResponse, error) {
	return nil, fmt.Errorf("implement me")
}

func (s service) getPurchaseOrderVersion(documentID, version []byte) (model *PurchaseOrderModel, err error) {
	var doc documents.Model = new(PurchaseOrderModel)
	err = s.repo.LoadByID(version, doc)
	if err != nil {
		return nil, err
	}
	model, ok := doc.(*PurchaseOrderModel)
	if !ok {
		return nil, err
	}

	if !bytes.Equal(model.CoreDocument.DocumentIdentifier, documentID) {
		return nil, centerrors.New(code.DocumentInvalid, "version is not valid for this identifier")
	}
	return model, nil
}

// GetLastVersion returns the latest version of the document
func (s service) GetCurrentVersion(documentID []byte) (documents.Model, error) {
	model, err := s.getPurchaseOrderVersion(documentID, documentID)
	if err != nil {
		return nil, centerrors.Wrap(err, "document not found")
	}
	nextVersion := model.CoreDocument.NextVersion
	for nextVersion != nil {
		temp, err := s.getPurchaseOrderVersion(documentID, nextVersion)
		if err != nil {
			return model, nil
		} else {
			model = temp
			nextVersion = model.CoreDocument.NextVersion
		}
	}
	return model, nil
}

// GetVersion returns the specific version of the document
func (s service) GetVersion(documentID []byte, version []byte) (documents.Model, error) {
	inv, err := s.getPurchaseOrderVersion(documentID, version)
	if err != nil {
		return nil, centerrors.Wrap(err, "document not found for the given version")
	}
	return inv, nil
}

// CreateProofs generates proofs for given document
func (s service) CreateProofs(documentID []byte, fields []string) (*documents.DocumentProof, error) {
	return nil, fmt.Errorf("implement me")
}

// CreateProofsForVersion generates proofs for specific version of the document
func (s service) CreateProofsForVersion(documentID, version []byte, fields []string) (*documents.DocumentProof, error) {
	return nil, fmt.Errorf("implement me")
}

// RequestDocumentSignature validates the document and returns the signature
func (s service) RequestDocumentSignature(model documents.Model) (*coredocumentpb.Signature, error) {
	return nil, fmt.Errorf("implement me")
}

// ReceiveAnchoredDocument validates the anchored document and updates it on DB
func (s service) ReceiveAnchoredDocument(model documents.Model, headers *p2ppb.CentrifugeHeader) error {
	return fmt.Errorf("implement me")
}