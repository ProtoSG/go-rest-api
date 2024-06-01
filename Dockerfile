# Usar una imagen base de Go
FROM golang:1.18

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar el archivo go.mod y go.sum y descargar las dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copiar el resto del código
COPY . .

# Compilar la aplicación
RUN go build -o main .

# Definir la entrada del contenedor
CMD ["./main"]

