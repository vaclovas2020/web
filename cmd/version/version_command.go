/*
Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved.
*/

/* Version subcommand implementation */
package version

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"
	"webimizer.dev/web"
)

type versionCmd struct {
}

func Register() {
	subcommands.Register(&versionCmd{}, "")
}

func (*versionCmd) Name() string     { return "version" }
func (*versionCmd) Synopsis() string { return "get current version info" }
func (*versionCmd) Usage() string {
	return `version:
	Get current Weblang VM runtime environment version
`
}

func (r *versionCmd) SetFlags(f *flag.FlagSet) {

}

func (r *versionCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	fmt.Printf("Version: %v\n", web.Version)
	return subcommands.ExitSuccess
}
