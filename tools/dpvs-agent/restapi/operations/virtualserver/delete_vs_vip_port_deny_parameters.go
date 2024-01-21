// Code generated by go-swagger; DO NOT EDIT.

package virtualserver

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"

	"github.com/dpvs-agent/models"
)

// NewDeleteVsVipPortDenyParams creates a new DeleteVsVipPortDenyParams object
//
// There are no default values defined in the spec.
func NewDeleteVsVipPortDenyParams() DeleteVsVipPortDenyParams {

	return DeleteVsVipPortDenyParams{}
}

// DeleteVsVipPortDenyParams contains all the bound params for the delete vs vip port deny operation
// typically these are obtained from a http.Request
//
// swagger:parameters DeleteVsVipPortDeny
type DeleteVsVipPortDenyParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	VipPort string
	/*
	  In: body
	*/
	ACL *models.ACLAddrList
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeleteVsVipPortDenyParams() beforehand.
func (o *DeleteVsVipPortDenyParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rVipPort, rhkVipPort, _ := route.Params.GetOK("VipPort")
	if err := o.bindVipPort(rVipPort, rhkVipPort, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.ACLAddrList
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("acl", "body", "", err))
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(r.Context())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.ACL = &body
			}
		}
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindVipPort binds and validates parameter VipPort from path.
func (o *DeleteVsVipPortDenyParams) bindVipPort(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.VipPort = raw

	return nil
}
