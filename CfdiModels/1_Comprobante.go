//Package cfdi :
package CfdiModels

import "encoding/xml"

//ComprobanteModel :
type ComprobanteModel struct {
	XMLName           xml.Name               `xml:"cfdi:Comprobante"`
	XmlnsXsi          string                 `xml:"xmlns:xsi,attr"`
	XmlnsIedu         string                 `xml:"xmlns:iedu,attr"`
	XsiSchemaLocation string                 `xml:"xsi:schemaLocation,attr"`
	XmlnsCfdi         string                 `xml:"xmlns:cfdi,attr"`
	MetodoPago        string                 `xml:"MetodoPago,attr"`
	TipoDeComprobante string                 `xml:"TipoDeComprobante,attr"`
	Total             string                 `xml:"Total,attr"`
	TipoCambio        string                 `xml:"TipoCambio,attr"`
	Moneda            string                 `xml:"Moneda,attr"`
	SubTotal          string                 `xml:"SubTotal,attr"`
	CondicionesDePago string                 `xml:"CondicionesDePago,attr"`
	Sello             string                 `xml:"Sello,attr"`
	NoCertificado     string                 `xml:"NoCertificado,attr"`
	FormaPago         string                 `xml:"FormaPago,attr"`
	LugarExpedicion   string                 `xml:"LugarExpedicion,attr"`
	Fecha             string                 `xml:"Fecha,attr"`
	Folio             string                 `xml:"Folio,attr"`
	Serie             string                 `xml:"Serie,attr"`
	Version           string                 `xml:"Version,attr"`
	Certificado       string                 `xml:"Certificado,attr"`
	Emisor            EmisorModel            `xml:"cfdi:Emisor"`
	Receptor          ReceptorModel          `xml:"cfdi:Receptor"`
	Conceptos         ConceptosModel         `xml:"cfdi:Conceptos"`
	ImpuestosTotales  *ImpuestosTotalesModel `xml:"cfdi:Impuestos"`
	Complemento       ComplementoModel       `xml:"cfdi:Complemento"`
}
