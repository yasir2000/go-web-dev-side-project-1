package main

import (
	"context"
	"errors"
	"sync"
	"time"
)

var sendError = false
var wg sync.WaitGroup

/**
 * Branch A
 */
func A1(ctx context.Context) error {
	if context.Canceled == ctx.Err() {
		println(`    -> A1 - cancelled`)

		return nil
	}

	time.Sleep(time.Millisecond * 100)
	println(`    -> A1 - 100ms`)

	subCtx, cancelFunc := context.WithCancel(ctx)
	c := make(chan error, 2)

	go func() { c <- A11(subCtx) }()
	go func() { c <- A12(subCtx) }()

	if err := <-c; err != nil {
		cancelFunc()
		<-c
		return err
	}

	err := <-c
	cancelFunc()

	return err
}

func A11(ctx context.Context) error {
	if context.Canceled == ctx.Err() {
		println(`        -> A11 - cancelled`)

		return nil
	}

	time.Sleep(time.Millisecond * 50)

	if sendError {
		println(`        -> A11 - error`)

		return errors.New(`error`)
	}

	println(`        -> A11 - 50ms`)

	return nil
}

func A12(ctx context.Context) error {
	for i := 0; i < 3; i++ {
		if context.Canceled == ctx.Err() {
			println(`        -> A12 - cancelled`)

			return nil
		}

		time.Sleep(time.Millisecond * 100)
	}
	println(`        -> A12 - 300ms`)

	return nil
}

/**
 * Branch B
 */
func B1(ctx context.Context) error {
	if context.Canceled == ctx.Err() {
		println(`    -> B1 cancelled`)

		return nil
	}

	time.Sleep(time.Millisecond * 100)
	println(`    -> B1 - 100ms`)

	return nil
}
func B2(ctx context.Context) error {
	if context.Canceled == ctx.Err() {
		println(`    -> B2 - cancelled`)
		println(`    -> B21 - cancelled`)

		return nil
	}

	time.Sleep(time.Millisecond * 100)
	println(`    -> B2 - 100ms`)

	return B21(ctx)
}
func B21(ctx context.Context) error {
	if context.Canceled == ctx.Err() {
		println(`         -> B21 - cancelled`)

		return nil
	}

	time.Sleep(time.Millisecond * 150)
	println(`         -> B21 - 150ms`)

	return nil
}

func main() {
	wg.Add(2)

	ctx, cancelFunc := context.WithCancel(context.Background())

	go func() {
		defer wg.Done()

		time.Sleep(time.Millisecond * 100)
		println(`A - 100ms`)

		if err := A1(ctx); err != nil {
			cancelFunc()
		}
	}()

	go func() {
		defer wg.Done()

		time.Sleep(time.Millisecond * 200)
		println(`B - 200ms`)

		if err := B1(ctx); err != nil {
			cancelFunc()
		}
		if err := B2(ctx); err != nil {
			cancelFunc()
		}
	}()

	wg.Wait()
}
