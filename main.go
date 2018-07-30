package main

import (
	_ "expvar"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"runtime/pprof"
	"strconv"
	"syscall"
	"time"
)

var (
	listenAddr string
)

func init() {
	flag.StringVar(&listenAddr, "listen", ":6060", "TCP network address on which the HTTP server will listen")
}

func main() {
	flag.Parse()

	log.Printf("Application started with PID %d", os.Getpid())

	http.HandleFunc("/", helloWorldHandler)
	http.HandleFunc("/goto", gotoHandler)
	go func() {
		log.Fatal(http.ListenAndServe(listenAddr, nil))
	}()

	log.Printf("HTTP server started on %s", listenAddr)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
	for {
		sig := <-c
		log.Printf("Received signal %s", sig.String())

		switch sig {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			log.Println("Application stopping...")
			return
		case syscall.SIGUSR1:
			pprof.Lookup("goroutine").WriteTo(os.Stderr, 2)
		case syscall.SIGUSR2:
			f, err := ioutil.TempFile(os.TempDir(), "go-app-cpu-profile-")
			if err != nil {
				log.Printf("failed to create temporary file: %v", err)
				continue
			}
			if err = pprof.StartCPUProfile(f); err != nil {
				log.Printf("failed to create temporary file: %v", err)
				continue
			}
			log.Print("Waiting 30 seconds to collect a CPU profile...")
			time.Sleep(30 * time.Second)
			pprof.StopCPUProfile()
			f.Close()
			log.Printf("Done collecting a 30 seconds CPU profile in %s", f.Name())
		}
	}
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	start := time.Now()
	log.Println("Starting processing request...")
	doSomeBusinessLogic()
	fmt.Fprint(w, "hello world")
	took := time.Since(start)
	log.Printf("Request processed in %s\n", took)
}

func gotoHandler(w http.ResponseWriter, r *http.Request) {
	step, err := strconv.Atoi(r.URL.Query().Get("step"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = switchToStep(Step(step))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "switched to step %d", step)
}
