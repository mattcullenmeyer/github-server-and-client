## GitHub Stargazer Server

The GitHub Stargazer Server provides an API to easily request how many stars a given GitHub respository has received. 

Get started by cloning the repo:
`git clone https://github.com/mattcullenmeyer/github-server-and-client.git`

Once cloned, you can build and run the server as a docker container. Run the following commands in the root directory of the project:
`docker build -t github-server .`
`docker run -it -p 8080:8080 github-server`

You can now make requests to the server in the form of `http://localhost:8080/api?repo=<username>/<password>`. See example below.
`http://localhost:8080/api?repo=mattcullenmeyer/github-server-and-client`