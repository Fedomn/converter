package main

import (
	"bufio"
	. "fedomn/converter/dispatcher"
	"fedomn/converter/util"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func parseArgsAndFile() {
	inputFile := ""
	flag.StringVar(&inputFile, "c", "", "-c config file name")
	flag.Parse()

	if inputFile != "" {
		inputIfs := util.ParseInputFile(inputFile)
		for _, each := range inputIfs {
			DefaultDispatcher.AddJob(JobInput{GalaxyJob, each})
		}
	}
}

func outputConsole() {
	go func() {
		galaxyOutputQueue := DefaultDispatcher.AcquireOutput(GalaxyJob)
		for output := range galaxyOutputQueue {
			if output.Output == "" {
				fmt.Printf("Input Command: %s\n\n", output.Input)
			} else {
				fmt.Printf("Input Question: %s\n", output.Input)
				fmt.Printf("Answer: %s \n\n", output.Output)
			}
		}
	}()
}

func makeExitSig() <-chan os.Signal {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	return sig
}

func makeConsoleInputChan() <-chan string {
	inputChan := make(chan string)
	go func() {
		f := bufio.NewScanner(os.Stdin)
		for f.Scan() {
			if f.Text() != "" {
				inputChan <- f.Text()
			}
		}
	}()
	return inputChan
}

func main() {
	DefaultDispatcher.Start()
	outputConsole()

	parseArgsAndFile()

	inputChan := makeConsoleInputChan()
	exitSig := makeExitSig()

	for {
		select {
		case input := <-inputChan:
			DefaultDispatcher.AddJob(JobInput{GalaxyJob, input})
		case <-exitSig:
			fmt.Println("converter exit: ByeBye!")
			os.Exit(0)
		}
	}
}
