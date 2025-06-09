package test

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

func waitForStatusOK(url string, retries int, delay time.Duration, t *testing.T) {
	for i := 0; i < retries; i++ {
		resp, err := http.Get(url)
		if err == nil && resp.StatusCode == 200 {
			resp.Body.Close()
			return
		}
		if resp != nil {
			resp.Body.Close()
		}
		time.Sleep(delay)
	}
	t.Fatalf("Did not get HTTP 200 from %s after %d retries", url, retries)
}

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
		waitForStatusOK(url, 10, 3*time.Second, t)
	})

	t.Run("WebServer2", func(t *testing.T) {
		const expectedPort = 8082
		url := fmt.Sprintf("http://localhost:%d", expectedPort)
		waitForStatusOK(url, 10, 3*time.Second, t)
	})
}
