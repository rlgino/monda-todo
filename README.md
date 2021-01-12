# TODO List

## Arquitectura

### App
Frontend hecho con ReactJS que consulta directo al backend

### src/api
Backend hecho en Golang puro y duro. Expone inicialmente 2 bounded contexts: task y user.
Estos BC están basados en una arquitectura hexagonal, cada uno con su respectivas capas de dominio,
aplicación e infraestructura.

## Stack
* Golang
* ReactJS