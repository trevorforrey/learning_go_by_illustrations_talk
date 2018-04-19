package main

// Example code not meant to run
// You can find an executable version of this on the file
// in the program : /exec/8_passInAndReturnChannel.go

func main() {
	baseData := [5]string{"rock", "ore", "ore", "rock", "ore"}
	generatedChannel := gen(baseData)
}

func gen(mine [5]string) <-chan string {
	out := make(chan string)
	go func() {
		// Write to output channel
	}()
	return out
}
