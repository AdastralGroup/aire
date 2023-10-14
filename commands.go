package main

import (
	"github.com/itchio/butler/cmd/apply"
	"github.com/itchio/butler/cmd/verify"
	"github.com/itchio/butler/mansion"
)

// Each of these specify their own arguments and flags in
// their own package.
func registerCommands(ctx *mansion.Context) {
	// documented commands

	verify.Register(ctx)
	apply.Register(ctx)
}
