package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/kohirens/stdlib/log"
	"io"
	"path/filepath"
	"time"
)

type ReleaseBody struct {
	Body         string `json:"body"`
	Name         string `json:"name"`
	TagName      string `json:"tag_name"`
	TargetCommit string `json:"target_commitish"`
}

type ReleasesResponse struct {
	Url             string    `json:"url"`
	HtmlUrl         string    `json:"html_url"`
	AssetsUrl       string    `json:"assets_url"`
	UploadUrl       string    `json:"upload_url"`
	TarballUrl      string    `json:"tarball_url"`
	ZipballUrl      string    `json:"zipball_url"`
	DiscussionUrl   string    `json:"discussion_url"`
	Id              int       `json:"id"`
	NodeId          string    `json:"node_id"`
	TagName         string    `json:"tag_name"`
	TargetCommitish string    `json:"target_commitish"`
	Name            string    `json:"name"`
	Body            string    `json:"body"`
	Draft           bool      `json:"draft"`
	Prerelease      bool      `json:"prerelease"`
	CreatedAt       time.Time `json:"created_at"`
	PublishedAt     time.Time `json:"published_at"`
	Author          Author    `json:"author"`
	Assets          []Asset   `json:"assets"`
}

// GetReleaseLatest Get a published release with the specified tag.
//
//	see: https://docs.github.com/en/rest/releases/releases?apiVersion=2022-11-28#get-the-latest-release
//	sample: https://api.github.com/repos/OWNER/REPO/releases/latest
func (gh *Client) GetReleaseLatest() (*Release, error) {
	url := fmt.Sprintf(epReleaseLatest, gh.Host, gh.Org, gh.Repository)

	res, err1 := gh.send("GET", url, nil)
	if err1 != nil {
		return nil, fmt.Errorf(stderr.CouldNotRequest, url, err1.Error())
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf(stderr.ReturnStatusCode, res.StatusCode)
	}

	bodyBits, err2 := io.ReadAll(res.Body)
	if err2 != nil {
		return nil, fmt.Errorf(stderr.CouldNotReadResponseBody, err2.Error())
	}

	rel := &Release{}
	if e := json.Unmarshal(bodyBits, rel); e != nil {
		return nil, fmt.Errorf(stderr.CouldNotDecodeJson, e.Error())
	}

	return rel, nil
}

// GetReleaseIdByTag Get a published release with the specified tag.
// see: https://docs.github.com/en/rest/releases/releases?apiVersion=2022-11-28#get-a-release-by-tag-name
// sample: https://api.github.com/repos/OWNER/REPO/releases/tags/TAG
func (gh *Client) GetReleaseIdByTag(version string) (*Release, error) {
	if version == "" {
		return nil, fmt.Errorf(stderr.VersionArgEmpty)
	}

	url := fmt.Sprintf(epReleaseId, gh.Host, gh.Org, gh.Repository, version)

	res, err1 := gh.send("GET", url, nil)
	if err1 != nil {
		return nil, fmt.Errorf(stderr.CouldNotRequest, url, err1.Error())
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf(stderr.ReturnStatusCode, res.StatusCode)
	}

	bodyBits, err2 := io.ReadAll(res.Body)
	if err2 != nil {
		return nil, fmt.Errorf(stderr.CouldNotReadResponseBody, err2.Error())
	}

	rel := &Release{}
	if e := json.Unmarshal(bodyBits, rel); e != nil {
		return nil, fmt.Errorf(stderr.CouldNotDecodeJson, e.Error())
	}

	return rel, nil
}

// TagAndRelease Users with push access to the repository can create a release.
//
//	see: https://docs.github.com/en/rest/releases/releases?apiVersion=2022-11-28#create-a-release
//	sample: https://api.github.com/repos/OWNER/REPO/releases
func (gh *Client) TagAndRelease(commit, name, tag string) (*ReleasesResponse, error) {
	uri := fmt.Sprintf(epRelease, gh.Host, gh.Org, gh.Repository)
	body := &ReleaseBody{
		Name:         name,
		TagName:      tag,
		TargetCommit: commit,
	}

	bodyBits, err1 := json.Marshal(body)
	if err1 != nil {
		return nil, fmt.Errorf(stderr.CouldNotJsonEncode, body, err1.Error())
	}

	log.Logf("attempting to publish a release to %v\n", uri)

	res, err2 := gh.send("POST", uri, bytes.NewReader(bodyBits))
	if err2 != nil {
		return nil, fmt.Errorf(stderr.CouldNotMakeRequest, err2)
	}

	if res.StatusCode != 201 {
		return nil, fmt.Errorf(
			stderr.ResponseStatusCode,
			res.StatusCode,
			uri,
			res.Status,
		)
	}

	b, err3 := io.ReadAll(res.Body)
	if err3 != nil {
		return nil, fmt.Errorf(stderr.CouldNotReadResponse, err3.Error())
	}

	rr := &ReleasesResponse{}
	err4 := json.Unmarshal(b, rr)
	if err4 != nil {
		return nil, fmt.Errorf(stderr.CouldNotJsonDecode, err4.Error())
	}

	return rr, nil
}

// UploadAsset The endpoint you call to upload release assets is specific to
// your release. Use the upload_url
//
//	see: https://docs.github.com/en/rest/releases/assets?apiVersion=2022-11-28#upload-a-release-asset
func (gh *Client) UploadAsset(assetPath string, release *Release) (*Asset, error) {

	basename := filepath.Base(assetPath)
	url := fmt.Sprintf(epUploadAsset, gh.Org, gh.Repository, release.Id) + "?name=" + basename

	body, errBody := bodyFromFile(assetPath)
	if errBody != nil {
		return nil, errBody
	}

	res, err2 := gh.send("POST", url, body)
	if err2 != nil {
		return nil, fmt.Errorf(stderr.CouldNotRequest, url, err2.Error())
	}

	if res.StatusCode != 201 {
		return nil, fmt.Errorf(stderr.ReturnStatusCode, res.StatusCode)
	}

	bodyBits, err2 := io.ReadAll(res.Body)
	if err2 != nil {
		return nil, fmt.Errorf(stderr.CouldNotReadResponseBody, err2.Error())
	}

	ast := &Asset{}
	if e := json.Unmarshal(bodyBits, ast); e != nil {
		return nil, fmt.Errorf(stderr.CouldNotDecodeJson, e.Error())
	}

	return ast, nil
}
