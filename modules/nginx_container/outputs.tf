output "container_name" {
  description = "El nombre del contenedor creado."
  value       = docker_container.nginx_server.name
}

output "external_port" {
  description = "El puerto externo del contenedor."
  value       = var.external_port
}