//Package cfdi :
package CfdiModels

import "encoding/xml"

//TrasladoTotalModel :
type TrasladoTotalModel struct {
	XMLName    xml.Name `xml:"cfdi:Traslado"`
	Impuesto   string   `xml:"Impuesto,attr"`
	TipoFactor string   `xml:"TipoFactor,attr"`
	TasaOCuota string   `xml:"TasaOCuota,attr"`
	Importe    string   `xml:"Importe,attr"`
}
