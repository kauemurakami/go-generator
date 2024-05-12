// Generator pattern encapsulates a goroutine call
// and returns a communication channel to be done with this goroutine
package main

import (
	"fmt"
	"time"
)

func main() {
	// usando esta função de uma maneira normal ela me retorna um canal
	// por esse canal que ela retorna pode-se fazer a leitura dos dados

	channel := write("hello world")

	for k := 0; k < 10; k++ {
		fmt.Println(<-channel) // output Hello world x10 e finaliza o programa

	}

}

// func generator que esconde uma rotina muito mais complexa
// abstraindo da função main
func write(text string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("Value received: %s", text)
			time.Sleep(time.Millisecond * 500)
		}
	}()
	return channel
}
