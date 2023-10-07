package main

func main() {
	ch := make(chan bool)
	close(ch)
}
