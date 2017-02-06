package gotest

import (
    "testing"
)
// go test -v

// 文件名必须是_test.go结尾,
// 必须import testing
// 所有测试用例函数必须是Test开头
// TestXxx()参数是testing.T
// 函数通过调用testing.T的Error,Errorf,FailNow,Fatal,FatalIf方法说明测试butongguo
// 通过Log方法记录测试信息
func Test_Division_1(t *testing.T) {
    if i, e :=  Division(6, 2); i != 3 || e != nil {
        t.Error("除法测试没有通过")
    } else {
        t.Log("第一个测试通过了")
    }
}

func Test_Division_2(t *testing.T) {
    t.Error("测试通不过")
}
