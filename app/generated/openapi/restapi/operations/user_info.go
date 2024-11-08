// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UserInfoHandlerFunc turns a function with the right signature into a user info handler
type UserInfoHandlerFunc func(UserInfoParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UserInfoHandlerFunc) Handle(params UserInfoParams) middleware.Responder {
	return fn(params)
}

// UserInfoHandler interface for that can handle valid user info params
type UserInfoHandler interface {
	Handle(UserInfoParams) middleware.Responder
}

// NewUserInfo creates a new http.Handler for the user info operation
func NewUserInfo(ctx *middleware.Context, handler UserInfoHandler) *UserInfo {
	return &UserInfo{Context: ctx, Handler: handler}
}

/* UserInfo swagger:route GET /api/me userInfo

UserInfo user info API

*/
type UserInfo struct {
	Context *middleware.Context
	Handler UserInfoHandler
}

func (o *UserInfo) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUserInfoParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}