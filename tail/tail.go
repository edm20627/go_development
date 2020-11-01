package main

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"reflect"
	"syscall"
)

func main() {
	if err := _main(); err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
}

func _main() error {
	if len(os.Args) < 2 {
		return errors.New("prog [file1 file2 ...]")
	}

	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	channels, err := makeChannelsForFiles(os.Args[1:])
	if err != nil {
		return err
	}

	cases, err := makeSelectCases(channels...)
	if err != nil {
		return err
	}

	go doSelect(cases)

	select {
	case <-sigch:
		return nil
	}

	return nil
}

func makeChannelsForFiles(files []string) ([]reflect.Value, error) {
	cs := make([]reflect.Value, len(files))
	for i, fn := range files {
		ch := make(chan []byte)

		f, err := os.Open(fn)
		if err != nil {
			return nil, err
		}
		go readFormFile(ch, f)
		cs[i] = reflect.ValueOf(ch)
	}
	return cs, nil
}

func readFormFile(ch chan []byte, f *os.File) {
	defer close(ch)
	defer f.Close()

	buf := make([]byte, 4096)
	for {
		if n, err := f.Read(buf); err == nil {
			ch <- buf[:n]
		}
	}
}

func makeSelectCases(cs ...reflect.Value) ([]reflect.SelectCase, error) {
	cases := make([]reflect.SelectCase, len(cs))
	for i, ch := range cs {
		if ch.Kind() != reflect.Chan {
			return nil, errors.New("argument must be a channel")
		}

		cases[i] = reflect.SelectCase{
			Chan: ch,
			Dir:  reflect.SelectRecv,
		}
	}

	return cases, nil
}

func doSelect(cases []reflect.SelectCase) {
	for {
		if chosen, recv, ok := reflect.Select(cases); ok {
			fmt.Printf("\n=== %s ===\n%s", os.Args[chosen+1], recv.Interface())
		}
	}
}
