package main

import (
	"reflect"
	"testing"
)

func TestConvSegundos(t *testing.T) {
	table := []struct {
		input  int
		output string
	}{
		{
			7650,
			"02:07:30",
		},
		// {
		// 	5800,
		// 	"01:36:40",
		// },
	}
	for _, test := range table {
		got := ConvSegundos(test.input)
		if !reflect.DeepEqual(got, test.output) {
			t.Errorf("Função retornou %#v, porém era para retornar %#v, dado os valores de %#v", got, test.output, test.input)
		}
	}
}
