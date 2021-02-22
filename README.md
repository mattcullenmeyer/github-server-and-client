## GitHub Stargazer Server

The GitHub Stargazer Server provides an API to easily request how many stars a given GitHub repository has received. 

### Getting Started

Get started by cloning the repo:  
`$ git clone https://github.com/mattcullenmeyer/github-stargazer-server.git`  

Once cloned, navigate to the root directory of the project:   
`$ cd github-stargazer-server`

### Running Server (Docker)

You can build and run the server as a docker container.  
Run the following commands in the root directory of the project:  
`$ docker build -t github-server .`  
`$ docker run -it -p 8080:8080 github-server`

Now that the server is running, you can make API requests using the following format:  
`http://localhost:8080/api?repo=<organization>/<password>`  
For example:  
`http://localhost:8080/api?repo=mattcullenmeyer/github-stargazer-server`

### Running Tests

The GitHub Stargazer Server includes unit tests that you can run to ensure the server is working properly.  

Run the tests by entering the following command in the root directory of the project.  
`$ go test`

### Running Server (Kubernetes)

The GitHub Stargazer Server can alterantively run on local Kubernetes with minikube.   

Run the following commands in the root directory of the project:   
`$ minikube start`   
`$ kubectl create -f deployment.yaml`   
`$ kubectl apply -f service.yaml`   

Now that the server is running, you need to determine the correct url path with the following command (be sure to keep the terminal open if you are on Windows).  
`$ minikube service stargazer-server`  

If the returned url is, say, `http://192.168.49.2:30893`, then a valid API request could look something like this:  
`http://192.168.49.2:30893/api?repo=mattcullenmeyer/github-stargazer-server`   