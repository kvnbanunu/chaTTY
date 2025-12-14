package server

import (
	"flag"
	"log/slog"
	"os"
)

// Args stores relevant fields used to initialize the server.
type Args struct {
	Port uint16
}

// ParseArgs parse command line arguments are handle accordingly.
func ParseArgs() *Args {
	debug := flag.Bool("d", false, "enable debug logging")
	port := flag.Uint("p", 8080, "port to listen on")
	flag.Parse()

	setLogLevel(*debug)

	return &Args{
		Port: uint16(*port),
	}
}

// setLogger sets the slog level to Debug if debug is true.
// Otherwise the default level is Info.
func setLogLevel(debug bool) {
	level := slog.LevelInfo
	if debug {
		level = slog.LevelDebug
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	}))
	slog.SetDefault(logger)
}
