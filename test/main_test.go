package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestTerraformDockerModularExample(t *testing.T) {
	t.Parallel()

	// Configura las opciones de Terraform para apuntar a nuestro código raíz
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// La ruta a nuestro código de Terraform
		TerraformDir: "../terraform",

		// Variables a pasar a nuestra configuración de Terraform usando -var-file
		VarFiles: []string{"terraform.tfvars"},
	})

	// Al final de la prueba, ejecuta 'terraform destroy' para limpiar
	defer terraform.Destroy(t, terraformOptions)

	// Ejecuta 'terraform init' y 'terraform apply'
	terraform.InitAndApply(t, terraformOptions)

	// Pruebas para el primer servidor web
	t.Run("WebServer1", func(t *testing.T) {
		const expectedName = "web-server-1"
		const expectedPort = 8081
		url := fmt.Sprintf("http://localhost:%d", expectedPort)

		// Verifica que el servidor Nginx responde con código 200 OK
		http_helper.HttpGetWithRetry(t, url, nil, 200, "Welcome to nginx!", 10, 3*time.Second)
	})

	// Pruebas para el segundo servidor web
	t.Run("WebServer2", func(t *testing.T) {
		const expectedName = "web-server-2"
		const expectedPort = 8082
		url := fmt.Sprintf("http://localhost:%d", expectedPort)

		// Verifica que el servidor Nginx responde con código 200 OK
		http_helper.HttpGetWithRetry(t, url, nil, 200, "Welcome to nginx!", 10, 3*time.Second)
	})
}
