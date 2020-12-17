package helm

import (
	"errors"
	"regexp"

	"github.com/chrisevett/helmsync/util"
)

func HelmPackage(chartPath string, version string) (string, error) {
	output, err := packageCreate(chartPath, version)
	if err != nil {
		return "", err
	}
	packageLocation, err := parsePath(output)
	return packageLocation, err
}

func packageCreate(chartPath string, version string) (string, error) {
	cmd := execCommand("helm", "package", chartPath, "--destination", "/tmp/", "--version", version)
	returnText, err := util.Command(cmd)
	return returnText, err
}

func parsePath(output string) (string, error) {
	re := regexp.MustCompile(`(\/.*)`)
	result := re.FindAllStringSubmatch(output, -1)
	if len(result) < 1 {
		return "", errors.New("no path matched")
	}
	return result[0][1], nil
}
