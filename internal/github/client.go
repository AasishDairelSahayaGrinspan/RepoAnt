package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const baseURL = "https://api.github.com"

type Repository struct {
	Name     string
	FullName string
	Owner    string
	Fork     bool
}

type repoResponse struct {
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Fork     bool   `json:"fork"`
	Owner    struct {
		Login string `json:"login"`
	} `json:"owner"`
	Permissions struct {
		Admin bool `json:"admin"`
	} `json:"permissions"`
}

type Client struct {
	token      string
	httpClient *http.Client
}

func NewClient(token string) *Client {
	return &Client{
		token: token,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (c *Client) doRequest(method, url string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	return c.httpClient.Do(req)
}

// CheckTokenScopes verifies the token has delete_repo scope
func (c *Client) CheckTokenScopes() (bool, string, error) {
	resp, err := c.doRequest("GET", baseURL+"/user")
	if err != nil {
		return false, "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return false, "", fmt.Errorf("failed to verify token")
	}

	scopes := resp.Header.Get("X-OAuth-Scopes")
	hasDeleteRepo := strings.Contains(scopes, "delete_repo")

	return hasDeleteRepo, scopes, nil
}

func (c *Client) ListRepositories() ([]Repository, error) {
	var allRepos []Repository
	page := 1

	for {
		url := fmt.Sprintf("%s/user/repos?per_page=100&page=%d&affiliation=owner", baseURL, page)
		resp, err := c.doRequest("GET", url)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		if resp.StatusCode == 401 {
			return nil, fmt.Errorf("unauthorized: invalid or expired token")
		}

		if resp.StatusCode == 403 {
			return nil, fmt.Errorf("forbidden: rate limit exceeded or insufficient permissions")
		}

		if resp.StatusCode != 200 {
			body, _ := io.ReadAll(resp.Body)
			return nil, fmt.Errorf("API error (%d): %s", resp.StatusCode, string(body))
		}

		var repos []repoResponse
		if err := json.NewDecoder(resp.Body).Decode(&repos); err != nil {
			return nil, fmt.Errorf("failed to parse response: %w", err)
		}

		if len(repos) == 0 {
			break
		}

		for _, r := range repos {
			allRepos = append(allRepos, Repository{
				Name:     r.Name,
				FullName: r.FullName,
				Owner:    r.Owner.Login,
				Fork:     r.Fork,
			})
		}

		page++
	}

	return allRepos, nil
}

func (c *Client) DeleteRepository(owner, name string) error {
	url := fmt.Sprintf("%s/repos/%s/%s", baseURL, owner, name)
	resp, err := c.doRequest("DELETE", url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 204:
		return nil
	case 401:
		return fmt.Errorf("unauthorized: invalid or expired token")
	case 403:
		// Read response body for more details
		body, _ := io.ReadAll(resp.Body)
		bodyStr := string(body)

		// Check for specific error messages
		if strings.Contains(bodyStr, "Must have admin rights") {
			return fmt.Errorf("you don't have admin rights to this repository. Only the owner can delete it")
		}
		if strings.Contains(bodyStr, "delete_repo") {
			return fmt.Errorf("your token is missing the 'delete_repo' scope. Please create a new token with this permission")
		}
		return fmt.Errorf("insufficient permissions. Make sure your token has 'delete_repo' scope and you own this repository.\nDetails: %s", bodyStr)
	case 404:
		return fmt.Errorf("repository not found or already deleted")
	default:
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("API error (%d): %s", resp.StatusCode, string(body))
	}
}
