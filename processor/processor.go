package processor

type (
	Processor interface {
		Process(context string) string
	}
)
