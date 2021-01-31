package main

import (
	"fmt"
  "net"
  "bufio"
  "os"
  "strings"
)

func main() {
  // Connect to UDP server
  resolver, err := net.ResolveUDPAddr("udp4", "127.0.0.1:34254")
  conn, err := net.DialUDP("udp4", nil, resolver)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Fprintf(conn, "Connect client: %s", conn.LocalAddr().String())
  fmt.Printf("The UDP server is %s", conn.RemoteAddr().String())
  fmt.Println()
  defer conn.Close()

  // Send message
  for {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("> ")
    text, _ := reader.ReadString('\n')
    text = strings.TrimSuffix(text, "\n")
    data := []byte(text)
    _, err := conn.Write(data)
    if err != nil {
      fmt.Println(err)
      return
    }

    // Receive reply
    // buffer := make([]byte, 2048)
    // res, _, err := conn.ReadFromUDP(buffer)
    // if err != nil {
    //   fmt.Println(err)
    //   return
    // }
    // fmt.Printf("Reply: %s\n", string(buffer[0:res]))
  }
}
