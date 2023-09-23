package git

import (
	"fmt"
	"github.com/kohirens/stdlib/path"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

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
		log.Panicf("error un-bundling %q to %q for a unit test", srcRepo, repoPath)
	}

	absPath, e2 := filepath.Abs(repoPath)
	if e2 != nil {
		panic(e2.Error())
	}

	return absPath
}
