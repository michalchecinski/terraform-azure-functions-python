package test

import (
	"os"
	"testing"

	"math/rand"

	"context"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/web/mgmt/web"
	"github.com/gruntwork-io/terratest/modules/azure"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/marstr/randname"
	"github.com/stretchr/testify/assert"
)

var (
	rg_name                    = randname.GenerateWithPrefix("tf-tests", 10)
	function_app_name          = randname.GenerateWithPrefix("tf-func-tests", 10)
	function_service_plan_name = randname.GenerateWithPrefix("tf-plan-tests", 10)
	function_storage_name      = randomString("tfstfunctests", 6)
	function_app_insights_name = randname.GenerateWithPrefix("tf-func-tests-ins", 10)
	storage_name               = randomString("tfstests", 10)
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

	client, err := getWebAppsClient()
	if err != nil {
		return
	}

	funcApp, err := client.Get(context.Background(), rg_name, function_app_name)
	if err != nil {
		return
	}

	actualIdentityType := funcApp.Identity.Type
	expectedIdentityType := web.ManagedServiceIdentityType("SystemAssigned")

	assert.Equal(t, expectedIdentityType, actualIdentityType)

	appSettings := funcApp.SiteConfig.AppSettings
	for i := 0; i < len(appSettings); i++ {
		if appSettings[i].Name == "storage_name" {
			actualStorageName := appSettings[i].Value
			break
		}
	}

	assert.Equal(t, storage_name, actualStorageName)
}

func randomString(s string, n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyz0123456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return s + string(b)
}

func getWebAppsClient() (client web.AppsClient, err error) {
	subscriptionId, found := os.LookupEnv("ARM_SUBSCRIPTION_ID")
	if found != true {
		return
	}
	client = web.NewAppsClient(subscriptionId)
	authorizer, err := azure.NewAuthorizer()
	if err != nil {
		return
	}
	client.Authorizer = *authorizer
	if err != nil {
		return
	}
	return
}
