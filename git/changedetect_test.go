package git

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var gitDiffFileParentOneChange = `
diff --git a/Makefile b/Makefile
deleted file mode 100644
index f867b0b..0000000
--- a/Makefile
+++ /dev/null
@@ -1,15 +0,0 @@
-CHARTSOURCES= $(shell find -L . -name 'Chart.yaml')
-CHARTDIRS= $(dir $(CHARTSOURCES))
-
-
-.PHONY: all
-all: 
-       foreach( var,list,text)
-       for DIR in $(CHARTDIRS) ; do \
-         make lint $(DIR); \
-       done
-
-.PHONY: lint 
-lint: 
-       helm lint $(DIR
-
diff --git a/istio/Chart.yaml b/istio/Chart.yaml
index 891c603..27ae81c 100644
--- a/istio/Chart.yaml
+++ b/istio/Chart.yaml
@@ -2,4 +2,4 @@ apiVersion: v1
 appVersion: "1.0"
 description: A Helm chart for Kubernetes
 name: istio
-version: 0.1.0
+version: 0.2.0
`

var gitDiffOneFolderGithubAction = `
diff --git a/.github/workflows/helmsync.yaml b/.github/workflows/helmsync.yaml
new file mode 100644
index 0000000..6a23728
--- /dev/null
+++ b/.github/workflows/helmsync.yaml
@@ -0,0 +1,24 @@
+name: Docker Image CI
+
+on: [push]
+
+jobs:
+  build:
+    runs-on: ubuntu-latest
+    steps:
+    - uses: actions/checkout@v1
+    - name: Generate build number
+      id: buildnumber
+      uses: einaregilsson/build-number@v1 
+      with:
+        # the github_token is built in
+        token: ${{secrets.github_token}}
+    - name: Login to ECR
+      id: ecr
+      uses: jwalton/gh-ecr-login@v1
+      with:
+        access-key-id: ${{ secrets.AWS_ACCESS_KEY_ID }}
+        secret-access-key: ${{ secrets.AWS_SECRET_ACCESS_KEY }}
+        region: us-west-1
+    - name: run helmsync
+      run: docker run -v $GITHUB_WORKSPACE:/code -e REPOPATH="/code" -e VERSIONNUMBER="0.1.${{ steps.buildnumber.outputs.build_number}}" -e ARTIFACTORYURL="${{ secrets.ARTIFACTORYURL }}" -e IGNOREINFO="TRUE" ${{ steps.ecr.outputs.account }}.dkr.ecr.us-west-1.amazonaws.com/sre/helmsync:0.1.15
diff --git a/istio/Chart.yaml b/istio/Chart.yaml
index 891c603..27ae81c 100644
--- a/istio/Chart.yaml
+++ b/istio/Chart.yaml
@@ -2,4 +2,4 @@ apiVersion: v1
 appVersion: "1.0"
 description: A Helm chart for Kubernetes
 name: istio
-version: 0.1.0
+version: 0.2.0
`

var gitDiffOneFolderChangedOutput = `
diff --git a/buttles/Chart.yaml b/buttles/Chart.yaml
index c708934..7eaf474 100644
--- a/buttles/Chart.yaml
+++ b/buttles/Chart.yaml
@@ -2,4 +2,4 @@ apiVersion: v1
 appVersion: "1.0"
 description: A Helm chart for Kubernetes
 name: buttles
-version: 0.1.0
+wersion: 0.2.0
`
var gitDiffManyFolderChangedOutput = `
diff --git a/buttles/Chart.yaml b/buttles/Chart.yaml
index c708934..7eaf474 100644
--- a/buttles/Chart.yaml
+++ b/buttles/Chart.yaml
@@ -2,4 +2,4 @@ apiVersion: v1
 appVersion: "1.0"
 description: A Helm chart for Kubernetes
 name: buttles
-version: 0.1.0
+wersion: 0.2.0
diff --git a/istio/Chart.yaml b/istio/Chart.yaml
index 891c603..2eb9568 100644
--- a/istio/Chart.yaml
+++ b/istio/Chart.yaml
@@ -2,4 +2,4 @@ apiVersion: v1
 appVersion: "1.0"
 description: A Helm chart for Kubernetes
 name: istio
-version: 0.1.0
+version: 0.w.0
diff --git a/istio/templates/buttles.yaml b/istio/templates/buttles.yaml
new file mode 100644
index 0000000..54e4d23
--- /dev/null
+++ b/istio/templates/buttles.yaml
@@ -0,0 +1 @@
+ksflksdjfdsfkjldfslkjdsfkj
`
var gitDiffNoFolderChangedOutput = ""

// copying and pasting code from old projects
// https://github.com/rendicott/stool/blob/master/verifiers/inspec/verifier.go
func fakeExecCommand(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestHelperProcess", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1"}
	return cmd
}

