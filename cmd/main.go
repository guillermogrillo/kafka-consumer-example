package main

import (
	"fmt"

	"github.com/guillermogrillo/kafka-consumer-example/pkg/consumer"
)

func main() {

	consumer.NewTollConsumer("localhost:9092", "default", []string{"toll-records"})
	fmt.Sprintln("stopped consumer")
}
