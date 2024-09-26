package main

import (
	"github.com/pin/tftp/v3"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"io"
	"os"
	"path/filepath"
	"time"
)

// readHandler is called when client starts file download from server
func readHandler(filename string, rf io.ReaderFrom) error {
	filename = filepath.Join("/var/lib/tftp", filename)
	file, err := os.Open(filename)
	if err != nil {
		log.Error().Err(err).Msgf("readHandler - failed opening file")
		return err
	}
	n, err := rf.ReadFrom(file)
	if err != nil {
		log.Error().Err(err).Msgf("readHandler - failed sending file")
		return err
	}
	log.Info().Msgf("%d bytes sent", n)
	return nil
}

// Hook for logging on every transfer completion or failure.
type logHook struct{}

func (h *logHook) OnSuccess(stats tftp.TransferStats) {
	log.Info().Msgf("Transfer of %s to %s complete", stats.Filename, stats.RemoteAddr)
}
func (h *logHook) OnFailure(stats tftp.TransferStats, err error) {
	log.Err(err).Msgf("Transfer of %s to %s failed", stats.Filename, stats.RemoteAddr)
}

func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	// use nil in place of handler to disable read or write operations
	s := tftp.NewServer(readHandler, nil)
	s.SetHook(&logHook{})

	// FYI: https://medium.com/@darpanmalhotra/exposing-tftp-server-as-kubernetes-service-part-7-ac26461354bc
	// Explanation from Sidero
	// A standard TFTP server implementation receives requests on port 69 and
	// allocates a new high port (over 1024) dedicated to that request. In single
	// port mode, the same port is used for transmit and receive. If the server
	// is started on port 69, all communication will be done on port 69.
	// This option is required since the Kubernetes service definition defines a
	// single port.
	s.EnableSinglePort()

	s.SetTimeout(5 * time.Second)  // optional
	err := s.ListenAndServe(":69") // blocks until s.Shutdown() is called
	if err != nil {
		log.Error().Err(err).Msgf("server failed listening")
		os.Exit(1)
	}
}
