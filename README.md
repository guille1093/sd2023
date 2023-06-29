# Almacenamiento Clave Valor Escalable Distribuido

## Tarea

Cree un servidor y un cliente gRPC que se comuniquen correctamente entre sí cumplimentando lo solicitado en el enunciado del laboratorio.

## Metas

- Escribir un archivo proto que defina un servicio RPC
- Generar archivos go desde un archivo proto
- Enviar solicitudes de gRPC del cliente al servidor
- Los datos a almacenar en el almacenamiento de tipo clave-valor distribuido se deben obtener de un conjunto de nodos que envían datos usando el protocolo MQTT.
- El almacenamiento clave-valor debe almacenar asociado a cada clave un conjunto de valores.
- La clave se corresponde al id de cada nodo. Los valores están compuestos de una marca de tiempo y un valor aleatorio.
- El valor aleatorio es enviado por los nodos usando MQTT, la marca de tiempo es obtenida al momento de llegada del mensaje MQTT. Los nodos deben usar un tópico datos/id_nodo.
- Los nodos que conforman el sistema de almacenamiento clave-valor debe escuchar en el tópico datos/# al llegar un dato deben analizar si les corresponde almacenar la clave y si es asi guardarla.
- Por otro lado deben usar de base el cliente para realizar consultas a la base de datos.

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

