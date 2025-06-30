package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

var ApplicationToken = ""

type FetchOptions struct {
	Url     string
	Method  string
	Headers map[string]string
	Body    map[string]any
	Query   map[string]string
	Token   string
}

// Generic Fetch function for the osu api calls
func Fetch[T any](opts FetchOptions) (T, error) {
	var result T

	// Parse URL
	u, err := url.Parse(opts.Url)
	if err != nil {
		return result, fmt.Errorf("invalid URL: %w", err)
	}

	// Add Query Parameters
	// If they were provided in opts, iterate over the map of params
	// add each key-value pair and encode it back to the url
	if opts.Query != nil {
		q := u.Query()
		for k, v := range opts.Query {
			q.Add(k, v)
		}
		u.RawQuery = q.Encode()
	}
	finalURL := u.String()

	// Request Body for POST/PUT requests
	// From type T to JSON
	var reqBody io.Reader
	if opts.Body != nil {
		bodyData, err := json.Marshal(opts.Body)
		if err != nil {
			return result, fmt.Errorf("could not marshal request body: %w", err)
		}
		reqBody = bytes.NewBuffer(bodyData)
	}

	// Create HTTP request
	req, err := http.NewRequest(opts.Method, finalURL, reqBody)
	if err != nil {
		return result, fmt.Errorf("could not create request: %w", err)
	}

	// Send and Recieve JSONs by default for now
	if opts.Body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")

	// Use ApplicationToken if opts.Token not provided
	token := opts.Token
	if token == "" {
		token = ApplicationToken
	}
	// Set the Bearer prefix
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	// Custom Headers if provided
	for k, v := range opts.Headers {
		req.Header.Set(k, v)
	}

	// Execute HTTP request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result, fmt.Errorf("request failed: %w", err)
	}
	defer res.Body.Close()

	// Read the Response Body
	resData, err := io.ReadAll(res.Body)
	if err != nil {
		return result, fmt.Errorf("could not read body: %w", err)
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return result, fmt.Errorf(
			"request failed with status %d: %s",
			res.StatusCode,
			string(resData),
		)

	}

	// Unmarshal from JSON to type T
	if err := json.Unmarshal(resData, &result); err != nil {
		return result, fmt.Errorf("could not unmarshal response: %w", err)
	}

	return result, nil
}
