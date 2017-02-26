package osusage

import (
    "os"
    "errors"
)

func IsFileExists(path string) (bool, error) {
    stat, err := os.Stat(path)

    if err == nil {
        // os.ModeType defined as dir/named pipe/socket/device/slink
        if stat.Mode() & os.ModeType == 0 {
            return true, nil
        }

        return false, errors.New(path+"exits but not regular file")
    }

    // file not exist
    if os.IsNotExist(err) {
        return false, nil
    }

    // maybe syscall error
    return false, err
}

func ReadFile(path string, b []byte) (int, error) {
    if ok, err := IsFileExists(path); ok && err == nil {
        file, err := os.Open(path)
        if err != nil {
            return 0, err
        }
        defer file.Close()

        return file.Read(b)
    }

    return 0, nil
}
