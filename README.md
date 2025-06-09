# Laboratorio Local de Terraform y Terratest con Docker 🐳

Este proyecto demuestra cómo construir y probar infraestructura como código (IaC) en un entorno local utilizando **Terraform** para la definición de la infraestructura, **Docker** como proveedor local y **Terratest** para las pruebas automatizadas.

La configuración es modular, lo que permite reutilizar y gestionar los componentes de la infraestructura de manera eficiente.

---

## 🤔 ¿Por qué este enfoque?

- **Aprendizaje sin costo:** Permite experimentar con Terraform y Terratest sin necesidad de una cuenta en un proveedor de nube (AWS, GCP, Azure), eliminando cualquier riesgo de facturación inesperada.
- **Desarrollo rápido y local:** Probar los cambios de infraestructura en tu propia máquina es mucho más rápido que desplegarlos en un entorno de nube remoto. Esto acelera el ciclo de desarrollo y depuración.
- **Pruebas de confianza:** Implementa un flujo de **Pruebas de Infraestructura**, donde puedes verificar que tu código de Terraform funciona como se espera antes de aplicarlo en producción.
- **Modularidad y reutilización:** Organizar el código de Terraform en módulos es una buena práctica que fomenta la reutilización, la mantenibilidad y la escalabilidad de tus proyectos de IaC.

---

## 🛠️ Tecnologías Utilizadas

- **Terraform:** Herramienta de código abierto para construir, cambiar y versionar infraestructura de manera segura y eficiente.
- **Terratest:** Biblioteca de Go para escribir pruebas automatizadas para tu código de infraestructura.
- **Docker:** Plataforma para desarrollar, enviar y ejecutar aplicaciones en contenedores. Actúa como nuestro "proveedor de nube" local.
- **Go:** Lenguaje de programación utilizado para escribir las pruebas de Terratest.

---

## 🚀 Guía de Implementación

### Requisitos Previos

- [Docker](https://www.docker.com/products/docker-desktop)
- [Terraform](https://learn.hashicorp.com/tutorials/terraform/install-cli)
- [Go (versión 1.18+)](https://golang.org/doc/install)

### Pasos para la Ejecución

1. **Clona el repositorio:**
    ```bash
    git clone <URL-del-repositorio>
    cd terraform-docker-modular-lab
    ```

2. **Inicializa el módulo de Go:**
    ```bash
    cd test
    go mod init terraform-docker-modular-lab/test
    go mod tidy
    ```

3. **Ejecuta las pruebas:**
    ```bash
    go test -v
    ```

---

## 🧑‍💻 Guía para Implementar en tu Proyecto

1. **Copia la estructura de carpetas** (`terraform`, `test`, etc.) a tu repositorio.
2. **Adapta los módulos de Terraform** según tus necesidades (puertos, imágenes, variables).
3. **Modifica o agrega pruebas en Go** en el directorio `test` para validar los recursos que implementes.
4. **Ejecuta las pruebas localmente** con `go test -v` antes de hacer push.
5. **Integra el flujo en tu CI/CD** (GitHub Actions, GitLab CI, etc.) para validar cada cambio automáticamente.

---

## ✅ Resultados Esperados

Al ejecutar `go test`, deberías observar el siguiente flujo en tu terminal:

1. Terratest invoca a `terraform init` y `terraform apply`.
2. Terraform, usando el módulo `nginx_container`, crea dos contenedores Docker de Nginx:
    - `web-server-1` expuesto en el puerto `8081`.
    - `web-server-2` expuesto en el puerto `8082`.
3. Terratest realiza una petición HTTP a `http://localhost:8081` y `http://localhost:8082` para verificar que los servidores Nginx responden correctamente.
   - **La validación solo comprueba que el código de estado HTTP sea 200, sin importar el contenido del body.**
4. La prueba confirma que los nombres de los contenedores son los esperados.
5. Finalmente, Terratest invoca a `terraform destroy` para eliminar los contenedores y limpiar el entorno.
6. La salida de la prueba mostrará **`PASS`**, indicando que todo funcionó correctamente.

---

## ❌ Casos Comunes de Error

- **Docker no está corriendo:**  
  El test fallará con errores de conexión. Solución: asegúrate de que el servicio Docker esté activo (`sudo systemctl start docker` o abre Docker Desktop).

- **El puerto ya está en uso:**  
  Si tienes otro servicio usando el puerto 8081 o 8082, Terraform fallará al crear el contenedor. Solución: libera los puertos o cambia los valores en las variables de Terraform.

- **No tienes permisos para usar Docker:**  
  Verás errores de permisos. Solución: agrega tu usuario al grupo `docker` o ejecuta con `sudo`.

- **Fallo en la prueba HTTP:**  
  Si el contenedor Nginx no inicia correctamente, la prueba de Terratest fallará tras varios reintentos. Solución: revisa los logs del contenedor con `docker logs <container_id>`.

- **Variables o archivos faltantes:**  
  Si falta `terraform.tfvars` o alguna variable obligatoria, Terraform fallará. Solución: revisa que todos los archivos requeridos estén presentes.

---

## 📈 Diagrama del Flujo

```mermaid
graph TD
    subgraph "Máquina Local"
        A[Usuario ejecuta 'go test'] --> B{Terratest};
        B -->|1. terraform init & apply| C[Terraform Core];
        C -->|2. Usa el provider de Docker| D[Docker Engine];
        D -->|3. Crea Contenedor Nginx 1| E[Container Nginx 'web-server-1'];
        D -->|4. Crea Contenedor Nginx 2| F[Container Nginx 'web-server-2'];
        B -->|5. Realiza pruebas| G[Verifica contenedores vía API Docker];
        G -->|6. OK/FAIL| B;
        B -->|7. terraform destroy| C;
        C -->|8. Elimina contenedores| D;
    end

    style A fill:#f9f,stroke:#333,stroke-width:2px
    style B fill:#ccf,stroke:#333,stroke-width:2px
    style C fill:#bbf,stroke:#333,stroke-width:2px
```

---