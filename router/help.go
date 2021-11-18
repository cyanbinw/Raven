package router

import (
	"github.com/gin-gonic/gin"
	"github.com/swirling-melodies/Raven/controllers"
	"reflect"
)

type Routers struct {
	controllers.TestRouters
	controllers.BillRouters
	controllers.DesireRouters
	controllers.InvestmentRouters
	controllers.LogRouters
	controllers.TargetRouters
	controllers.UserRouters
	controllers.WorkRouters
}

func NewRouters() *Routers {
	i := &Routers{}
	i.TestRouters.SetWork()
	i.BillRouters.SetWork()
	i.DesireRouters.SetWork()
	i.InvestmentRouters.SetWork()
	i.LogRouters.SetWork()
	i.TargetRouters.SetWork()
	i.UserRouters.SetWork()
	i.WorkRouters.SetWork()

	/*	data := reflect.ValueOf(i).Elem()
		for item :=0; item < data.NumField(); item++{
			data.Field(item).Interface().(controllers.IRouters).SetWork()
		}*/

	return i
}

func isBlank(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.String:
		return value.Len() == 0
	case reflect.Bool:
		return !value.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return value.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return value.IsNil()
	}
	return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
}

func loadRouter(item []controllers.Path, c *gin.RouterGroup) {
	for _, i := range item {
		if i.Request == controllers.Get {
			c.GET(i.Route, i.Action)
		} else if i.Request == controllers.Post {
			c.POST(i.Route, i.Action)
		} else if i.Request == controllers.Put {
			c.PUT(i.Route, i.Action)
		} else if i.Request == controllers.Delete {
			c.DELETE(i.Route, i.Action)
		} else {

		}
	}
}
