package gomodule

import (
	"github.com/google/blueprint"
	"runtime"
)

var (
	pctx = blueprint.NewPackageContext("github.com/BohdanShmalko/Implementation-assembly-system/build/gomodule")

	goBuild, goVendor, goTest, jsBundle blueprint.Rule
)

func setOSRulesTestBinary() {
	var buildCmd, vendorCmd, testCmd, bundleCmd string
	os := runtime.GOOS
	switch os {
	case "windows":
		buildCmd = "cmd /c cd $workDir && go build -o $outputPath $pkg"
		vendorCmd = "cmd /c cd $workDir && go mod vendor"
		testCmd = "cmd /c cd $workDir && go test -v $testPkg > $testLogPath"
		bundleCmd = "cmd /c cd $workDir && npx webpack $srcs -o $workDir/out/js --no-stats && cd $workDir/out/js && (if exist $output.js del $output.js) && rename main.js $output.js && if $obfuscate==true npx javascript-obfuscator $output.js --output $output.js"
	case "darwin":
		buildCmd = "cd $workDir && go build -o $outputPath $pkg"
		vendorCmd = "cd $workDir && go mod vendor"
		testCmd = "cd $workDir && go test -v $testPkg > $testLogPath"
		bundleCmd = "cd $workDir" //TODO
	case "linux":
		buildCmd = "cd $workDir && go build -o $outputPath $pkg"
		vendorCmd = "cd $workDir && go mod vendor"
		testCmd = "cd $workDir && go test -v $testPkg > $testLogPath"
		bundleCmd = "cd $workDir" //TODO
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

	jsBundle = pctx.StaticRule("js_bundle", blueprint.RuleParams{
		Command:     bundleCmd,
		Description: "bundle js files $srcs",
	}, "workDir", "srcs", "output", "obfuscate")
}

func init() {
	setOSRulesTestBinary()
}
