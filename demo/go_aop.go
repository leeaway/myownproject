package demo

import (
	"bou.ke/monkey"
	"fmt"
	"reflect"
)

//连接点
type JoinPoint struct {
	Receiver interface{}
	Method   reflect.Method
	Params   []reflect.Value
	Result   []reflect.Value
}

func NewJoinPoint(receiver interface{}, params []reflect.Value, method reflect.Method) *JoinPoint {
	point := &JoinPoint{
		Receiver: receiver,
		Params:   params,
		Method:   method,
	}

	fn := method.Func
	fnType := fn.Type()
	nout := fnType.NumOut()
	point.Result = make([]reflect.Value, nout)
	for i := 0; i < nout; i++ {
		//默认返回空值
		point.Result[i] = reflect.Zero(fnType.Out(i))
	}

	return point
}

//切面接口
type AspectInterface interface {
	Before(point *JoinPoint) bool
	After(point *JoinPoint)
	Finally(point *JoinPoint)
	GetAspectExpress() string
}

//切面列表
var aspectMap = make(map[string]AspectInterface, 0)

//注册切点
func RegisterPoint(pointType reflect.Type) {
	pkgPth := pointType.PkgPath()
	receiverName := pointType.Name()
	if pointType.Kind() == reflect.Ptr {
		pkgPth = pointType.Elem().PkgPath()
		receiverName = pointType.Elem().Name()
	}
	for i := 0; i < pointType.NumMethod(); i++ {
		method := pointType.Method(i)
		//方法位置字符串"包名.接收者.方法名"，用于匹配代理
		methodLocation := fmt.Sprintf("%s.%s.%s", pkgPth, receiverName, method.Name)
		var guard *monkey.PatchGuard
		var proxy = func(in []reflect.Value) []reflect.Value {
			guard.Unpatch()
			defer guard.Restore()
			receiver := in[0]
			point := NewJoinPoint(receiver, in[1:], method)
			defer finallyProcessed(point, methodLocation)
			if !beforeProcessed(point, methodLocation) {
				return point.Result
			}
			point.Result = receiver.MethodByName(method.Name).Call(in[1:])
			afterProcessed(point, methodLocation)
			return point.Result
		}
		//动态创建代理函数
		proxyFn := reflect.MakeFunc(method.Func.Type(), proxy)
		//利用monkey框架替换代理函数
		guard = monkey.PatchInstanceMethod(pointType, method.Name, proxyFn.Interface())
	}
}

//注册切面,methodLocation为包名.接收者.方法名
func RegisterAspect(methodLocation string, aspect AspectInterface) {
	aspectMap[methodLocation] = aspect
}

//前置处理
func beforeProcessed(point *JoinPoint, methodLocation string) bool {
	aspect, ok := aspectMap[methodLocation]
	if !ok || !aspect.Before(point) {
		return false
	}
	return true
}

//后置处理
func afterProcessed(point *JoinPoint, methodLocation string) {
	aspect, ok := aspectMap[methodLocation]
	if !ok {
		return
	}
	aspect.After(point)
}

//最终处理
func finallyProcessed(point *JoinPoint, methodLocation string) {
	aspect, ok := aspectMap[methodLocation]
	if !ok {
		return
	}
	aspect.Finally(point)
}
