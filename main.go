package main

func main() {

	server := NewServer(":3000")
	server.Handle("/", "GET", HandleRoot)
	server.Handle("/create", "POST", PostRequest)
	server.Handle("/user", "POST", UserPostRequest)
	server.Handle("/api", "GET", server.AddMiddleware(HandleHome, CheckAuth(), Loggin()))
	server.Listen()
}
