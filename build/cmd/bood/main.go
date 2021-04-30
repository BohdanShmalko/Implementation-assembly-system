package main

import (
	"flag"
	"github.com/BohdanShmalko/Implementation-assembly-system/build/gomodule"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/google/blueprint"
	"github.com/roman-mazur/bood"
)

var (
	dryRun           = flag.Bool("dry-run", false, "Generate ninja build file but don't start the build")
	verbose          = flag.Bool("v", false, "Display debugging logs")
	integrationTests = flag.Bool("integration-tests", false, "select one module to build")
)

func NewContext() *blueprint.Context {
	ctx := bood.PrepareContext()
	if *integrationTests {
		ctx.RegisterModuleType("integration_tests", gomodule.IntegrationFactory)
		ctx.RegisterModuleType("go_testedbinary", gomodule.TestBinMockFactory)
		ctx.RegisterModuleType("js_bundler", gomodule.JsBundleMockFactory)
	} else {
		ctx.RegisterModuleType("integration_tests", gomodule.IntegrationMockFactory)
		ctx.RegisterModuleType("go_testedbinary", gomodule.TestBinFactory)
		ctx.RegisterModuleType("js_bundler", gomodule.JsBundleFactory)
	}
	return ctx
}

func main() {
	flag.Parse()

	config := bood.NewConfig()
	if !*verbose {
		config.Debug = log.New(ioutil.Discard, "", 0)
	}
	ctx := NewContext()

	ninjaBuildPath := bood.GenerateBuildFile(config, ctx)

	if !*dryRun {
		config.Info.Println("Starting the build now")

		cmd := exec.Command("ninja", append([]string{"-f", ninjaBuildPath}, flag.Args()...)...)
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			config.Info.Fatal("Error invoking ninja build. See logs above.")
		}
	}
}
