# GO

<!-- TOC -->
* [GO](#go)
  * [Overview](#overview)
* [🖥️ Technologies](#-technologies)
* [Packages y Adapters:](#packages-y-adapters)
* [🚀 Deploy](#-deploy)
  * [🪼 Github:](#-github)
<!-- TOC -->

## Overview

# 🖥️ Technologies

| Language     | Version  | Language | Platform | Package |                     Source                     |
|:-------------|:--------:|:--------:|----------|:-------:|:----------------------------------------------:|
| go           | go1.22.5 |    ✅     |          |         |                                                |
| ssh          | v0.25.0  |          |          |    ✅    |       [package](golang.org/x/crypto/ssh)       |
| sftp         | v1.13.6  |          |          |    ✅    |         [package](github.com/pkg/sftp)         |
| viper        | v1.19.0  |          |          |    ✅    |       [package](github.com/spf13/viper)        |
| mapstructure |  v2.0.0  |          |          |    ✅    | [package](github.com/go-viper/mapstructure/v2) |
| docker       |  4.33.1  |          |          |    ✅    |      [platform](https://www.docker.com/)       |

# Packages y Adapters:

- **fundamentals**: Programming basic language Golang.
- **ssh**:
  El paquete ssh implementa un cliente y un servidor SSH.
  SSH es un protocolo de seguridad de transporte, un protocolo de autenticación y una familia de protocolos de
  aplicación. El protocolo de nivel de aplicación más típico es un shell remoto y este se implementa específicamente.
  Sin embargo, la naturaleza multiplexada de SSH está expuesta a los usuarios que desean brindar soporte a otros.
- **sftp**:
  El sftp paquete proporciona soporte para operaciones del sistema de archivos en servidores ssh remotos que utilizan el
  subsistema SFTP. También implementa un servidor SFTP para servir archivos desde el sistema de archivos.

# 🚀 Deploy

## 🪼 Github:

GitHub es una plataforma en línea para alojar y gestionar código fuente usando Git, un sistema de control de versiones.
Permite a los desarrolladores colaborar en proyectos, gestionar versiones del código, realizar revisiones y automatizar
flujos de trabajo. Sus principales características incluyen repositorios de código, control de versiones, pull requests,
issues, y GitHub Actions para automatización. También ofrece herramientas para la documentación y la creación de sitios
web estáticos. Es ampliamente utilizada para facilitar la colaboración y organización en el desarrollo de software.

GitHub Actions:

GitHub Actions es una herramienta integrada en GitHub para automatizar flujos de trabajo de desarrollo, como la
integración continua (CI) y la entrega continua (CD). Permite definir y ejecutar tareas automáticamente en respuesta a
eventos en un repositorio, como "push" o solicitudes de extracción.