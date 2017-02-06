package gotest

import (
    "testing"
)
// 执行go test -test.run webbench_test.go -test.bench=".*"
func Benchmark_Division(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Division(4, 5)
    }
}

func Benchmark_TimeConsumingFunction(b *testing.B) {
    b.StopTimer() // 停止压力测试额时间计数

    // 初始化操作
    // 不影响时间

    b.StartTimer()
    for i := 0; i < b.N; i++ {
        Division(4, 5)
    }

}
