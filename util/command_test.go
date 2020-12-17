package util

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
	"os/exec"
	"testing"
)

func fakeExecCommand(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestHelperProcess", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1"}
	return cmd
}

func TestHelperProcess(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}
	args := os.Args

	for len(args) > 0 {
		if args[0] == "--" {
			args = args[1:]
			break
		}
		args = args[1:]
	}
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "No command\n")
		os.Exit(2)
	}

	cmd, args := args[0], args[1:]

	switch cmd {
	case "helm":
		if args[0] == "lint" {
			if args[1] == "../lintInfo/" {
				fmt.Fprintf(os.Stdout, helmLintInfo)
				os.Exit(0)
			} else if args[1] == "../osError/" {
				fmt.Fprintf(os.Stderr, "this is a helm error")
				os.Exit(-1)
			} else {
				fmt.Fprintf(os.Stderr, "invalid input for testing")
				os.Exit(3)
			}
		}
	default:
		fmt.Fprintf(os.Stderr, "Unknown command %q\n", cmd)
		os.Exit(2)
	}
}

var helmLintInfo = `
==> Linting buttles
[INFO] Chart.yaml: icon is recommended

1 chart(s) linted, no failures
`

var _ = Describe("Command", func() {
	Context("when a command is successful", func() {
		It("returns the stdout ", func() {
			execCommand = fakeExecCommand
			defer func() { execCommand = exec.Command }()
			cmd := execCommand("helm", "lint", "../lintInfo/")
			stdout, err := Command(cmd)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(stdout).Should(Equal(helmLintInfo))
		})
	})
	Context("when a command errors", func() {
		It("returns the error code and the stderr", func() {
			execCommand = fakeExecCommand
			defer func() { execCommand = exec.Command }()

			cmd := execCommand("helm", "lint", "../osError/")
			stderr, err := Command(cmd)
			Expect(err).Should(HaveOccurred())
			Expect(stderr).Should(Equal("this is a helm error"))
		})
	})
})
