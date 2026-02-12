package main

import (
	"context"
	"fmt"
	"time"
)

func doSomething(ctx context.Context) {
	ctx, cancelCtx := context.WithCancel(ctx)
	printCh := make(chan int)
	go doAnother(ctx, printCh)
	time.Sleep(1000 * time.Millisecond)
	for num := 1; num <= 3; num++ {
		printCh <- num
	}
	defer cancelCtx()
	time.Sleep(1000 * time.Millisecond)
	fmt.Printf("doSomething: finished\n")
}

func doAnother(ctx context.Context, printCh <-chan int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		case num := <-printCh:
			fmt.Printf("Num: %d\n", num)
		}
	}
}

func main() {
	ctx := context.Background()
	doSomething(ctx)
}
