//Package cfdi :
package CfdiModels

import "encoding/xml"

//ConceptosModel :
type ConceptosModel struct {
	XMLName      xml.Name        `xml:"cfdi:Conceptos"`
	ListConcepto []ConceptoModel `xml:"cfdi:Concepto"`
}
