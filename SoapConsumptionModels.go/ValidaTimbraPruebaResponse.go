//Package SoapConsumption paquete encargado de proporcionar las estructuras para los responses del soap
package SoapConsumptionModels

import "encoding/xml"

//ValidaTimbraPruebaResponse es un struct relacionado con los atributos del comprobate XML
type ValidaTimbraResponse struct {
	XMLName            xml.Name           `xml:"ValidaTimbraPruebaResponse"` //Cambiar ValidaTimbraResponse
	XMLNS              string             `xml:"xmlns,attr"`
	ValidaTimbraResult ValidaTimbraResult `xml:"ValidaTimbraPruebaResult"` //Cambiar ValidaTimbraResult"
	FileXml            []byte
}
