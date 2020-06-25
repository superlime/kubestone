package iperf2

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestIperf2Controller(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Iperf2 Controller Suite")
}
