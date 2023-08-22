package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v35/github"
	cors "github.com/itsjamie/gin-cors"
)

// Repository holds information about a GitHub repository
type Repository struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

func main() {
	// Initialize a new Gin router
	router := gin.Default()

	// Apply CORS middleware to allow requests from different origins
	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Content-Type",
		ExposedHeaders:  "",
		Credentials:     true,
		ValidateHeaders: false,
	}))

	// Define a route to fetch and return repository data
	router.GET("/repositories", func(c *gin.Context) {
		// Fetch repositories from the 'github' organization
		orgRepos, err := fetchOrganizationRepositories("github") // Replace with your actual organization name
		if err != nil {
			log.Printf("Error fetching organization repositories: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch repositories"})
			return
		}
		// Return repository data as a JSON response
		c.JSON(http.StatusOK, orgRepos)
	})

	// Start the server on localhost:4200
	router.Run("localhost:4200")
}

// fetchOrganizationRepositories fetches repositories from a GitHub organization
func fetchOrganizationRepositories(organization string) ([]Repository, error) {
	ctx := context.Background()

	// Create a GitHub API client
	client := github.NewClient(nil)

	// Set options for listing repositories
	opt := &github.RepositoryListByOrgOptions{
		ListOptions: github.ListOptions{PerPage: 10},
	}

	// Fetch repositories from the organization
	repos, _, err := client.Repositories.ListByOrg(ctx, organization, opt)
	if err != nil {
		return nil, err
	}

	// Prepare the list of repositories to be returned
	var orgRepos []Repository
	for _, repo := range repos {
		orgRepos = append(orgRepos, Repository{
			Name:        *repo.Name,
			Description: *repo.Description,
			URL:         *repo.HTMLURL,
		})
	}

	return orgRepos, nil
}
