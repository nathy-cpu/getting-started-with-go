package helloworld

func greetingPrefix(language string) (prefix string) {
	const (
		spanish = "Spanish"
		french  = "French"

		englishHelloPrefix = "Hello, "
		spanishHelloPrefix = "Hola, "
		frenchHelloPrefix  = "Bonjour, "
	)

	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}
