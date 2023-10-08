package git

import (
	"bytes"
	"fmt"
	"github.com/kohirens/stdlib/log"
	"github.com/kohirens/stdlib/path"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

// Checkout Open an existing repo and checkout commit by full ref-name
func Checkout(repoLocalPath, ref string) (string, string, error) {
	log.Logf("pulling latest\n")
	_, e1 := gitCmd(repoLocalPath, "fetch", "--all", "-p")
	if e1 != nil {
		return "", "", fmt.Errorf(stderr.GitFetchFailed, repoLocalPath, ref, e1.Error())
	}

	log.Infof(stdout.RefInfo, ref)
	log.Infof(stdout.GitCheckout, ref)

	_, e2 := gitCmd(repoLocalPath, "checkout", ""+ref)
	if e2 != nil {
		return "", "", fmt.Errorf(stderr.GitCheckoutFailed, e2.Error())
	}

	repoDir, e8 := filepath.Abs(repoLocalPath)
	if e8 != nil {
		return "", "", e8
	}

	latestCommitHash, e4 := HeadCommitHash(repoDir)
	if e4 != nil {
		return "", "", e4
	}

	return repoDir, latestCommitHash, nil
}

// Clone Use git application to clone a repo from one location to a local
// directory.
func Clone(repoUri, repoDir, refName string) (string, string, error) {
	log.Infof("branch to clone is %q", refName)
	log.Infof("git clone %s", repoUri)

	var sco []byte
	var e1 error

	if isRemoteRepo(repoUri) {
		branchName := refName
		re := regexp.MustCompile("^refs/[^/]+/(.*)$")
		if re.MatchString(refName) {
			branchName = re.ReplaceAllString(refName, "${1}")
		}

		log.Logf("cloning branch name: %v", branchName)

		// git clone --depth 1 --branch <tag_name> <repo_url>
		// NOTE: Branch cannot be a full ref but can be short ref name or a tag.
		sco, e1 = gitCmd(".", "clone", "--depth", "1", "--branch", branchName, repoUri, repoDir)
	} else {
		// git clone <repo_url>
		sco, e1 = gitCmd(".", "clone", repoUri, repoDir)
		if e1 != nil {
			return "", "", fmt.Errorf(stderr.Cloning, repoUri, e1.Error())
		}

		log.Infof("clone output \n%s", sco)

		// get current branch
		cb, e4 := gitCmd(repoDir, "branch", "--show-current", refName)
		if e4 != nil {
			return "", "", fmt.Errorf(stderr.CurrentBranch, repoUri, e4.Error(), cb)
		}
		// skip if already on desired branch
		if strings.Trim(bytes.NewBuffer(cb).String(), "\r\n") != refName {
			// git checkout <ref_name>
			co, e3 := gitCmd(repoDir, "checkout", "-b", refName)
			if e3 != nil {
				return "", "", fmt.Errorf(stderr.Checkout, repoUri, e3.Error(), co)
			}

			log.Infof("checkout output \n%s", co)
		}
	}

	latestCommitHash, e2 := HeadCommitHash(repoDir)
	if e2 != nil {
		return "", "", fmt.Errorf(stderr.GettingCommitHash, repoDir, e2.Error())
	}

	return repoDir, latestCommitHash, nil
}

// CloneFromBundle Set up a repository from a git bundle.
func CloneFromBundle(bundleName, tmpDir, bundleDir, ps string) string {
	repoPath := tmpDir + ps + bundleName

	// It may have already been unbundled.
	fileInfo, err1 := os.Stat(repoPath)
	if (err1 == nil && fileInfo.IsDir()) || os.IsExist(err1) {
		absPath, e2 := filepath.Abs(repoPath)
		if e2 == nil {
			return absPath
		}
		return repoPath
	}

	wd, e := os.Getwd()
	if e != nil {
		panic(fmt.Sprintf("%v failed to get working directory", e.Error()))
	}

	srcRepo := wd + ps + bundleDir + ps + bundleName + ".bundle"
	// It may not exist.
	if !path.Exist(srcRepo) {
		panic(fmt.Sprintf("%v bundle not found", srcRepo))
	}

	cmd := exec.Command("git", "clone", "-b", "main", srcRepo, repoPath)
	_, _ = cmd.CombinedOutput()
	if ec := cmd.ProcessState.ExitCode(); ec != 0 {
		log.Panf("error un-bundling %q to %q for a unit test", srcRepo, repoPath)
	}

	absPath, e2 := filepath.Abs(repoPath)
	if e2 != nil {
		panic(e2.Error())
	}

	return absPath
}

// HeadCommitHash Returns the HEAD commit hash.
func HeadCommitHash(repoDir string) (string, error) {
	latestCommitHash, e1 := gitCmd(repoDir, "rev-parse", "HEAD")
	if e1 != nil {
		return "", fmt.Errorf(stderr.GettingCommitHash, repoDir, e1.Error())
	}

	return strings.Trim(string(latestCommitHash), "\n"), nil
}

// LatestTag Will return the latest tag or an empty string from a repository.
func LatestTag(repoDir string) (string, error) {
	tags, e1 := RemoteTags(repoDir)
	if e1 != nil {
		return "", fmt.Errorf(stderr.GetLatestTag, repoDir, e1.Error())
	}

	return tags[0], nil
}

// RemoteTags Get the remote tags on a repo using git ls-remote.
func RemoteTags(repo string) ([]string, error) {
	// Even without cloning or fetching, you can check the list of tags on the upstream repo with git ls-remote:
	sco, e1 := gitCmd(repo, "ls-remote", "--sort=-version:refname", "--tags")
	if e1 != nil {
		return nil, fmt.Errorf(stderr.GetRemoteTags, e1.Error())
	}

	reTags := regexp.MustCompile("[a-f0-9]+\\s+refs/tags/(\\S+)")
	mat := reTags.FindAllSubmatch(sco, -1)
	if mat == nil {
		return nil, fmt.Errorf("%s", "no tags found")
	}

	ret := make([]string, len(mat))
	for i, v := range mat {
		log.Dbugf(stdout.RemoteTagDbug1, string(v[1]))
		ret[i] = string(v[1])
	}

	return ret, nil
}

// gitCmd run a git command.
func gitCmd(repoPath string, args ...string) ([]byte, error) {
	cmd := exec.Command("git", args...)
	cmd.Env = os.Environ()
	cmd.Dir = repoPath
	cmdStr := cmd.String()

	log.Infof(stdout.RunningCommand, cmdStr)

	cmdOut, cmdErr := cmd.CombinedOutput()
	exitCode := cmd.ProcessState.ExitCode()

	if cmdErr != nil {
		return nil, fmt.Errorf(stderr.RunGitFailed, args, cmdErr.Error(), cmdOut)
	}

	if exitCode != 0 {
		return nil, fmt.Errorf(stderr.GitExitErrCode, args, exitCode)
	}

	return cmdOut, nil
}

// isRemoteRepo return true if Git repository is a remote URL or false if local.
func isRemoteRepo(repoLocation string) bool {
	if len(repoLocation) < 1 {
		return false
	}
	// git@github.com:kohirens/tmpltoap.git
	// https://github.com/kohirens/tmpltoapp.git
	isGitUri := regexp.MustCompile("^(git|http|https)://.+$")
	if isGitUri.MatchString(repoLocation) {
		return true
	}

	return false
}
