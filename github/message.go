package github

var stderr = struct {
	CouldNotBuildRequest     string
	CouldNotDecodeJson       string
	CouldNotReadFile         string
	CouldNotReadResponseBody string
	CouldNotRequest          string
	ReturnStatusCode         string
	VersionArgEmpty          string
	BranchExists             string
	CouldNotGetRequest       string
	CouldNotJsonEncode       string
	CouldNotJsonDecode       string
	CouldNotMakeRequest      string
	CouldNotMergePullRequest string
	CouldNotPingMergeStatus  string
	CouldNotPrepareRequest   string
	CouldNotPostRequest      string
	CouldNotReadResponse     string
	MergeWaitTimeout         string
	ResponseStatusCode       string
}{
	CouldNotBuildRequest:     "could not build a request: %s",
	CouldNotDecodeJson:       "could not decode JSON: %s",
	CouldNotReadFile:         "could not read %s: %s",
	CouldNotReadResponseBody: "could not read response body: %s",
	CouldNotRequest:          "problem with request to %s: %s",
	ReturnStatusCode:         "unexpected return status code %d",
	VersionArgEmpty:          "version argument cannot be an empty string",
	BranchExists:             "the branch %q exists on %s, please delete it manually, then re-run this job so it can complete successfully",
	CouldNotGetRequest:       "could not GET request: %v",
	CouldNotJsonEncode:       "could not encode %t to JSON: %v",
	CouldNotJsonDecode:       "could not decode JSON: %s",
	CouldNotMakeRequest:      "could not make %s request: %s",
	CouldNotMergePullRequest: "unable to merge pr %d: %s",
	CouldNotPingMergeStatus:  "unable to ping pull request %d merge status: %s",
	CouldNotPrepareRequest:   "could not prepare a request: %v",
	CouldNotPostRequest:      "could not POST request: %v",
	CouldNotReadResponse:     "could not read response body: %v",
	MergeWaitTimeout:         "pr %d has not merged after %d seconds",
	ResponseStatusCode:       "got a %d response from %v: %s",
}

var stdout = struct {
	UrlRequest        string
	CheckMergeStatus  string
	MakePullRequest   string
	MergeResponse     string
	PullRequestMade   string
	PullRequestMerged string
	SendMergeRequest  string
}{
	UrlRequest:        "%s'ing to %s",
	CheckMergeStatus:  "checking pr %d merged status",
	MakePullRequest:   "making a pull request to %s",
	MergeResponse:     "merge status: %v",
	PullRequestMade:   "pull request %d has been made",
	PullRequestMerged: "pr %d has has been merged",
	SendMergeRequest:  "sending a merge request to %s",
}
