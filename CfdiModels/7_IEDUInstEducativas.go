//Package cfdi :
package CfdiModels

import "encoding/xml"

//IEDUInstEducativasModel :
type IEDUInstEducativasModel struct {
	XMLName           xml.Name `xml:"iedu:instEducativas"`
	XsiSchemaLocation string   `xml:"xsi:schemaLocation,attr"`
	Version           string   `xml:"Version,attr"`
	NombreAlumno      string   `xml:"NombreAlumno,attr"`
	CURP              string   `xml:"CURP,attr"`
	NivelEducativo    string   `xml:"NivelEducativo,attr"`
	AutRVOE           string   `xml:"AutRVOE,attr"`
}
