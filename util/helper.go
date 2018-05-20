package util

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"reflect"
	"runtime"
	"syscall"
	"testing"
)

func Assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d: "+msg+"\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}

func Equals(tb testing.TB, msg string, wat, got interface{}) {
	if !reflect.DeepEqual(wat, got) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d: %s \n\n\twat: %#v\n\n\tgot: %#v\n\n", filepath.Base(file), line, msg, wat, got)
		tb.FailNow()
	}
}

func Ok(tb testing.TB, msg string, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d: %s \n\n unexpected error: %s\n\n", filepath.Base(file), line, msg, err.Error())
		tb.FailNow()
	}
}

type validator func(interface{}) bool

func Contains(source interface{}, find interface{}) bool {
	sourceVal := reflect.ValueOf(source)
	if sourceVal.Type().Kind() != reflect.Slice {
		return false
	}

	for i := 0; i < sourceVal.Len(); i++ {
		val := sourceVal.Index(i).Interface()
		if val == find {
			return true
		}
	}
	return false
}

func ContainsBy(source interface{}, fn validator) bool {
	sourceVal := reflect.ValueOf(source)
	if sourceVal.Type().Kind() != reflect.Slice {
		return false
	}

	for i := 0; i < sourceVal.Len(); i++ {
		val := sourceVal.Index(i).Interface()
		if match := fn(val); match == true {
			return true
		}
	}
	return false
}

func FindBy(source interface{}, fn validator) (interface{}, error) {
	sliceVal := reflect.ValueOf(source)
	if sliceVal.Type().Kind() != reflect.Slice {
		return nil, fmt.Errorf("parameter 1 must be a slice")
	}

	for i := 0; i < sliceVal.Len(); i++ {
		val := sliceVal.Index(i).Interface()
		if match := fn(val); match == true {
			return val, nil
		}
	}
	return nil, nil
}

func GracefulExit(signals ...os.Signal) {
	defer fmt.Println("converter exit: ByeBye!")
	sig := make(chan os.Signal, 1)
	if len(signals) == 0 {
		signals = append(signals, os.Interrupt, syscall.SIGTERM)
	}
	signal.Notify(sig, signals...)
	<-sig
}

func ParseInputFile(filePath string) []string {
	ifs := make([]string, 0)
	fd, err := os.Open(filePath)
	if err != nil {
		fmt.Println("input file not exist")
		os.Exit(1)
	}
	scanner := bufio.NewScanner(fd)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		if scanner.Text() != "" {
			ifs = append(ifs, scanner.Text())
		}
	}
	return ifs
}
