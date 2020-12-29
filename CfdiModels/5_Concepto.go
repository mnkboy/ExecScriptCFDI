//Package cfdi :
package CfdiModels

import "encoding/xml"

//ConceptoModel :
type ConceptoModel struct {
	XMLName             xml.Name                 `xml:"cfdi:Concepto"`
	ClaveProdServ       string                   `xml:"ClaveProdServ,attr"`
	NoIdentificacion    string                   `xml:"NoIdentificacion,attr"`
	Cantidad            string                   `xml:"Cantidad,attr"`
	ClaveUnidad         string                   `xml:"ClaveUnidad,attr"`
	Unidad              string                   `xml:"Unidad,attr"`
	Descripcion         string                   `xml:"Descripcion,attr"`
	ValorUnitario       string                   `xml:"ValorUnitario,attr"`
	Importe             string                   `xml:"Importe,attr"`
	Impuestos           *ImpuestosModel          `xml:"cfdi:Impuestos"`
	ComplementoConcepto ComplementoConceptoModel `xml:"cfdi:ComplementoConcepto"`
}
