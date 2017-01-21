package main

import (
    "fmt"
)

type Human struct {
    name string
    age int
    phone string
}

type Student struct {
    Human
    school string
    loan float32
}

type Employee struct {
    Human
    company string
    money float32
}

type Men interface {
    SayHi()
    Sing(lyrics string)
    Guzzle(beerStein string)
}

type YoungChap interface {
    SayHi()
    Sing(song string)
    BorrowMoney(amount float32)
}

type EderlyGent interface {
    SayHi()
    Sing(song string)
    SpendSalary(amount float32)
}

func (h *Human) SayHi() {
    fmt.Println(h.name, ":", h.phone)
}

func (h *Human) Sing(lyrics string) {
    fmt.Println("lyrics:", lyrics)
}

func (h *Human) Guzzle(beerStein string) {
    fmt.Println("beerStein:", beerStein)
}

// 重载Human SayHi()
func (s *Student) SayHi() {
    fmt.Println("student:", s.name, ":", s.phone)
}

func (s *Student) BorrowMoney(amount float32) {
    s.loan += amount
}

func (e *Employee) SpendSalary(amount float32) {
    e.money -= amount
}

// interface{} 相当于 void*
func main() {
    s := Student{Human{"zzh", 25, "13261558131"}, "SDUT", 1024.0,}
    e := Employee{Human{"zzh", 27, "13261558131"}, "Inte", 4096.0,}

    e.SayHi()
    s.SayHi()

    e.Sing("e la la la")
    s.Sing("s la la la")

    e.Guzzle("eeee")
    s.Guzzle("ssss")

    e.SpendSalary(4096.0)
    s.BorrowMoney(4096.0)
}

