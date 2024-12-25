package main

import (
    "fmt"
    "flag"
    "net"
)

func main() {
    clientMessage := flag.String("message", "test", "Custom Message from User")

    flag.Parse()

    // if able to connect then server is up, run as client
    if !testConnection() {
        fmt.Println("running as client")
        musicClient(*clientMessage)

    } else {
        fmt.Println("running as server")
        musicServer()

    }
}

func musicServer() {
    listener, err := net.Listen("tcp", ":8080")
    if err != nil {
        fmt.Println("Error listening:", err)
        return
    }
    defer listener.Close()

    fmt.Println("Listening on port 8080")

    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Error accepting connection", err)
            continue
        }
        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    // fmt.Println("handleConnection")
    defer conn.Close()

    buf := make([]byte, 1024)
    n, err := conn.Read(buf)

    if err != nil {
        fmt.Println("Error readings:", err)
        return
    }

    message := string(buf[:n])

    fmt.Println("Message = ", message)

    _, err = conn.Write([]byte(message))

    if err != nil {
        fmt.Println("write failed")
        return
    }
}

func musicClient(message string) bool {
    conn, err := net.Dial("tcp", "localhost:8080")

    if err != nil {
        fmt.Println("could not connect to server, ", err)
        return false
    }

    defer conn.Close()

    byteMessage := []byte(message)

    for {
        _, err = conn.Write(byteMessage)
        if err != nil {
            fmt.Println("could not write message to server")
            return false
        }

        buf := make([]byte, 1024)

        n, err := conn.Read(buf)
        if err != nil {
            fmt.Println("could not read from server")
            return false
        }

        fmt.Println("Message from server: ", string(buf[:n]))
    }
}


func testConnection() bool {
    // fmt.Println("testing if server is running")
    listener, err := net.Listen("tcp", ":8080")
    if err != nil {
        return false
    }
    // fmt.Println("connection was able to be established")
    listener.Close()
    return true
}
