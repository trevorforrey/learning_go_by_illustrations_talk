package main

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
