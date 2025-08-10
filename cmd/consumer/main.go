package main
import (
    "context"
    "fmt"
    "github.com/segmentio/kafka-go"
)
func main(){
    r := kafka.NewReader(kafka.ReaderConfig{
        Brokers: []string{"localhost:9092"},
        GroupID: "demo-group",
        Topic: "events",
    })
    defer r.Close()
    for i:=0;i<10;i++{
        m, err := r.ReadMessage(context.Background())
        if err!=nil { panic(err) }
        fmt.Printf("Message: %s\n", string(m.Value))
    }
}
