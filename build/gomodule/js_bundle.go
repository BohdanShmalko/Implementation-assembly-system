package gomodule

import (
	"fmt"
	"github.com/google/blueprint"
	"github.com/roman-mazur/bood"
	"path"
	"strconv"
)

type jsBundleModule struct {
	blueprint.SimpleName

	properties struct {
		Srcs      []string
		Obfuscate bool
	}
}

func (tb *jsBundleModule) GenerateBuildActions(ctx blueprint.ModuleContext) {
	name := ctx.ModuleName()
	config := bood.ExtractConfig(ctx)
	config.Debug.Printf("Bundle js files to '%s.js'", "js", name)

	outputPath := path.Join(config.BaseOutputDir, name)

	var srcs string
	if len(tb.properties.Srcs) > 0 {
		for _, file := range tb.properties.Srcs {
			srcs += fmt.Sprintf(" %s/%s", ctx.ModuleDir(), file)
		}
	}

	if len(tb.properties.Srcs) > 0 {
		ctx.Build(pctx, blueprint.BuildParams{
			Description: fmt.Sprintf("Bundle js files : %s", tb.properties.Srcs),
			Rule:        jsBundle,
			Outputs:     []string{outputPath},
			Implicits:   tb.properties.Srcs,
			Args: map[string]string{
				"srcs":      srcs,
				"workDir":   ctx.ModuleDir(),
				"output":    name,
				"obfuscate": strconv.FormatBool(tb.properties.Obfuscate),
			},
		})
	}
}

func JsBundleFactory() (blueprint.Module, []interface{}) {
	mType := &jsBundleModule{}
	return mType, []interface{}{&mType.SimpleName.Properties, &mType.properties}
}
