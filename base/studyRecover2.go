package main
import (
    "fmt"
    "time"
)

func main() {
    // 这个 defer 并不会执行
    defer func() {
        // recover() 可以将捕获到的panic信息打印
        if err := recover(); err != nil {
            fmt.Println(err)
        }
    }()

    go func() {
        panic("遇到报错")
    }()
    go func() {
        defer func() {
        // recover() 可以将捕获到的panic信息打印
        if err := recover(); err != nil {
            fmt.Println(err)
        }
    }()

    }()
    
    fmt.Println("正常执行")
    
    time.Sleep(2 * time.Second)
}