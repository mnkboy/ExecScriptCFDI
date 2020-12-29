//Package cfdi :
package CfdiModels

import "encoding/xml"

//TimbreFiscalDigitalModel :
type TimbreFiscalDigitalModel struct {
	XMLName           xml.Name `xml:"tfd:TimbreFiscalDigital"`
	XsiSchemaLocation string   `xml:"xsi:schemaLocation,attr"`
	Version           string   `xml:"Version,attr"`
	UUID              string   `xml:"UUID,attr"`
	FechaTimbrado     string   `xml:"FechaTimbrado,attr"`
	RfcProvCertif     string   `xml:"RfcProvCertif,attr"`
	SelloCFD          string   `xml:"SelloCFD,attr"`
	NoCertificadoSAT  string   `xml:"NoCertificadoSAT,attr"`
	SelloSAT          string   `xml:"SelloSAT,attr"`
	XmlnsTfd          string   `xml:"xmlns:tfd,attr"`
}
