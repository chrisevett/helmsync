package helm

import (
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("HelmLint", func() {
	Context("when the lint result is ok", func() {
		It("returns 0", func() {
			execCommand = fakeExecCommand
			defer func() { execCommand = exec.Command }()

			result, err := HelmLint("../lintOk/", true)
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(Equal(0))

		})
	})
	Context("when the lint result is info", func() {
		It("returns 0 with default config", func() {
			execCommand = fakeExecCommand
			defer func() { execCommand = exec.Command }()

			result, err := HelmLint("../lintInfo/", true)
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(Equal(0))
		})
		It("returns the number of info if configured not to tolerate it", func() {
			execCommand = fakeExecCommand
			defer func() { execCommand = exec.Command }()

			result, err := HelmLint("../lintInfo/", false)
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(Equal(1))
		})
	})
	Context("when the lint result is error", func() {
		It("returns the number of errors", func() {
			execCommand = fakeExecCommand
			defer func() { execCommand = exec.Command }()

			result, err := HelmLint("../lintManyErrors/", true)
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(Equal(2))
		})
	})
	Context("when the lint result is a mixture of error and info", func() {
		It("returns the number of errors and info if configured", func() {
			execCommand = fakeExecCommand
			defer func() { execCommand = exec.Command }()

			result, err := HelmLint("../lintManyErrors/", false)
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(Equal(3))
		})
	})
	Context("when there is a helm error", func() {
		It("returns the error", func() {
			execCommand = fakeExecCommand
			defer func() { execCommand = exec.Command }()

			result, err := HelmLint("../osError/", true)
			Expect(err).To(HaveOccurred())
			Expect(result).To(Equal(-1))
		})
	})
})

var _ = Describe("lintPackage", func() {
	Context("when a helm command is successful", func() {
		It("returns the stdout from helm", func() {
			execCommand = fakeExecCommand
			defer func() { execCommand = exec.Command }()

			stdout, err := lintPackage("../lintOk/")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(stdout).Should(Equal(helmLintOk))
		})
	})
	Context("when a git command errors", func() {
		It("returns the error code and the stderr", func() {
			execCommand = fakeExecCommand
			defer func() { execCommand = exec.Command }()

			stderr, err := lintPackage("../osError/")
			Expect(err).Should(HaveOccurred())
			Expect(stderr).Should(Equal("this is a helm error"))
		})
	})
})

var _ = Describe("parseError", func() {
	Context("when there is an error", func() {
		It("returns 1", func() {
			actual := parseError(helmLintOneError)
			Expect(actual).Should(Equal(1))
		})
	})
	Context("when there is no error", func() {
		It("returns 0", func() {
			actual := parseError(helmLintOk)
			Expect(actual).Should(Equal(0))
		})
	})
	Context("when there are many errors", func() {
		It("returns the number of errors", func() {
			actual := parseError(helmLintManyErrors)
			Expect(actual).Should(Equal(2))
		})

	})
})

var _ = Describe("parseInfo", func() {
	Context("when there is no info", func() {
		It("returns 0", func() {
			execCommand = fakeExecCommand
			defer func() { execCommand = exec.Command }()
			actual := parseInfo(helmLintOk)
			Expect(actual).Should(Equal(0))
		})
	})
	Context("when there is an info", func() {
		It("returns the count of info", func() {
			execCommand = fakeExecCommand
			defer func() { execCommand = exec.Command }()
			actual := parseInfo(helmLintInfo)
			Expect(actual).Should(Equal(1))
		})
	})
})
