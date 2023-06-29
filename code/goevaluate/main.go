package main

import (
	"fmt"
	"runtime"

	"github.com/Knetic/govaluate"
)

type Valuate struct {
	exFunc  map[string]govaluate.ExpressionFunction
	express *govaluate.EvaluableExpression
}

func (v *Valuate) AddFunc(funcName string, exFunc govaluate.ExpressionFunction) {
	if v.exFunc == nil {
		v.exFunc = make(map[string]govaluate.ExpressionFunction)
	}
	v.exFunc[funcName] = exFunc
}

func (v *Valuate) Calculate(expr string) (bool, error) {
	exp, err := govaluate.NewEvaluableExpressionWithFunctions(expr, v.exFunc)
	if err != nil {
		return false, fmt.Errorf("表达式错误：" + err.Error())
	}
	res, err := exp.Evaluate(map[string]interface{}{})
	if err != nil {
		return false, fmt.Errorf("表达式错误：" + err.Error())
	}
	if _, ok := res.(bool); !ok {
		return false, fmt.Errorf("表达式错误：返回值非bool")
	}
	return res.(bool), nil
}
func main() {
	runtime.GOMAXPROCS(14)
	fmt.Println(runtime.NumCPU())
	var valuate Valuate
	valuate.AddFunc("count", countTrueFunc)
	res, err := valuate.Calculate("count(true,false,true)==1")
	fmt.Println(res, err)
}

func countTrueFunc(arguments ...interface{}) (interface{}, error) {
	cnt := float64(0)
	for _, a := range arguments {
		if a == nil {
			continue
		}
		if _, ok := a.(bool); !ok {
			fmt.Printf("类型错误: CountTrue, %v", a)
			return nil, fmt.Errorf("表达式错误：CountTrue参数类型错误")
		}
		if a == true {
			cnt++
		}
	}
	return cnt, nil
}
