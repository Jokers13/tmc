// Package server provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package server

const (
	BearerAuthScopes = "BearerAuth.Scopes"
)

// Defines values for GetCompletionsParamsKind.
const (
	FetchNames GetCompletionsParamsKind = "fetchNames"
	Names      GetCompletionsParamsKind = "names"
)

// AttachmentLinks defines model for AttachmentLinks.
type AttachmentLinks struct {
	Content string `json:"content"`
}

// AttachmentsList defines model for AttachmentsList.
type AttachmentsList = []AttachmentsListEntry

// AttachmentsListEntry defines model for AttachmentsListEntry.
type AttachmentsListEntry struct {
	Links     *AttachmentLinks `json:"links,omitempty"`
	MediaType string           `json:"mediaType"`
	Name      string           `json:"name"`
}

// AuthorsResponse defines model for AuthorsResponse.
type AuthorsResponse struct {
	Data []string `json:"data"`
}

// ErrorResponse defines model for ErrorResponse.
type ErrorResponse struct {
	Code     *string `json:"code,omitempty"`
	Detail   *string `json:"detail,omitempty"`
	Instance *string `json:"instance,omitempty"`
	Status   int     `json:"status"`
	Title    string  `json:"title"`
	Type     *string `json:"type,omitempty"`
}

// ImportThingModelResponse defines model for ImportThingModelResponse.
type ImportThingModelResponse struct {
	Data ImportThingModelResult `json:"data"`
}

// ImportThingModelResult defines model for ImportThingModelResult.
type ImportThingModelResult struct {
	Code    *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	TmID    string  `json:"tmID"`
}

// InfoResponse defines model for InfoResponse.
type InfoResponse struct {
	Details *[]string   `json:"details,omitempty"`
	Name    string      `json:"name"`
	Version InfoVersion `json:"version"`
}

// InfoVersion defines model for InfoVersion.
type InfoVersion struct {
	Implementation string  `json:"implementation"`
	Specification  *string `json:"specification,omitempty"`
}

// InventoryEntry defines model for InventoryEntry.
type InventoryEntry struct {
	Attachments *AttachmentsList     `json:"attachments,omitempty"`
	Links       *InventoryEntryLinks `json:"links,omitempty"`

	// Repo The name of the source repository where the inventory entry or version resides.
	// May be left empty when there is only a single repository served by the backend and thus there is not need for
	// disambiguation. See also '/repos'
	Repo               *SourceRepository       `json:"repo,omitempty"`
	SchemaAuthor       SchemaAuthor            `json:"schema:author"`
	SchemaManufacturer SchemaManufacturer      `json:"schema:manufacturer"`
	SchemaMpn          string                  `json:"schema:mpn"`
	TmName             string                  `json:"tmName"`
	Versions           []InventoryEntryVersion `json:"versions"`
}

// InventoryEntryLinks defines model for InventoryEntryLinks.
type InventoryEntryLinks struct {
	Self string `json:"self"`
}

// InventoryEntryResponse defines model for InventoryEntryResponse.
type InventoryEntryResponse struct {
	Data []InventoryEntry `json:"data"`
}

// InventoryEntryVersion defines model for InventoryEntryVersion.
type InventoryEntryVersion struct {
	Attachments *AttachmentsList            `json:"attachments,omitempty"`
	Description string                      `json:"description"`
	Digest      string                      `json:"digest"`
	ExternalID  string                      `json:"externalID"`
	Links       *InventoryEntryVersionLinks `json:"links,omitempty"`

	// Repo The name of the source repository where the inventory entry or version resides.
	// May be left empty when there is only a single repository served by the backend and thus there is not need for
	// disambiguation. See also '/repos'
	Repo      *SourceRepository `json:"repo,omitempty"`
	Timestamp string            `json:"timestamp"`
	TmID      string            `json:"tmID"`
	Version   ModelVersion      `json:"version"`
}

// InventoryEntryVersionLinks defines model for InventoryEntryVersionLinks.
type InventoryEntryVersionLinks struct {
	Content string `json:"content"`
	Self    string `json:"self"`
}

