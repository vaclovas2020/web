/*
Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved.
*/

/* Run subcommand implementation */
package run

import (
	"context"
	"flag"
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/google/subcommands"
	"webimizer.dev/web"
	"webimizer.dev/web/bytecode/class"
)

type runCmd struct {
	configFile  string
	gitUrl      string
	gitUser     string
	gitToken    string
	gitWebHook  string
	gitLocalDir string
}

func Register() {
	subcommands.Register(&runCmd{}, "")
}

func (*runCmd) Name() string     { return "run" }
func (*runCmd) Synopsis() string { return "run weblang application" }
func (*runCmd) Usage() string {
	return `run [-config-file] [-git-url] [-git-user] [-git-token] [-git-webhook] [-git-localdir]
  Run Weblang application in safe VM environment.
`
}

func (r *runCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&r.configFile, "config-file", "weblang.yml", "config file location")
	f.StringVar(&r.gitUrl, "git-url", "", "git repository url")
	f.StringVar(&r.gitUser, "git-user", "", "git username")
	f.StringVar(&r.gitToken, "git-token", "", "git token (password)")
	f.StringVar(&r.gitWebHook, "git-webhook", "", "git webhook url")
	f.StringVar(&r.gitLocalDir, "git-localdir", "", "git local repository dir")
}

func (r *runCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if r.configFile == "" {
		println("Error: please provide correct -config-file flag")
		return subcommands.ExitFailure
	}

	fmt.Println("----------------------")
	fmt.Printf("Welcome to %v %s/%s %s %s (bytecode version %v)\n\n", web.Version, runtime.GOOS, runtime.GOARCH, runtime.Compiler, runtime.Version(), class.ByteCodeVersion)
	fmt.Println("Copyright (c) 2022-2024 Vaclovas Lapinskis. All rights reserved.")
	fmt.Println("License: BSD-3-Clause License")
	fmt.Println("----------------------")

	time.Sleep(time.Second)

	vm := web.VM{}
	vm.InitVM(r.configFile)
	if r.gitUrl != "" && r.gitUser != "" && r.gitToken != "" && r.gitWebHook != "" && r.gitLocalDir != "" {
		vm.GitPreperWebHook(r.gitUrl, r.gitUser, r.gitToken, r.gitLocalDir, r.gitWebHook)
	}
	log.Fatal(vm.StartServer())
	return subcommands.ExitSuccess
}
