// Code generated by go-swagger; DO NOT EDIT.

package fx_api

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// Client.GetFXDeal creates a new GetFXDealRequest object
// with the default values initialized.
func (c *Client) GetFXDeal() *GetFXDealRequest {
	var ()
	return &GetFXDealRequest{

		FxDealID: c.Defaults.GetStrfmtUUID("GetFXDeal", "fx_deal_id"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type GetFXDealRequest struct {

	/*FxDealID      FX Deal ID      */

	FxDealID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *GetFXDealRequest) FromJson(j string) *GetFXDealRequest {

	return o
}

func (o *GetFXDealRequest) WithFxDealID(fxDealID strfmt.UUID) *GetFXDealRequest {

	o.FxDealID = fxDealID

	return o
}

//////////////////
// WithContext adds the context to the get f x deal Request
func (o *GetFXDealRequest) WithContext(ctx context.Context) *GetFXDealRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the get f x deal Request
func (o *GetFXDealRequest) WithHTTPClient(client *http.Client) *GetFXDealRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *GetFXDealRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param fx_deal_id
	if err := r.SetPathParam("fx_deal_id", o.FxDealID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
