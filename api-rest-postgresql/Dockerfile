# Usa la imagen base de Go
FROM golang:latest

# Establece el directorio de trabajo
WORKDIR /app

# Copia los archivos go.mod y go.sum primero para aprovechar el cache de Docker
COPY go.mod ./
COPY go.sum ./

# Descarga las dependencias
RUN go mod download

# Copia el resto de los archivos al contenedor
COPY . .

# Cambia al directorio donde está el main.go
WORKDIR /app/cmd

# Compila la aplicación
RUN go build -o main .

# Expone el puerto en el que la aplicación se ejecutará
EXPOSE 8080

# Especifica el comando de inicio
CMD ["./main"]
