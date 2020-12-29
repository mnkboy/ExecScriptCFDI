package main

import (
	"encoding/xml"
	CfdiModels "execScriptCFDI/CfdiModels"
	"execScriptCFDI/Seal"
	"execScriptCFDI/SoapConsumption"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/tiaguinho/gosoap"
)

func main() {
	//Generamos un comprobante
	Comprobante := CfdiModels.ComprobanteModel{}
	Comprobante = *llenaComprobante(&Comprobante)

	//Guardamos el archivo el comprobante
	file, _ := xml.MarshalIndent(Comprobante, "", "\t")
	_ = ioutil.WriteFile("./Seal/cfdiTimbradoIEDU.xml", file, 0644)

	//Llamamos a la funcion de sellado
	resultado := Seal.Sella("Seal")

	//Verificamos si la comprobacion es exitosa
	if strings.Contains(resultado, "Verified OK") {
		fmt.Println("Se ha comprobado exitosamente.")
	} else {
		//Verificamos si fallo la comprobacion
		if strings.Contains(resultado, "Verification Failure") {
			fmt.Println("\nLa comprobacion ha fallado.")
		} else {
			//Si es otro error imprimimos CMD
			fmt.Println("\nError: \n" + resultado)
			fmt.Println("Error")
		}
	}

	//Abrimos el archivo XML
	xmlFile, err := ioutil.ReadFile("./Seal/new_cfdiTimbradoIEDU.xml")
	if err != nil {
		panic(err)
	}

	cfdiString := string(xmlFile) // convert content to a 'string'
	cfdiString = strings.TrimSpace(cfdiString)

	ok, msg := consumeSoap(cfdiString)

	if !ok {
		fmt.Println("Error:", msg)
	}

}

