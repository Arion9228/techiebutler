package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// MuxRouter represents a router using the Gorilla Mux package.
type MuxRouter struct{}

var (
	muxRouterInstance = mux.NewRouter() // muxRouterInstance is the instance of the Gorilla Mux router.
)

// NewMuxRouter creates a new instance of MuxRouter.
func NewMuxRouter() *MuxRouter {
	return &MuxRouter{}
}

// GET registers a handler for the GET HTTP method.
func (*MuxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxRouterInstance.HandleFunc(uri, f).Methods("GET")
}

// POST registers a handler for the POST HTTP method.
func (*MuxRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxRouterInstance.HandleFunc(uri, f).Methods("POST")
}

// PUT registers a handler for the PUT HTTP method.
func (*MuxRouter) PUT(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxRouterInstance.HandleFunc(uri, f).Methods("PUT")
}

// DELETE registers a handler for the DELETE HTTP method.
func (*MuxRouter) DELETE(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxRouterInstance.HandleFunc(uri, f).Methods("DELETE")
}

// SERVE starts the HTTP server on the specified port.
func (*MuxRouter) SERVE(port string) {
	fmt.Printf("Mux HTTP server running on port %v", port)
	http.ListenAndServe(port, muxRouterInstance)
}
