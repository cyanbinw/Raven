package test

import (
	"github.com/swirling-melodies/Raven/common"
	"testing"
)

func TestReadJSON(t *testing.T) {
	data, err := common.ReadJSON("../initialData/investmentTypeinitialData.json")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(data)
}
