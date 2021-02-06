## GitHub Star Server and Client

Creates a server with an API and a client that interacts with that API in order to determine how many stars any given list of github repositories has received.

* `go mod init github.com/mattcullenmeyer/github-server-and-client`
* `docker build -t github-server .`
* `docker run -d -p 8080:8080 github-server`