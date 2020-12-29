//Package cfdi :
package CfdiModels

import "encoding/xml"

//TrasladosTotalesModel :
type TrasladosTotalesModel struct {
	XMLName  xml.Name           `xml:"cfdi:Traslados"`
	Traslado TrasladoTotalModel `xml:"cfdi:Traslado"`
}
