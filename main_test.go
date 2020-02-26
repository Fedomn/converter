package main_test

import (
	"fmt"
	"testing"

	. "github.com/fedomn/converter/dispatcher"
	. "github.com/fedomn/converter/util"
)

func TestGalaxyGuider(t *testing.T) {
	var baseInfo = []struct {
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
	}

	var tests = map[string]string{
		"how much is pish tegj glob glob ?":                                       "pish tegj glob glob is 42",
		"how many Credits is glob prok Silver ?":                                  "glob prok Silver is 68 Credits",
		"how many Credits is glob prok Gold ?":                                    "glob prok Gold is 57800 Credits",
		"how many Credits is glob prok Iron ?":                                    "glob prok Iron is 782 Credits",
		"how much wood could a woodchuck chuck if a woodchuck could chuck wood ?": "I have no idea what you are talking about",
	}

	DefaultDispatcher.Start()

	for _, info := range baseInfo {
		DefaultDispatcher.AddJob(JobInput{GalaxyJob, info.input})
		// acquire会阻塞 直到有一个job处理完
		DefaultDispatcher.AcquireOutput(GalaxyJob)
	}

	go func() {
		for key := range tests {
			DefaultDispatcher.AddJob(JobInput{GalaxyJob, key})
		}
	}()

	for i := 0; i < len(tests); i++ {
		output := DefaultDispatcher.AcquireOutput(GalaxyJob)
		Equals(t, fmt.Sprintf("input: %s", output.Input), tests[output.Input], output.Output)
		fmt.Printf("\033[32minput\033[0m: %s  \033[32moutput\033[0m: %s\n", output.Input, output.Output)
	}
}
