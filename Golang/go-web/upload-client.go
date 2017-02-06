package main

import (
    "bytes"
    "fmt"
    "io"
    "io/ioutil"
    "mime/multipart"
    "net/http"
    "os"
)

func postFile(fileName string, targetURL string) error {
    bodyBuf := &bytes.Buffer{}
    bodyWriter := multipart.NewWriter(bodyBuf)

    fileWriter, err := bodyWriter.CreateFormFile("uploadfile", fileName)
    if err != nil {
        fmt.Println("error wriing to buffer")
        return err
    }

    fh, err := os.Open(fileName)
    if err != nil {
        fmt.Println("error opening file")
        return err
    }
    defer fh.Close()

    _, err = io.Copy(fileWriter, fh)
    if err != nil {
        return err
    }

    contentType := bodyWriter.FormDataContentType()
    bodyWriter.Close()

    resp, err := http.Post(targetURL, contentType, bodyBuf)
    if err != nil {
        return err
    }

    defer resp.Body.Close()
    respBody, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return err
    }

    fmt.Println(resp.Status)
    fmt.Println(string(respBody))
    return nil
}

func main() {
    targetURL := "http://localhost:9090/upload"
    fileName := "./hello-web.go"
    postFile(fileName, targetURL)
}
