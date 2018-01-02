package ontid

import (
. "github.com/ONT_TEST/testframework"
)

func TestONT() {
	//Register Asset test
	TFramework.RegTestCase("TestONTID", TestONTID)
	

	//Benchmark
	//TFramework.RegBenchTestCase("BenchmarkTransaction", BenchmarkTransaction)
}
