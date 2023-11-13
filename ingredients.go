// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package trebllesamplesdk

import (
	"bytes"
	"context"
	"fmt"
	"github.com/speakeasy-sdks/Treblle-sample-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/Treblle-sample-sdk/pkg/models/sdkerrors"
	"github.com/speakeasy-sdks/Treblle-sample-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/Treblle-sample-sdk/pkg/utils"
	"io"
	"net/http"
	"strings"
)

// Ingredients - The ingredients endpoints.
type Ingredients struct {
	sdkConfiguration sdkConfiguration
}

func newIngredients(sdkConfig sdkConfiguration) *Ingredients {
	return &Ingredients{
		sdkConfiguration: sdkConfig,
	}
}

// ListIngredients - Get a list of ingredients.
// Get a list of ingredients, if authenticated this will include stock levels and product codes otherwise it will only include public information.
func (s *Ingredients) ListIngredients(ctx context.Context, ingredients []string) (*operations.ListIngredientsResponse, error) {
	request := operations.ListIngredientsRequest{
		Ingredients: ingredients,
	}

	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url := strings.TrimSuffix(baseURL, "/") + "/ingredients"

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", s.sdkConfiguration.UserAgent)

	if err := utils.PopulateQueryParams(ctx, req, request, nil); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.ListIngredientsResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out []shared.Ingredient
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.Classes = out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out sdkerrors.APIError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}
			return nil, &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out shared.Error
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.Error = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	}

	return res, nil
}
