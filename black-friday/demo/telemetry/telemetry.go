package telemetry

import (
	"encoding/json"
	"net/http"
	"net/http/pprof"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Serve(addr string) error {
	s := http.NewServeMux()
	s.HandleFunc("/healthz", func(rw http.ResponseWriter, req *http.Request) {
		statusCheck(rw)
	})
	s.HandleFunc("/readiness", func(rw http.ResponseWriter, req *http.Request) {
		statusCheck(rw)
	})
	s.Handle("/metrics", promhttp.Handler())
	s.Handle("/debug/pprof/", http.HandlerFunc(pprof.Index))
	s.Handle("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
	s.Handle("/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
	s.Handle("/debug/pprof/symbol", http.HandlerFunc(pprof.Symbol))
	s.Handle("/debug/pprof/trace", http.HandlerFunc(pprof.Trace))

	return http.ListenAndServe(addr, s)
}

func statusCheck(rw http.ResponseWriter) {
	statusCode := http.StatusOK
	status := map[string]string{
		"status": "UP",
	}
	rw.WriteHeader(statusCode)
	if err := json.NewEncoder(rw).Encode(status); err != nil {
		rw.Write([]byte("failed to run status check"))
	}
}
