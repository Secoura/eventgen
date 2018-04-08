package processor

type Processor interface {
	process(template string) string
}

func ProcessTemplate(p Processor, template string) string {
	return p.process(template)
}
