package main

import (
	"flag"
	"fmt"
	"os"
	"sync"
	"testing"

	"github.com/maticnetwork/reddit-bakeoff-poc/helper"
	"github.com/maticnetwork/reddit-bakeoff-poc/performance-tests/reddit"
)

// Default network to execute tests
var (
	network = flag.String("network", "development", "Name of network to run tests against")
	once    sync.Once
)

// TestMain will exec each test, one by one
func TestMain(m *testing.M) {

	// setup
	setUp(network)

	// exec test and this returns an exit code to pass to os
	retCode := m.Run()

	// exec tearDown function
	tearDown("Stop test execution")

	// If exit code is distinct of zero,
	// the test will be failed (red)
	os.Exit(retCode)
}

// setUp configurations for network
func setUp(network *string) {
	once.Do(func() {
		flag.Parse()
		fmt.Println("Executing tests", "network", *network)

		// set global default for conf
		helper.InitAndSetConf(network)

		// set reddit conf
		helper.InitAndSetRedditConf(network)
	})
}

// tearDown function
func tearDown(key string) {
	fmt.Println(key)
}

func Benchmark(b *testing.B) {
	b.Run("Reddit", reddit.BenchmarkReddit)
}
