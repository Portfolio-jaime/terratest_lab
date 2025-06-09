variable "container_name" {
  description = "El nombre del contenedor Docker."
  type        = string
}

variable "external_port" {
  description = "El puerto externo a mapear al puerto 80 del contenedor."
  type        = number
}

variable "image_id" {
  description = "El ID de la imagen Docker a utilizar."
  type        = string
}