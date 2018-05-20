package main_test

import (
	. "fedomn/converter/dispatcher"
	. "fedomn/converter/util"
	"fmt"
	"testing"
)

func TestGalaxyGuider(t *testing.T) {
	var tests = []struct {
		input  string
		output string
	}{
		{"glob is I", ""},
		{"prok is V", ""},
		{"pish is X", ""},
		{"tegj is L", ""},
		{"glob glob Silver is 34 Credits", ""},
		{"glob prok Gold is 57800 Credits", ""},
		{"pish pish Iron is 3910 Credits", ""},
		{"how much is pish tegj glob glob ?", "pish tegj glob glob is 42"},
		{"how many Credits is glob prok Silver ?", "glob prok Silver is 68 Credits"},
		{"how many Credits is glob prok Gold ?", "glob prok Gold is 57800 Credits"},
		{"how many Credits is glob prok Iron ?", "glob prok Iron is 782 Credits"},
		{"how much wood could a woodchuck chuck if a woodchuck could chuck wood ?", "I have no idea what you are talking about"},
	}

	DefaultDispatcher.Start()

	go func() {
		for _, tt := range tests {
			DefaultDispatcher.AddJob(JobInput{GalaxyJob, tt.input})
		}
	}()

	outputQueue := DefaultDispatcher.AcquireOutput(GalaxyJob)
	for _, tt := range tests {
		msg := fmt.Sprintf("input : %s", tt.input)
		output := <-outputQueue
		Equals(t, msg, tt.input, output.Input)
		Equals(t, msg, tt.output, output.Output)
	}
}
