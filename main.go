package main

import (
	"fmt"
	"runtime"

	log "github.com/sirupsen/logrus"
)

func main() {
	expected := "go1.25.0"
	goVersion := runtime.Version()
	log.WithFields(log.Fields{
		"expected": expected,
		"runtime":  goVersion,
	}).Info("Go version: %s\n", goVersion)

	if goVersion != expected {
		panic(fmt.Errorf("expected version: %s, got: %s", expected, goVersion))
	}
}
