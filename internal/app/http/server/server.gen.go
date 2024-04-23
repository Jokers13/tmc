// Package server provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oapi-codegen/runtime"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get completions for shell completion script
	// (GET /.completions)
	GetCompletions(w http.ResponseWriter, r *http.Request, params GetCompletionsParams)
	// Get the contained authors of the inventory
	// (GET /authors)
	GetAuthors(w http.ResponseWriter, r *http.Request, params GetAuthorsParams)
	// Get the overall health of the service
	// (GET /healthz)
	GetHealth(w http.ResponseWriter, r *http.Request)
	// Returns the liveliness of the service
	// (GET /healthz/live)
	GetHealthLive(w http.ResponseWriter, r *http.Request)
	// Returns the readiness of the service
	// (GET /healthz/ready)
	GetHealthReady(w http.ResponseWriter, r *http.Request)
	// Returns whether the service is initialized
	// (GET /healthz/startup)
	GetHealthStartup(w http.ResponseWriter, r *http.Request)
	// Get the inventory of the catalog
	// (GET /inventory)
	GetInventory(w http.ResponseWriter, r *http.Request, params GetInventoryParams)
	// Get an inventory entry by it's name
	// (GET /inventory/{name})
	GetInventoryByName(w http.ResponseWriter, r *http.Request, name string)
	// Get the versions of an inventory entry
	// (GET /inventory/{name}/.versions)
	GetInventoryVersionsByName(w http.ResponseWriter, r *http.Request, name string)
	// Get the contained manufacturers of the inventory
	// (GET /manufacturers)
	GetManufacturers(w http.ResponseWriter, r *http.Request, params GetManufacturersParams)
	// Get the contained mpns (manufacturer part numbers) of the inventory
	// (GET /mpns)
	GetMpns(w http.ResponseWriter, r *http.Request, params GetMpnsParams)
	// Push a new Thing Model
	// (POST /thing-models)
	PushThingModel(w http.ResponseWriter, r *http.Request)
	// Delete a Thing Model by ID
	// (DELETE /thing-models/{tmIDOrName})
	DeleteThingModelById(w http.ResponseWriter, r *http.Request, tmIDOrName string, params DeleteThingModelByIdParams)
	// Get the content of a Thing Model by its ID or fetch name
	// (GET /thing-models/{tmIDOrName})
	GetThingModelById(w http.ResponseWriter, r *http.Request, tmIDOrName string, params GetThingModelByIdParams)
	// Get the attachments of a Thing Model or an inventory entry
	// (GET /thing-models/{tmIDOrName}/.attachments)
	GetThingModelAttachmentListByName(w http.ResponseWriter, r *http.Request, tmIDOrName string)
	// Get the actual content of an attachment of a Thing Model or an inventory entry
	// (DELETE /thing-models/{tmIDOrName}/.attachments/{attachmentFileName})
	DeleteThingModelAttachmentByName(w http.ResponseWriter, r *http.Request, tmIDOrName string, attachmentFileName string)
	// Get the actual content of an attachment of a Thing Model or an inventory entry
	// (GET /thing-models/{tmIDOrName}/.attachments/{attachmentFileName})
	GetThingModelAttachmentByName(w http.ResponseWriter, r *http.Request, tmIDOrName string, attachmentFileName string)
	// Get the actual content of an attachment of a Thing Model or an inventory entry
	// (PUT /thing-models/{tmIDOrName}/.attachments/{attachmentFileName})
	PutThingModelAttachmentByName(w http.ResponseWriter, r *http.Request, tmIDOrName string, attachmentFileName string)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// GetCompletions operation middleware
