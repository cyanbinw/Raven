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

type TestReuter struct {
	Controller string
	Work       []Path
}

func (t *TestReuter) GetController() {
	t.Controller = "Test"
}

func (t *TestReuter) SetWork() {
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

func (t TestReuter) GetWork() []Path {
	return t.Work
}

type BillReuter struct {
	Controller string
	Work       []Path
}

func (t *BillReuter) GetController() {
	t.Controller = "Bill"
}

func (t *BillReuter) SetWork() {
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

func (t BillReuter) GetWork() []Path {
	return t.Work
}

type DesireReuter struct {
	Controller string
	Work       []Path
}

func (t *DesireReuter) GetController() {
	t.Controller = "Desire"
}

func (t *DesireReuter) SetWork() {
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

func (t DesireReuter) GetWork() []Path {
	return t.Work
}

type InvestmentReuter struct {
	Controller string
	Work       []Path
}

func (t *InvestmentReuter) GetController() {
	t.Controller = "Investment"
}

func (t *InvestmentReuter) SetWork() {
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

func (t InvestmentReuter) GetWork() []Path {
	return t.Work
}

type LogReuter struct {
	Controller string
	Work       []Path
}

func (t *LogReuter) GetController() {
	t.Controller = "Log"
}

func (t *LogReuter) SetWork() {
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

func (t LogReuter) GetWork() []Path {
	return t.Work
}

type TargetReuter struct {
	Controller string
	Work       []Path
}

func (t *TargetReuter) GetController() {
	t.Controller = "Target"
}

func (t *TargetReuter) SetWork() {
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

func (t TargetReuter) GetWork() []Path {
	return t.Work
}

type UserReuter struct {
	Controller string
	Work       []Path
}

func (t *UserReuter) GetController() {
	t.Controller = "User"
}

func (t *UserReuter) SetWork() {
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

func (t UserReuter) GetWork() []Path {
	return t.Work
}

type WorkReuter struct {
	Controller string
	Work       []Path
}

func (t *WorkReuter) GetController() {
	t.Controller = "Work"
}

func (t *WorkReuter) SetWork() {
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

func (t WorkReuter) GetWork() []Path {
	return t.Work
}

type IReuter interface {
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
