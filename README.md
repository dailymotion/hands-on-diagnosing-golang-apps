# Hands-on Diagnosing Go Applications

Let's play with [Go](https://golang.org/) applications: the goal of this *hands-on* is to learn how to diagnose a misbehaving Go application, with different kind of tools, from generic unix tools, to very specific go tools.

The first part of the hands-on is an explanation of the tools we'll be using:
- [some standard unix tools](unix-tools/README.md)
- [sending signals to a running application](signals/README.md)
- [generating and inspecting core dumps](core-files/README.md)
- [go HTTP handlers](http-handlers/README.md) (collecting profiles/traces)
- [analyzing pprof profiles](pprof-profiles/README.md)
- [analyzing execution traces](traces/README.md)

Then, once you feel comfortable with these tools, you can build or download the sample application that we'll be using to practice:
- build it simply by running `go build` (with go >= 1.10) - it will produce a binary named `hands-on-diagnosing-golang-apps`
- [download a binary from the latest release](https://github.com/dailymotion-leo/hands-on-diagnosing-golang-apps/releases/latest)

Start it (just execute the binary), and hit the endpoint at `:6060` by default.

At this point, the application should have a correct behaviour. The [default endpoint](http://localhost:6060/) should display a `hello world` message correctly. This would be a good time to play with some of the tools we learned about, to see what exactly it means for our application to have a "correct" behaviour.

When you're ready, you can start the exercice:
- Start by the **step 1**, by hitting the [/goto?step=1](http://localhost:6060/goto?step=1) endpoint, that will activate the first "bad" behaviour. Now, you'll need to use the tools we just explored, to identify what is missbehaving in the application. Go, do it. And then come back here. Once it's done (or if you failed to identify the issue), you can read the [solution](step-1/README.md).
- Then, switch to the **step 2**, by hitting the [/goto?step=2](http://localhost:6060/goto?step=2) endpoint, that will activate the second "bad" behaviour. Once you've done, you can read the [solution](step-2/README.md).
- Then, switch to the **step 3**, by hitting the [/goto?step=3](http://localhost:6060/goto?step=3) endpoint, that will activate the third "bad" behaviour. Once you've done, you can read the [solution](step-3/README.md).
- Then, switch to the **step 4**, by hitting the [/goto?step=4](http://localhost:6060/goto?step=4) endpoint, that will activate the fourth "bad" behaviour. Once you've done, you can read the [solution](step-4/README.md).
- Then, switch to the **step 5**, by hitting the [/goto?step=5](http://localhost:6060/goto?step=5) endpoint, that will activate the fifth "bad" behaviour. Once you've done, you can read the [solution](step-5/README.md).
- Then, switch to the **step 6**, by hitting the [/goto?step=6](http://localhost:6060/goto?step=6) endpoint, that will activate the sixth "bad" behaviour. Once you've done, you can read the [solution](step-6/README.md).