func (siw *ServerInterfaceWrapper) GetCompletions(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetCompletionsParams

	// ------------- Optional query parameter "kind" -------------

	err = runtime.BindQueryParameter("form", true, false, "kind", r.URL.Query(), &params.Kind)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "kind", Err: err})
		return
	}

	// ------------- Optional query parameter "toComplete" -------------

	err = runtime.BindQueryParameter("form", true, false, "toComplete", r.URL.Query(), &params.ToComplete)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "toComplete", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetCompletions(w, r, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetAuthors operation middleware
func (siw *ServerInterfaceWrapper) GetAuthors(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetAuthorsParams

	// ------------- Optional query parameter "filter.manufacturer" -------------

	err = runtime.BindQueryParameter("form", true, false, "filter.manufacturer", r.URL.Query(), &params.FilterManufacturer)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "filter.manufacturer", Err: err})
		return
	}

	// ------------- Optional query parameter "filter.mpn" -------------

	err = runtime.BindQueryParameter("form", true, false, "filter.mpn", r.URL.Query(), &params.FilterMpn)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "filter.mpn", Err: err})
		return
	}

	// ------------- Optional query parameter "search" -------------

	err = runtime.BindQueryParameter("form", true, false, "search", r.URL.Query(), &params.Search)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "search", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetAuthors(w, r, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetHealth operation middleware
func (siw *ServerInterfaceWrapper) GetHealth(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetHealth(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetHealthLive operation middleware
func (siw *ServerInterfaceWrapper) GetHealthLive(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetHealthLive(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetHealthReady operation middleware
func (siw *ServerInterfaceWrapper) GetHealthReady(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetHealthReady(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetHealthStartup operation middleware
func (siw *ServerInterfaceWrapper) GetHealthStartup(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetHealthStartup(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetInventory operation middleware
func (siw *ServerInterfaceWrapper) GetInventory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetInventoryParams

	// ------------- Optional query parameter "filter.author" -------------

	err = runtime.BindQueryParameter("form", true, false, "filter.author", r.URL.Query(), &params.FilterAuthor)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "filter.author", Err: err})
		return
	}

	// ------------- Optional query parameter "filter.manufacturer" -------------

	err = runtime.BindQueryParameter("form", true, false, "filter.manufacturer", r.URL.Query(), &params.FilterManufacturer)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "filter.manufacturer", Err: err})
		return
	}

	// ------------- Optional query parameter "filter.mpn" -------------

	err = runtime.BindQueryParameter("form", true, false, "filter.mpn", r.URL.Query(), &params.FilterMpn)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "filter.mpn", Err: err})
		return
	}

	// ------------- Optional query parameter "filter.name" -------------

	err = runtime.BindQueryParameter("form", true, false, "filter.name", r.URL.Query(), &params.FilterName)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "filter.name", Err: err})
		return
	}

	// ------------- Optional query parameter "search" -------------

	err = runtime.BindQueryParameter("form", true, false, "search", r.URL.Query(), &params.Search)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "search", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetInventory(w, r, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetInventoryByName operation middleware
func (siw *ServerInterfaceWrapper) GetInventoryByName(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "name" -------------
	var name string

	err = runtime.BindStyledParameterWithOptions("simple", "name", mux.Vars(r)["name"], &name, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "name", Err: err})
		return
	}

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetInventoryByName(w, r, name)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetInventoryVersionsByName operation middleware
func (siw *ServerInterfaceWrapper) GetInventoryVersionsByName(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "name" -------------
	var name string

	err = runtime.BindStyledParameterWithOptions("simple", "name", mux.Vars(r)["name"], &name, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "name", Err: err})
		return
	}

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetInventoryVersionsByName(w, r, name)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetManufacturers operation middleware
func (siw *ServerInterfaceWrapper) GetManufacturers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetManufacturersParams

	// ------------- Optional query parameter "filter.author" -------------

	err = runtime.BindQueryParameter("form", true, false, "filter.author", r.URL.Query(), &params.FilterAuthor)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "filter.author", Err: err})
		return
	}

	// ------------- Optional query parameter "filter.mpn" -------------

	err = runtime.BindQueryParameter("form", true, false, "filter.mpn", r.URL.Query(), &params.FilterMpn)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "filter.mpn", Err: err})
		return
	}

	// ------------- Optional query parameter "search" -------------

	err = runtime.BindQueryParameter("form", true, false, "search", r.URL.Query(), &params.Search)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "search", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetManufacturers(w, r, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetMpns operation middleware
func (siw *ServerInterfaceWrapper) GetMpns(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetMpnsParams

	// ------------- Optional query parameter "filter.author" -------------

	err = runtime.BindQueryParameter("form", true, false, "filter.author", r.URL.Query(), &params.FilterAuthor)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "filter.author", Err: err})
		return
	}

	// ------------- Optional query parameter "filter.manufacturer" -------------

	err = runtime.BindQueryParameter("form", true, false, "filter.manufacturer", r.URL.Query(), &params.FilterManufacturer)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "filter.manufacturer", Err: err})
		return
	}

	// ------------- Optional query parameter "search" -------------

	err = runtime.BindQueryParameter("form", true, false, "search", r.URL.Query(), &params.Search)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "search", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetMpns(w, r, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PushThingModel operation middleware
func (siw *ServerInterfaceWrapper) PushThingModel(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PushThingModel(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// DeleteThingModelById operation middleware
func (siw *ServerInterfaceWrapper) DeleteThingModelById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "tmIDOrName" -------------
	var tmIDOrName string

	err = runtime.BindStyledParameterWithOptions("simple", "tmIDOrName", mux.Vars(r)["tmIDOrName"], &tmIDOrName, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "tmIDOrName", Err: err})
		return
	}

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params DeleteThingModelByIdParams

	// ------------- Required query parameter "force" -------------

	if paramValue := r.URL.Query().Get("force"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "force"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "force", r.URL.Query(), &params.Force)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "force", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteThingModelById(w, r, tmIDOrName, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetThingModelById operation middleware
func (siw *ServerInterfaceWrapper) GetThingModelById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "tmIDOrName" -------------
	var tmIDOrName string

	err = runtime.BindStyledParameterWithOptions("simple", "tmIDOrName", mux.Vars(r)["tmIDOrName"], &tmIDOrName, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "tmIDOrName", Err: err})
		return
	}

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetThingModelByIdParams

	// ------------- Optional query parameter "restoreId" -------------

	err = runtime.BindQueryParameter("form", true, false, "restoreId", r.URL.Query(), &params.RestoreId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "restoreId", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetThingModelById(w, r, tmIDOrName, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetThingModelAttachmentListByName operation middleware
func (siw *ServerInterfaceWrapper) GetThingModelAttachmentListByName(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "tmIDOrName" -------------
	var tmIDOrName string

	err = runtime.BindStyledParameterWithOptions("simple", "tmIDOrName", mux.Vars(r)["tmIDOrName"], &tmIDOrName, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "tmIDOrName", Err: err})
		return
	}

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetThingModelAttachmentListByName(w, r, tmIDOrName)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// DeleteThingModelAttachmentByName operation middleware
func (siw *ServerInterfaceWrapper) DeleteThingModelAttachmentByName(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "tmIDOrName" -------------
	var tmIDOrName string

	err = runtime.BindStyledParameterWithOptions("simple", "tmIDOrName", mux.Vars(r)["tmIDOrName"], &tmIDOrName, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "tmIDOrName", Err: err})
		return
	}

	// ------------- Path parameter "attachmentFileName" -------------
	var attachmentFileName string

	err = runtime.BindStyledParameterWithOptions("simple", "attachmentFileName", mux.Vars(r)["attachmentFileName"], &attachmentFileName, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "attachmentFileName", Err: err})
		return
	}

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteThingModelAttachmentByName(w, r, tmIDOrName, attachmentFileName)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetThingModelAttachmentByName operation middleware
func (siw *ServerInterfaceWrapper) GetThingModelAttachmentByName(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "tmIDOrName" -------------
	var tmIDOrName string

	err = runtime.BindStyledParameterWithOptions("simple", "tmIDOrName", mux.Vars(r)["tmIDOrName"], &tmIDOrName, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "tmIDOrName", Err: err})
		return
	}

	// ------------- Path parameter "attachmentFileName" -------------
	var attachmentFileName string

	err = runtime.BindStyledParameterWithOptions("simple", "attachmentFileName", mux.Vars(r)["attachmentFileName"], &attachmentFileName, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "attachmentFileName", Err: err})
		return
	}

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetThingModelAttachmentByName(w, r, tmIDOrName, attachmentFileName)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PutThingModelAttachmentByName operation middleware
func (siw *ServerInterfaceWrapper) PutThingModelAttachmentByName(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "tmIDOrName" -------------
	var tmIDOrName string

	err = runtime.BindStyledParameterWithOptions("simple", "tmIDOrName", mux.Vars(r)["tmIDOrName"], &tmIDOrName, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "tmIDOrName", Err: err})
		return
	}

	// ------------- Path parameter "attachmentFileName" -------------
	var attachmentFileName string

	err = runtime.BindStyledParameterWithOptions("simple", "attachmentFileName", mux.Vars(r)["attachmentFileName"], &attachmentFileName, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "attachmentFileName", Err: err})
		return
	}

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PutThingModelAttachmentByName(w, r, tmIDOrName, attachmentFileName)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, GorillaServerOptions{})
}

