package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/docker"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformDockerExample(t *testing.T) {
	t.Parallel()

	// Configura las opciones de Terraform para apuntar a nuestro código
	terraformOptions := &terraform.Options{
		// La ruta a nuestro código de Terraform
		TerraformDir: "../terraform",
	}

	// Al final de la prueba, ejecuta 'terraform destroy' para limpiar los recursos
	defer terraform.Destroy(t, terraformOptions)

	// Ejecuta 'terraform init' y 'terraform apply'
	terraform.InitAndApply(t, terraformOptions)

	// Valida que el contenedor Docker se está ejecutando
	opts := &docker.RunOptions{Command: []string{"ps", "-a"}}
	output := docker.Run(t, opts)
	assert.Contains(t, output, "tutorial")
}
