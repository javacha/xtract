package search

import (
	"fmt"
	"strings"
	"tools/src/valueExtractor"
)

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
	value := strings.Trim(restoBuff[0:idx+1], " ")
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
