//Package cfdi :
package CfdiModels

import "encoding/xml"

//ComplementoModel :
type ComplementoModel struct {
	XMLName             xml.Name                 `xml:"cfdi:Complemento"`
	TimbreFiscalDigital TimbreFiscalDigitalModel `xml:"tfd:TimbreFiscalDigital"`
}
