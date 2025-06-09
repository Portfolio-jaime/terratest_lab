terraform {
  required_providers {
    docker = {
      source  = "kreuzwerker/docker"
      version = "~> 3.0.1"
    }
  }
}

provider "docker" {}

# --- AÑADIMOS EL RECURSO DE IMAGEN AQUÍ ---
resource "docker_image" "nginx" {
  name         = "nginx:latest"
  keep_locally = true # Cambiamos a 'true' para que no intente borrarla en cada 'apply'
}

module "nginx_servers" {
  for_each = var.servers
  source   = "../modules/nginx_container"

  container_name = each.key
  external_port  = each.value.port
  # --- CAMBIO AQUÍ: Pasamos el ID de la imagen ---
  image_id       = docker_image.nginx.image_id
}