package gomodule

import (
	"github.com/google/blueprint"
	"runtime"
)

var (
	pctx = blueprint.NewPackageContext("github.com/BohdanShmalko/Implementation-assembly-system/build/gomodule")

	goBuild, goVendor, goTest blueprint.Rule
)

func setOSRulesTestBinary() {
	var buildCmd, vendorCmd, testCmd string
	os := runtime.GOOS
	switch os {
	case "windows":
		buildCmd = "cmd /c cd $workDir && go build -o $outputPath $pkg"
		vendorCmd = "cmd /c cd $workDir && go mod vendor"
		testCmd = "cmd /c cd $workDir && go test -v $testPkg > $testLogPath"
	case "darwin":
		buildCmd = "cd $workDir && go build -o $outputPath $pkg"
		vendorCmd = "cd $workDir && go mod vendor"
		testCmd = "cd $workDir && go test -v $testPkg > $testLogPath"
	case "linux":
		buildCmd = "cd $workDir && go build -o $outputPath $pkg"
		vendorCmd = "cd $workDir && go mod vendor"
		testCmd = "cd $workDir && go test -v $testPkg > $testLogPath"
	default:
		panic("not compatible with your operating system")
	}

	goBuild = pctx.StaticRule("binaryBuild", blueprint.RuleParams{
		Command:     buildCmd,
		Description: "build go command $pkg",
	}, "workDir", "outputPath", "pkg")

	goVendor = pctx.StaticRule("vendor", blueprint.RuleParams{
		Command:     vendorCmd,
		Description: "vendor dependencies of $name",
	}, "workDir", "name")

	goTest = pctx.StaticRule("test", blueprint.RuleParams{
		Command:     testCmd,
		Description: "test go pkg $testPkg",
	}, "workDir", "testLogPath", "testPkg")
}

func init()  {
	setOSRulesTestBinary()
}
