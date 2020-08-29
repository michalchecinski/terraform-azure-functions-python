package test

import (
	"testing"

	"math/rand"

	"github.com/marstr/randname"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

var (
	rg_name                    = randname.GenerateWithPrefix("tf-tests", 10)
	function_app_name          = randname.GenerateWithPrefix("tf-func-tests", 10)
	function_service_plan_name = randname.GenerateWithPrefix("tf-plan-tests", 10)
	function_storage_name      = RandomString("tfstfunctests", 6)
	function_app_insights_name = randname.GenerateWithPrefix("tf-func-tests-ins", 10)
	storage_name               = RandomString("tfstests", 10)
)

func TestTerraform(t *testing.T) {

	azureRegions := []string{"westeurope", "North Europe", "East US", "Central US"}
	randomIndex := rand.Intn(len(azureRegions))

	terraformOptions := &terraform.Options{
		TerraformDir: "..",
		Vars: map[string]interface{}{
			"location":                   azureRegions[randomIndex],
			"rg_name":                    rg_name,
			"function_app_name":          function_app_name,
			"function_service_plan_name": function_service_plan_name,
			"function_storage_name":      function_storage_name,
			"function_app_insights_name": function_app_insights_name,
			"storage_name":               storage_name,
		},
	}

	// Clean up resources with "terraform destroy" at the end of the test.
	defer terraform.Destroy(t, terraformOptions)

	// Run "terraform init" and "terraform apply". Fail the test if there are any errors.
	terraform.InitAndApply(t, terraformOptions)
}

func RandomString(s string, n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyz0123456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return s + string(b)
}
