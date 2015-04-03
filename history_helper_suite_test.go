package history_helper_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestHistoryHelper(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "HistoryHelper Suite")
}
