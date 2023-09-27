package main

import (
	"strings"
	"testing"

	"github.com/bxcodec/faker/v4"
	"github.com/stretchr/testify/assert"
)

func Test_Problem271(t *testing.T) {

	var mockData []string
	err := faker.FakeData(&mockData)
	if err != nil {
		t.Fatal("create mock data failed")
	}
	mockDataWithSymbol := make([]string, len(mockData))
	for idx, val := range mockData {
		mockDataWithSymbol[idx] = strings.Replace(val, "a", "#", -1)
		mockDataWithSymbol[idx] = strings.Replace(val, "c", "#", -1)
		mockDataWithSymbol[idx] = strings.Replace(val, "x", ";", -1)
		mockDataWithSymbol[idx] = strings.Replace(val, "z", ":", -1)
	}
	qs := [][]string{
		{"weeeeeeeee", "say", ":", "yes"},
		mockData,
		mockDataWithSymbol,
	}

	codec := Codec{}
	for _, q := range qs {
		decodeData := codec.Decode(codec.Encode(q))
		assert.Equal(t, len(q), len(decodeData))
		for idx := range q {
			assert.Equal(t, q[idx], decodeData[idx])
		}
	}
}
