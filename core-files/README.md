# Generating and inspecting core dumps

"Core dumps" are files that contains the whole "state" of a running process. They are usually generated when the process "crash", so that the developers can use them to inspect the state and discover the source of the crash.

## Generating core dumps

There are a few prerequisites before trying to generate a core dump:
- make sure the `core file size` limit is big enough
  - you can check the current value with `ulimit -c`
  - and set a value with `ulimit -c unlimited` - might be big enough ;-)
- set the `GOTRACEBACK` environment variable to `crash`
  - see the [signals](../signals/README.md) section for more about the `GOTRACEBACK` env var

There are a few different ways to generate a core dump:
- wait for a crash of the process ;-)
- trigger a core dump (and exit the process) by sending the `SIGABRT` signal to the process
  - `kill -s SIGABRT PID` - see the [signals](../signals/README.md) section for more
- use the `gcore` command to dump a core file from a running process (without killing the process)
  - `sudo gcore PID`

A note about core dumps storage:
- on MacOS, by default, they are stored in the `/cores` directory
- on Linux, by default, they are stored in a location defined in the `/proc/sys/kernel/core_pattern` file - or in the current directory, name `core.PID`

## Inspecting core dumps

Usually we use [GDB](https://www.gnu.org/software/gdb/) to inspect core dumps, and it works for inspecting core dumps generated from Go applications too.

But there is a better solution: [Delve](https://github.com/derekparker/delve). It is a debugger made specifically for Go applications. It can be used to debug live process, as well as inspect core dumps.

First, you need to install it, by following the [installation guide](https://github.com/derekparker/delve/tree/master/Documentation/installation).

Then, use the `dlv core` command to inspect a core file. Note that you will need to give it the path of the executable from which the core file was dumped - see the [delve documentation](https://github.com/derekparker/delve/blob/master/Documentation/usage/dlv_core.md).

Note that there are some issues on Mac, but it works well on Linux.

Another way to use delve, that works on Mac, is to attach it directly to a running process, with the `dlv attach` command:

```
$ dlv attach 85340
Type 'help' for list of commands.
(dlv)
```

Once you are inside a delve session, you can start by printing the list of goroutines:

```
(dlv) goroutines
[36 goroutines]
  Goroutine 1 - User: ./go/src/github.com/dailymotion-leo/Console-API/main.go:236 main.main (0x1c5e31a)
  Goroutine 2 - User: /usr/local/opt/go/libexec/src/runtime/proc.go:292 runtime.gopark (0x102eaea)
  Goroutine 3 - User: /usr/local/opt/go/libexec/src/runtime/proc.go:292 runtime.gopark (0x102eaea)
  Goroutine 4 - User: /usr/local/opt/go/libexec/src/runtime/proc.go:292 runtime.gopark (0x102eaea)
  Goroutine 5 - User: /usr/local/opt/go/libexec/src/runtime/sigqueue.go:139 os/signal.signal_recv (0x10435b7)
  Goroutine 7 - User: /usr/local/opt/go/libexec/src/runtime/sema.go:510 sync.runtime_notifyListWait (0x104051b)
  Goroutine 9 - User: /usr/local/opt/go/libexec/src/runtime/asm_amd64.s:363 runtime.systemstack_switch (0x10589c0)
  Goroutine 10 - User: /usr/local/opt/go/libexec/src/runtime/proc.go:292 runtime.gopark (0x102eaea)
  Goroutine 11 - User: /usr/local/opt/go/libexec/src/runtime/proc.go:292 runtime.gopark (0x102eaea)
  Goroutine 12 - User: /usr/local/opt/go/libexec/src/runtime/proc.go:292 runtime.gopark (0x102eaea)
  Goroutine 18 - User: ./go/src/github.com/dailymotion-leo/Console-API/vendor/github.com/DataDog/dd-trace-go/tracer/tracer.go:362 github.com/dailymotion-leo/Console-API/vendor/github.com/DataDog/dd-trace-go/tracer.(*Tracer).worker (0x154be8f)
  Goroutine 19 - User: /usr/local/opt/go/libexec/src/runtime/asm_amd64.s:363 runtime.systemstack_switch (0x10589c0)
  Goroutine 20 - User: ./go/src/github.com/dailymotion-leo/Console-API/vendor/go.opencensus.io/stats/view/worker.go:146 github.com/dailymotion-leo/Console-API/vendor/go.opencensus.io/stats/view.(*worker).start (0x18c13ad)
  Goroutine 21 - User: /usr/local/opt/go/libexec/src/runtime/proc.go:292 runtime.gopark (0x102eaea)
  Goroutine 22 - User: /usr/local/opt/go/libexec/src/runtime/proc.go:292 runtime.gopark (0x102eaea)
  Goroutine 23 - User: /usr/local/opt/go/libexec/src/runtime/proc.go:292 runtime.gopark (0x102eaea)
  Goroutine 24 - User: /usr/local/opt/go/libexec/src/runtime/proc.go:292 runtime.gopark (0x102eaea)
  Goroutine 26 - User: ./go/src/github.com/dailymotion-leo/Console-API/server/server.go:73 github.com/dailymotion-leo/Console-API/server.RunHTTPServer.func1 (0x1c5c880)
* Goroutine 27 - User: ./go/src/github.com/dailymotion-leo/Console-API/database/views_updater.go:89 github.com/dailymotion-leo/Console-API/database.refreshOnNotification (0x1692e37)
  Goroutine 28 - User: ./go/src/github.com/dailymotion-leo/Console-API/services/webpush.go:172 github.com/dailymotion-leo/Console-API/services.(*liveWebPushService).Run (0x1ab0a6d)
  Goroutine 29 - User: ./go/src/github.com/dailymotion-leo/Console-API/adconfig/exporter.go:107 github.com/dailymotion-leo/Console-API/adconfig.(*AdConfigExporter).Run (0x1ac523a)
  Goroutine 30 - User: /usr/local/opt/go/libexec/src/runtime/netpoll.go:173 internal/poll.runtime_pollWait (0x10291b7)
  Goroutine 31 - User: /usr/local/opt/go/libexec/src/runtime/proc.go:292 runtime.gopark (0x102eaea)
  Goroutine 34 - User: /usr/local/opt/go/libexec/src/runtime/proc.go:292 runtime.gopark (0x102eaea)
  Goroutine 50 - User: ./go/src/github.com/dailymotion-leo/Console-API/vendor/github.com/DataDog/dd-trace-go/tracer/tracer.go:362 github.com/dailymotion-leo/Console-API/vendor/github.com/DataDog/dd-trace-go/tracer.(*Tracer).worker (0x154be8f)
  Goroutine 51 - User: /usr/local/opt/go/libexec/src/runtime/netpoll.go:173 internal/poll.runtime_pollWait (0x10291b7)
  Goroutine 52 - User: /usr/local/opt/go/libexec/src/runtime/asm_amd64.s:363 runtime.systemstack_switch (0x10589c0)
  Goroutine 53 - User: /usr/local/opt/go/libexec/src/database/sql/sql.go:935 database/sql.(*DB).connectionOpener (0x13ab2f9)
  Goroutine 54 - User: /usr/local/opt/go/libexec/src/database/sql/sql.go:948 database/sql.(*DB).connectionResetter (0x13ab45a)
  Goroutine 84 - User: /usr/local/opt/go/libexec/src/runtime/sema.go:510 sync.runtime_notifyListWait (0x104051b)
  Goroutine 85 - User: /usr/local/opt/go/libexec/src/runtime/sema.go:510 sync.runtime_notifyListWait (0x104051b)
  Goroutine 86 - User: /usr/local/opt/go/libexec/src/runtime/sema.go:510 sync.runtime_notifyListWait (0x104051b)
  Goroutine 87 - User: ./go/src/github.com/dailymotion-leo/Console-API/server/server.go:73 github.com/dailymotion-leo/Console-API/server.RunHTTPServer.func1 (0x1c5c880)
  Goroutine 102 - User: ./go/src/github.com/dailymotion-leo/Console-API/vendor/github.com/lib/pq/notify.go:773 github.com/dailymotion-leo/Console-API/vendor/github.com/lib/pq.(*Listener).listenerConnLoop (0x1638575)
  Goroutine 103 - User: /usr/local/opt/go/libexec/src/runtime/netpoll.go:173 internal/poll.runtime_pollWait (0x10291b7)
  Goroutine 106 - User: /usr/local/opt/go/libexec/src/runtime/sema.go:510 sync.runtime_notifyListWait (0x104051b)
```

And then deep into the stack trace of a single goroutine:

```
(dlv) goroutine 27 stack
0  0x000000000102eaea in runtime.gopark
   at /usr/local/opt/go/libexec/src/runtime/proc.go:292
1  0x000000000103e5c0 in runtime.selectgo
   at /usr/local/opt/go/libexec/src/runtime/select.go:392
2  0x0000000001692e37 in github.com/dailymotion-leo/Console-API/database.refreshOnNotification
   at ./go/src/github.com/dailymotion-leo/Console-API/database/views_updater.go:89
3  0x000000000105b551 in runtime.goexit
   at /usr/local/opt/go/libexec/src/runtime/asm_amd64.s:2361
```

And you can print the source code of a specific frame:

```
(dlv) goroutine 27 frame 2 list
Goroutine 27 frame 2 at /Users/v.behar/go/src/github.com/dailymotion-leo/Console-API/database/views_updater.go:89 (PC: 0x1692e37)
    84:			queue.cond.L.Unlock()
    85:			queue.cond.Signal()
    86:		}
    87:
    88:		for {
=>  89:			select {
    90:			case n := <-dl.listener.Notify:
    91:				// ignore nil notifications
    92:				// after a connection loss and re-established a nil notification is sent
    93:				if n == nil {
    94:					break
```

Display the arguments of a function:

```
(dlv) goroutine 27 frame 2 args -v
ctx = context.Context(*context.cancelCtx) *{
	Context: context.Context(*context.emptyCtx) *0,
	mu: sync.Mutex {state: 0, sema: 0},
	done: chan struct {} {
		...}
dl = *github.com/dailymotion-leo/Console-API/database.databaseViewUpdater {
	listener: *github.com/dailymotion-leo/Console-API/vendor/github.com/lib/pq.Listener {
		Notify: chan *github.com/dailymotion-leo/Console-API/vendor/github.com/lib/pq.Notification {
			qcount: 0,
			dataqsiz: 32,
			buf: *[32]*struct github.com/dailymotion-leo/Console-API/vendor/github.com/lib/pq.Notification [
				...
			],
			elemsize: 8,
			closed: 0,
			elemtype: *runtime._type {
				...},
			sendx: 0,
			recvx: 0,
			recvq: waitq<*github.com/dailymotion-leo/Console-API/vendor/github.com/lib/pq.Notification> {
				first: *(*sudog<*github.com/dailymotion-leo/Console-API/vendor/github.com/lib/pq.Notification>)(0xc42041c3c0),
				last: *(*sudog<*github.com/dailymotion-leo/Console-API/vendor/github.com/lib/pq.Notification>)(0xc42041c3c0),},
			sendq: waitq<*github.com/dailymotion-leo/Console-API/vendor/github.com/lib/pq.Notification> {
				first: *sudog<*github.com/dailymotion-leo/Console-API/vendor/github.com/lib/pq.Notification> nil,
				last: *sudog<*github.com/dailymotion-leo/Console-API/vendor/github.com/lib/pq.Notification> nil,},
			lock: runtime.mutex {key: 0},},
		name: "postgres://127.0.0.1:5432/console?sslmode=disable",
		minReconnectInterval: net/http.http2prefaceTimeout,
		maxReconnectInterval: Minute,
		dialer: github.com/dailymotion-leo/Console-API/vendor/github.com/lib/pq.Dialer(github.com/dailymotion-leo/Console-API/vendor/github.com/lib/pq.defaultDialer) *(*"github.com/dailymotion-leo/Console-API/vendor/github.com/lib/pq.Dialer")(0xc4205085d8),
		eventCallback: github.com/dailymotion-leo/Console-API/database.InitDBListener.func1,
		lock: (*sync.Mutex)(0xc4205085f0),
		isClosed: false,
		reconnectCond: *(*sync.Cond)(0xc4201513c0),
		cn: *(*github.com/dailymotion-leo/Console-API/vendor/github.com/lib/pq.ListenerConn)(0xc4201515c0),
		connNotificationChan: <-chan *github.com/dailymotion-leo/Console-API/vendor/github.com/lib/pq.Notification {
			qcount: 0,
			dataqsiz: 32,
			buf: *[32]*struct github.com/dailymotion-leo/Console-API/vendor/github.com/lib/pq.Notification [
				...
			],
			elemsize: 8,
			closed: 0,
			elemtype: *runtime._type {
				...},
			sendx: 0,
			recvx: 0,
			recvq: waitq<*github.com/dailymotion-leo/Console-API/vendor/github.com/lib/pq.Notification> {
				first: *(*sudog<*github.com/dailymotion-leo/Console-API/vendor/github.com/lib/pq.Notification>)(0xc42014e420),
				last: *(*sudog<*github.com/dailymotion-leo/Console-API/vendor/github.com/lib/pq.Notification>)(0xc42014e420),},
			sendq: waitq<*github.com/dailymotion-leo/Console-API/vendor/github.com/lib/pq.Notification> {
				first: *sudog<*github.com/dailymotion-leo/Console-API/vendor/github.com/lib/pq.Notification> nil,
				last: *sudog<*github.com/dailymotion-leo/Console-API/vendor/github.com/lib/pq.Notification> nil,},
			lock: runtime.mutex {key: 0},},
		channels: map[string]struct {} [...],},
	queues: map[string]github.com/dailymotion-leo/Console-API/database.channelQueue [
		"deal_condition": (*github.com/dailymotion-leo/Console-API/database.channelQueue)(0xc420164088),
		"bidder_list": (*github.com/dailymotion-leo/Console-API/database.channelQueue)(0xc4201640b0),
	],}
~r2 = (unreadable invalid interface type: could not find str field)
```

The local variables:

```
(dlv) goroutine 27 frame 2 locals -v
queue = github.com/dailymotion-leo/Console-API/database.channelQueue {
	shouldRefresh: *false,
	viewNameToRefresh: "mv.refresh_bidder_list()",
	cond: *sync.Cond {
		noCopy: sync.noCopy {},
		L: sync.Locker(*sync.Mutex) ...,
		notify: (*sync.notifyList)(0xc420572110),
		checker: 842356171056,},
	isShuttingDown: *false,}
```

and you can print an expression

```
(dlv) goroutine 27 frame 2 p queue.viewNameToRefresh
"mv.refresh_bidder_list()"
```

or its type:

```
(dlv) goroutine 27 frame 2 whatis queue.viewNameToRefresh
string
```

So it's very easy to dive deep into the state of a process.

You can also inspect the package-globals vars:

```
(dlv) vars -v github.com/dailymotion-leo/Console-API/api
github.com/dailymotion-leo/Console-API/api.AppVersion = github.com/dailymotion-leo/Console-API/api.Version {
	Version: "1.5.1-45-0-ga29b008",
	APIVersion: "v1",
	APILevel: 13,
	GoVersion: "go1.10.3",
	GitRefs: []string len: 4, cap: 4, [
		"heads/master",
		"remotes/origin/HEAD",
		"remotes/origin/master",
		"tags/1.5.1-45",
	],
	MainGitRef: "tags/1.5.1-45",
	BuildTime: "2018-07-26T09:17:55Z",
	AdconfigSchema: "11",
	AdconfigSchemas: []string len: 3, cap: 3, ["10","11","9"],}
github.com/dailymotion-leo/Console-API/api.AppInfo = github.com/dailymotion-leo/Console-API/api.Info {
	AppName: "console-api",
	Hostname: "ip-192-168-48-199.us-west-2.compute.internal",
	Environment: "local",
	StartTime: "2018-07-26T11:06:30Z",
	PID: 96539,
	PPID: 412,
	UID: 502,
	EUID: 502,
	GID: 20,
	EGID: 20,
	Groups: []int len: 15, cap: 15, [20,12,61,79,80,81,98,501,701,33,100,204,395,398,399],
	WorkingDirectory: "/Users/v.behar/go/src/github.com/dailymotion-leo/Console-API",}
github.com/dailymotion-leo/Console-API/api/build.GitVersion = "1.5.1-45-0-ga29b008"
github.com/dailymotion-leo/Console-API/api/build.GitRefs = "heads/master,remotes/origin/HEAD,remotes/origin/master,tags/1.5....+5 more"
github.com/dailymotion-leo/Console-API/api/build.BuildTime = "2018-07-26T09:17:55Z"
github.com/dailymotion-leo/Console-API/api.initdone· = 2
github.com/dailymotion-leo/Console-API/api/build.initdone· = 2
```

See the [delve CLI documentation](https://github.com/derekparker/delve/tree/master/Documentation/cli) for the list of commands that can be used inside a delve session.

## Resources

- <https://rakyll.org/coredumps/>
- <https://fntlnz.wtf/post/gopostmortem/>

## Next

You can now head over to the next section, on [go HTTP handlers](../http-handlers/README.md) and collecting profiles/traces.
