# Etapa 1: Construir a aplicação
FROM golang:1.23.1-alpine AS builder

# Definir o diretório de trabalho dentro do container
WORKDIR /app

# Copiar os arquivos de código-fonte para o diretório de trabalho
COPY . .

# Construir a aplicação
RUN go build -o app

# Etapa 2: Criar a imagem final e rodar a aplicação
FROM alpine:latest

# Definir o diretório de trabalho para o container final
WORKDIR /app

# Copiar o binário construído da etapa anterior
COPY --from=builder /app/app .

# Expor a porta 8080
EXPOSE 8080

# Comando para rodar a aplicação
CMD ["./app"]