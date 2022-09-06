package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"strconv"

	"github.com/chen-huicheng/GSWGo/stl"
)

type Hello struct {
	N int
	f float32
}

func test1(mod interface{}, v int) interface{} {
	modType := reflect.TypeOf(mod).Elem()
	fmt.Println(modType)
	dst := reflect.New(modType).Interface()
	fmt.Println(dst)
	return dst
}

func test2() {
	type t struct {
		N int
		s string
	}
	n := t{42, "hello"}
	fmt.Println("before", n)
	rv := reflect.ValueOf(&n).Elem()
	rv.FieldByName("N").SetInt(7)
	fmt.Println("after", n)
}
func main() {
	test2()
	a := test1(&Hello{}, 2)
	fmt.Println(a)
	var i interface{}
	hello := Hello{1, 2.0}
	i = hello
	// 反射第一定律：反射可以将interface类型变量转换成反射对象
	fmt.Println("type:", reflect.TypeOf(i))
	fmt.Println("value:", reflect.ValueOf(i))
	// 反射第二定律：反射可以将反射对象还原成interface对象
	rv := reflect.ValueOf(i)
	ii := rv.Interface()
	if hello1, ok := ii.(Hello); ok {
		fmt.Println(hello1)
	}
	if i == ii {
		fmt.Println("interface{hello} == reflect.ValueOf(hello).Interface()")
	} else {
		fmt.Println("!=")
	}

	// 反射第三定律：反射对象可修改，value值必须是可设置的  即指针类型
	if rv.CanSet() {
		fmt.Println("hello can set")
	}

	rvx := reflect.ValueOf(&hello).Elem()
	// reflect.ValueOf(&hello).Elem().FieldByName("i").SetInt(7)
	if rvx.CanSet() {
		fmt.Println("before set", rvx, hello)
		rvx.FieldByName("N").SetInt(5)
		fmt.Println("end    set", rvx, hello)
	}
	test()
}

func test() {
	fmt.Println("\ntest func")
	t := reflect.TypeOf(3)
	fmt.Println(t.String())
	fmt.Println(t)

	var w io.Writer
	// w.Write([]byte("hello"))  //panic: runtime error: invalid memory address or nil pointer dereference

	w = os.Stdout
	w.Write([]byte("hello\n"))

	w = new(bytes.Buffer)
	w.Write([]byte("world\n"))

	fmt.Println(w)

	var a interface{}
	a = 3
	fmt.Println(a)
}

func SetValue(rv reflect.Value, value string) error {
	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		item, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		rv.SetInt(item)

	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		item, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return err
		}
		rv.SetUint(item)
	case reflect.String:
		rv.SetString(value)
	case reflect.Float32, reflect.Float64:
		item, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		rv.SetFloat(item)
	default: // float, complex, bool, chan, func, interface
		return fmt.Errorf("unsupported type: %s", rv.Type())
	}
	return nil

}
func InitValue(mod interface{}) error {
	fmt.Println(reflect.TypeOf(mod), reflect.TypeOf(mod).Elem())
	modType := reflect.TypeOf(mod).Elem()
	nf := modType.NumField()
	rv := reflect.ValueOf(mod).Elem()
	for i := 0; i < nf; i++ {
		fmt.Println(modType.Field(i).Tag)
		value := modType.Field(i).Tag.Get("init")
		// modType.Name()
		if value == "" {
			log.Printf("InitMod: struct '%s' field '%s' not 'init' tag ", modType.Name(), modType.Field(i).Name)
			continue
		}
		if value == "magicKey" {
			InitValue(rv.Field(i).Addr().Interface())
			continue
		}
		err := SetValue(rv.Field(i), value)
		if err != nil {
			return err
		}
	}
	return nil
	// rv := reflect.ValueOf(mod).Elem()
}
func InitLog() {
	// 创建、追加、读写，777，所有权限
	f, err := os.OpenFile("log.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		return
	}
	log.SetOutput(f)
	log.SetFlags(log.Llongfile | log.LstdFlags)
}
func main2() {
	InitLog()
	user := stl.User{}
	if err := InitValue(&user); err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)

	stu := stl.Student{}
	if err := InitValue(&stu); err != nil {
		fmt.Println(err)
	}
	fmt.Println(stu)
}
