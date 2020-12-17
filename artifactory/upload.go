package artifactory

import (
	"errors"
	"log"
	"net/http"
	"os"
	"regexp"
)

func Upload(artifactoryUrl string, filePath string) error {
	if validateLocalPackagePath(filePath) {
		err, fileUrl := artifactoryUrlFromPackageName(artifactoryUrl, filePath)
		if err != nil {
			return err
		}
		return artifactoryPush(fileUrl, filePath)
	} else {
		return errors.New("could not stat file " + filePath)
	}
}

func artifactoryPush(artifactoryUrl string, filePath string) error {
	data, err := os.Open(filePath)
	if err != nil {
		log.Printf("Could not open file in artifactoryPush %s", filePath)
		return err
	}
	defer data.Close()
	req, err := http.NewRequest("PUT", artifactoryUrl, data)
	if err != nil {
		return err
	}

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return err
	}
	defer res.Body.Close()
	return err

}

func validateLocalPackagePath(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}
	return true

}

func artifactoryUrlFromPackageName(artifactoryUrl string, filePath string) (error, string) {
	re := regexp.MustCompile(`\/([^\/]*\d*?\.\d*?\.\d*?\.tgz$)`)
	// re := regexp.MustCompile(`\/([^\/]*\.$)`)
	result := re.FindAllStringSubmatch(filePath, -1)
	if len(result) < 1 {
		return errors.New("no path matched"), ""
	}

	return nil, artifactoryUrl + result[0][1]
}
