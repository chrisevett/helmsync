package helm

import (
	"os/exec"
	"regexp"

	"github.com/chrisevett/helmsync/util"
)

var execCommand = exec.Command

func HelmLint(chartPath string, ignoreInfo bool) (int, error) {
	var helmErrors, helmInfos int
	output, err := lintPackage(chartPath)

	if err != nil {
		return -1, err
	}

	helmErrors = parseError(output)
	if ignoreInfo == false {
		helmInfos = parseInfo(output)
	}

	return (helmErrors + helmInfos), nil
}

func lintPackage(chartPath string) (string, error) {
	cmd := execCommand("helm", "lint", chartPath)
	returnText, err := util.Command(cmd)
	return returnText, err
}

func parseError(output string) int {
	re := regexp.MustCompile(`\[ERROR\]`)
	result := re.FindAllStringSubmatch(output, -1)
	return len(result)
}

func parseInfo(output string) int {
	re := regexp.MustCompile(`\[INFO\]`)
	result := re.FindAllStringSubmatch(output, -1)
	return len(result)
}
