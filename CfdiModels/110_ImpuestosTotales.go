//Package cfdi :
package CfdiModels

import "encoding/xml"

//ImpuestosTotalesModel :
type ImpuestosTotalesModel struct {
	XMLName                   xml.Name              `xml:"cfdi:Impuestos"`
	TotalImpuestosTrasladados string                `xml:"TotalImpuestosTrasladados,attr"`
	TrasladosTotales          TrasladosTotalesModel `xml:"cfdi:Traslados"`
}
