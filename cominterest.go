/*
*************************
(C) 2017 Fabian Salamanca
Compund interest
Interés compuesto
*************************
*/
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type cliente struct {
	nombre      string
	apellido    string
	edad        int
	estadoCivil string
}

func (clien *cliente) client() {
	fmt.Println("==== CUSTOMER ====")
	fmt.Printf("Name: %s %s \n\r", clien.nombre, clien.apellido)
	fmt.Printf("Age: %d \n\r\n\r", clien.edad)
}

type credito struct {
	intanual float32
	capi     float64
	capf     float64
	n        int
}

func (cred *credito) calculo() {
	var capcal float32 = 1.0
	if cred.intanual > 1 {
		capcal = 1.0 + (cred.intanual / 100)
	} else {
		capcal = 1.0 + cred.intanual
	}
	cred.capf = math.Pow(float64(capcal), float64(cred.n/12)) * cred.capi
	fmt.Printf("Initial capital = $ %.2f \n\r", cred.capi)
	fmt.Printf("Final capital = $ %.2f \n\r", cred.capf)
	fmt.Printf("Period in years = %d   Monthly interest = %.2f \n\r\n\r", cred.n/12, cred.intanual/12)
	mensualidad := cred.capf / float64(cred.n)
	saldoi := cred.capf
	for i := 0; i < cred.n; i++ {
		saldoi -= mensualidad
		fmt.Printf("Monthly payment %d = $ %.2f  |  Pending capital = $ %.2f \n\r", i+1, mensualidad, saldoi)
	}
}

//TODO: Output calculus in a 2D slice

func cuestionario(texto string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please input your ", texto)
	fmt.Print(":**> ")
	salida, _ := reader.ReadString('\n')
	salida = strings.Trim(salida, "\n")
	fmt.Println("")
	return salida
}

func main() {

	fmt.Println("")
	fmt.Println("--------------------------")
	fmt.Println("Compound Interest CLI Tool")
	fmt.Println("--------------------------")
	fmt.Println("")

	nombre := cuestionario("name")
	apellido := cuestionario("last name")
	edad := cuestionario("age")
	precio := cuestionario("initial capital (price)")
	intanual := cuestionario("annual interest %")
	periodo := cuestionario("credit period in months")
	iedad, _ := strconv.ParseInt(edad, 10, 32)

	client := cliente{nombre: nombre, apellido: apellido, edad: int(iedad)}
	client.client()

	fprecio, _ := strconv.ParseFloat(precio, 64)
	ffintanual, _ := strconv.ParseFloat(intanual, 64)
	fintanual := float32(ffintanual)
	iperiodo, _ := strconv.ParseInt(periodo, 10, 32)
	cred := credito{capi: fprecio, intanual: fintanual, n: int(iperiodo)}
	cred.calculo()

	fmt.Println("")
}
