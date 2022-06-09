package test

import (
	"github.com/swirling-melodies/Raven/common"
	"github.com/swirling-melodies/Raven/database"
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

func TestInitRedis(t *testing.T) {
	database.V8Example()
}
