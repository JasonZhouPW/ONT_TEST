package main

import (
	"github.com/Ontology/common/log"
	. "github.com/ONT_TEST/ont"
	_ "github.com/ONT_TEST/testcase"
	. "github.com/ONT_TEST/testframework"
	"flag"
	log4 "github.com/alecthomas/log4go"
	"strings"
	"time"
)

var (
	ONTJsonRpcAddress string
	ONTWSAddress      string
	CycleTestMode     bool
	CycleTestInterval int
	BenchTestMode     bool
	BenchThreadNum    int
	BenchLastTime     int
)

func init() {
	flag.StringVar(&ONTJsonRpcAddress, "rpc", "http://localhost:20336", "The address of dna jsonrpc")
	flag.StringVar(&ONTWSAddress, "ws", "ws://localhost:20335", "The address of dna web socket")
	flag.BoolVar(&CycleTestMode, "c", false, "Is cycle test mode")
	flag.IntVar(&CycleTestInterval, "ci", 10, "Interval between test in cycle mode")
	flag.BoolVar(&BenchTestMode, "b", false, "Is benchmark test mode")
	flag.IntVar(&BenchThreadNum, "bn", 50, "Thread num in benchmark mode")
	flag.IntVar(&BenchLastTime, "bt", 30, "Last time in benchmark mode")
	flag.Parse()
}

func parseAddress(addresses string) []string {
	return strings.Split(strings.Trim(addresses, ";"), ";")
}

func main() {
	log4.LoadConfiguration("./etc/log4go.xml")
	log.Init()

	ont := NewOntology(parseAddress(ONTJsonRpcAddress), parseAddress(ONTWSAddress))
	ontClient := NewOntClient()
	ontClient.Init()

	TFramework.SetOnt(ont)
	TFramework.SetOntClient(ontClient)
	TFramework.Start(&TestFrameworkOptions{
		CycleTestMode:     CycleTestMode,
		CycleTestInterval: CycleTestInterval,
		BenchTestMode:     BenchTestMode,
		BenchThreadNum:    BenchThreadNum,
		BenchLastTime:     BenchLastTime,
	})

	time.Sleep(time.Second)
}
