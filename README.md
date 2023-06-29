# Laboratorio 4 – Almacenamiento Clave Valor Escalable

## Tarea

Cree un servidor y un cliente gRPC que se comuniquen correctamente entre sí cumplimentando lo solicitado en el enunciado del laboratorio.

## Metas

- Escribir un archivo proto que defina un servicio RPC
- Generar archivos go desde un archivo proto
- Enviar solicitudes de gRPC del cliente al servidor

### Parte 1:
1. Instalar protoc en su máquina local.
2. Cree su archivo protobuf. Asegúrese de nombrar el paquete igual que sus otros archivos.
3. Ejecute <code>make build</code> en el directorio del proyecto para generar los archivos go.


### Parte 2: Servidor y cliente

Implementar lo faltante.

## Construir binarios

- Vaya a `cmd/cliente` o `cmd/servidor` y ejecute `go build .`. Esto generará archivos binarios de cliente o servidor respectivamente.

## Pruebas

- `go test ./...` debería ejecutar todas las pruebas que desarrollen, las mismas deben verificar los casos planteados en el laboratorio.

