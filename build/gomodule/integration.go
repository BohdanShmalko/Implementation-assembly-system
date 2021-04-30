package gomodule

import (
	"fmt"
	"github.com/google/blueprint"
	"github.com/roman-mazur/bood"
	"path"
	"regexp"
)

type integrationModule struct {
	blueprint.SimpleName

	properties struct {
		Srcs        []string
		SrcsExclude []string
		TestPkg     string
		VendorFirst bool
	}
}

func (tb *integrationModule) GenerateBuildActions(ctx blueprint.ModuleContext) {
	name := ctx.ModuleName()
	config := bood.ExtractConfig(ctx)
	testLogPath := path.Join(config.BaseOutputDir, "integrationTestLog.txt")

	var inputs []string
	var testInputs []string
	inputErors := false
	for _, src := range tb.properties.Srcs {
		if matches, err := ctx.GlobWithDeps(src, tb.properties.SrcsExclude); err == nil {
			for _, pathName := range matches {
				if isTest, _ := regexp.Match("^.*_test.go$", []byte(pathName)); isTest {
					testInputs = append(testInputs, pathName)
				} else {
					inputs = append(inputs, pathName)
				}
			}
			inputs = append(inputs, matches...)
		} else {
			ctx.PropertyErrorf("srcs", "Cannot resolve files that match pattern %s", src)
			inputErors = true
		}
	}
	if inputErors {
		return
	}

	if tb.properties.VendorFirst {
		vendorDirPath := path.Join(ctx.ModuleDir(), "vendor")
		ctx.Build(pctx, blueprint.BuildParams{
			Description: fmt.Sprintf("Vendor dependencies of %s", name),
			Rule:        goVendor,
			Outputs:     []string{vendorDirPath},
			Implicits:   []string{path.Join(ctx.ModuleDir(), "go.mod")},
			Optional:    true,
			Args: map[string]string{
				"workDir": ctx.ModuleDir(),
				"name":    name,
			},
		})
		inputs = append(inputs, vendorDirPath)
	}

	if len(tb.properties.TestPkg) > 0 {
		ctx.Build(pctx, blueprint.BuildParams{
			Description: fmt.Sprintf("Test module %s", tb.properties.TestPkg),
			Rule:        goTest,
			Outputs:     []string{testLogPath},
			Implicits:   append(testInputs, inputs...),
			Args: map[string]string{
				"testLogPath": testLogPath,
				"workDir":     ctx.ModuleDir(),
				"testPkg":     tb.properties.TestPkg,
			},
		})
	}
}

func IntegrationFactory() (blueprint.Module, []interface{}) {
	mType := &integrationModule{}
	return mType, []interface{}{&mType.SimpleName.Properties, &mType.properties}
}

type integrationMockModule struct {
	blueprint.SimpleName

	properties struct {
		Srcs        []string
		SrcsExclude []string
		TestPkg     string
		VendorFirst bool
	}
}

func (tb *integrationMockModule) GenerateBuildActions(ctx blueprint.ModuleContext) {}

func IntegrationMockFactory() (blueprint.Module, []interface{}) {
	mType := &integrationMockModule{}
	return mType, []interface{}{&mType.SimpleName.Properties, &mType.properties}
}
