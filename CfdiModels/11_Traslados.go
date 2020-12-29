//Package cfdi :
package CfdiModels

import "encoding/xml"

//TrasladosModel :
type TrasladosModel struct {
	XMLName  xml.Name      `xml:"cfdi:Traslados"`
	Traslado TrasladoModel `xml:"cfdi:Traslado"`
}
