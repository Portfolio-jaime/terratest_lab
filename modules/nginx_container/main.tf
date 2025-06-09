terraform {
  required_providers {
    docker = {
      source  = "kreuzwerker/docker"
      version = "~> 3.0.1"
    }
  }
}

# --- RECURSO docker_image ELIMINADO DE AQUÍ ---

resource "docker_container" "nginx_server" {
  # --- CAMBIO AQUÍ ---
  image = var.image_id
  name  = var.container_name
  ports {
    internal = 80
    external = var.external_port
  }
}