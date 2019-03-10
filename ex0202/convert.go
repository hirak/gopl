package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/hirak/gopl/ex0201/tempconv"
)

const (
	length      = "length"
	weight      = "weight"
	temperature = "temperature"
)

func main() {
	var t string
	flag.StringVar(&t, "t", "temperature", "temperature, length, weight")
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		for _, arg := range args {
			process(t, arg)
		}
	} else {
		scan := bufio.NewScanner(os.Stdin)
		for scan.Scan() {
			process(t, scan.Text())
		}
	}
}

func process(t, s string) {
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "convert: %v\n", err)
		os.Exit(1)
	}

	switch t {
	case length:
		m := Meter(v)
		f := Foot(v)
		fmt.Printf("%s = %s, %s = %s\n", m, m.ToF(), f, f.ToM())
	case weight:
		k := Kilogram(v)
		p := Pound(v)
		fmt.Printf("%s = %s, %s = %s\n", k, k.ToP(), p, p.ToK())
	case temperature:
		f := tempconv.Fahrenheit(v)
		c := tempconv.Celsius(v)
		fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
	default:
		fmt.Printf("unknown type: %q\n", t)
		os.Exit(1)
	}
}

type (
	Meter float64
	Foot  float64

	Pound    float64
	Kilogram float64
)

func (m Meter) String() string {
	return fmt.Sprintf("%gm", m)
}

func (m Meter) ToF() Foot {
	return Foot(m / 0.3048)
}

func (f Foot) String() string {
	return fmt.Sprintf("%gft", f)
}

func (f Foot) ToM() Meter {
	return Meter(f * 0.3048)
}

func (p Pound) String() string {
	return fmt.Sprintf("%glb", p)
}

func (p Pound) ToK() Kilogram {
	return Kilogram(p * 0.45359237)
}

func (g Kilogram) String() string {
	return fmt.Sprintf("%gkg", g)
}

func (g Kilogram) ToP() Pound {
	return Pound(g / 0.45359237)
}
