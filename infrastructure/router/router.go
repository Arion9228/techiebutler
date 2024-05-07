package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type MuxRouter struct{}

var (
	muxRouterInstance = mux.NewRouter()
)

func NewMuxRouter() *MuxRouter {
	return &MuxRouter{}
}

func (*MuxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxRouterInstance.HandleFunc(uri, f).Methods("GET")
}

func (*MuxRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxRouterInstance.HandleFunc(uri, f).Methods("POST")
}

func (*MuxRouter) PUT(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxRouterInstance.HandleFunc(uri, f).Methods("PUT")
}

func (*MuxRouter) DELETE(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxRouterInstance.HandleFunc(uri, f).Methods("DELETE")
}

func (*MuxRouter) SERVE(port string) {
	fmt.Printf("Mux HTTP server running on port %v", port)
	http.ListenAndServe(port, muxRouterInstance)
}
