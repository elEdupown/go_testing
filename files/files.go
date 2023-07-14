package files

import (
	"github.com/elEdupown/go_testing/ejercicios"
	"os"
	"fmt"
	//"bufio"
	"io/ioutil"
)

var filename string = "./files/txt/tabla.txt"

func GrabarTabla() {
	var texto string = ejercicios.Tabla()

	archivo, err := os.Create(filename)

	if err != nil {
		fmt.Println("Hubo un error al crear el archivo"+err.Error())
		return
	}

	fmt.Fprintln(archivo, texto)
	archivo.Close()
}

func SumaTabla(){
	var texto string = ejercicios.Tabla()

	if !Append(filename, texto){
		fmt.Println("Error al concatenar contenido")
	}
}

func Append(filen string, texto string) bool{
	arch, err := os.OpenFile(filen, os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		fmt.Println("Hubo un error al abrir el archivo"+err.Error())
		return false
	}

	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Hubo un error al escribir el archivo"+err.Error())
		return false
	}

	arch.Close()
	return true
}

func LeerArchivo() string{
	archivo, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Hubo un error al leer el archivo"+err.Error())
		return ""
	}

	return string(archivo)
}