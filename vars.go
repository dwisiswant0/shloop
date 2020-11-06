package main

import "sync"

var (
	i       int
	wg      sync.WaitGroup
	err     error
	count   int
	thread  int
	verbose bool
)
