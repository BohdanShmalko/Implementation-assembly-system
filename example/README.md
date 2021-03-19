Example for `build.bood` :

```
go_testedbinary {
  name: "exampleResult",
  pkg: "github.com/BohdanShmalko/Implementation-assembly-system/example",
  testPkg: "github.com/BohdanShmalko/Implementation-assembly-system/example",
  srcs: ["**/*.go", "../go.mod"]
}

js_bundler {
  name : "result",
  srcs : ["mainScript.js", "script1.js", "script2.js"],
  obfuscate : true
}
```