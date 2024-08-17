<style>
.curso {
    color: #1E90FF;
}
.seccion {
    color: #FF8C00;
}
.clase {
    color: #32CD32;
}
</style>

# <span class="curso"> Curso Básico de Programación en Go</span>




## <span class="seccion"> Hola mundo en Go </span>

### <span class="clase"> 1. Introducción al Curso de Golang </span>
¿Qué aprenderas? <br>
• Variables y constantes <br>
• Funciones y estructuras básicas <br>
• Gorutines <br>
• GO modules <br>
• Librerías para Backend <br>
• Tips <br>

### <span class="clase"> 2. ¿Qué es y por qué usar GO? </span>
• Fue creado para el manejo de los precesos pesados pero que fuera amigable como python <br>
• Compilado y estáticamente tipado <br>
• Creado en Google por Robert Griesemer, Rob Pike y Ken Thompson <br>
• Anunciado en Noviembre 2009 <br>
• Primera versi+ón en Marzo 2012 <br>
• Go/Golang <br>
• Programadores: Gophers <br><br>
¿Porqué aprender Go? <br>
• Gran velocidad de compilación <br>
• Alto rendimiento para tareas pesadas <br>
• Soporte nativo para concurrencia <br>
• Top 5 más amados y mejores pagados en encuestas <br>
• Obliga a implementar buenas practicas <br>
• Comunidad muy receptiva <br><br>
Quiénes usan Go? <br>
• MercadoLibre: 70,000 request en maquinbas con  20MB RAM <br>
• Tiwtch: Usuarios concurrentes <br>
• Twitter: Procesar analitica de la App con Kafka <br>
• Uber: Posición conductores y pasajeros en tiempo real <br>
• Docker y Kubernetes: Para despliegues de Apps <br>

### <span class="clase"> 3. Instalar Go en Linux </span>
Descargar el codigo nativo (go1.15.6.linux.amd64.tar.gz) <br>
https://go.dev/dl/


```bash 
code .bashrc
```

```bash 
#in the file ".bashrc"
#Entorno donde se desarrollaran los programas
export GOPATH=/home/tito/go      
#Donde se instalaran los binarios
export GOBIN=/$GOPATH/bin
#Donde va a estar el ejecutador de go
export GOTROOT=/$GOPATH/bin

#export a las variables de entorno
export PATH=$PATH:$GOBIN:$GORROT/bin
```

```bash 
. . bashrc
go version
```

### <span class="clase"> 5. Instalar Go en Window </span>
```bash 
C:\Users\adria\go   #En el directorio crear la carpeta src que contendra los proyectos
```

### <span class="clase"> 6. Nuestras primeras líneas de código con Go </span>
El archivo principal debe llamarse `main.go` <br>
Para compilar un programa en linux `go build src/main.go` y para ejecutar `./main` tambien se puede utilizar `go run archivo.go` <br>
Para ejecutar en window se usa `go run archivo.go` en la terminal



## <span class="seccion"> Variables, funciones y documentación </span>

### <span class="clase"> 7. Variables, constantes y zero values </span>