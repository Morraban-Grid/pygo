//go:build ignore

package main

import (
	"fmt"
	"time"
)

//NotificationService define el contrato de cualquier servicio de mensajeria
type NotificationService interface{
	//Definimos el método send
	//para que cualquier servicio de mensajeria pueda enviar un mensaje 
	//a un destiantario específico
	Send(recipient string, message string) error
}

// Implementación 1 - Email
type EmailProvider struct{
	SMTPServer string
}

func (e EmailProvider) Send(to string, message string) error {
	//Simulación de lógica de envio de email
	fmt.Printf("Enviando email a %s via %s: %s\n", to, e.SMTPServer, message)
	return nil
}

// Implementación 2 - SMS
type SMSProvider struct{
	APIKey string
}

func (s SMSProvider) Send(phone string, msg string) error{
	//Simulación de la lógica de envio de SMS
	fmt.Printf("Enviando SMS a %s usando API Key %s: %s\n", phone, s.APIKey, msg)
	return nil
}

// Lógica de negocio desacoplada

//AlterManager no sabe que proveedor usará, solo sabe que puede "Enviar"
type AlertManager struct{
	notifier NotificationService
}

func (a *AlertManager) NotifyUser(contact string, text string) error{
	err := a.notifier.Send(contact, text)
	if err != nil{
		fmt.Println("Error al enviar la notificación: ", err)
	}
}

func main(){
	//Podemos intercambiar las implementaciones facilmente
	emailSvc := EmailProvider{SMTPServer: "smtp.gmail.com"}
	smsSvc := SMSProvider{APIKey: "secret_key_123"}

	//Uso con Email
	app := AlertManager{Notifier: emailSvc}
	app.NotifyUser("user@example.com", "Tu servidor se ha reiniciado")

	//Uso con SMS
	app.notifier = smsSvc
	app.NotifyUser("+51900800700", "Alerta de seguridad detectada")
}

