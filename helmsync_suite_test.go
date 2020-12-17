package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestHelmsync(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Helmsync Suite")
}
