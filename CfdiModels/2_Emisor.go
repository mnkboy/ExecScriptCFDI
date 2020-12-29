//Package cfdi :
package CfdiModels

import "encoding/xml"

//EmisorModel :
type EmisorModel struct {
	XMLName       xml.Name `xml:"cfdi:Emisor"`
	Rfc           string   `xml:"Rfc,attr"`
	Nombre        string   `xml:"Nombre,attr"`
	RegimenFiscal string   `xml:"RegimenFiscal,attr"`
}
