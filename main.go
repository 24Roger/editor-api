package main

func main() {
	mux := Route()
	server := Server(mux)
	server.Run()
}
