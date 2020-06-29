package crashproof

import (
	"runtime/debug"
	"sync"
)

// ConcurrentCodeCrashCatcherDelegate -
type ConcurrentCodeCrashCatcherDelegate func(err interface{}, packageName string, callStack []StackFrame)

// ConcurrentCodeCrashCatcher -
var ConcurrentCodeCrashCatcher ConcurrentCodeCrashCatcherDelegate

// Go -
func Go(concurrentCode func()) {
	GoWithArgs(func(args ...interface{}) {
		concurrentCode()
	}, nil)
}

// GoWithArgs -
func GoWithArgs(concurrentCode func(args ...interface{}), args ...interface{}) {
	go func() {
		defer func() {
			debug.SetPanicOnFault(true)
			if err := recover(); err != nil {
				packageName, callStack := GetCallStack(err)
				if ConcurrentCodeCrashCatcher != nil {
					ConcurrentCodeCrashCatcher(err, packageName, callStack)
				}
			}
		}()

		if concurrentCode != nil {
			concurrentCode(args...)
		}
	}()
}

// RunAppAndCatchCrashes -
func RunAppAndCatchCrashes(appCode func()) {
	wg := sync.WaitGroup{}

	wg.Add(1)
	Go(func() {
		if appCode != nil {
			appCode()
		}
		wg.Done()
	})

	wg.Wait()
}
