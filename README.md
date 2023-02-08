# go-mod-test

This is a test to confirm my understanding of the `replace` directive.

```
+-----+
|  a  |
+-----+
   |     
   v     
+-----+
|  b  |  replace github.com/tingstad/go-mod-test/c => github.com/tingstad/go-mod-test/c v1.0.0
+-----+  require github.com/tingstad/go-mod-test/c v1.2.0
   |     
   v     
+-----+
|  c  |
+-----+
```

```shell
$ (cd b/ && go run main.go)
hello got: version 1
$ (cd a/ && go run main.go)
hello got: version 2
```

The test confirms my hypothesis that dependendies' replace directives are indeed ignored.

[Go Modules Reference](https://go.dev/ref/mod):
> The _main module_ is the module containing the directory where the go command is invoked.

> replace directives only apply in the main module’s go.mod file and are ignored in other modules.

> The content of a module (including its go.mod file) may be replaced using a replace directive in a main module’s go.mod file or a workspace’s go.work file. A replace directive may apply to a specific version of a module or to all versions of a module.
> 
> Replacements change the module graph
