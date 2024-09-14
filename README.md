# Aplicação de Teste

## Curso Descomplicando Terraform

### Instrutor: Rafael Gomes (Gomex)

Esta é uma aplicação simples em Go para servir como exemplo de produto para o projeto final do curso.

#### Instruções

Para compilar a aplicação e rodar a aplicação:

```
$ go build -o app main.go
$ ./app
Servidor rodando em http://<localhost ou ip da máquina>:8080

```

Para gerar a imagem Docker e rodar o container (pode-se usar o docker ou o podman):

```
$ docker build -t go-web-app .
[1/2] STEP 1/4: FROM golang:1.23.1-alpine AS builder
[1/2] STEP 2/4: WORKDIR /app
--> Using cache 2c10c0a5ccaf687c770315ca3b222bb0e8a65a16ef544140ee4cdcb075f5f0a5
--> 2c10c0a5ccaf
[1/2] STEP 3/4: COPY . .
--> 79c854a7a526
[1/2] STEP 4/4: RUN go build -o app
--> e5b74c4e9fb7
[2/2] STEP 1/5: FROM alpine:latest
[2/2] STEP 2/5: WORKDIR /app
--> Using cache 7df5113bdd79f2deb3ffccf1d539302fec43c11166300b75d04f610b5e9c32ff
--> 7df5113bdd79
[2/2] STEP 3/5: COPY --from=builder /app/app .
--> de46cfdd257c
[2/2] STEP 4/5: EXPOSE 8080
--> 813644312b7c
[2/2] STEP 5/5: CMD ["./app"]
[2/2] COMMIT go-web-app
--> 021a66c71d48
Successfully tagged localhost/go-web-app:latest
021a66c71d48fa2775f45ca9da0448c57672236cd86559f224476cd12421eba0

$ docker run -p 8080:8080 go-web-app
Servidor rodando em http://<localhost ou ip da máquina>:8080
```


