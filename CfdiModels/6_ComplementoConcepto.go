//Package cfdi :
package CfdiModels

import "encoding/xml"

//ComplementoConceptoModel :
type ComplementoConceptoModel struct {
	XMLName            xml.Name                `xml:"cfdi:ComplementoConcepto"`
	IEDUInstEducativas IEDUInstEducativasModel `xml:"iedu:instEducativas"`
}
