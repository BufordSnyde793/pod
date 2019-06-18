package generic

import (
	"github.com/centrifuge/go-centrifuge/anchors"
	"github.com/centrifuge/go-centrifuge/documents"
)

// UpdateValidator returns a validator group that should be run before updating the generic document
func UpdateValidator(repo anchors.AnchorRepository) documents.ValidatorGroup {
	return documents.ValidatorGroup{
		documents.UpdateVersionValidator(repo),
	}
}
