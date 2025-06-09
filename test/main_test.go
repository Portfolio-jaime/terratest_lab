package test

import (
	"fmt"
	"testing"
	"time"

	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestTerraformDockerModularExample(t *testing.T) {
	t.Parallel()

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../terraform",
		VarFiles:     []string{"terraform.tfvars"},
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	t.Run("WebServer1", func(t *testing.T) {
		const expectedPort = 8081
		url := fmt.Sprintf("http://localhost:%d", expectedPort)
		httph := http_helper.HttpGetWithRetry
		httph(t, url, nil, 200, "Welcome to nginx!", 10, 3*time.Second)
		http_helper.HttpGetWithRetry(t, url, nil, 200, "Welcome to nginx!", 10, 3*time.Second)
	})

	t.Run("WebServer2", func(t *testing.T) {
		const expectedPort = 8082
		url := fmt.Sprintf("http://localhost:%d", expectedPort)
		http_helper.HttpGetWithRetry(t, url, nil, 200, "Welcome to nginx!", 10, 3*time.Second)
	})
}
