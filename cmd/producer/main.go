package main
import (
    "context"
    "fmt"
    "time"
    "github.com/segmentio/kafka-go"
)
func main(){
    w := &kafka.Writer{
        Addr: kafka.TCP("localhost:9092"),
        Topic: "events",
        Balancer: &kafka.LeastBytes{},
    }
    defer w.Close()
    for i:=0;i<10;i++{
        err := w.WriteMessages(context.Background(), kafka.Message{Value: []byte(fmt.Sprintf("event-%d", i))})
        if err!=nil { panic(err) }
        time.Sleep(200 * time.Millisecond)
    }
    fmt.Println("Produced 10 messages")
}
