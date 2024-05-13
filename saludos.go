package saludos

import (
	"errors"
	"fmt"
	"math/rand"
)

// la función HELLO devuelve un saludo

func Hola(nombre string) (string, error) {
	// devuelve un saludo con el nombre en el mensaje
	// crear un mensaje de error a devolver, cuando la cadena esta vacia
	if nombre == "" {
		return "", errors.New("el nombre esta vacio")
	}
	// forma normal, estatica de un mensasje
	//mensaje := fmt.Sprintf("hola %v! ¡Bienvenido!", nombre)
	// forma aleatoria de devolver un mensaje
	mensaje := fmt.Sprintf(SaludosAleatorios(), nombre)
	return mensaje, nil
}
func SaludosAleatorios() string {
	SaludA := []string{
		"hola %v! ¡Bienvenido!",
		"Qué bueno verte %v",
		"¡Saludo, %v! ¡Encantado de verte!",
	}
	return SaludA[rand.Intn(len(SaludA))]

}
