// Package compat provides a compatibility layer that routes legacy command-line
// flags (--cap=N --topics, --cap=N --overview) through the new chapter system.
//
// This enables gradual migration: legacy flag syntax still works even after
// the old menu system has been removed.
package compat

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fabianoflorentino/aprendago/internal/chapter"
)

// Route handles legacy --cap=N --topics/--overview flags using the new
// chapter system.
//
// Returns true if the request was handled (caller should not show an error).
// Returns false if there is no --cap flag (not compat's domain).
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
		fmt.Printf("capítulo %d não encontrado\n", capNum)
		return true
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
			fmt.Printf("Use o novo formato: aprendago cap %d <topico>\n", capNum)
			return true
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
