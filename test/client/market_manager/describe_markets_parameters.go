// Code generated by go-swagger; DO NOT EDIT.

package market_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDescribeMarketsParams creates a new DescribeMarketsParams object
// with the default values initialized.
func NewDescribeMarketsParams() *DescribeMarketsParams {
	var ()
	return &DescribeMarketsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDescribeMarketsParamsWithTimeout creates a new DescribeMarketsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDescribeMarketsParamsWithTimeout(timeout time.Duration) *DescribeMarketsParams {
	var ()
	return &DescribeMarketsParams{

		timeout: timeout,
	}
}

// NewDescribeMarketsParamsWithContext creates a new DescribeMarketsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDescribeMarketsParamsWithContext(ctx context.Context) *DescribeMarketsParams {
	var ()
	return &DescribeMarketsParams{

		Context: ctx,
	}
}

// NewDescribeMarketsParamsWithHTTPClient creates a new DescribeMarketsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDescribeMarketsParamsWithHTTPClient(client *http.Client) *DescribeMarketsParams {
	var ()
	return &DescribeMarketsParams{
		HTTPClient: client,
	}
}

/*DescribeMarketsParams contains all the parameters to send to the API endpoint
for the describe markets operation typically these are written to a http.Request
*/
type DescribeMarketsParams struct {

	/*Limit*/
	Limit *int64
	/*MarketID*/
	MarketID []string
	/*Offset*/
	Offset *int64
	/*Owner*/
	Owner []string
	/*Reverse*/
	Reverse *bool
	/*SearchWord*/
	SearchWord *string
	/*SortKey*/
	SortKey *string
	/*Status*/
	Status []string
	/*UserID*/
	UserID []string
	/*Visibility*/
	Visibility []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the describe markets params
func (o *DescribeMarketsParams) WithTimeout(timeout time.Duration) *DescribeMarketsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the describe markets params
func (o *DescribeMarketsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the describe markets params
func (o *DescribeMarketsParams) WithContext(ctx context.Context) *DescribeMarketsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the describe markets params
func (o *DescribeMarketsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the describe markets params
func (o *DescribeMarketsParams) WithHTTPClient(client *http.Client) *DescribeMarketsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the describe markets params
func (o *DescribeMarketsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the describe markets params
func (o *DescribeMarketsParams) WithLimit(limit *int64) *DescribeMarketsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the describe markets params
func (o *DescribeMarketsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithMarketID adds the marketID to the describe markets params
func (o *DescribeMarketsParams) WithMarketID(marketID []string) *DescribeMarketsParams {
	o.SetMarketID(marketID)
	return o
}

// SetMarketID adds the marketId to the describe markets params
func (o *DescribeMarketsParams) SetMarketID(marketID []string) {
	o.MarketID = marketID
}

// WithOffset adds the offset to the describe markets params
func (o *DescribeMarketsParams) WithOffset(offset *int64) *DescribeMarketsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the describe markets params
func (o *DescribeMarketsParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithOwner adds the owner to the describe markets params
func (o *DescribeMarketsParams) WithOwner(owner []string) *DescribeMarketsParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the describe markets params
func (o *DescribeMarketsParams) SetOwner(owner []string) {
	o.Owner = owner
}

// WithReverse adds the reverse to the describe markets params
func (o *DescribeMarketsParams) WithReverse(reverse *bool) *DescribeMarketsParams {
	o.SetReverse(reverse)
	return o
}

// SetReverse adds the reverse to the describe markets params
func (o *DescribeMarketsParams) SetReverse(reverse *bool) {
	o.Reverse = reverse
}

// WithSearchWord adds the searchWord to the describe markets params
func (o *DescribeMarketsParams) WithSearchWord(searchWord *string) *DescribeMarketsParams {
	o.SetSearchWord(searchWord)
	return o
}

// SetSearchWord adds the searchWord to the describe markets params
func (o *DescribeMarketsParams) SetSearchWord(searchWord *string) {
	o.SearchWord = searchWord
}

// WithSortKey adds the sortKey to the describe markets params
func (o *DescribeMarketsParams) WithSortKey(sortKey *string) *DescribeMarketsParams {
	o.SetSortKey(sortKey)
	return o
}

// SetSortKey adds the sortKey to the describe markets params
func (o *DescribeMarketsParams) SetSortKey(sortKey *string) {
	o.SortKey = sortKey
}

// WithStatus adds the status to the describe markets params
func (o *DescribeMarketsParams) WithStatus(status []string) *DescribeMarketsParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the describe markets params
func (o *DescribeMarketsParams) SetStatus(status []string) {
	o.Status = status
}

// WithUserID adds the userID to the describe markets params
func (o *DescribeMarketsParams) WithUserID(userID []string) *DescribeMarketsParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the describe markets params
func (o *DescribeMarketsParams) SetUserID(userID []string) {
	o.UserID = userID
}

// WithVisibility adds the visibility to the describe markets params
func (o *DescribeMarketsParams) WithVisibility(visibility []string) *DescribeMarketsParams {
	o.SetVisibility(visibility)
	return o
}

// SetVisibility adds the visibility to the describe markets params
func (o *DescribeMarketsParams) SetVisibility(visibility []string) {
	o.Visibility = visibility
}

// WriteToRequest writes these params to a swagger request
func (o *DescribeMarketsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	valuesMarketID := o.MarketID

	joinedMarketID := swag.JoinByFormat(valuesMarketID, "multi")
	// query array param market_id
	if err := r.SetQueryParam("market_id", joinedMarketID...); err != nil {
		return err
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	valuesOwner := o.Owner

	joinedOwner := swag.JoinByFormat(valuesOwner, "multi")
	// query array param owner
	if err := r.SetQueryParam("owner", joinedOwner...); err != nil {
		return err
	}

	if o.Reverse != nil {

		// query param reverse
		var qrReverse bool
		if o.Reverse != nil {
			qrReverse = *o.Reverse
		}
		qReverse := swag.FormatBool(qrReverse)
		if qReverse != "" {
			if err := r.SetQueryParam("reverse", qReverse); err != nil {
				return err
			}
		}

	}

	if o.SearchWord != nil {

		// query param search_word
		var qrSearchWord string
		if o.SearchWord != nil {
			qrSearchWord = *o.SearchWord
		}
		qSearchWord := qrSearchWord
		if qSearchWord != "" {
			if err := r.SetQueryParam("search_word", qSearchWord); err != nil {
				return err
			}
		}

	}

	if o.SortKey != nil {

		// query param sort_key
		var qrSortKey string
		if o.SortKey != nil {
			qrSortKey = *o.SortKey
		}
		qSortKey := qrSortKey
		if qSortKey != "" {
			if err := r.SetQueryParam("sort_key", qSortKey); err != nil {
				return err
			}
		}

	}

	valuesStatus := o.Status

	joinedStatus := swag.JoinByFormat(valuesStatus, "multi")
	// query array param status
	if err := r.SetQueryParam("status", joinedStatus...); err != nil {
		return err
	}

	valuesUserID := o.UserID

	joinedUserID := swag.JoinByFormat(valuesUserID, "multi")
	// query array param user_id
	if err := r.SetQueryParam("user_id", joinedUserID...); err != nil {
		return err
	}

	valuesVisibility := o.Visibility

	joinedVisibility := swag.JoinByFormat(valuesVisibility, "multi")
	// query array param visibility
	if err := r.SetQueryParam("visibility", joinedVisibility...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