// InventoryEntryVersionResponse defines model for InventoryEntryVersionResponse.
type InventoryEntryVersionResponse struct {
	Data InventoryEntryVersion `json:"data"`
}

// InventoryEntryVersionsResponse defines model for InventoryEntryVersionsResponse.
type InventoryEntryVersionsResponse struct {
	Data []InventoryEntryVersion `json:"data"`
}

// InventoryResponse defines model for InventoryResponse.
type InventoryResponse struct {
	Data []InventoryEntry `json:"data"`
	Meta *Meta            `json:"meta,omitempty"`
}

// ManufacturersResponse defines model for ManufacturersResponse.
type ManufacturersResponse struct {
	Data []string `json:"data"`
}

// Meta defines model for Meta.
type Meta struct {
	Page *MetaPage `json:"page,omitempty"`
}

// MetaPage defines model for MetaPage.
type MetaPage struct {
	Elements int `json:"elements"`
}

// ModelVersion defines model for ModelVersion.
type ModelVersion struct {
	Model string `json:"model"`
}

// MpnsResponse defines model for MpnsResponse.
type MpnsResponse struct {
	Data []string `json:"data"`
}

// RepoDescription defines model for RepoDescription.
type RepoDescription struct {
	Description *string `json:"description,omitempty"`
	Name        string  `json:"name"`
}

// ReposResponse defines model for ReposResponse.
type ReposResponse struct {
	Data []RepoDescription `json:"data"`
}

// SchemaAuthor defines model for SchemaAuthor.
type SchemaAuthor struct {
	SchemaName string `json:"schema:name"`
}

// SchemaManufacturer defines model for SchemaManufacturer.
type SchemaManufacturer struct {
	SchemaName string `json:"schema:name"`
}

// SourceRepository The name of the source repository where the inventory entry or version resides.
// May be left empty when there is only a single repository served by the backend and thus there is not need for
// disambiguation. See also '/repos'
type SourceRepository = string

// AttachmentFileName defines model for AttachmentFileName.
type AttachmentFileName = string

// FetchName defines model for FetchName.
type FetchName = string

// ForceImport defines model for ForceImport.
type ForceImport = bool

// RepoConstraint defines model for RepoConstraint.
type RepoConstraint = string

// RepoDisambiguator defines model for RepoDisambiguator.
type RepoDisambiguator = string

// TMID defines model for TMID.
type TMID = string

// TMName defines model for TMName.
type TMName = string

// UnauthorizedError defines model for UnauthorizedError.
type UnauthorizedError = ErrorResponse

// GetCompletionsParams defines parameters for GetCompletions.
type GetCompletionsParams struct {
	// Kind Kind of data to complete
	Kind *GetCompletionsParamsKind `form:"kind,omitempty" json:"kind,omitempty"`

	// Args Current args
	Args *[]string `form:"args,omitempty" json:"args,omitempty"`

	// ToComplete Data to complete
	ToComplete *string `form:"toComplete,omitempty" json:"toComplete,omitempty"`
}

// GetCompletionsParamsKind defines parameters for GetCompletions.
type GetCompletionsParamsKind string

// GetAuthorsParams defines parameters for GetAuthors.
type GetAuthorsParams struct {
	// FilterManufacturer Filters the authors according to whether they have inventory entries
	// which belong to at least one of the given manufacturers with an exact match.
	// The filter works additive to other filters.
	FilterManufacturer *string `form:"filter.manufacturer,omitempty" json:"filter.manufacturer,omitempty"`

	// FilterMpn Filters the authors according to whether they have inventory entries
	// which belong to at least one of the given mpn (manufacturer part number) with an exact match.
	// The filter works additive to other filters.
	FilterMpn *string `form:"filter.mpn,omitempty" json:"filter.mpn,omitempty"`

	// Search Filters the authors according to whether they have inventory entries
	// where their content matches the given search.
	// The search works additive to other filters.
	Search *string `form:"search,omitempty" json:"search,omitempty"`
}

