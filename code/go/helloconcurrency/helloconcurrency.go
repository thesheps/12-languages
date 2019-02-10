package helloconcurrency

// Hello grabs the string Hello World in a pointlessly concurrent way.
func Hello() string {
	hello := ""
	channel := make(chan string)

	for i := 0; i < 12; i++ {
		go getLetter(channel, i)
	}

	for i := 0; i < 12; i++ {
		hello += <- channel
	}

	return hello
}

func getLetter(channel chan string, index int) {
	channel <- string("Hello World!"[index])
}
