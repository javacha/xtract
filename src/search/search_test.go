package search

import (
	"testing"
)

var testingBuffer string = "timeFake=\"yo\" time=\"2024-04-05 11:39:27.864\" level=INFO module=helpers head_rest=\"User-Agent:[Mozilla/4.0 [en] (WinNT; I)] X-Forwarded-For:[76.249.26.144, 172.29.99.68, 172.29.98.125, 10.131.0.1] X-Forwarded-Port:[443] X-Forwarded-Proto:[https] Accept:[image/gif, */*] Vnd.bbva.propagated-Url:[http://intra-apigateway-play.arg.igrupobbva/coelsa-cedip/consultas/ListaCedip/33********69?filtros=%22documento_tenedor%22+%5Beq%5D+__33663519269__+AND+%28%22estado%22+%5Beq%5D+__ACTIVO__+OR+%22estado%22+%5Beq%5D+__ACTIVO-PENDIENTE__+OR+%22estado%22+%5Beq%5D+__DEPOSITADO__+OR+%22estado%22+%5Beq%5D+__PRESENTADO__%29&cantPag=20&pag=3] Vnd.bbva.user-Id:[oficiosembargosjud] X-Real-Ip:[172.29.99.68] Content-Type:[application/x-www-form-urlencoded] Forwarded:[for=172.29.98.125;host=intra-apigateway-play.arg.igrupobbva;proto=https;proto-version=] Vnd.bbva.user-Scope:[] Accept-Encoding:[gzip] Authorization:[Basic b2ZpY********qNHVk] Cookie:[cookiesession1=678B28D71AE730AFA66D153E1F7AEE1E]  X-Forwarded-Host:[intra-apigateway-play.arg.igrupobbva]\""

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestValueWithComillas(t *testing.T) {
	t.Log("arranvamos con ", testingBuffer)
	want := "\"2024-04-05 11:39:27.864\""
	msg, err := GetSmartValue(testingBuffer, "time")
	if want != msg || err != nil {
		t.Fatalf(`TestValueWithComillas DIFIERE %#q <=> %#q`, msg, want)
	}
}

func TestValuePlain(t *testing.T) {
	t.Log("arranvamos con ", testingBuffer)
	want := "INFO"
	msg, err := GetSmartValue(testingBuffer, "level")
	if want != msg || err != nil {
		t.Fatalf(`TestValuePlain | obtenido %#q <=> esperado %#q`, msg, want)
	}
}

func TestValueWithCorchetes(t *testing.T) {
	t.Log("arranvamos con ", testingBuffer)

	want := "[Mozilla/4.0 [en] (WinNT; I)]"
	msg, err := GetSmartValue(testingBuffer, "User-Agent")
	if want != msg || err != nil {
		t.Fatalf(`TestValueWithCorchetes | obtenido %#q <=> esperado %#q`, msg, want)
	}
}
