package hello

const (
	spanish            = "Spanish"
	french             = "French"
	hindi              = "हिंदी"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	hindiHelloPrefix   = "नमस्ते, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case hindi:
		prefix = hindiHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
