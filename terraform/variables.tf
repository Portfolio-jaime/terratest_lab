variable "servers" {
  description = "Un mapa de servidores web a crear."
  type = map(object({
    port = number
  }))
  default = {}
}