var _ = Describe("Changedetect", func() {

})

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
	case "git":
		if args[0] == "-C" {
			if args[1] == "../testOneChange/" {
				fmt.Fprintf(os.Stdout, gitDiffOneFolderChangedOutput)
				os.Exit(0)
			} else if args[1] == "../testManyChanges/" {
				fmt.Fprintf(os.Stdout, gitDiffManyFolderChangedOutput)
				os.Exit(0)
			} else if args[1] == "../testZeroChanges/" {
				fmt.Fprintf(os.Stdout, gitDiffNoFolderChangedOutput)
				os.Exit(0)
			} else if args[1] == "../testDotInPath/" {
				fmt.Fprintf(os.Stdout, gitDiffOneFolderGithubAction)
				os.Exit(0)
			} else if args[1] == "../testParentFileChange/" {
				fmt.Fprintf(os.Stdout, gitDiffFileParentOneChange)
				os.Exit(0)
			} else if args[1] == "../testGitError/" {
				fmt.Fprintf(os.Stderr, "this is a git error")
				os.Exit(127)
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

var _ = Describe("Given we are diffing a git repo", func() {
	Context("when there is one change", func() {
		It("returns one changed folder", func() {
			execCommand = fakeExecCommand
			defer func() { execCommand = exec.Command }()

			folders, err := ChangedFolders("../testDotInPath/")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(folders).Should(Equal([]string{"istio"}))
		})
	})
	Context("when there is one change wtih a .github change", func() {
		It("returns one changed folder", func() {
			execCommand = fakeExecCommand
			defer func() { execCommand = exec.Command }()

			folders, err := ChangedFolders("../testOneChange/")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(folders).Should(Equal([]string{"buttles"}))
		})
	})
	Context("when there is one change to a parent file eg readme.md and a change", func() {
		It("returns one changed folder and ignores the file", func() {
			execCommand = fakeExecCommand
			defer func() { execCommand = exec.Command }()

			folders, err := ChangedFolders("../testParentFileChange/")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(folders).Should(Equal([]string{"istio"}))
		})
	})
	Context("when there are many changes", func() {
		It("returns a deduplicated list of all folders with changes", func() {
			execCommand = fakeExecCommand
			defer func() { execCommand = exec.Command }()

			folders, err := ChangedFolders("../testManyChanges/")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(folders).Should(Equal([]string{"buttles", "istio"}))
		})
	})
	Context("when there are no changes", func() {
		It("returns no changed", func() {
			execCommand = fakeExecCommand
			defer func() { execCommand = exec.Command }()

			folders, err := ChangedFolders("../testZeroChanges/")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(folders).Should(Equal([]string{}))
		})
	})
	Context("when there is an error", func() {
		It("returns returns an error message", func() {
			execCommand = fakeExecCommand
			defer func() { execCommand = exec.Command }()

			folders, err := ChangedFolders("../testGitError/")
			Expect(err).Should(HaveOccurred())
			Expect(folders).Should(Equal([]string{}))
		})
	})
})

var _ = Describe("gitDiff", func() {
	Context("when a git command is successful", func() {
		It("returns the stdout from git", func() {
			execCommand = fakeExecCommand
			defer func() { execCommand = exec.Command }()

			stdout, err := gitDiff("../testOneChange/")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(stdout).Should(Equal(gitDiffOneFolderChangedOutput))
		})
	})
	Context("when a git command errors", func() {
		It("returns the error code and the stderr", func() {
			execCommand = fakeExecCommand
			defer func() { execCommand = exec.Command }()

			stderr, err := gitDiff("../testGitError/")
			Expect(err).Should(HaveOccurred())
			Expect(stderr).Should(Equal("this is a git error"))
		})
	})
})

var _ = Describe("parseFolders", func() {
	Context("when no folders are present", func() {
		It("returns an empty slice", func() {
			execCommand = fakeExecCommand
			defer func() { execCommand = exec.Command }()
			expected := []string{}

			actual := parseFolders(gitDiffNoFolderChangedOutput)
			Expect(actual).Should(Equal(expected))
		})
	})
	Context("when one folder is present", func() {
		It("returns a single folder name", func() {
			execCommand = fakeExecCommand
			defer func() { execCommand = exec.Command }()
			expected := []string{"buttles"}

			actual := parseFolders(gitDiffOneFolderChangedOutput)
			Expect(actual).Should(Equal(expected))
		})
	})
	Context("when one folder is present when a github action change is made", func() {
		It("returns a single folder name", func() {
			execCommand = fakeExecCommand
			defer func() { execCommand = exec.Command }()
			expected := []string{"istio"}

			actual := parseFolders(gitDiffOneFolderGithubAction)
			Expect(actual).Should(Equal(expected))
		})
	})
	Context("when many folders are present", func() {
		It("returns each folder without duplicates", func() {
			execCommand = fakeExecCommand
			defer func() { execCommand = exec.Command }()
			expected := []string{"buttles", "istio"}

			actual := parseFolders(gitDiffManyFolderChangedOutput)
			Expect(actual).Should(Equal(expected))
		})
	})
})
