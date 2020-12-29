package SoapConsumption

import (
	"net/http"

	"execScriptCFDI/SoapConsumptionModels.go"

	"github.com/tiaguinho/gosoap"
)

//ConsumeSOAPService es una funcion encargada de consumir un servicio SOAP ofrecido por un tercero
func ConsumeSOAPService(urlSOAP string, soapService string, params gosoap.Params) (bool, string, *SoapConsumptionModels.ValidaTimbraResponse) {
	//Establecemos la conexion con el servicio SOAP
	//soap, err := gosoap.SoapClient(urlSOAP)
	soap, err := gosoap.SoapClient(urlSOAP, &http.Client{})

	if err != nil {
		return false, err.Error(), nil
	}

	//Llamos al servicio y le pasamos los parametros body
	//err = soap.Call(soapService, params)
	res, err := soap.Call(soapService, params)
	//Comprobamos si no hubo ningun error
	if err != nil {
		return false, err.Error(), nil
	}

	//Declaramos una variable de tipo Envelope, que es una estructura con formato para recibir
	response := SoapConsumptionModels.ValidaTimbraResponse{}

	//Unmarshaleamos la respuesta a la variable

	res.Unmarshal(&response)

	if response.ValidaTimbraResult.Estatus == "ERROR" {
		//Devolvemos un xml con la respuesta debidamente estructurada
		return false, response.ValidaTimbraResult.MensajeError, &response
	}

	//Devolvemos un xml con la respuesta debidamente estructurada
	return true, response.ValidaTimbraResult.MensajeError, &response

}
