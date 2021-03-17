package main

import "flag"

type flags struct {
	debug bool
}

func ParseFlags() flags {
	debugMode := flag.Bool("debug", false, "run in debug mode")
	flag.Parse()

	return flags{
		debug: *debugMode,
	}
}
