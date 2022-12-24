/*
Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved.
*/

/* Package for subcommandss implementations */
package cmd

import (
	"context"
	"flag"

	"github.com/google/subcommands"
	"webimizer.dev/web/cmd/license"
	"webimizer.dev/web/cmd/run"
	"webimizer.dev/web/cmd/version"
)

func RegisterAndExecute() {
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	license.Register()
	version.Register()
	run.Register()
	flag.Parse()
	ctx := context.Background()
	subcommands.Execute(ctx)
}
