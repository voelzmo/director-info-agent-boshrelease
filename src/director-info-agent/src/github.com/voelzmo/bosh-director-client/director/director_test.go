package director_test

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/voelzmo/bosh-director-client/director"
)

var dir director.Director

var _ = Describe("Director", func() {
	Context("When initialized with a target", func() {

		BeforeEach(func() {
			target := os.Getenv("BDC_ITEST_TARGET")
			rootCAPath := os.Getenv("BDC_ITEST_ROOT_CA_PATH")
			clientName := os.Getenv("BDC_ITEST_CLIENT_NAME")
			clientSecret := os.Getenv("BDC_ITEST_CLIENT_SECRET")
			dir = director.NewDirector(target, rootCAPath, clientName, clientSecret)
		})

		It("#status gets the director's status ", func() {
			status := dir.Status()
			Expect(status.Name).To(Equal("my-bosh"))
		})

		It("#login returns a token and not an error", func() {
			login := dir.Login()
			Expect(login.Error).To(BeEmpty())
			Expect(login.AccessToken).NotTo(BeEmpty())
		})

		It("#deployments return a list of deployments with their stemcells", func() {
			deployments := dir.Deployments()
			Expect(deployments).NotTo(BeEmpty())
			Expect(deployments[0].Name).To(Equal("test"))
		})
	})
})
