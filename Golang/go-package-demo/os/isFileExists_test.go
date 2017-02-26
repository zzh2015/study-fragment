package osusage

import (
    "testing"
)

func Test_IsFileExists(t *testing.T) {
    if ok, err := IsFileExists("./isFileExists.go"); !ok || err != nil {
        t.Error("testing error")
    } else {
        t.Log("testing success")
    }
}

func Test_ReadFile(t *testing.T) {
    buf := make([]byte, 2048)

    if n, err := ReadFile("./isFileExists.go", buf); n == 0 || err != nil {
        t.Error("read file content error")
    } else {
        t.Log(string(buf[:]))
    }
}
