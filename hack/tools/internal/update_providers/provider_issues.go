//go:build tools
// +build tools

/*
Copyright 2023 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// main is the main package for the open issues utility.
package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const (
	baseURL = "https://api.github.com"
	// TODO: Update v1.x.0-beta.0 for each release
	issueTitle = "CAPI v1.x.0-beta.0 has been released and is ready for testing"
)

var (
	repoList = []string{
		"kubernetes-sigs/cluster-api-addon-provider-helm",
		"kubernetes-sigs/cluster-api-provider-aws",
		"kubernetes-sigs/cluster-api-provider-azure",
		"kubernetes-sigs/cluster-api-provider-cloudstack",
		"kubernetes-sigs/cluster-api-provider-digitalocean",
		"kubernetes-sigs/cluster-api-provider-gcp",
		"kubernetes-sigs/cluster-api-provider-kubemark",
		"kubernetes-sigs/cluster-api-provider-kubevirt",
		"kubernetes-sigs/cluster-api-provider-ibmcloud",
		"kubernetes-sigs/cluster-api-provider-nested",
		"oracle/cluster-api-provider-oci",
		"kubernetes-sigs/cluster-api-provider-openstack",
		"kubernetes-sigs/cluster-api-operator",
		"kubernetes-sigs/cluster-api-provider-packet",
		"kubernetes-sigs/cluster-api-provider-vsphere",
	}
)

// Issue is the struct for the issue.
type Issue struct {
	// Title is the title of the issue to be created on provider repositories.
	Title string `json:"title"`

	// Body is the body of the issue to be created on provider repositories.
	Body string `json:"body"`
}

// IssueResponse is the struct for the issue response.
type IssueResponse struct {
	// HTMLURL is the URL of the issue.
	HTMLURL string `json:"html_url"`
}

func main() {
	githubToken, keySet := os.LookupEnv("GITHUB_ISSUE_OPENER_TOKEN")
	if !keySet || githubToken == "" {
		fmt.Println("GitHub personal access token is required.")
		fmt.Println("Refer to README.md in folder for more information.")
		os.Exit(1)
	}

	// always start in dry run mode unless explicitly set to false
	dryRun := true
	if os.Getenv("PROVIDER_ISSUES_DRY_RUN") != "false" {
		fmt.Printf("\n")
		fmt.Println("###############################################")
		fmt.Println("This script will run in dry run mode.")
		fmt.Println("To run it for real, set the PROVIDER_ISSUES_DRY_RUN environment variable to \"false\".")
		fmt.Println("###############################################")
		fmt.Printf("\n")
	} else {
		dryRun = false
	}

	fmt.Println("List of CAPI Providers:")
	fmt.Println("-", strings.Join(repoList, "\n- "))
	fmt.Printf("\n")

	issueBody := getIssueBody()
	fmt.Println("Issue body:")
	fmt.Println(issueBody)

	// if dry run, exit
	if dryRun {
		fmt.Printf("\n")
		fmt.Println("DRY RUN: issue(s) body will not be posted.")
		fmt.Println("Exiting...")
		fmt.Printf("\n")
		os.Exit(0)
	}

	// else, ask for confirmation
	fmt.Printf("\n")
	fmt.Println("Issues will be posted to the above repositories.")
	continueOrAbort()

	for _, repo := range repoList {
		issue := Issue{
			Title: issueTitle,
			Body:  issueBody,
		}

		issueJSON, err := json.Marshal(issue)
		if err != nil {
			fmt.Printf("Failed to marshal issue: %s\n", err)
			os.Exit(1)
		}

		url := fmt.Sprintf("%s/repos/%s/issues", baseURL, repo)
		req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, url, bytes.NewBuffer(issueJSON))
		if err != nil {
			fmt.Printf("Failed to create request: %s\n", err)
			os.Exit(1)
		}

		req.Header.Set("Accept", "application/vnd.github+json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", githubToken))
		req.Header.Set("X-Github-Api-Version", "2022-11-28")
		req.Header.Set("User-Agent", "provider_issues")
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("Failed to send request: %s\n", err)
			os.Exit(1)
		}

		if resp.StatusCode != http.StatusCreated {
			fmt.Printf("Failed to create issue for repository '%s'\nStatus code: %d\nStatus:%s\n\n\n", repo, resp.StatusCode, resp.Status)
		} else {
			responseBody, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Printf("Failed to read response body: %s\n", err)
				err := resp.Body.Close()
				if err != nil {
					fmt.Printf("Failed to close response body: %s\n", err)
				}
				os.Exit(1)
			}

			var issueResponse IssueResponse
			err = json.Unmarshal(responseBody, &issueResponse)
			if err != nil {
				fmt.Printf("Failed to unmarshal issue response: %s\n", err)
				err := resp.Body.Close()
				if err != nil {
					fmt.Printf("Failed to close response body: %s\n", err)
				}
				os.Exit(1)
			}

			fmt.Printf("\nIssue created for repository '%s'\nURL: %s\n", repo, issueResponse.HTMLURL)
		}
		err = resp.Body.Close()
		if err != nil {
			fmt.Printf("Failed to close response body: %s\n", err)
		}
	}
}

func continueOrAbort() {
	fmt.Println("Continue? (y/n)")
	var response string
	_, err := fmt.Scanln(&response)
	if err != nil {
		fmt.Printf("Failed to read response: %s\n", err)
		os.Exit(1)
	}
	if response != "y" {
		fmt.Println("Aborting...")
		os.Exit(0)
	}
}

func getIssueBody() string {
	// do not indent the body
	// indenting the body will result in the body being posted as a code snippet
	issueBody := `
<!-- body -->
<!-- TODO: remove all TODOs before running this utility -->
<!-- TODO: update CAPI release semver -->
CAPI v1.x.0-beta.0 has been released and is ready for testing.
Looking forward to your feedback before CAPI 1.x.0 release!

## For quick reference

<!-- body -->
<!-- TODO: CAPI release notes -->
- [CAPI v1.x.0-beta.0 release notes](https://github.com/kubernetes-sigs/cluster-api/releases/tag/v1.x.0-beta.0)
- [Shortcut to CAPI git issues](https://github.com/kubernetes-sigs/cluster-api/issues)

## Following are the planned dates for the upcoming releases

<!-- TODO: update CAPI release timeline -->
|Release|Expected Date|
|-----|-----|
|v1.5.0-beta.x | Tuesday 5th July 2023|
|release-1.5 branch created (Begin [Code Freeze])|Tuesday 11th July 2023|
|v1.5.0-rc.0 released|Tuesday 11th July 2023|
| release-1.5 jobs created | Tuesday 11th July 2023 |
| v1.5.0-rc.x released | Tuesday 18th July 2023 |
| v1.5.0 released | Tuesday 25th July 2023 |

<!-- [List of CAPI providers](https://github.com/kubernetes-sigs/cluster-api/blob/main/docs/release/release-tasks.md#communicate-beta-to-providers) -->
<!-- body -->
`
	return issueBody
}
