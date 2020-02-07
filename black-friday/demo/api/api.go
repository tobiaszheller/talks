package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func Serve(addr string) error {
	s := http.NewServeMux()
	s.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		start := time.Now()
		simulateWork()
		log.Printf("Work took %s\n", time.Since(start).String())
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte("ok"))
	})
	log.Printf("Running http server on %q\n", addr)
	return http.ListenAndServe(addr, s)
}

func simulateWork() {
	log.Printf("I am working\n")
	const (
		numberOfIterations = 100
		mapSize            = 100
	)
	for i := 0; i < numberOfIterations; i++ {
		m := map[string]string{}
		for j := 0; j < mapSize; j++ {
			m[uuid.New().String()] = uuid.New().String()
		}
		bb, _ := json.Marshal(m)
		_ = bb
	}
}