// GetInventoryParams defines parameters for GetInventory.
type GetInventoryParams struct {
	// Repo Source repository name. Optionally constrains the results to only those from given named repository. See '/repos'
	Repo *RepoConstraint `form:"repo,omitempty" json:"repo,omitempty"`

	// FilterAuthor Filters the inventory by one or more authors having exact match.
	// The filter works additive to other filters.
	FilterAuthor *string `form:"filter.author,omitempty" json:"filter.author,omitempty"`

	// FilterManufacturer Filters the inventory by one or more manufacturers having exact match.
	// The filter works additive to other filters.
	FilterManufacturer *string `form:"filter.manufacturer,omitempty" json:"filter.manufacturer,omitempty"`

	// FilterMpn Filters the inventory by one ore more mpn (manufacturer part number) having exact match.
	// The filter works additive to other filters.
	FilterMpn *string `form:"filter.mpn,omitempty" json:"filter.mpn,omitempty"`

	// FilterName Filters the inventory by inventory entry name having a prefix match of full path parts.
	// The filter works additive to other filters.
	FilterName *string `form:"filter.name,omitempty" json:"filter.name,omitempty"`

	// Search Filters the inventory according to whether the content of the inventory entries matches the given search.
	// The search works additive to other filters.
	Search *string `form:"search,omitempty" json:"search,omitempty"`
}

// GetInventoryByFetchNameParams defines parameters for GetInventoryByFetchName.
type GetInventoryByFetchNameParams struct {
	// Repo Source repository name. Optionally constrains the results to only those from given named repository. See '/repos'
	Repo *RepoConstraint `form:"repo,omitempty" json:"repo,omitempty"`
}

// GetInventoryByNameParams defines parameters for GetInventoryByName.
type GetInventoryByNameParams struct {
	// Repo Source repository name. Optionally constrains the results to only those from given named repository. See '/repos'
	Repo *RepoConstraint `form:"repo,omitempty" json:"repo,omitempty"`
}

// GetInventoryByIDParams defines parameters for GetInventoryByID.
type GetInventoryByIDParams struct {
	// Repo Source repository name. Optionally constrains the results to only those from given named repository. See '/repos'
	Repo *RepoConstraint `form:"repo,omitempty" json:"repo,omitempty"`
}

// GetManufacturersParams defines parameters for GetManufacturers.
type GetManufacturersParams struct {
	// FilterAuthor Filters the manufacturers according to whether they belong to at least one of the given authors with an exact match.
	// The filter works additive to other filters.
	FilterAuthor *string `form:"filter.author,omitempty" json:"filter.author,omitempty"`

	// FilterMpn Filters the manufacturers according to whether they have inventory entries
	// which belong to at least one of the given mpn (manufacturer part number) with an exact match.
	// The filter works additive to other filters.
	FilterMpn *string `form:"filter.mpn,omitempty" json:"filter.mpn,omitempty"`

	// Search Filters the manufacturers according to whether they have inventory entries
	// where their content matches the given search.
	// The search works additive to other filters.
	Search *string `form:"search,omitempty" json:"search,omitempty"`
}

// GetMpnsParams defines parameters for GetMpns.
type GetMpnsParams struct {
	// FilterAuthor Filters the mpns according to whether they belong to at least one of the given authors with an exact match.
	// The filter works additive to other filters.
	FilterAuthor *string `form:"filter.author,omitempty" json:"filter.author,omitempty"`

	// FilterManufacturer Filters the mpns according to whether they belong to at least one of the given manufacturers with an exact match.
	// The filter works additive to other filters.
	FilterManufacturer *string `form:"filter.manufacturer,omitempty" json:"filter.manufacturer,omitempty"`

	// Search Filters the mpns according to whether their inventory entry content matches the given search.
	// The search works additive to other filters.
	Search *string `form:"search,omitempty" json:"search,omitempty"`
}

// ImportThingModelJSONBody defines parameters for ImportThingModel.
type ImportThingModelJSONBody = map[string]interface{}

// ImportThingModelParams defines parameters for ImportThingModel.
type ImportThingModelParams struct {
	// Repo Source/target repository name. The parameter is required when repository is ambiguous. See '/repos'
	Repo *RepoDisambiguator `form:"repo,omitempty" json:"repo,omitempty"`

	// Force flag to force the import, ignoring any conflicts with existing data
	Force *ForceImport `form:"force,omitempty" json:"force,omitempty"`

	// OptPath optional path parts to append to the target path (and id) of imported TM, after the mandatory path structure
	OptPath *string `form:"optPath,omitempty" json:"optPath,omitempty"`
}

