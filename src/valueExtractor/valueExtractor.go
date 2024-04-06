package valueExtractor

var valueDeliminter string
var delimiterDeep int

func SetDelimiter(char string) {
	valueDeliminter = char
}

func isEndDelimiter(char string) bool {
	if valueDeliminter == "\"" && char == "\"" {
		return true
	}

	if valueDeliminter == " " && char == " " {
		return true
	}

	if valueDeliminter == "[" {
		if char == "[" {
			delimiterDeep++
		}
		if char == "]" {
			delimiterDeep--
			if delimiterDeep < 0 {
				return true
			}

		}
	}
	return false
}
