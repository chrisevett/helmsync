package helm

import (
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("HelmPackage", func() {
	Context("when the package result is ok", func() {
		It("returns the packagelocation", func() {
			execCommand = fakeExecCommand
			defer func() { execCommand = exec.Command }()

			result, err := HelmPackage("../goodchart/", "0.2.0")
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(Equal(goodHelmPath))
		})
	})
	Context("when the package result is not ok", func() {
		It("returns an error", func() {
			execCommand = fakeExecCommand
			defer func() { execCommand = exec.Command }()

			result, err := HelmPackage("../badchart", "0.2.0")
			Expect(err).To(HaveOccurred())
			Expect(result).To(Equal(""))
		})
	})
})
var _ = Describe("packageCreate", func() {
	Context("when the package command is successful", func() {
		It("returns the stdout from helm", func() {
			execCommand = fakeExecCommand
			defer func() { execCommand = exec.Command }()

			stdout, err := packageCreate("../goodchart/", "0.2.0")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(stdout).Should(Equal(helmPackageGoodChart))
		})
	})
	Context("when the chart is malformed", func() {
		It("returns an error", func() {
			execCommand = fakeExecCommand
			defer func() { execCommand = exec.Command }()

			output, err := packageCreate("../badchart/", "0.2.0")
			Expect(err).Should(HaveOccurred())
			Expect(output).Should(Equal(""))
		})
	})
	Context("when there is an os error", func() {
		It("returns an error", func() {
			execCommand = fakeExecCommand
			defer func() { execCommand = exec.Command }()

			output, err := lintPackage("../osError/")
			Expect(err).Should(HaveOccurred())
			Expect(output).Should(Equal("this is a helm error"))
		})
	})
})

var _ = Describe("parsePath", func() {
	Context("when there is a valid path", func() {
		It("returns the path", func() {
			actual, err := parsePath(helmPackageGoodChart)
			Expect(actual).Should(Equal(goodHelmPath))
			Expect(err).ShouldNot(HaveOccurred())
		})
	})
	Context("when there is not a valid path", func() {
		It("returns an error", func() {
			_, err := parsePath(helmPackageBadChart)
			Expect(err).Should(HaveOccurred())
		})
	})
})
