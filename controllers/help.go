package controllers

import (
	"github.com/gin-gonic/gin"
	"reflect"
	"strings"
)

const (
	Get    RequestEnum = 1
	Post               = 2
	Put                = 3
	Delete             = 4
)

type RequestEnum int

type ReturnData struct {
	Successful bool        `json:"successful"`
	Error      string      `json:"error"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

type Path struct {
	Action  gin.HandlerFunc
	Request RequestEnum
	Route   string
}

type TestRouters struct {
	Controller string
	Work       []Path
}

func (t *TestRouters) GetController() {
	t.Controller = "Test"
}

func (t *TestRouters) SetWork() {
	/*if t.Controller == "" {
		t.GetController()
	}*/
	d := reflect.ValueOf(t).Elem()
	for i := 0; i < d.NumMethod(); i++ {
		k := d.Type()
		data := Path{}
		if !strings.Contains(k.Method(i).Name, "RE") {
			continue
		}
		name := k.Method(i)
		v := d.Method(i)
		z := v.Interface()
		strList := strings.Split(name.Name, "RE")
		if len(strList) > 0 {
			path := "/" + t.Controller + "/" + strList[0]
			data.Action = z.(func(ctx *gin.Context))

			data.Request = GetEnumValue(strList[1])
			data.Route = path
			t.Work = append(t.Work, data)
		}
	}
}

func (t TestRouters) GetWork() []Path {
	return t.Work
}

type BillRouters struct {
	Controller string
	Work       []Path
}

func (t *BillRouters) GetController() {
	t.Controller = "Bill"
}

func (t *BillRouters) SetWork() {
	if t.Controller == "" {
		t.GetController()
	}
	d := reflect.ValueOf(t).Elem()
	for i := 0; i < d.NumMethod(); i++ {
		k := d.Type()
		data := Path{}
		if !strings.Contains(k.Method(i).Name, "RE") {
			continue
		}
		name := k.Method(i)
		v := d.Method(i)
		z := v.Interface()
		strList := strings.Split(name.Name, "RE")
		if len(strList) > 0 {
			path := "/" + t.Controller + "/" + strList[0]
			data.Action = z.(func(ctx *gin.Context))
			data.Request = GetEnumValue(strList[1])
			data.Route = path
			t.Work = append(t.Work, data)
		}
	}
}

func (t BillRouters) GetWork() []Path {
	return t.Work
}

type DesireRouters struct {
	Controller string
	Work       []Path
}

func (t *DesireRouters) GetController() {
	t.Controller = "Desire"
}

func (t *DesireRouters) SetWork() {
	if t.Controller == "" {
		t.GetController()
	}
	d := reflect.ValueOf(t).Elem()
	for i := 0; i < d.NumMethod(); i++ {
		k := d.Type()
		data := Path{}
		if !strings.Contains(k.Method(i).Name, "RE") {
			continue
		}
		name := k.Method(i)
		v := d.Method(i)
		z := v.Interface()
		strList := strings.Split(name.Name, "RE")
		if len(strList) > 0 {
			path := "/" + t.Controller + "/" + strList[0]
			data.Action = z.(func(ctx *gin.Context))
			data.Request = GetEnumValue(strList[1])
			data.Route = path
			t.Work = append(t.Work, data)
		}
	}
}

func (t DesireRouters) GetWork() []Path {
	return t.Work
}

type InvestmentRouters struct {
	Controller string
	Work       []Path
}

func (t *InvestmentRouters) GetController() {
	t.Controller = "Investment"
}

func (t *InvestmentRouters) SetWork() {
	if t.Controller == "" {
		t.GetController()
	}
	d := reflect.ValueOf(t).Elem()
	for i := 0; i < d.NumMethod(); i++ {
		k := d.Type()
		data := Path{}
		if !strings.Contains(k.Method(i).Name, "RE") {
			continue
		}
		name := k.Method(i)
		v := d.Method(i)
		z := v.Interface()
		strList := strings.Split(name.Name, "RE")
		if len(strList) > 0 {
			path := "/" + t.Controller + "/" + strList[0]
			data.Action = z.(func(ctx *gin.Context))
			data.Request = GetEnumValue(strList[1])
			data.Route = path
			t.Work = append(t.Work, data)
		}
	}
}

func (t InvestmentRouters) GetWork() []Path {
	return t.Work
}

type LogRouters struct {
	Controller string
	Work       []Path
}

func (t *LogRouters) GetController() {
	t.Controller = "Log"
}

func (t *LogRouters) SetWork() {
	if t.Controller == "" {
		t.GetController()
	}
	d := reflect.ValueOf(t).Elem()
	for i := 0; i < d.NumMethod(); i++ {
		k := d.Type()
		data := Path{}
		if !strings.Contains(k.Method(i).Name, "RE") {
			continue
		}
		name := k.Method(i)
		v := d.Method(i)
		z := v.Interface()
		strList := strings.Split(name.Name, "RE")
		if len(strList) > 0 {
			path := "/" + t.Controller + "/" + strList[0]
			data.Action = z.(func(ctx *gin.Context))
			data.Request = GetEnumValue(strList[1])
			data.Route = path
			t.Work = append(t.Work, data)
		}
	}
}

func (t LogRouters) GetWork() []Path {
	return t.Work
}

type TargetRouters struct {
	Controller string
	Work       []Path
}

func (t *TargetRouters) GetController() {
	t.Controller = "Target"
}

func (t *TargetRouters) SetWork() {
	if t.Controller == "" {
		t.GetController()
	}
	d := reflect.ValueOf(t).Elem()
	for i := 0; i < d.NumMethod(); i++ {
		k := d.Type()
		data := Path{}
		if !strings.Contains(k.Method(i).Name, "RE") {
			continue
		}
		name := k.Method(i)
		v := d.Method(i)
		z := v.Interface()
		strList := strings.Split(name.Name, "RE")
		if len(strList) > 0 {
			path := "/" + t.Controller + "/" + strList[0]
			data.Action = z.(func(ctx *gin.Context))
			data.Request = GetEnumValue(strList[1])
			data.Route = path
			t.Work = append(t.Work, data)
		}
	}
}

func (t TargetRouters) GetWork() []Path {
	return t.Work
}

type UserRouters struct {
	Controller string
	Work       []Path
}

func (t *UserRouters) GetController() {
	t.Controller = "User"
}

func (t *UserRouters) SetWork() {
	if t.Controller == "" {
		t.GetController()
	}
	d := reflect.ValueOf(t).Elem()
	for i := 0; i < d.NumMethod(); i++ {
		k := d.Type()
		data := Path{}
		if !strings.Contains(k.Method(i).Name, "RE") {
			continue
		}
		name := k.Method(i)
		v := d.Method(i)
		z := v.Interface()
		strList := strings.Split(name.Name, "RE")
		if len(strList) > 0 {
			path := "/" + t.Controller + "/" + strList[0]
			data.Action = z.(func(ctx *gin.Context))
			data.Request = GetEnumValue(strList[1])
			data.Route = path
			t.Work = append(t.Work, data)
		}
	}
}

func (t UserRouters) GetWork() []Path {
	return t.Work
}

type WorkRouters struct {
	Controller string
	Work       []Path
}

func (t *WorkRouters) GetController() {
	t.Controller = "Work"
}

func (t *WorkRouters) SetWork() {
	if t.Controller == "" {
		t.GetController()
	}
	d := reflect.ValueOf(t).Elem()
	for i := 0; i < d.NumMethod(); i++ {
		k := d.Type()
		data := Path{}
		if !strings.Contains(k.Method(i).Name, "RE") {
			continue
		}
		name := k.Method(i)
		v := d.Method(i)
		z := v.Interface()
		strList := strings.Split(name.Name, "RE")
		if len(strList) > 0 {
			path := "/" + t.Controller + "/" + strList[0]
			data.Action = z.(func(ctx *gin.Context))
			data.Request = GetEnumValue(strList[1])
			data.Route = path
			t.Work = append(t.Work, data)
		}
	}
}

func (t WorkRouters) GetWork() []Path {
	return t.Work
}

type IRouters interface {
	GetController()
	SetWork()
	GetWork() []Path
}

func GetEnumValue(str string) RequestEnum {
	switch str {
	case "Get":
		return Get
	case "Post":
		return Post
	case "Put":
		return Put
	case "Delete":
		return Delete
	default:
		return 0
	}
}
