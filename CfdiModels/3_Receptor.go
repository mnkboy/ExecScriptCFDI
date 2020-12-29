//Package cfdi :
package CfdiModels

import "encoding/xml"

//ReceptorModel :
type ReceptorModel struct {
	XMLName xml.Name `xml:"cfdi:Receptor"`
	Rfc     string   `xml:"Rfc,attr"`
	Nombre  string   `xml:"Nombre,attr"`
	UsoCFDI string   `xml:"UsoCFDI,attr"`
}
