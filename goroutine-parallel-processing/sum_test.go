package main

import (
	"testing"
)

func BenchmarkSum(b *testing.B) {
	arr := generateRandomArray()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Sum(arr)
	}
}

func BenchmarkSumTwoParallel(b *testing.B) {
	arr := generateRandomArray()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SumTwoParallel(arr)
	}
}

func BenchmarkSumMaxParallel(b *testing.B) {
	arr := generateRandomArray()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SumMaxParallel(arr)
	}
}

// 执行结果：
/*

F:\airdroid_code\go\src\go_learn_demo\goroutine-parallel-processing>go test -bench=^BenchmarkSum$
testing: warning: no tests to run
BenchmarkSum-4              2000            669835 ns/op
PASS
ok      go_learn_demo/goroutine-parallel-processing     1.991s

F:\airdroid_code\go\src\go_learn_demo\goroutine-parallel-processing>go test -bench=^BenchmarkSumTwoParallel$
testing: warning: no tests to run
BenchmarkSumTwoParallel-4           5000            421753 ns/op
PASS
ok      go_learn_demo/goroutine-parallel-processing     2.645s

F:\airdroid_code\go\src\go_learn_demo\goroutine-parallel-processing>go test -bench=^BenchmarkSumMaxParallel$
testing: warning: no tests to run
BenchmarkSumMaxParallel-4           3000            361045 ns/op
PASS
ok      go_learn_demo/goroutine-parallel-processing     1.677s
*/


