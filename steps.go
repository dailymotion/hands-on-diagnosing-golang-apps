package main

import (
	"bytes"
	"context"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"runtime"
	"sync"
	"time"
)

type Step int

const (
	NoStep                 = Step(0)
	MutexStep              = Step(1)
	CPUStep                = Step(2)
	MemoryLeakStep         = Step(3)
	MemoryAllocStep        = Step(4)
	GoroutinesLeakStep     = Step(5)
	FileDescriptorLeakStep = Step(6)
)

var (
	currentStep Step
	mutex       sync.Mutex           // for the "mutex" step
	buffer      bytes.Buffer         // for the "memory leak" step
	cancelFuncs []context.CancelFunc // for the "goroutines leak" step
	tmpDirPath  string               // for the "file descriptor leak" step
	closers     []io.Closer          // for the "file descriptor leak" step
)

func switchToStep(step Step) error {
	switch step {
	case NoStep, MutexStep, CPUStep, MemoryLeakStep, MemoryAllocStep, GoroutinesLeakStep, FileDescriptorLeakStep:
	default:
		return fmt.Errorf("Invalid step %d", step)
	}

	switch currentStep {
	case NoStep:
	case MutexStep:
		mutex.Unlock()
	case CPUStep:
	case MemoryLeakStep:
		buffer = bytes.Buffer{}
		runtime.GC()
	case MemoryAllocStep:
	case GoroutinesLeakStep:
		shutdownGoroutines()
	case FileDescriptorLeakStep:
		closeFiles()
		os.RemoveAll(tmpDirPath)
	}

	switch step {
	case NoStep:
	case MutexStep:
		mutex.Lock()
	case CPUStep:
	case MemoryLeakStep:
	case MemoryAllocStep:
	case GoroutinesLeakStep:
	case FileDescriptorLeakStep:
		tmpDirPath, _ = ioutil.TempDir(os.TempDir(), "go-app-dir-")
		if len(tmpDirPath) > 0 {
			os.MkdirAll(tmpDirPath, os.ModePerm)
		}
	}

	currentStep = step
	return nil
}

func doSomeBusinessLogic() {
	switch currentStep {

	case MutexStep:
		mutex.Lock()
		defer mutex.Unlock()

	case CPUStep:
		eatCPU(5 * time.Second)

	case MemoryLeakStep:
		leakMemory(10000000)

	case MemoryAllocStep:
		allocateMemory(10000000)

	case GoroutinesLeakStep:
		startGoroutines(5)

	case FileDescriptorLeakStep:
		createFiles(5)
	}
}

func eatCPU(duration time.Duration) {
	for start := time.Now(); time.Since(start) < duration; {
	}
}

func leakMemory(size int64) {
	io.CopyN(&buffer, rand.Reader, size)
}

func allocateMemory(size int64) {
	{
		buf := bytes.NewBuffer(make([]byte, 0, size))
		io.CopyN(buf, rand.Reader, size)
	}
	runtime.GC()
}

func startGoroutines(number int) {
	for i := 0; i < number; i++ {
		ctx, cancelFunc := context.WithCancel(context.Background())
		go func(ctx context.Context) {
			<-ctx.Done()
		}(ctx)
		cancelFuncs = append(cancelFuncs, cancelFunc)
	}
}

func shutdownGoroutines() {
	for _, cancelFunc := range cancelFuncs {
		cancelFunc()
	}
	cancelFuncs = []context.CancelFunc{}
}

func createFiles(number int) {
	for i := 0; i < number; i++ {
		f, err := ioutil.TempFile(tmpDirPath, "go-app-file-")
		if err != nil {
			panic(err)
		}
		closers = append(closers, f)
	}
}

func closeFiles() {
	for _, closer := range closers {
		closer.Close()
	}
}
