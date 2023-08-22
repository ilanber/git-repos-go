# GitHub Repository Fetcher

This Go program sets up a web server to fetch and display GitHub repositories from a specified organization.

## Getting Started

1. Install Go: [Install Go](https://golang.org/doc/install)
2. Clone this repository: `git clone https://github.com/your-username/github-repo-fetcher.git`
3. Navigate to the project directory: `cd github-repo-fetcher`

## Prerequisites

Before running the program, you need to install the required dependencies. Open a terminal and run:

```go
go get -u github.com/gin-gonic/gin github.com/google/go-github/v35/github github.com/itsjamie/gin-cors
```

## Usage
1. Open a terminal and navigate to the project directory.
2. Run the program: go run main.go
3. The server will start and listen on http://localhost:4200

## Fetch Repositories
To fetch repositories from the 'github' organization, follow these steps:

1. Open Postman or any similar tool.
2. Create a GET request to http://localhost:4200/repositories.
3. Send the request.

You will receive JSON data containing repository details.

## Custom Organization
You can customize the organization name by modifying the 'fetchOrganizationRepositories' function call in the 'main' function.

```go
orgRepos, err := fetchOrganizationRepositories("github") // Replace with your actual organization name
```

## Drill Down

1. Web Server:
The code sets up a web server using the Gin framework (github.com/gin-gonic/gin). A web server listens for incoming HTTP requests and responds with appropriate data. Think of it like a restaurant that listens to customers' orders and serves them food.

2. Routing:
The server defines routes, which are URLs that clients can visit to get specific responses. In this code, there's a route defined at /repositories that fetches GitHub repositories and sends the data back to the client. Think of routes as different sections or pages in a restaurant menu.

3. CORS Middleware:
The code uses CORS middleware (github.com/itsjamie/gin-cors) to handle Cross-Origin Resource Sharing. This allows the server to receive requests from different origins (domains). It's like giving permission for customers from other neighborhoods to visit your restaurant.

4. Data Structure:
The Repository struct defines a structure for organizing information about a GitHub repository. It's like a template for how data should be stored and presented.

5. GitHub API Interaction:
The code uses the github.NewClient function from the github.com/google/go-github/v35/github package to create a client for interacting with the GitHub API. This allows the program to communicate with GitHub's servers and fetch repository data.

6. Fetching Repositories:
The fetchOrganizationRepositories function is responsible for fetching repositories from a specified GitHub organization. It utilizes the GitHub client to make API calls to get the repository data.

7. Context:
The context.Background() creates a context, which is like a container for request-specific values. It's used when making API calls to GitHub.

8. Looping and Data Manipulation:
The for loop inside the fetchOrganizationRepositories function iterates through the fetched repositories and constructs a list of Repository objects with relevant information.

9. Error Handling:
Throughout the code, you'll notice error handling using if err != nil conditions. This ensures that if something goes wrong (like a failed API call), the program doesn't crash and provides appropriate error messages.

10. JSON Response:
The code uses the c.JSON function from Gin to send back the fetched repository data as a JSON response to the client. It's like serving the ordered food to the customer's table.