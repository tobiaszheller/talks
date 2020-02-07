package main

import (
	"log"

	"github.com/kelseyhightower/envconfig"

	"github.com/tobiaszheller/talks/black-friday/demo/api"
	"github.com/tobiaszheller/talks/black-friday/demo/telemetry"
)

type config struct {
	TelemetryAddr string `envconfig:"TELEMETRY_ADDR"`
	HttpAddr      string `envconfig:"HTTP_ADDR"`
}

func main() {
	var cfg config
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatal(err)
	}

	go runTelemetry(cfg.TelemetryAddr)
	api.Serve(cfg.HttpAddr)
}

func runTelemetry(addr string) {
	log.Printf("Running telemetry on %q\n", addr)
	if err := telemetry.Serve(addr); err != nil {
		log.Fatal(err)
	}
}
