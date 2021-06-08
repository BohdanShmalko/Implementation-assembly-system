package gomodule

import (
	"bytes"
	"fmt"
	"github.com/google/blueprint"
	"github.com/roman-mazur/bood"
	"strings"
	"testing"
)

func Test_JsBundle(t *testing.T) {
	ctx := blueprint.NewContext()

	ctx.MockFileSystem(map[string][]byte{
		"Blueprints": []byte(`
			js_bundler {
  				name : "smthJsResult",
  				srcs : ["smth_js1.js", "smth_js2.js"],
  				obfuscate : true
			}
		`),
		"smth_js1.js": nil,
		"smth_js2.js": nil,
	})

	ctx.RegisterModuleType("js_bundler", JsBundleFactory)

	config := bood.NewConfig()

	_, err := ctx.ParseBlueprintsFiles(".", config)
	if len(err) != 0 {
		t.Fatalf("ParseBlueprintsFiles error : %s", err)
	}

	_, err = ctx.PrepareBuildActions(config)
	if len(err) != 0 {
		t.Errorf("PrepareBuildActions error : %s", err)
	}

	ninjaContentBuffer := new(bytes.Buffer)
	if err := ctx.WriteBuildFile(ninjaContentBuffer); err != nil {
		t.Errorf("WriteBuildFile error : %s", err)
	} else {
		ninjaContent := ninjaContentBuffer.String()
		pcgName := "out/smthJsResult"
		//if runtime.GOOS == "windows" {
		//	pcgName += ".exe"
		//}
		if !strings.Contains(ninjaContent, fmt.Sprintf("%s:", pcgName)) {
			t.Errorf("out/smthJsResult does not exist")
		}
		if !strings.Contains(ninjaContent, "smth_js1.js") {
			t.Errorf("smth_js1.js does not exist")
		}
	}
}