type GorillaServerOptions struct {
	BaseURL          string
	BaseRouter       *mux.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r *mux.Router) http.Handler {
	return HandlerWithOptions(si, GorillaServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r *mux.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, GorillaServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options GorillaServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = mux.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.HandleFunc(options.BaseURL+"/thing-models/{tmIDOrName:.+}/.attachments/{attachmentFileName}", wrapper.PutThingModelAttachmentByName).Methods("PUT")

	r.HandleFunc(options.BaseURL+"/thing-models/{tmIDOrName:.+}/.attachments/{attachmentFileName}", wrapper.GetThingModelAttachmentByName).Methods("GET")

	r.HandleFunc(options.BaseURL+"/thing-models/{tmIDOrName:.+}/.attachments/{attachmentFileName}", wrapper.DeleteThingModelAttachmentByName).Methods("DELETE")

	r.HandleFunc(options.BaseURL+"/thing-models/{tmIDOrName:.+}/.attachments", wrapper.GetThingModelAttachmentListByName).Methods("GET")

	r.HandleFunc(options.BaseURL+"/thing-models/{tmIDOrName:.+}", wrapper.GetThingModelById).Methods("GET")

	r.HandleFunc(options.BaseURL+"/thing-models/{tmIDOrName:.+}", wrapper.DeleteThingModelById).Methods("DELETE")

	r.HandleFunc(options.BaseURL+"/thing-models", wrapper.PushThingModel).Methods("POST")

	r.HandleFunc(options.BaseURL+"/mpns", wrapper.GetMpns).Methods("GET")

	r.HandleFunc(options.BaseURL+"/manufacturers", wrapper.GetManufacturers).Methods("GET")

	r.HandleFunc(options.BaseURL+"/inventory/{name:.+}/.versions", wrapper.GetInventoryVersionsByName).Methods("GET")

	r.HandleFunc(options.BaseURL+"/inventory/{name:.+}", wrapper.GetInventoryByName).Methods("GET")

	r.HandleFunc(options.BaseURL+"/inventory", wrapper.GetInventory).Methods("GET")

	r.HandleFunc(options.BaseURL+"/healthz/startup", wrapper.GetHealthStartup).Methods("GET")

	r.HandleFunc(options.BaseURL+"/healthz/ready", wrapper.GetHealthReady).Methods("GET")

	r.HandleFunc(options.BaseURL+"/healthz/live", wrapper.GetHealthLive).Methods("GET")

	r.HandleFunc(options.BaseURL+"/healthz", wrapper.GetHealth).Methods("GET")

	r.HandleFunc(options.BaseURL+"/authors", wrapper.GetAuthors).Methods("GET")

	r.HandleFunc(options.BaseURL+"/.completions", wrapper.GetCompletions).Methods("GET")

	return r
}
