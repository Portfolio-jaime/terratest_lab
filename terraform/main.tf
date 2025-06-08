terraform {
  required_providers {
    docker = {
      source  = "kreuzwerker/docker"
      version = "~> 3.0.1"
    }
  }
}

provider "docker" {}

module "nginx_servers" {
  for_each = var.servers
  source   = "../modules/nginx_container"

  container_name = each.key
  external_port  = each.value.port
}