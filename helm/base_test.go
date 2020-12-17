package helm

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func fakeExecCommand(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestHelperProcess", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = append(os.Environ(), "GO_WANT_HELPER_PROCESS=1")
	//	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1"}
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
			if args[1] == "../lintManyErrors/" {
				fmt.Fprintf(os.Stdout, helmLintManyErrors)
				os.Exit(0)
			} else if args[1] == "../lintOneError/" {
				fmt.Fprintf(os.Stdout, helmLintOneError)
				os.Exit(0)
			} else if args[1] == "../lintOk/" {
				fmt.Fprintf(os.Stdout, helmLintOk)
				os.Exit(0)
			} else if args[1] == "../lintInfo/" {
				fmt.Fprintf(os.Stdout, helmLintInfo)
				os.Exit(0)
			} else if args[1] == "../osError/" {
				fmt.Fprintf(os.Stderr, "this is a helm error")
				os.Exit(-1)
			} else {
				fmt.Fprintf(os.Stderr, "invalid input for testing")
				os.Exit(3)
			}
		} else if args[0] == "package" {
			if args[1] == "../goodchart/" {
				fmt.Fprintf(os.Stdout, helmPackageGoodChart)
				os.Exit(0)
			} else if args[1] == "../badchart/" {
				fmt.Fprintf(os.Stdout, helmPackageBadChart)
				os.Exit(1)
			} else if args[1] == "../osError/" {
				fmt.Fprintf(os.Stderr, "missing helm")
				os.Exit(1)
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

var helmLintManyErrors = `
==> Linting istio
[ERROR] Chart.yaml: version '0.w.0' is not a valid SemVer
[INFO] Chart.yaml: icon is recommended
[ERROR] templates/buttles.yaml: unable to parse YAML
        error unmarshaling JSON: while decoding JSON: json: cannot unmarshal string into Go value of type rules.K8sYamlStruct

Error: 1 chart(s) linted, 1 chart(s) failed
`

var helmLintOneError = `
==> Linting istio
[ERROR] Chart.yaml: version '0.w.0' is not a valid SemVer
[INFO] Chart.yaml: icon is recommended

Error: 1 chart(s) linted, 1 chart(s) failed
`

var helmLintOk = `
==> Linting buttles
Lint OK

1 chart(s) linted, no failures
`

var helmPackageBadChart = "Error: chart metadata (Chart.yaml) missing"

var helmPackageGoodChart = "Successfully packaged chart and saved it to: /Users/Chris.Evett/src/github/CamelotVG/sre-helm-charts/buttles-0.2.0.tgz"

var goodHelmPath = "/Users/Chris.Evett/src/github/CamelotVG/sre-helm-charts/buttles-0.2.0.tgz"
