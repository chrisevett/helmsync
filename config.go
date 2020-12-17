package main

import (
	"errors"
	"net/url"
	"os"
	"regexp"
	"strconv"
)

type Config struct {
	repoPath       string
	versionNumber  string
	artifactoryUrl string
	ignoreInfo     bool
	lintOnly       bool
}

func ParseConfigs() (Config, error) {
	path, present := os.LookupEnv("REPOPATH")
	//if !(present && validatePath(path)) {
	if !(present) {
		return Config{}, errors.New("Error: REPOPATH is not set")
	}
	version, present := os.LookupEnv("VERSIONNUMBER")
	if !(present && validateSemver(version)) {
		return Config{}, errors.New("Error: VERSIONNUMBER is not set")
	}
	url, present := os.LookupEnv("ARTIFACTORYURL")
	if !(present && validateUrl(url)) {
		return Config{}, errors.New("Error: ARTIFACTORYURL is not set")
	}
	ignoreString, present := os.LookupEnv("IGNOREINFO")
	ignoreInfo, err := strconv.ParseBool(ignoreString)
	if !present || (err != nil) {
		return Config{}, errors.New("Error: IGNOREINFO is not set")
	}

	lintonlyString, present := os.LookupEnv("LINTONLY")
	lintOnly, err := strconv.ParseBool(lintonlyString)
	if !present || (err != nil) {
		return Config{}, errors.New("Error: LINTONLY is not set")
	}
	return Config{repoPath: path, versionNumber: version, artifactoryUrl: url, ignoreInfo: ignoreInfo, lintOnly: lintOnly}, nil
	return Config{}, errors.New("Error: ignoreInfo is not set")
}

func validatePath(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	return true
}

func validateSemver(version string) bool {
	re := regexp.MustCompile(`\d*\.\d*\.\d*`)
	result := re.FindAllStringSubmatch(version, -1)
	return len(result) > 0
}

func validateUrl(myurl string) bool {
	_, err := url.Parse("https://example.org")
	if err != nil {
		return false
	}

	return true
}
