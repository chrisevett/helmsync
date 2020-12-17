package main

import (
	"log"

	"github.com/chrisevett/helmsync/artifactory"
	"github.com/chrisevett/helmsync/git"
	"github.com/chrisevett/helmsync/helm"
)

func main() {
	config, err := ParseConfigs()
	if err != nil {
		log.Printf("error parsing configuration")
		logAndExit(err)
	}

	folders, err := git.ChangedFolders(config.repoPath)
	if err != nil {
		log.Printf("error detecting git changes")
		logAndExit(err)
	}

	totalFailures := 0
	for _, folder := range folders {
		fullpath := config.repoPath + "/" + folder
		failures, err := helm.HelmLint(fullpath, config.ignoreInfo)
		totalFailures += failures
		if err != nil {
			log.Printf("error linting helm chart: %s", folder)
			logAndExit(err)
		}
	}

	log.Printf("linting failures %d", totalFailures)
	if totalFailures > 0 {
		log.Printf("Too many failures. Failing.")
		logAndExit(err)
	}

	for _, folder := range folders {
		fullpath := config.repoPath + "/" + folder
		location, err := helm.HelmPackage(fullpath, config.versionNumber)
		if err != nil {
			log.Printf("error packaging helm chart %s.", folder)
			logAndExit(err)
		}

		err = artifactory.Upload(config.artifactoryUrl, location)
		if err != nil {
			log.Printf("error pushing helm chart to artifactory %s.", folder)
			logAndExit(err)
		}
	}

	log.Printf("Success.")

}

func logAndExit(err error) {
	if err != nil {
		log.Printf("%+v", err)
		log.Fatal("Failing...")
	}
}
