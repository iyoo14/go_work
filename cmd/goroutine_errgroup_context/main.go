package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"time"
)

func main() {
	eg, ctx := errgroup.WithContext(context.Background())
	limit := make(chan struct{}, 8)
	for i := 1; i < 10; i++ {
		i := i
		limit <- struct{}{}
		proc(eg, ctx, limit, i)
	}
	if err := eg.Wait(); err != nil {
		log.Fatal(err)
	}
	log.Println("success")
}

func proc(eg *errgroup.Group, ctx context.Context, limit chan struct{}, n int) {
	eg.Go(func() error {
		defer func() {
			<-limit
		}()
		select {
		case <-ctx.Done():
			fmt.Println("Canceled:", n)
			return nil
		default:
			err := doProc(n)
			if err != nil {
				fmt.Println("Error:", n, err)
				return err
			}
			fmt.Println("fin", n)
			return nil
		}

	})

}

func doProc(n int) error {
	time.Sleep(2 * time.Second) // 長い処理
	if n%311 == 0 {
		return errors.New("n bad number")
	}
	return nil
}
