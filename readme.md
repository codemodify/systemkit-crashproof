# ![](https://fonts.gstatic.com/s/i/materialiconsoutlined/flare/v4/24px.svg) Crash Proof Go Apps
[![](https://img.shields.io/github/v/release/codemodify/systemkit-crashproof?style=flat-square)](https://github.com/codemodify/systemkit-crashproof/releases/latest)
![](https://img.shields.io/github/languages/code-size/codemodify/systemkit-crashproof?style=flat-square)
![](https://img.shields.io/github/last-commit/codemodify/systemkit-crashproof?style=flat-square)
[![](https://img.shields.io/badge/license-0--license-brightgreen?style=flat-square)](https://github.com/codemodify/TheFreeLicense)

![](https://img.shields.io/github/workflow/status/codemodify/systemkit-crashproof/qa?style=flat-square)
![](https://img.shields.io/github/issues/codemodify/systemkit-crashproof?style=flat-square)
[![](https://goreportcard.com/badge/github.com/codemodify/systemkit-crashproof?style=flat-square)](https://goreportcard.com/report/github.com/codemodify/systemkit-crashproof)

[![](https://img.shields.io/badge/godoc-reference-brightgreen?style=flat-square)](https://godoc.org/github.com/codemodify/systemkit-crashproof)
![](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)
![](https://img.shields.io/gitter/room/codemodify/systemkit-crashproof?style=flat-square)

![](https://img.shields.io/github/contributors/codemodify/systemkit-crashproof?style=flat-square)
![](https://img.shields.io/github/stars/codemodify/systemkit-crashproof?style=flat-square)
![](https://img.shields.io/github/watchers/codemodify/systemkit-crashproof?style=flat-square)
![](https://img.shields.io/github/forks/codemodify/systemkit-crashproof?style=flat-square)

### What it offers
- The one and only crash proofing mechanics for your Go app.
- Catches crashes/panics deep into your concurrent code
- What's the catch: DON'T use `go func(){}()`, use `crashproof.Go(func(){})` as in setup below


### Sample
```go
func main() {
	crashproof.ConcurrentCodeCrashCatcher = reportCrash
	crashproof.RunAppAndCatchCrashes(func() {
		...
		// YOUR CONCURRENT APP HERE
		crashproof.Go(func(){
			panic("OOPS")
		})
		...
	})
}

func reportCrash(err interface{}, packageName string, callStack []crashproof.StackFrame) {
	fmt.Fprintf(os.Stderr, "\n\nCRASH: %v\n\npackage %s\n\nstack: %v\n\n", err, packageName, callStack)
}
```
