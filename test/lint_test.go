package test_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	lint "yaml-linter/internal"
)

const (
	assetsDir  = "./cicd-test-assets"
	schemaName = "../internal/schemas/cicd-schema.json"
)

var (
	validator  = lint.NewValidator()
	testAssets = map[string]string{
		"build-docker":                           assetsDir + "/build-docker.yaml",
		"deploy-ansible":                         assetsDir + "/deploy-ansible.yaml",
		"deploy-ansible-with-testrail-reporting": assetsDir + "/deploy-ansible-with-testrail-reporting.yaml",
		"publish-docker":                         assetsDir + "/publish-docker.yaml",
		"notify-bitbucket":                       assetsDir + "/notify-bitbucket.yaml",
	}
)

var _ = Describe("Linter Test Suite for cicd-schema.json", func() {

	BeforeEach(func() {
		validator = lint.NewValidator()
		validator.SchemaPath = schemaName
	})

	Describe("Build", func() {
		Describe("Build using Docker Driver", func() {
			It("should pass when spec is valid", func() {
				validator.DataPath = testAssets["build-docker"]
				Expect(validator.Validate()).To(ContainSubstring("SUCCESS"))
			})
		})
	})

	Describe("Deploy", func() {
		Describe("Deploy using Ansible Driver", func() {
			It("should pass when spec is valid", func() {
				validator.DataPath = testAssets["deploy-ansible"]
				Expect(validator.Validate()).To(ContainSubstring("SUCCESS"))
			})
		})
	})

	Describe("Deploy", func() {
		Describe("Deploy using Ansible With Testrail Reporting Driver", func() {
			It("should pass when spec is valid", func() {
				validator.DataPath = testAssets["deploy-ansible-with-testrail-reporting"]
				Expect(validator.Validate()).To(ContainSubstring("SUCCESS"))
			})
		})
	})

	Describe("Publish", func() {
		Describe("Publish using Artifactory Docker Driver", func() {
			It("should pass when spec is valid", func() {
				validator.DataPath = testAssets["publish-docker"]
				Expect(validator.Validate()).To(ContainSubstring("SUCCESS"))
			})
		})
	})

	Describe("Notify", func() {
		Describe("Notify using Bitbucket Driver", func() {
			It("should pass when spec is valid", func() {
				validator.DataPath = testAssets["notify-bitbucket"]
				Expect(validator.Validate()).To(ContainSubstring("SUCCESS"))
			})
		})
	})

})
