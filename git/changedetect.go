package git

import (
	"log"
	"os/exec"
	"regexp"

	"github.com/chrisevett/helmsync/util"
)

var execCommand = exec.Command

func ChangedFolders(repoPath string) ([]string, error) {
	stdout, err := gitDiff(repoPath)
	if err != nil {
		return []string{}, err
	}
	log.Printf("stdout %s", stdout)
	folders := parseFolders(stdout)
	return folders, err
}

func gitDiff(repoPath string) (string, error) {
	cmd := execCommand("git", "-C", repoPath, "diff", "master")
	returnText, err := util.Command(cmd)
	return returnText, err
}

func parseFolders(gitOutput string) []string {
	log.Print("Parsing folders...")
	folders := []string{}
	// this is to keep track of dupes
	// based on https://www.golangprograms.com/remove-duplicate-values-from-slice.html
	keys := make(map[string]bool)

	re := regexp.MustCompile(`diff --git a/([^\.].*?)/.* b`)
	res := re.FindAllStringSubmatch(gitOutput, -1)

	for i := range res {
		entry := res[i][1]
		// only add unique entries
		if _, value := keys[entry]; !value {
			keys[entry] = true
			folders = append(folders, entry)
		}
	}

	log.Printf("parsed the following folders %v", folders)
	return folders
}
