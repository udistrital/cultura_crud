# cultura_crud

API CRUD para la gestión de información de las actividades culturales realizadas por el Bienestar institucional

## Especificaciones Técnicas

### Tecnologías Implementadas y Versiones

- [Golang](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/golang.md)
- [BeeGo](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/beego.md)
- [Docker](https://docs.docker.com/engine/install/ubuntu/)
- [Docker Compose](https://docs.docker.com/compose/)

### Variables de Entorno

```shell
# parametros de api
CULTURA_CRUD_HTTP_PORT=[Puerto de exposición del API]
CULTURA_CRUD_RUN_MODE=[Modo de ejecución del API]
# paramametros de bd #
CULTURA_CRUD_PGUSER=[Usuario de BD]
CULTURA_CRUD_PGPASS=[Contraseña del usaurio de BD]
CULTURA_CRUD_PGHOST=[URL, Dominio o EndPoint de la BD]
CULTURA_CRUD_PGPORT=[Puerto de la BD]
CULTURA_CRUD_PGDB=[Nombre de Base de Datos]
PARAMETROCULTURA_CRUD_PGSCHEMA=[Nombre del Esquema de Base de Datos]
```

**NOTA:** Las variables se pueden ver en el fichero conf/app.conf y están identificadas con `CULTURA_CRUD_...`

### Ejecución del Proyecto

```shell
#1. Obtener el repositorio con Go
go get github.com/udistrital/cultura_crud

#2. Moverse a la carpeta del repositorio
cd $GOPATH/src/github.com/udistrital/cultura_crud

# 3. Moverse a la rama **develop**
git pull origin develop && git checkout develop

# 4. alimentar todas las variables de entorno que utiliza el proyecto.
CULTURA_CRUD_HTTP_PORT=8080 CULTURA_CRUD_PGHOST=127.0.0.1:27017 CULTURA_CRUD_SOME_VARIABLE=some_value bee run
```

### Ejecución Dockerfile

```shell
# Implementado para despliegue del Sistema de integración continua CI.
```

### Ejecución docker-compose

```shell
#1. Clonar el repositorio
git clone -b develop https://github.com/udistrital/cultura_crud

#2. Moverse a la carpeta del repositorio
cd cultura_crud

#3. Crear un fichero con el nombre **custom.env**
touch custom.env

#4. Crear la network **back_end** para los contenedores
docker network create back_end

#5. Ejecutar el compose del contenedor
docker-compose up --build

#6. Comprobar que los contenedores estén en ejecución
docker ps
```

### Ejecución Pruebas

Pruebas unitarias

```shell
# En Proceso
```

## Modelo de Datos

[PGmodeler](database/cultura.dbm) -
[SVG](database/cultura.svg)

## Estado CI

| Develop | Release 1.0.0 | Master |
| -- | -- | -- |
| [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/cultura_crud/status.svg?ref=refs/heads/develop)](https://hubci.portaloas.udistrital.edu.co/udistrital/cultura_crud) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/cultura_crud/status.svg?ref=refs/heads/release/1.0.0)](https://hubci.portaloas.udistrital.edu.co/udistrital/cultura_crud) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/cultura_crud/status.svg?ref=refs/heads/master)](https://hubci.portaloas.udistrital.edu.co/udistrital/cultura_crud) |

## Licencia

This file is part of cultura_crud.

cultura_crud is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

cultura_crud is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with parametros_crud. If not, see https://www.gnu.org/licenses/.
