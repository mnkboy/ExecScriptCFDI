//Package SoapConsumption paquete encargado de proporcionar las estructuras para los responses del soap
package SoapConsumptionModels

import "encoding/xml"

//ValidaTimbraResult es un struct relacionado con los atributos del comprobate XML
type ValidaTimbraResult struct {
	XMLName       xml.Name `xml:"ValidaTimbraPruebaResult"` //Cambiar ValidaTimbraResult
	IDOperacion   string   `xml:"IDOperacion"`
	Estatus       string   `xml:"Estatus"`
	OriginalChain string   `xml:"CadenaOriginal"`
	Xml           string   `xml:"Xml"`
	QR            string   `xml:"QR"`
	Saldo         string   `xml:"Saldo"`
	MensajeError  string   `xml:"MensajeError"`
	CodigoError   string   `xml:"CodigoError"`
}
