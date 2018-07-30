# Go HTTP handlers

A very easy way to diagnose a running Go applications that exposes an HTTP server, is to use some HTTP handlers provided by the Go stdlib:
- [expvar](https://golang.org/pkg/expvar/)
- [pprof](https://golang.org/pkg/net/http/pprof/)

## expvar

[expvar](https://golang.org/pkg/expvar/) exposes metrics - both from the Go runtime and application-related ones - at `/debug/vars`.

So if the Go application you are trying to diagnose registers the `expvar` handler, when hitting the `/debug/vars` endpoint, you will get something like the following JSON:

```
{
    "cmdline": [
        "output/console-api"
    ],
    "memstats": {
        "Alloc": 4180632,
        "BuckHashSys": 1446937,
        "BySize": [
            {
                "Frees": 0,
                "Mallocs": 0,
                "Size": 0
            },
            {
                "Frees": 977,
                "Mallocs": 1478,
                "Size": 8
            },
            {
                "Frees": 17579,
                "Mallocs": 23462,
                "Size": 16
            },
            ...
            {
                "Frees": 1,
                "Mallocs": 3,
                "Size": 18432
            },
            {
                "Frees": 1,
                "Mallocs": 4,
                "Size": 19072
            }
        ],
        "DebugGC": false,
        "EnableGC": true,
        "Frees": 51909,
        "GCCPUFraction": -1.1315822002750068e-09,
        "GCSys": 483328,
        "HeapAlloc": 4180632,
        "HeapIdle": 1073152,
        "HeapInuse": 6397952,
        "HeapObjects": 30323,
        "HeapReleased": 0,
        "HeapSys": 7471104,
        "LastGC": 1532608014139450911,
        "Lookups": 36,
        "MCacheInuse": 13888,
        "MCacheSys": 16384,
        "MSpanInuse": 97736,
        "MSpanSys": 114688,
        "Mallocs": 82232,
        "NextGC": 8318016,
        "NumForcedGC": 0,
        "NumGC": 4,
        "OtherSys": 1483999,
        "PauseEnd": [
            1532603190880047748,
            1532603190914008250,
            1532603190938432997,
            1532608014139450911,
            ...
        ],
        "PauseNs": [
            285293,
            52119,
            101770,
            166289,
            ...
        ],
        "PauseTotalNs": 605471,
        "StackInuse": 917504,
        "StackSys": 917504,
        "Sys": 11933944,
        "TotalAlloc": 7814768
    }
}
```

You should have a look at the [Go runtime documentation](https://golang.org/pkg/runtime/) to understand the meaning of each metric - `HeapAlloc`, `HeapIdle`, `HeapInuse`, ...

## pprof

[pprof](https://golang.org/pkg/net/http/pprof/) exposes different kind of profiles under `/debug/pprof`:
- **CPU profile** at `/debug/pprof/profile`, for profiling the CPU usage
- **Heap profile** at `/debug/pprof/heap`, for profiling the memory usage
- **goroutine profile** at `/debug/pprof/goroutine`
- **goroutine blocking profile** at `/debug/pprof/block`

There are also a few other endpoints:
- collect an **execution trace** at `/debug/pprof/trace`
- show the **command line** at `/debug/pprof/cmdline`

### Goroutines list

If you hit `/debug/pprof/goroutines?debug=2`, you will get something like:

```
goroutine 165 [running]:
runtime/pprof.writeGoroutineStacks(0x215c600, 0xc4202e61c0, 0x1012779, 0xc4205ba150)
	/usr/local/opt/go/libexec/src/runtime/pprof/pprof.go:650 +0xa7
runtime/pprof.writeGoroutine(0x215c600, 0xc4202e61c0, 0x2, 0xc4204b4000, 0x215c3a0)
	/usr/local/opt/go/libexec/src/runtime/pprof/pprof.go:639 +0x44
runtime/pprof.(*Profile).WriteTo(0x28aa020, 0x215c600, 0xc4202e61c0, 0x2, 0xc4202e61c0, 0xc420216c00)
	/usr/local/opt/go/libexec/src/runtime/pprof/pprof.go:310 +0x3e4
net/http/pprof.handler.ServeHTTP(0x1fb9962, 0x9, 0x21687e0, 0xc4202e61c0, 0xc420216d00)
	/usr/local/opt/go/libexec/src/net/http/pprof/pprof.go:243 +0x20d
github.com/dailymotion-leo/Console-API/vendor/github.com/gorilla/mux.(*Router).ServeHTTP(0xc420246070, 0x21687e0, 0xc4202e61c0, 0xc420216d00)
	/Users/v.behar/go/src/github.com/dailymotion-leo/Console-API/vendor/github.com/gorilla/mux/mux.go:162 +0xed
net/http.serverHandler.ServeHTTP(0xc4201f2750, 0x21687e0, 0xc4202e61c0, 0xc420216b00)
	/usr/local/opt/go/libexec/src/net/http/server.go:2694 +0xbc
net/http.(*conn).serve(0xc42067e1e0, 0x216a2a0, 0xc420151900)
	/usr/local/opt/go/libexec/src/net/http/server.go:1830 +0x651
created by net/http.(*Server).Serve
	/usr/local/opt/go/libexec/src/net/http/server.go:2795 +0x27b

goroutine 1 [chan receive, 118 minutes]:
main.main()
	/Users/v.behar/go/src/github.com/dailymotion-leo/Console-API/main.go:236 +0x164a

goroutine 5 [syscall, 118 minutes]:
os/signal.signal_recv(0x0)
	/usr/local/opt/go/libexec/src/runtime/sigqueue.go:139 +0xa7
os/signal.loop()
	/usr/local/opt/go/libexec/src/os/signal/signal_unix.go:22 +0x22
created by os/signal.init.0
	/usr/local/opt/go/libexec/src/os/signal/signal_unix.go:28 +0x41

goroutine 7 [semacquire, 118 minutes]:
sync.runtime_notifyListWait(0xc420042990, 0x0)
	/usr/local/opt/go/libexec/src/runtime/sema.go:510 +0x10b
sync.(*Cond).Wait(0xc420042980)
	/usr/local/opt/go/libexec/src/sync/cond.go:56 +0x80
github.com/dailymotion-leo/Console-API/vendor/github.com/cihub/seelog.(*asyncLoopLogger).processItem(0xc420146120, 0x0)
	/Users/v.behar/go/src/github.com/dailymotion-leo/Console-API/vendor/github.com/cihub/seelog/behavior_asynclooplogger.go:50 +0x91
github.com/dailymotion-leo/Console-API/vendor/github.com/cihub/seelog.(*asyncLoopLogger).processQueue(0xc420146120)
	/Users/v.behar/go/src/github.com/dailymotion-leo/Console-API/vendor/github.com/cihub/seelog/behavior_asynclooplogger.go:63 +0x44
created by github.com/dailymotion-leo/Console-API/vendor/github.com/cihub/seelog.NewAsyncLoopLogger
	/Users/v.behar/go/src/github.com/dailymotion-leo/Console-API/vendor/github.com/cihub/seelog/behavior_asynclooplogger.go:40 +0x9d
```

### Collecting a Heap profile

Collecting a heap profile allows to understand what is holding on heap memory. It's just a matter of hitting `/debug/pprof/heap`, and retrieve the binary profile file returned in response. We'll see in the next section how to [analyze a profile](../pprof-profiles/README.md).

### Collecting a Goroutine profile

Collecting a goroutine profile allows to understand what are doing the goroutines. It's just a matter of hitting `/debug/pprof/goroutine`, and retrieve the binary profile file returned in response. We'll see in the next section how to [analyze a profile](../pprof-profiles/README.md).

### Collecting a CPU profile

Collecting a CPU profile is different from collecting a Heap/Goroutine profile, because you need to ask for a specific duration (in seconds). For example, to collect a 30 seconds CPU profile, you need to request `/debug/pprof/profile?seconds=30`. You will get a binary profile file in response, after 30 seconds.

Note that you need to make sure that the `http.Server`'s `WriteTimeout` is big enough to avoid timeout-ing while generating the profile data.

## Collecting an execution trace

An execution trace is different from a CPU profile, because it requires a different tool to be analyzed - see the [analyzing traces](../traces/README.md) section.
But collecting an execution trace is very similar to collecting a CPU profile: you need to hit a specific endpoint with a duration for the trace, for example `/debug/pprof/trace?seconds=30` for a 30 seconds trace.

Same as for the CPU profile, you need to make sure that the `http.Server`'s `WriteTimeout` is big enough to avoid timeout-ing while generating the trace data.

## Next

You can now head over to the next section, on [analyzing pprof profiles](../pprof-profiles/README.md).
