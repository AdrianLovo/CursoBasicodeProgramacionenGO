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
```go
helloMessage := "Hello"
worldMessage := "world"

//println
fmt.Println(helloMessage, worldMessage);	//Agrega salto de linea de forma automatica

//printf
nombre := "Platzi"
cursos := 500
fmt.Printf("%s tiene más de %d cursos \n", nombre, cursos)	//%s cadena y %d numero
fmt.Printf("%v tiene más de %v cursos \n", nombre, cursos)	//%v si no sabes el tipo de dato

//sprintf
message := fmt.Sprintf("%s tiene más de %d cursos", nombre, cursos)
fmt.Println(message)
```

### <span class="clase"> 11. Uso de funciones </span>
```go
func normalFUnction(message string){
	fmt.Println(message)
}

func tripleArgument(a, b int, c string){	// (a int, b int, c string) pude ser de esta forma tambien
	fmt.Println(a, b, c)	
}

func returnValue(a int) int{	//retornando 1 volor
	return a * 2
}

func dobleReturn(a int) (c, d int){		//retorna dos valores
	return a, a * 2
}

func main() { //funcion principal
	normalFUnction("Hola mundo")
	tripleArgument(1, 2, "Hola")
	value := returnValue(2)
	fmt.Println("value:", value)

	value1, value2 := dobleReturn(2)	//recibiendo dos valroes
	fmt.Println("value1 y value2",value1, value2)

	only1, _ := dobleReturn(8)	//recibiendo dos valroes
	fmt.Println("only1",only1)
}
```

### <span class="clase"> 12. Go doc: La forma de ver documentación </span>
Documentación de go: <br>
Princpal: https://go.dev/doc/  <br>
Paquetes: https://pkg.go.dev/ <br>



## <span class="seccion"> Estructuras de control de flujo y condicionales </span>

### <span class="clase"> 13. El poder de los ciclos en Goland: for, for while y for forever </span>
En Go solamente existen ciclos for
```go
	//for condicional
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	fmt.Printf("\n");

	//for while
	counter := 0
	for counter < 10{
		fmt.Println(counter)
		counter++
	}

	//for forever
	counterForever:= 0
	for {
		fmt.Println(counterForever)
		counterForever++
	}
```

### <span class="clase"> 14. Operadores lógicos y de comparación </span>
Son operadores que nos permiten hacer una comparación de condiciones y en caso de cumplirse como sino ejecutarán un código determinado. Si se cumple es VERDADERO/TRUE y si no se cumple son FALSO/FALSE <br><br>
Operadores de comparación <br>
• valor1 == valor2: Retorna TRUE si valor1 y valor2 son exactamente iguales <br>
• valor1 != valor2: Retorna TRUE si valor1 es diferente de valor2 <br>
• valor1 < valor2: Retorna TRUE si valor1 es menor que valor2 <br>
• valor1 > valor2: Retorna TRUE si valor1 es mayor que valor2 <br>
• valor1 >= valor2: Retorna TRUE si valor1 es igual o mayor que valor2 <br>
• valor1 <= valor2: Retorna TRUE si valor1 es menor o igual que valor2 <br><br>
Operadores lógicos <br>
Operador AND: <br>
Este operador indica que todas las condiciones declaradas deben cumplirse para poderse marcar como TRUE. En Go, se utiliza este símbolo && <br>
• Ejemplo1: 1>0 && 2 >0 Esto retornará TRUE porque tanto la primera como la segunda condición son verdaderas <br> 
• Ejemplo2: 2<0 && 1 > 0 Esto retornará FALSE porque una de las condiciones no es verdadera <br>
Operador OR: <br>
Este operador indica que al menos una de las condiciones debe cumplirse para marcarse como TRUE. En Go, se representa con el símbolo || <br>
• Ejemplo: 2<0 || 1 > 0 Esto retornará TRUE porque la segunda condición se cumple, a pesar que la primera no <br>
Operador NOT: <br>
Este operador retornará el opuesto al boleano que está dentro de la variable <br>
```go
myBool :=  true
fmt.Println(!myBool)
```

### <span class="clase"> 15. El condicional if </span>
Documentacion: https://pkg.go.dev/strconv
```go
valor1 := 2
valor2 := 2

if(valor1 == 1){
    fmt.Println("Es 1")
}else{
    fmt.Println("No es 1")
}

// and
if (valor1 == 1 && valor2 ==2){
    fmt.Println("Es verdad")
}else{
    fmt.Println("No es verdad")
}

// or
if(valor1== 0 || valor2 ==2){
    fmt.Println("Es verdad, OR")
}else{
    fmt.Println("No es verdad, OR")
}

// convertir texto a numero
value, err := strconv.Atoi("53")
if(err != nil){		// indica si no tiene error
    log.Fatal(err)
}
fmt.Println("Value:", value)
```

### <span class="clase"> 16. Múltiple condiciones anidadas con Switch </span>
Cuando hay multiples if es mejor cambiar switch
```go
// iterar sobre una misma variable
switch modulo := 5 % 2;modulo {
    case 0:
        fmt.Println("Es par")
    default:
        fmt.Println("Es impar")
}	


// sin codicion, cuando se va a evaluar multiples condiciones
value := 50
switch{
case value > 100:
    fmt.Println("Es mayor a 100")
case value < 0:
    fmt.Println("Es menor a 0")
default:
    fmt.Println("No condicion")
}
```

### <span class="clase"> 17.  </span>
defer: ejecuta el codigo al final, se puede usar más de una linea pero se recomienda utilizar sólo una
```go

```




## <span class="seccion">  Estructuras de datos básicas <span>