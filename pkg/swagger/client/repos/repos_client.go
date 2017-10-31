// Code generated by go-swagger; DO NOT EDIT.

package repos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new repos API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for repos API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteReposRepoID deletes a repo

Delete a single repo identified via its id
*/
func (a *Client) DeleteReposRepoID(params *DeleteReposRepoIDParams) (*DeleteReposRepoIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteReposRepoIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteReposRepoID",
		Method:             "DELETE",
		PathPattern:        "/repos/{repoId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteReposRepoIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteReposRepoIDNoContent), nil

}

/*
GetRepos gets some repos

Returns a list containing all repos.
*/
func (a *Client) GetRepos(params *GetReposParams) (*GetReposOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetReposParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetRepos",
		Method:             "GET",
		PathPattern:        "/repos",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetReposReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetReposOK), nil

}

/*
GetReposRepoID gets a repo

Returns a single repo by its id
*/
func (a *Client) GetReposRepoID(params *GetReposRepoIDParams) (*GetReposRepoIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetReposRepoIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetReposRepoID",
		Method:             "GET",
		PathPattern:        "/repos/{repoId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetReposRepoIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetReposRepoIDOK), nil

}

/*
PostRepos creates a repo

Adds a new repo to the repos list.
*/
func (a *Client) PostRepos(params *PostReposParams) (*PostReposNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostReposParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostRepos",
		Method:             "POST",
		PathPattern:        "/repos",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostReposReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostReposNoContent), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}