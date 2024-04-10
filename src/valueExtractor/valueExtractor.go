package valueExtractor

var startDelimiter string
var endDelimiter string
var delimiterDeep int = 1

func SetDelimiter(char string) {
	startDelimiter = char
	if char == "[" {
		endDelimiter = "]"
	} else {
		endDelimiter = char
	}
}

func IsEndDelimiter(char string) bool {

	if endDelimiter == char {
		delimiterDeep--
		if delimiterDeep == 0 {
			return true
		}
	}
	if startDelimiter == char {
		delimiterDeep++
	}

	return false
}

func CharIsKeySeparator(char string) bool {
	var keySeparators = [2]string{":", "="}
	for _, c := range keySeparators {
		if char == c {
			return true
		}
	}
	return false
}

func CharIsValueDelimiter(char string) bool {
	var keySeparators = [2]string{"[", "\""}
	for _, c := range keySeparators {
		if char == c {
			return true
		}
	}
	return false
}
