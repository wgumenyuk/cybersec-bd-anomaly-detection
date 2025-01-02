package common

import "net/http"

var Methods = []string{
	http.MethodGet,
	http.MethodPost,
	http.MethodPut,
	http.MethodDelete,
}

var Status = []int{
	http.StatusOK,
	http.StatusNotFound,
	http.StatusBadRequest,
	http.StatusPermanentRedirect,
	http.StatusUnauthorized,
	http.StatusForbidden,
	http.StatusInternalServerError,
}

var GetEndpoints = []string{
	"/static/js/index.js",
	"/static/css/index.css",
	"/static/logo.png",
	"/static/favicon.ico",
	"/",
	"/about",
	"/blog",
	"/contact",
	"/api/v1/users",
	"/api/v1/search",
	"/api/v1/products",
	"/api/v1/orders",
}

var PostEndpoints = []string{
	"/api/v1/login",
	"/api/v1/register",
	"/api/v1/orders",
}

var PutEndpoints = []string{
	"/api/v1/users",
	"/api/v1/orders",
}

var DeleteEndpoints = []string{
	"/api/v1/users",
	"/api/v1/orders",
}
