/*
Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved.
*/

/* Run subcommand implementation */
package license

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/google/subcommands"
)

type licenseCmd struct {
}

func Register() {
	subcommands.Register(&licenseCmd{}, "")
}

func (*licenseCmd) Name() string     { return "license" }
func (*licenseCmd) Synopsis() string { return "Application license" }
func (*licenseCmd) Usage() string {
	return `license
  Read application license text
`
}

func (r *licenseCmd) SetFlags(f *flag.FlagSet) {
}

func (r *licenseCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	text, err := ioutil.ReadFile("LICENSE")
	if err != nil {
		fmt.Println(err)
		return subcommands.ExitFailure
	}
	fmt.Println(string(text))
	return subcommands.ExitSuccess
}
