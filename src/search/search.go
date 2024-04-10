package search

import (
	"fmt"
	"strings"
	"tools/src/valueExtractor"
)

// ////////////////////////////////////////////////////////////////////////////////
//
// Retorna el valor de la clave indicada
// Dado          {  "total": 4306,  "decoding": 0  }
// Invocamos       "total:"
// Retorna         4306
func GetValueFromKey(buff, key, valueDeliminter string) string {
	inicioKey := strings.Index(buff, key)
	longitudKey := len(key)
	restoBuff := buff[inicioKey+longitudKey:]
	posFinValor := strings.Index(restoBuff, valueDeliminter)
	return strings.Trim(restoBuff[:posFinValor], " ")
}

// ////////////////////////////////////////////////////////////////////////////////
//
// Retorna el valor de la clave json indicada
// Dado          {  "total": 4306,  "decoding": 0  }
// Invocamos       total
// Retorna         4306
// Asume que la clave tiene " " y que viene un : separando la clave/valor
func GetJsonValue(buff, key, valueDeliminter string) string {
	inicioKey := strings.Index("\""+buff+"\"", key)
	longitudKey := len(key)
	restoBuff := buff[inicioKey+longitudKey:]
	posDosPuntos := strings.Index(restoBuff, ":")
	posFinValor := strings.Index(restoBuff, valueDeliminter)
	return strings.Trim(restoBuff[posDosPuntos+1:posFinValor], " ")
}

// ////////////////////////////////////////////////////////////////////////////////
//
// Retorna el valor ......
// Dado          X-Forwarded-Port:[443] X-Forwarded-Proto:[https] Accept:[image/gif,
// Invocamos     buff,  X-Forwarded-Port:,  [ , ]
// Retorna       443
// Asume que la clave tiene " " y que viene un : separando la clave/valor
func GetValueFromKeyWithDelimiters(buff, key, startDeliminter string, endDeliminter string, printDelimiters bool) string {
	inicioKey := strings.Index(buff, key)
	longitudKey := len(key)
	restoBuff := buff[inicioKey+longitudKey:]
	posIniValor := strings.Index(restoBuff, startDeliminter) + 1
	posFinValor := strings.Index(restoBuff, endDeliminter)
	if printDelimiters {
		posIniValor--
		posFinValor++
	}
	return strings.Trim(restoBuff[posIniValor:posFinValor], " ")
}

// =-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-= //

func GetSmartValue(buff, key string) (string, error) {
	restoBuff, err := locateValueBuffer(buff, key)
	if err != nil {
		return "", err
	}

	primerCaracterEsDelimitador := restoBuff[:1]
	if !valueExtractor.CharIsValueDelimiter(primerCaracterEsDelimitador) {
		valueExtractor.SetDelimiter(" ")
	} else {
		valueExtractor.SetDelimiter(primerCaracterEsDelimitador)
	}

	idx := 1
	stillInValue := true
	for stillInValue {
		if valueExtractor.IsEndDelimiter(string(restoBuff[idx])) {
			break
		} else {
			idx++
		}
	}
	value := restoBuff[0 : idx+1]
	return value, nil

}

// devuelve el buffer restante luego de ubicar la key y quitarla
// ejemplo
//
//	param 1   time="2024-04-05 11:39:27.864" level=INFO module=helpers....
//	param 2   time
//
// retorna
//
//	"2024-04-05 11:39:27.864" level=INFO module=helpers,
func locateValueBuffer(buff, key string) (string, error) {
	inicioKey := strings.Index(buff, key)
	if inicioKey < 0 {
		return "", fmt.Errorf("key %s no existe", key)
	}
	longitudKey := len(key)
	restoBuff := buff[inicioKey+longitudKey:]

	primerCaracterResto := restoBuff[:1]
	if !valueExtractor.CharIsKeySeparator(primerCaracterResto) {
		return locateValueBuffer(restoBuff, key)
	}
	return restoBuff[1:], nil
}
