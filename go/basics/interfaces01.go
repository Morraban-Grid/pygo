//go:build ignore 

package main

import "fmt"

// Interface que define un método para procesar pagos
type MetodoPago interface {
	// El método Pagar toma un monto de tipo float64 y devuelve una cadena 
	Pagar(monto float64) string
}

// Estructura para representar una tarjeta de crédito
type Tarjeta struct {
	Numero string
	Banco string
}

// Método de la estructura Tarjeta que implementa la interfaz MetodoPago
func (t Tarjeta) Pagar(monto float64) string {
	return "Pagando $" + fmt.Sprintf("%.2f", monto) + " con tarjeta " + t.Banco
}

// Estructura para representar una cuenta de PayPal
type PayPal struct {
	Email string
}

// Método de la estructura PayPal que implementa la interfaz MetodoPago
func (p PayPal) Pagar(monto float64) string {
	return "Pagando $" + fmt.Sprintf("%.2f", monto) + " mediante la cuenta " + p.Email
}

// Función que procesa una transacción utilizando un método de pago que implementa la interfaz MetodoPago
func ProcesarTransaccion(m MetodoPago, monto float64) {
	fmt.Println("============================")
	fmt.Println("Iniciando transacción ...")
	// Creamos una variable de tipo string para almacenar el resultado del método Pagar
	resultado := m.Pagar(monto)
	fmt.Println(resultado)
	fmt.Println("Transacción finalizada.")
	fmt.Println("============================")
} 

func main() {
	tarjeta01 := Tarjeta{
		Numero: "1234-5678-9012-3456",
		Banco: "BCP",
	}

	paypal01 := PayPal{
		Email: "maximo.123@gmail.com",
	}

	ProcesarTransaccion(tarjeta01, 150.75)
	ProcesarTransaccion(paypal01, 200.00)
}