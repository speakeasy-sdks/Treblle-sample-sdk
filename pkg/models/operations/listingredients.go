// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/Treblle-sample-sdk/pkg/models/shared"
	"net/http"
)

type ListIngredientsRequest struct {
	// A list of ingredients to filter by. If not provided all ingredients will be returned.
	Ingredients []string `queryParam:"style=form,explode=false,name=ingredients"`
}

func (o *ListIngredientsRequest) GetIngredients() []string {
	if o == nil {
		return nil
	}
	return o.Ingredients
}

type ListIngredientsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// An unknown error occurred interacting with the API.
	Error *shared.Error
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// A list of ingredients.
	Classes []shared.Ingredient
}

func (o *ListIngredientsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListIngredientsResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *ListIngredientsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListIngredientsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListIngredientsResponse) GetClasses() []shared.Ingredient {
	if o == nil {
		return nil
	}
	return o.Classes
}
