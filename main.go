package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httputil"
	"strconv"
	"strings"
)

func main() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/status/{status:[0-9]+}", statusCodeHandler)
	rtr.HandleFunc("/dump", dumpRequestHandler)
	rtr.HandleFunc("/ip", remoteAddrHandler)
	rtr.HandleFunc("/useragent", userAgentHandler)
	http.ListenAndServe(":8080", globals(rtr))
}

func globals(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		h.ServeHTTP(w, r)
	})
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
	ipAddress := r.RemoteAddr
	fwdAddress := r.Header.Get("X-Forwarded-For")
	if fwdAddress != "" {
		ipAddress = fwdAddress // If it's a single IP, then awesome!
		ips := strings.Split(fwdAddress, ", ")
		if len(ips) > 1 {
			ipAddress = ips[0]
		}
	}
	fmt.Fprintln(w, ipAddress)
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
