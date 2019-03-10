package tempconv

import "testing"

func TestCelsius_String(t *testing.T) {
	testcases := map[Celsius]string{
		-10:  "-10℃",
		100:  "100℃",
		36.5: "36.5℃",
	}

	for val, want := range testcases {
		if got := val.String(); got != want {
			t.Errorf("Celsius.String want:%v, got:%v", want, got)
		}
	}
}

func TestFahrenheit_String(t *testing.T) {
	testcases := map[Fahrenheit]string{
		-10:  "-10℉",
		100:  "100℉",
		36.5: "36.5℉",
	}

	for val, want := range testcases {
		if got := val.String(); got != want {
			t.Errorf("Fahrenheit.String want:%v, got:%v", want, got)
		}
	}
}

func TestKelvin_String(t *testing.T) {
	testcases := map[Kelvin]string{
		-10:  "-10K",
		100:  "100K",
		36.5: "36.5K",
	}

	for val, want := range testcases {
		if got := val.String(); got != want {
			t.Errorf("Kelvin.String want:%v, got:%v", want, got)
		}
	}
}

func TestTranslate(t *testing.T) {
	c := Celsius(10)
	f := CToF(c)
	if want := Fahrenheit(50); f != want {
		t.Fatalf("CToF invalid want:%v, got:%v", want, f)
	}

	k := FToK(f)
	if want := Kelvin(283.15); k != want {
		t.Fatalf("FToK invalid want:%v, got:%v", want, k)
	}

	f = KToF(k)
	if want := Fahrenheit(50); f != want {
		t.Fatalf("KToF invalid want:%v, got:%v", want, f)
	}

	c = KToC(k)
	if want := Celsius(10); c != want {
		t.Fatalf("KToC invalid want:%v, got:%v", want, c)
	}
}
