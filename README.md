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
```go
//declaracion de constantes
const pi float64 = 3.14 //declarando el tipo de dato
const pi2 = 3.1415      //sin declarar el tipo

fmt.Println("pi", pi)
fmt.Println("pi2", pi2)

//declaracion de variables enteras
base := 12 // : indican que se va a crear por primera vez y deduce el tipo que sera
var altura int = 14
var area int

fmt.Println(base, altura, area)

//zero values: tiene por defecto un valor 
var a int			//0
var b float64		//0
var c string		//vacio
var d bool			//false

fmt.Println(a, b, c, d)
```

### <span class="clase"> 8. Operadores aritméticos</span>
```go
x :=10
y := 50

//suma
result := x + y
fmt.Println("Suma:", result)

//resta
result = y - x
fmt.Println("Resta:", result)

//multiplicacion
result = x * y
fmt.Println("Multiplicación:", result)

//division
result = y / x
fmt.Println("División:", result)

//modulo
result = y % x
fmt.Println("Modulo:", result)

//incrementar
x++
fmt.Println("Incremental:", x)

//decrementar
x--
fmt.Println("Decremental:", x)
```

### <span class="clase"> 9. Tipos de datos primitivos </span>
Números enteros: <br>
• int = Depende del OS (32 o 64 buits) <br>
• int8 = 8 bits = -128 a 127 <br>
• int16 = 16 bits = -2<sup>15</sup> a 2<sup>15</sup> <br>
• int32 = 32 bits = -2<sup>31</sup> a 2<sup>31</sup> <br>
• int64 = 64 bits = -2<sup>63</sup> a 2<sup>63</sup> <br><br>
Números enteros positivos: <br>
• uint = Depende del OS (32 o 64 buits) <br>
• uint8 = 8 bits = 0 a 2<sup>8</sup>-1 <br>
• uint16 = 16 bits = 0 a 2<sup>16</sup>-1 <br>
• uint32 = 32 bits = 0 a 2<sup>32</sup>-1 <br>
• uint64 = 64 bits = 0 a 2<sup>64</sup>-1 <br><br>
Números decimales: <br>
• float32 = 32 bits = +/-1.18e<sup>-38</sup> a +/-3.4e<sup>38</sup> <br>
• float64 = 64 bits = +/-1.18e<sup>-38</sup> a +/-3.4e<sup>38</sup> <br><br>
Textos y boleanos
• string = "" <br>
• bool = true o false <br><br>
Números complejos: <br>
• Complex64 = Real e imaginario float32 <br>
• Complex128 = Real e imaginario float64 <br>
Ejemplo: c:= 10 + 8i <br>


### <span class="clase"> 10. Paquete fmt: Algo más que imprimir en consola </span>