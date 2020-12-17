package artifactory

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("aritfactoryUrlFromPackageName", func() {
	Context("when the package path is valid", func() {
		It("returns a url with the package name", func() {

			packagePath := "/path/to/package/my-chart.4.2.0.tgz"

			expected := "http://artifactory.buts.org/repo/my-chart.4.2.0.tgz"
			artifactoryUrl := "http://artifactory.buts.org/repo/"

			err, actual := artifactoryUrlFromPackageName(artifactoryUrl, packagePath)
			Expect(actual).Should(Equal(expected))
			Expect(err).ShouldNot(HaveOccurred())
		})
		It("returns a url when the path is one deep", func() {

			packagePath := "/tmp/istio.0.1.10.tgz"

			expected := "http://artifactory.buts.org/repo/istio.0.1.10.tgz"
			artifactoryUrl := "http://artifactory.buts.org/repo/"

			err, actual := artifactoryUrlFromPackageName(artifactoryUrl, packagePath)
			Expect(actual).Should(Equal(expected))
			Expect(err).ShouldNot(HaveOccurred())
		})
		Context("when the package path is invalid", func() {
			It("returns an error", func() {

				packagePath := "/fart/chart"

				artifactoryUrl := "http://artifactory.buts.org/repo/"

				err, _ := artifactoryUrlFromPackageName(artifactoryUrl, packagePath)
				Expect(err).Should(HaveOccurred())
			})
		})
	})
})
