//Package cfdi :
package CfdiModels

import "encoding/xml"

//ImpuestosModel :
type ImpuestosModel struct {
	XMLName   xml.Name       `xml:"cfdi:Impuestos"`
	Traslados TrasladosModel `xml:"cfdi:Traslados"`
}