func llenaComprobante(comprobante *CfdiModels.ComprobanteModel) *CfdiModels.ComprobanteModel {
	//Comprobante
	comprobante.XmlnsXsi = "http://www.w3.org/2001/XMLSchema-instance"
	comprobante.XmlnsIedu = "http://www.sat.gob.mx/iedu"
	comprobante.XsiSchemaLocation = "http://www.sat.gob.mx/cfd/3 http://www.sat.gob.mx/sitio_internet/cfd/3/cfdv33.xsd http://www.sat.gob.mx/iedu http://www.sat.gob.mx/sitio_internet/cfd/iedu/iedu.xsd http://www.sat.gob.mx/iedu http://www.sat.gob.mx/sitio_internet/cfd/iedu/iedu.xsd"
	comprobante.MetodoPago = "PUE"
	comprobante.TipoDeComprobante = "I"
	comprobante.Total = "4640.00"
	comprobante.TipoCambio = "1"
	comprobante.Moneda = "MXN"
	comprobante.SubTotal = "4000.00"
	comprobante.CondicionesDePago = "Contado"
	comprobante.Sello = ""
	comprobante.NoCertificado = ""
	comprobante.FormaPago = "01"
	comprobante.LugarExpedicion = "85040"
	comprobante.Fecha = "2018-01-30T18:33:40"
	comprobante.Folio = "12345"
	comprobante.Serie = "A"
	comprobante.Version = "3.3"
	comprobante.Certificado = ""
	comprobante.XmlnsCfdi = "http://www.sat.gob.mx/cfd/3"

	//Emisor
	// comprobante.Emisor.Rfc = "MAG041126GT8"
	comprobante.Emisor.Rfc = "XAXX010101000"
	comprobante.Emisor.Nombre = "EMCORSOFT, SC"
	comprobante.Emisor.RegimenFiscal = "601"

	//Receptor
	comprobante.Receptor.Rfc = "USO110603I26"
	comprobante.Receptor.Nombre = "UMBRALL SOFTWARE SA DE CV"
	comprobante.Receptor.UsoCFDI = "G01"

	//Concepto
	concepto := CfdiModels.ConceptoModel{}
	concepto.ClaveProdServ = "01010101"
	concepto.NoIdentificacion = "998877660011"
	concepto.Cantidad = "1.000000"
	concepto.ClaveUnidad = "H87"
	concepto.Unidad = "Pieza"
	concepto.Descripcion = "CABLE NUMERO 5"
	concepto.ValorUnitario = "4000.000000"
	concepto.Importe = "4000.000000"

	//ComplementoConcepto
	ComplementoConcepto := CfdiModels.ComplementoConceptoModel{}

	//IEDUInstEducativas
	IEDUInstEducativas := CfdiModels.IEDUInstEducativasModel{}

	IEDUInstEducativas.XsiSchemaLocation = "http://www.sat.gob.mx/iedu http://www.sat.gob.mx/sitio_internet/cfd/iedu/iedu.xsd"
	IEDUInstEducativas.Version = "1.0"
	IEDUInstEducativas.NombreAlumno = "Nombre del alumno"
	IEDUInstEducativas.CURP = "ROCA901231HSRRXK44"
	IEDUInstEducativas.NivelEducativo = "Secundaria"
	IEDUInstEducativas.AutRVOE = "11222"

	//Agregamos el iedu al complemento
	ComplementoConcepto.IEDUInstEducativas = IEDUInstEducativas

	//Agregamos el complemento al concepto
	concepto.ComplementoConcepto = ComplementoConcepto
	concepto.ComplementoConcepto = ComplementoConcepto
	concepto.ComplementoConcepto = ComplementoConcepto
	concepto.ComplementoConcepto = ComplementoConcepto
	concepto.ComplementoConcepto = ComplementoConcepto
	concepto.ComplementoConcepto = ComplementoConcepto

	//Impuestos
	Impuestos := CfdiModels.ImpuestosModel{}

	Impuestos.Traslados.Traslado.Base = "4000"
	Impuestos.Traslados.Traslado.Impuesto = "002"
	Impuestos.Traslados.Traslado.TipoFactor = "Tasa"
	Impuestos.Traslados.Traslado.TasaOCuota = "0.160000"
	Impuestos.Traslados.Traslado.Importe = "640"

	concepto.Impuestos = &Impuestos

	//Agregamos el concepto a la lista
	comprobante.Conceptos.ListConcepto = append(comprobante.Conceptos.ListConcepto, concepto)

	//=============== SEGUNDO ITEM ===============
	//Concepto
	concepto = CfdiModels.ConceptoModel{}
	concepto.ClaveProdServ = "02020202"
	concepto.NoIdentificacion = "998877660022"
	concepto.Cantidad = "2.000000"
	concepto.ClaveUnidad = "H87"
	concepto.Unidad = "Pieza"
	concepto.Descripcion = "CABLE NUMERO 22"
	concepto.ValorUnitario = "8000.000000"
	concepto.Importe = "8000.000000"

	//ComplementoConcepto
	ComplementoConcepto = CfdiModels.ComplementoConceptoModel{}

	//IEDUInstEducativas
	IEDUInstEducativas = CfdiModels.IEDUInstEducativasModel{}

	IEDUInstEducativas.XsiSchemaLocation = "http://www.sat.gob.mx/iedu http://www.sat.gob.mx/sitio_internet/cfd/iedu/iedu.xsd"
	IEDUInstEducativas.Version = "1.0"
	IEDUInstEducativas.NombreAlumno = "Nombre del alumno"
	IEDUInstEducativas.CURP = "ROCA901231HSRRXK44"
	IEDUInstEducativas.NivelEducativo = "Secundaria"
	IEDUInstEducativas.AutRVOE = "11222"

	//Agregamos el iedu al complemento
	ComplementoConcepto.IEDUInstEducativas = IEDUInstEducativas

	//Agregamos el complemento al concepto
	concepto.ComplementoConcepto = ComplementoConcepto
	concepto.ComplementoConcepto = ComplementoConcepto
	concepto.ComplementoConcepto = ComplementoConcepto
	concepto.ComplementoConcepto = ComplementoConcepto
	concepto.ComplementoConcepto = ComplementoConcepto
	concepto.ComplementoConcepto = ComplementoConcepto

	//Impuestos
	Impuestos = CfdiModels.ImpuestosModel{}

	Impuestos.Traslados.Traslado.Base = "8000"
	Impuestos.Traslados.Traslado.Impuesto = "002"
	Impuestos.Traslados.Traslado.TipoFactor = "Tasa"
	Impuestos.Traslados.Traslado.TasaOCuota = "0.160000"
	Impuestos.Traslados.Traslado.Importe = "1280.00"

	concepto.Impuestos = &Impuestos

	//Agregamos el concepto a la lista
	comprobante.Conceptos.ListConcepto = append(comprobante.Conceptos.ListConcepto, concepto)

	//TimbreFiscalDigital
	TimbreFiscalDigital := CfdiModels.TimbreFiscalDigitalModel{}

	TimbreFiscalDigital.XsiSchemaLocation = "http://www.sat.gob.mx/TimbreFiscalDigital http://www.sat.gob.mx/sitio_internet/cfd/TimbreFiscalDigital/TimbreFiscalDigitalv11.xsd"
	TimbreFiscalDigital.Version = "1.1"
	TimbreFiscalDigital.UUID = "B241EBEE-819A-AAAA-AAAA-526444336600"
	TimbreFiscalDigital.FechaTimbrado = "2018-01-30T19:33:40"
	TimbreFiscalDigital.RfcProvCertif = "IAD121214B34"
	TimbreFiscalDigital.SelloCFD = "aXfE8qEVguBnXY4kHa9iWiC+8VFLVJzIiS8UKp/7VO7FytpxLg0Em2wLTRvFWfvIMBtNzxYuLpqo6c5Xd383yfOVpTFTFHF/FEUTt0TL5kWxBkSLb5A2PD3WshOg4FQxezWAXkITsOpFljIwrHoatv31qcBT5vkZJ3S8rCc0HFAGPYsx09fArcLv1xAHhIQ0aQzLJJ2/69kAgVGFKnuzjS5rsx2nVCfIqkG/IF8U7j5uei+lmLPLA/BmuAQRIt4iEvhhd2FxvemJEESBWYp4m2vDer92D68ViET2VMCFmJSFIDiuXsNeVoYcOhSyFfE1ggUp8L4HbUDvFZ/+u+y6eg=="
	TimbreFiscalDigital.NoCertificadoSAT = "20001000000300022323"
	TimbreFiscalDigital.SelloSAT = "EytLGdSR2C3NjLHNu3L9ZKfwqhwqNqRmeuAmMLZcVQlgJfdHjJXES20gQb452DSLSmRToJbFHbbZ2pyJQbcA3fvIRV0VsziMGjGZPyG1OLNIulqWO7qs0CSPvRJy02l5TfsRrLTkjaEfHTD1zj51W6fIJnz5YtERWn+5FN0ISPvWf3Is/h8gzsjnPpkOZsyU+kr8FUBdvKvu8MFMelA5LQHfhbwTUvm/KvTTNO4TOFP4G0CLScZ0evHyR1FiVjd9sOSDiXYifdmgJdVPSsJtCD3L4YEHIWKB+Z4ijkNuslqfmUl3sZTAeDeoosr+sdxjVET4p5wicKhHQkmWqVDfcg=="
	TimbreFiscalDigital.XmlnsTfd = "http://www.sat.gob.mx/TimbreFiscalDigital"

	//Agregamos el tibre fiscal digital al complemento
	Complemento := CfdiModels.ComplementoModel{}
	Complemento.TimbreFiscalDigital = TimbreFiscalDigital

	//Agregamos el complemento al comprobante
	comprobante.Complemento = Complemento

	//Agregamos ImpuestosTotales
	ImpuestosTotalesModel := CfdiModels.ImpuestosTotalesModel{}

	ImpuestosTotalesModel.TrasladosTotales.Traslado.Impuesto = "002"
	ImpuestosTotalesModel.TrasladosTotales.Traslado.TipoFactor = "Tasa"
	ImpuestosTotalesModel.TrasladosTotales.Traslado.TasaOCuota = "0.160000"
	ImpuestosTotalesModel.TrasladosTotales.Traslado.Importe = "1920.00"

	//Agregamos los impuestos al comprobante
	comprobante.ImpuestosTotales = &ImpuestosTotalesModel
	comprobante.ImpuestosTotales.TotalImpuestosTrasladados = ImpuestosTotalesModel.TrasladosTotales.Traslado.Importe

	return comprobante
}

func consumeSoap(cfdiString string) (bool, string) {
	//Declaramos la url del servidor que nos proveera el servicio de timbrado de facturas
	urlSOAP := "https://invoiceone.mx/fachadacore/fachadacore.asmx?wsdl"
	//Declaramos el nombre del metodo que utilizaremos para timbrar la factura
	soapService := "ValidaTimbraPrueba" //Cambiar a ValidaTimbra para produccion

	//Establecemos los parametros para utilizar el servicio soap
	params := gosoap.Params{
		"VersionDLL":     "SisteMexico",
		"SoftwareID":     "QV1TAXSH",
		"Agente":         "QV1TAXSH",
		"PasswordEmisor": "",
		"Xml":            cfdiString,
	}

	//Llamamos al metodo para consumir el serivicio soap de invoice one
	res, messageSoap, respuestaSoap := SoapConsumption.ConsumeSOAPService(urlSOAP, soapService, params)

	//Si hay error devolvemos error
	if !res {
		//Imprimimos la respuesta del servicio
		return res, messageSoap

	}

	//Imrimimos la respuesta del servicio
	fmt.Println(respuestaSoap)

	//Imprimimos la respuesta del servicio
	return res, messageSoap

}
