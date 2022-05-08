package e2e_test

import (
	"fmt"
	"net/http"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestE2e(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "E2e Suite")
}

var _ = BeforeSuite(func() {
	appURL := "http://localhost:8081/helloworld"
	resp, err := http.Get(appURL)
	Expect(err).ShouldNot(HaveOccurred())
	defer resp.Body.Close()
	Expect(resp).Should(HaveHTTPStatus(http.StatusOK))
	fmt.Println("Test pre cecks successfully completed for running e2e test")
})
