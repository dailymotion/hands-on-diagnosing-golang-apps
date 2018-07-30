# Signals

A standard way of interacting with a running process is to send "signals" to it. The classic one that everybody knows is the `KILL` signal, that force-kill a process.

You can send a signal using the `kill` command. You can pass it the signal to send with the `-s` option:

```
$ kill -s TERM 123
```

will send the default `TERM` (or `SIGTERM`) signal to the process with PID `123`.

## SIGABRT signal

Here, we touch the first thing that is specific to go applications: by default, the go runtime handles the `SIGABRT` (or `ABRT`) signal by printing the stack trace to stderr and killing the process.

So if you ever find yourself in front of a "stucked" go application that doesn't answer anymore, instead of "just" killing it, you should at least send it the `SIGABRT` signal so that it prints its stack trace. That might help the developers to investigate the issue.

You can control the behaviour of the special `SIGABRT` signal by setting the `GOTRACEBACK` environment variable (before starting the go application, of course) to one of the following value:
- `none`: exit without printing anything (no stack trace)
- `single` (the default): print the stack trace for the current goroutine, or all goroutines if there are no current goroutine (or if the failure is internal to the runtime), and then exit
- `all`: print the stack trace for all user-created goroutines, and then exit
- `system`: print the stack trace for all user-created + system goroutines (including the frames for runtime functions), and then exit
- `crash`: print the stack trace for all user-created + system goroutines (including the frames for runtime functions), and then "crash" - which, on Unix system, will trigger a core dump (more on that later)

So if you set `GOTRACEBACK=crash` before starting the application, it will produce a core dump when receiving the `SIGABRT` signal.

You can read more about it in the Go documentation:
- <https://golang.org/pkg/os/signal/>
- <https://golang.org/pkg/runtime/>

## Custom signals handling

Applications can also implement custom handling of specific signals.

For example, [nginx](http://nginx.org/) handles some signals in a custom way:
- `HUP` (or `SIGHUP`): reload the configuration
- `USR1` (or `SIGUSR1`): re-opening the log files
- ...

There are 2 special signals, that applications can handle as they want:
- `SIGUSR1`
- `SIGUSR2`

For example to dump the stack traces.

## Next

You can now head over to the next section, on [generating and inspecting core dumps](../core-files/README.md).
