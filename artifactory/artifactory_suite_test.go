package artifactory_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestArtifactory(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Artifactory Suite")
}
