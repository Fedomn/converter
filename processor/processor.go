package processor

type (
	Processor interface {
		Process(context string) (input, output string)
	}
)
