
# SSL Commands

SERVER_CN=localhost

Generar el certificado de autoridad, Generate Certificate Authority

```shell
openssl genrsa --passout pass:1111 -des3 -out ca.key 4069
```
Genera del certificado de confianza: Trust Certificate (ca.crt)

```shell
openssl req --passin pass:1111 -new -x509 -days 100 -key ca.key -out ca.crt -subj "//CN=${SERVER_CN}"
```

Genera la llave privada del servidor

```shell
openssl genrsa -passout pass:1111 -des3 -out server.key 4096
```

Genera certificado para firmar peticiones

```shell
openssl req --passin pass:1111 -new -key server.key -out server.csr -subj "//CN=${SERVER_CN}"
```

Firma el certificado con el CA

```shell
openssl x509 -req -passin pass:1111 -days 100 -in server.csr -CA ca.crt -CAkey ca.key -set_serial 01 -out server.crt
```

Convertir el certificado a un formato pem que podemos usar en grpc

```shell
openssl pkcs8 -topk8 -nocrypt -passin pass:1111 -in server.key -out server.pem
```