package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func init() {
	flag.Parse()
	rand.Seed(time.Now().UnixNano())

	i = 0

	count, err = strconv.Atoi(os.Getenv("SHLOOP_COUNT"))
	if err != nil {
		count = 5
	}

	thread, err = strconv.Atoi(os.Getenv("SHLOOP_THREAD"))
	if err != nil {
		thread = 1
	}

	verbose, err = strconv.ParseBool(os.Getenv("SHLOOP_VERBOSE"))
	if err != nil {
		verbose = false
	}
}

func main() {
	args := make(chan string)

	for t := 0; t < thread; t++ {
		wg.Add(1)
		go func() {
			for cmd := range args {
				run(parse(cmd), verbose)
			}

			defer wg.Done()
		}()
	}

	for c := 0; c < count; c++ {
		args <- strings.Join(flag.Args(), " ")
	}

	close(args)
	wg.Wait()
}

func run(arg string, v bool) {
	if v {
		log.Printf("cmd: %s\n", arg)
	}

	cmd := exec.Command("bash", "-c", arg)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
