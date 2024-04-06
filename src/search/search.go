package search

import (
	"fmt"
	"strings"
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

func extractKeyAndRemainderBuffer(buff, key string) (string, string) {
	inicioKey := strings.Index(buff, key)
	longitudKey := len(key)
	restoBuff := buff[inicioKey+longitudKey:]

	primerCaracterResto := restoBuff[1:]
	if charIsKeySeparator(primerCaracterResto) {

	}

	return "a", "b"
}

func GetValue(buff, key string) string {
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

func Tester() {
	fmt.Println(nextCharIs(":[172.29.99.68] Content-Type:[application/x-www-form-urlencoded]"))
}

func nextCharIs(buff string) bool {
	siguiente := buff[:1]
	return charIsKeySeparator(siguiente)
}

func charIsKeySeparator(char string) bool {
	var keySeparators = [2]string{":", "="}
	for _, n := range keySeparators {
		if char == n {
			return true
		}
	}
	return false
}
