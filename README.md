## GitHub Stargazer Server

The GitHub Stargazer Server provides an API to easily request how many stars a given GitHub respository has received. 

### Running the Server

Get started by cloning the repo:  
`$ git clone https://github.com/mattcullenmeyer/github-stargazer-server.git`  

Once cloned, you can build and run the server as a docker container.  
Run the following commands in the root directory of the project:  
`$ docker build -t github-server .`  
`$ docker run -it -p 8080:8080 github-server`

Now that the server is running, you can make API requests using the following format:  
`$ http://localhost:8080/api?repo=<organization>/<password>`  
For example:  
`$ http://localhost:8080/api?repo=mattcullenmeyer/github-stargazer-server`

### Running Tests

The Stargazer Server includes unit tests that you can run to ensure the server is working properly.  

Run the tests by entering the following command in the root directory of the project.  
`$ go test`