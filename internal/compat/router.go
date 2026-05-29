// Package compat provides a compatibility layer that routes legacy command-line
// flags (--cap=N --topics, --cap=N --overview) through the new chapter system
// when the requested chapter has been migrated, falling back to the legacy
// menu system otherwise.
//
// This enables incremental migration: each chapter can be migrated independently
// without breaking existing usage patterns.
package compat

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fabianoflorentino/aprendago/internal/chapter"
)

// Route tries to handle legacy --cap=N --topics/--overview flags
// using the new chapter system.
//
// Returns true if the request was handled (caller should skip legacy fallback).
// Returns false if the request should fall through to legacy menu.Options().
func Route(args []string) bool {
	capNum := -1
	var flags []string

	for _, arg := range args {
		if strings.HasPrefix(arg, "--cap=") {
			n, err := strconv.Atoi(strings.TrimPrefix(arg, "--cap="))
			if err == nil {
				capNum = n
			}
			continue
		}
		flags = append(flags, arg)
	}

	if capNum < 0 {
		return false // no --cap flag, not compat's domain
	}

	ch := chapter.Get(capNum)
	if ch == nil {
		return false // chapter not registered yet, let legacy handle
	}

	// Determine action from remaining flags
	action := "overview" // default when no specific flag

	for _, f := range flags {
		switch f {
		case "--topics":
			action = "topics"
		case "--overview":
			action = "overview"
		default:
			// Individual topic flags (--bem-vindo, --estudos, etc.)
			// have complex name-to-section-title mapping handled by the
			// legacy system. Let legacy handle these.
			return false
		}
	}

	switch action {
	case "topics":
		topics, err := ch.Topics()
		if err != nil {
			fmt.Printf("Erro: %v\n", err)
			return true
		}
		fmt.Println()
		for _, t := range topics {
			fmt.Printf("  %s\n", t)
		}
	case "overview":
		if err := ch.Overview(); err != nil {
			fmt.Printf("Erro: %v\n", err)
		}
	}

	return true
}
