//Package cfdi :
package CfdiModels

import "encoding/xml"

//TrasladoModel :
type TrasladoModel struct {
	XMLName    xml.Name `xml:"cfdi:Traslado"`
	Base       string   `xml:"Base,attr"`
	Impuesto   string   `xml:"Impuesto,attr"`
	TipoFactor string   `xml:"TipoFactor,attr"`
	TasaOCuota string   `xml:"TasaOCuota,attr"`
	Importe    string   `xml:"Importe,attr"`
}
