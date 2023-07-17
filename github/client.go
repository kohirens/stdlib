package github

import (
	"fmt"
	"github.com/kohirens/stdlib/log"
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"
)

const (
	BaseUri           = "https://%s/repos/%s/%s"
	epBranches        = BaseUri + "/branches/%s"
	epPulls           = BaseUri + "/pulls"
	epPullMerge       = BaseUri + "/pulls/%d/merge"
	epRelease         = BaseUri + "/releases"
	epReleaseLatest   = BaseUri + "/releases/latest"
	publicServer      = "github.com"
	epUploadAsset     = "https://uploads.github.com/repos/%s/%s/releases/%d/assets"
	epReleaseId       = BaseUri + "/releases/tags/%s"
	HeaderApiAccept   = "application/vnd.github+json"
	HeaderApiPostType = "application/octet-stream"
)

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	ChglogConfigFile string
	Client           HttpClient
	Domain           string
	MergeMethod      string
	Org              string
	RepositoryUri    string
	Repository       string
	Token            string
	Username         string
	Host             string
}

var (
	HeaderApiVersion = "2022-11-28"
)

// NewClient Return a GitHub API client.
func NewClient(org, repository, token, host string, client HttpClient) *Client {
	if host == publicServer { // patch for public GitHub
		host = "api." + host
	} else {
		// patch for Enterprise server
		host = host + "/api/v3"
	}

	return &Client{
		Client:     client,
		Org:        org,
		Repository: repository,
		Token:      token,
		Username:   "git",
		Host:       host,
	}
}

func (gh *Client) DoesBranchExistRemotely(branch string) bool {
	uri := fmt.Sprintf(epBranches, gh.Host, gh.Org, gh.Repository, branch)

	res, err1 := gh.send("GET", uri, nil)
	if err1 != nil {
		log.Logf(stderr.CouldNotGetRequest, err1.Error())
		return false
	}

	bodyBits, err2 := io.ReadAll(res.Body)
	if err2 != nil {
		log.Logf(stderr.CouldNotReadResponse, err2.Error())
		return false
	}

	fmt.Println(bodyBits)

	return res.StatusCode == 200
}

func (gh *Client) WaitForPrToMerge(prNumber int, waitSeconds int) error {
	uri := fmt.Sprintf(epPullMerge, gh.Host, gh.Org, gh.Repository, prNumber)

	log.Logf(stdout.CheckMergeStatus, prNumber)

	res, err2 := gh.send("GET", uri, nil)
	if err2 != nil {
		return fmt.Errorf(stderr.CouldNotMakeRequest, "GET", err2.Error())
	}

	for i := 0; i < waitSeconds; i++ {
		time.Sleep(1 * time.Second)

		log.Infof("checking if pr %d was merged\n", prNumber)

		res, err2 = gh.send("GET", uri, nil)
		if err2 != nil {
			return fmt.Errorf(stderr.CouldNotPingMergeStatus, prNumber, err2.Error())
		}

		if res.StatusCode == 204 {
			break
		}
	}

	if res.StatusCode != 204 {
		return fmt.Errorf(stderr.MergeWaitTimeout, prNumber, waitSeconds)
	}

	return nil
}

func NewRequest(uri, method, token string) (*http.Request, error) {
	req, err1 := http.NewRequest(method, uri, nil)
	if err1 != nil {
		return nil, fmt.Errorf(stderr.CouldNotPrepareRequest, err1.Error())
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Accept", "application/vnd.github+json")

	return req, nil
}

func ParseRepositoryUri(uri string) (string, string, string) {
	//https://github.com/kohirens/version-release-orb
	//git@github.com:kohirens/version-release-orb.git
	re := regexp.MustCompile(`^(https://|git@)([^/:]+)[/:]([^/]+)/(.+)`)
	m := re.FindAllStringSubmatch(uri, -1)

	if m != nil {
		return m[0][2], m[0][3], strings.Replace(m[0][4], ".git", "", 1)
	}

	return "", "", ""
}

func (gh *Client) send(method, url string, body io.Reader) (*http.Response, error) {
	req, err1 := http.NewRequest(method, url, body)
	if err1 != nil {
		return nil, fmt.Errorf(stderr.CouldNotBuildRequest, err1.Error())
	}

	req.Header.Set("Authorization", "Bearer "+gh.Token)
	req.Header.Set("Accept", HeaderApiAccept)
	req.Header.Set("X-GitHub-Api-Version", HeaderApiVersion)
	if method == "POST" {
		req.Header.Set("Content-Type", HeaderApiPostType)
	}

	log.Infof(stdout.UrlRequest, method, url)

	res, err2 := gh.Client.Do(req)
	if err2 != nil {
		return nil, fmt.Errorf(stderr.CouldNotRequest, url, err2.Error())
	}

	return res, nil
}
