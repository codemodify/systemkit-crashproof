package crashproof

import (
	"encoding/json"
	"runtime"
	"strings"
)

// StackFrame -
type StackFrame struct {
	File string `json:"file"`
	Line int    `json:"line"`
	Func string `json:"function"`
}

// String - `stringer` interface
func (thisRef StackFrame) String() string {
	data, _ := json.Marshal(thisRef)
	return string(data)
}

// GetCallStack -
func GetCallStack(e interface{}) (string, []StackFrame) {
	const depth = 32
	const skip = 5
	var pcs [depth]uintptr

	callStackSize := runtime.Callers(skip, pcs[:])
	callStackFrames := runtime.CallersFrames(pcs[:callStackSize])

	firstPackageName := ""
	stackFrames := []StackFrame{}

	for {
		frame, ok := callStackFrames.Next()
		if !ok {
			break
		}

		pkg, fn := splitPackageFuncName(frame.Function)
		if firstPackageName == "" && pkg != "runtime" {
			firstPackageName = pkg
		}

		if stackFilter(pkg, fn, frame.File, frame.Line) {
			stackFrames = stackFrames[:0]
			continue
		}

		stackFrames = append(stackFrames, StackFrame{
			File: frame.File,
			Line: frame.Line,
			Func: fn,
		})
	}

	return firstPackageName, stackFrames
}

func splitPackageFuncName(funcName string) (string, string) {
	var packageName string
	if ind := strings.LastIndex(funcName, "/"); ind > 0 {
		packageName += funcName[:ind+1]
		funcName = funcName[ind+1:]
	}
	if ind := strings.Index(funcName, "."); ind > 0 {
		packageName += funcName[:ind]
		funcName = funcName[ind+1:]
	}
	return packageName, funcName
}

func stackFilter(packageName, funcName string, file string, line int) bool {
	return packageName == "runtime" && funcName == "panic"
}
