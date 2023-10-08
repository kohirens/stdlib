package git

var stderr = struct {
	Checkout          string
	Cloning           string
	CurrentBranch     string
	GetLatestTag      string
	GetRemoteTags     string
	GettingCommitHash string
	GitCheckoutFailed string
	GitExitErrCode    string
	GitFetchFailed    string
	RunGitFailed      string
}{
	Checkout:          "checkout failed for branch %q",
	Cloning:           "error cloning %v: %s",
	CurrentBranch:     "failed to get current for %s",
	GetLatestTag:      "failed to get latest tag from %v: %v",
	GetRemoteTags:     "could not get remote tags, please check for a typo, it exist, and is readable: %v",
	GettingCommitHash: "error getting commit hash %v: %s",
	GitCheckoutFailed: "git checkout failed: %s",
	GitExitErrCode:    "git %v returned exit code %q",
	GitFetchFailed:    "fetch failed on %s and %s; %s",
	RunGitFailed:      "error running git %v: %v%s",
}

var stdout = struct {
	GitCheckout    string
	RefInfo        string
	RemoteTagDbug1 string
	RunningCommand string
}{
	GitCheckout:    "git checkout %s",
	RefInfo:        "ref = %v ",
	RemoteTagDbug1: "remote tag: %v",
	RunningCommand: "running command %s",
}
