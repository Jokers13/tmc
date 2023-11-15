package commands

import (
	"log/slog"

	"github.com/web-of-things-open-source/tm-catalog-cli/internal/model"
	"github.com/web-of-things-open-source/tm-catalog-cli/internal/validation"
)

// ValidateThingModel validates the presence of the mandatory fields in the TM to be imported.
// Returns parsed *model.ThingModel, where the author name, manufacturer name, and mpn have been sanitized for use in filenames
func ValidateThingModel(raw []byte) (*model.ThingModel, error) {
	log := slog.Default()

	tm, err := validation.ParseRequiredMetadata(raw)
	if err != nil {
		return nil, err
	}
	log.Info("required Thing Model metadata is present")

	err = validation.ValidateAsTM(raw)
	if err != nil {
		return tm, err
	}
	log.Info("passed validation against JSON schema for Thing Models")

	validated, err := validation.ValidateAsModbus(raw)
	if validated {
		if err != nil {
			log.Info("failed [optional] validation against JSON schema for Modbus protocol binding", "error", err)
		} else {
			log.Info("passed validation against JSON schema for Modbus protocol binding")
		}
	}

	return tm, nil
}
