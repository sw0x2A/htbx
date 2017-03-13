package htbx

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httputil"
	"strconv"
)

func init() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/status/{status:[0-9]+}", statusCodeHandler)
	rtr.HandleFunc("/dump", dumpRequestHandler)
	rtr.HandleFunc("/ip", remoteAddrHandler)
	rtr.HandleFunc("/useragent", userAgentHandler)
	http.Handle("/", rtr)
}

func dumpRequestHandler(w http.ResponseWriter, r *http.Request) {
	rd, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(w, string(rd))
}

func remoteAddrHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(w, r.RemoteAddr)
}

func userAgentHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(w, r.Header.Get("User-Agent"))
}

func statusCodeHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ss := params["status"]
	si, err := strconv.Atoi(ss)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	st := http.StatusText(si)
	if st == "" {
		st = ss
	}
	w.Header().Set("Content-Type", "text/plain")
	http.Error(w, st, si)
	return
}