// GetThingModelByFetchNameParams defines parameters for GetThingModelByFetchName.
type GetThingModelByFetchNameParams struct {
	// Repo Source repository name. Optionally constrains the results to only those from given named repository. See '/repos'
	Repo *RepoConstraint `form:"repo,omitempty" json:"repo,omitempty"`

	// RestoreId restore the TM's original external id, if it had one
	RestoreId *bool `form:"restoreId,omitempty" json:"restoreId,omitempty"`
}

// DeleteTMNameAttachmentParams defines parameters for DeleteTMNameAttachment.
type DeleteTMNameAttachmentParams struct {
	// Repo Source/target repository name. The parameter is required when repository is ambiguous. See '/repos'
	Repo *RepoDisambiguator `form:"repo,omitempty" json:"repo,omitempty"`
}

// GetTMNameAttachmentParams defines parameters for GetTMNameAttachment.
type GetTMNameAttachmentParams struct {
	// Repo Source/target repository name. The parameter is required when repository is ambiguous. See '/repos'
	Repo *RepoDisambiguator `form:"repo,omitempty" json:"repo,omitempty"`

	// Concat Fetch a concatenation of the attachment to a TM name and homonymous attachments to all of its versions
	Concat *bool `form:"concat,omitempty" json:"concat,omitempty"`
}

// PutTMNameAttachmentParams defines parameters for PutTMNameAttachment.
type PutTMNameAttachmentParams struct {
	// Repo Source/target repository name. The parameter is required when repository is ambiguous. See '/repos'
	Repo *RepoDisambiguator `form:"repo,omitempty" json:"repo,omitempty"`

	// Force flag to force the import, ignoring any conflicts with existing data
	Force *ForceImport `form:"force,omitempty" json:"force,omitempty"`
}

// DeleteThingModelByIdParams defines parameters for DeleteThingModelById.
type DeleteThingModelByIdParams struct {
	// Repo Source/target repository name. The parameter is required when repository is ambiguous. See '/repos'
	Repo *RepoDisambiguator `form:"repo,omitempty" json:"repo,omitempty"`

	// Force flag to force the deletion. must be set to "true"
	Force string `form:"force" json:"force"`
}

// GetThingModelByIdParams defines parameters for GetThingModelById.
type GetThingModelByIdParams struct {
	// Repo Source repository name. Optionally constrains the results to only those from given named repository. See '/repos'
	Repo *RepoConstraint `form:"repo,omitempty" json:"repo,omitempty"`

	// RestoreId restore the TM's original external id, if it had one
	RestoreId *bool `form:"restoreId,omitempty" json:"restoreId,omitempty"`
}

// DeleteThingModelAttachmentByNameParams defines parameters for DeleteThingModelAttachmentByName.
type DeleteThingModelAttachmentByNameParams struct {
	// Repo Source/target repository name. The parameter is required when repository is ambiguous. See '/repos'
	Repo *RepoDisambiguator `form:"repo,omitempty" json:"repo,omitempty"`
}

// GetThingModelAttachmentByNameParams defines parameters for GetThingModelAttachmentByName.
type GetThingModelAttachmentByNameParams struct {
	// Repo Source/target repository name. The parameter is required when repository is ambiguous. See '/repos'
	Repo *RepoDisambiguator `form:"repo,omitempty" json:"repo,omitempty"`
}

// PutTMIDAttachmentParams defines parameters for PutTMIDAttachment.
type PutTMIDAttachmentParams struct {
	// Repo Source/target repository name. The parameter is required when repository is ambiguous. See '/repos'
	Repo *RepoDisambiguator `form:"repo,omitempty" json:"repo,omitempty"`

	// Force flag to force the import, ignoring any conflicts with existing data
	Force *ForceImport `form:"force,omitempty" json:"force,omitempty"`
}

// ImportThingModelJSONRequestBody defines body for ImportThingModel for application/json ContentType.
type ImportThingModelJSONRequestBody = ImportThingModelJSONBody
