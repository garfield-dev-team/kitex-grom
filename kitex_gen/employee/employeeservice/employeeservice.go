// Code generated by Kitex v0.9.1. DO NOT EDIT.

package employeeservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	employee "github.com/garfield-dev-team/kitex_gorm/kitex_gen/employee"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"CreateEmployee": kitex.NewMethodInfo(
		createEmployeeHandler,
		newEmployeeServiceCreateEmployeeArgs,
		newEmployeeServiceCreateEmployeeResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"UpdateEmployee": kitex.NewMethodInfo(
		updateEmployeeHandler,
		newEmployeeServiceUpdateEmployeeArgs,
		newEmployeeServiceUpdateEmployeeResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"DeleteEmployee": kitex.NewMethodInfo(
		deleteEmployeeHandler,
		newEmployeeServiceDeleteEmployeeArgs,
		newEmployeeServiceDeleteEmployeeResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetEmployee": kitex.NewMethodInfo(
		getEmployeeHandler,
		newEmployeeServiceGetEmployeeArgs,
		newEmployeeServiceGetEmployeeResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"ListEmployees": kitex.NewMethodInfo(
		listEmployeesHandler,
		newEmployeeServiceListEmployeesArgs,
		newEmployeeServiceListEmployeesResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	employeeServiceServiceInfo                = NewServiceInfo()
	employeeServiceServiceInfoForClient       = NewServiceInfoForClient()
	employeeServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return employeeServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return employeeServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return employeeServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "EmployeeService"
	handlerType := (*employee.EmployeeService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "employee",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.9.1",
		Extra:           extra,
	}
	return svcInfo
}

func createEmployeeHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*employee.EmployeeServiceCreateEmployeeArgs)
	realResult := result.(*employee.EmployeeServiceCreateEmployeeResult)
	success, err := handler.(employee.EmployeeService).CreateEmployee(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newEmployeeServiceCreateEmployeeArgs() interface{} {
	return employee.NewEmployeeServiceCreateEmployeeArgs()
}

func newEmployeeServiceCreateEmployeeResult() interface{} {
	return employee.NewEmployeeServiceCreateEmployeeResult()
}

func updateEmployeeHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*employee.EmployeeServiceUpdateEmployeeArgs)
	realResult := result.(*employee.EmployeeServiceUpdateEmployeeResult)
	success, err := handler.(employee.EmployeeService).UpdateEmployee(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newEmployeeServiceUpdateEmployeeArgs() interface{} {
	return employee.NewEmployeeServiceUpdateEmployeeArgs()
}

func newEmployeeServiceUpdateEmployeeResult() interface{} {
	return employee.NewEmployeeServiceUpdateEmployeeResult()
}

func deleteEmployeeHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*employee.EmployeeServiceDeleteEmployeeArgs)
	realResult := result.(*employee.EmployeeServiceDeleteEmployeeResult)
	success, err := handler.(employee.EmployeeService).DeleteEmployee(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newEmployeeServiceDeleteEmployeeArgs() interface{} {
	return employee.NewEmployeeServiceDeleteEmployeeArgs()
}

func newEmployeeServiceDeleteEmployeeResult() interface{} {
	return employee.NewEmployeeServiceDeleteEmployeeResult()
}

func getEmployeeHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*employee.EmployeeServiceGetEmployeeArgs)
	realResult := result.(*employee.EmployeeServiceGetEmployeeResult)
	success, err := handler.(employee.EmployeeService).GetEmployee(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newEmployeeServiceGetEmployeeArgs() interface{} {
	return employee.NewEmployeeServiceGetEmployeeArgs()
}

func newEmployeeServiceGetEmployeeResult() interface{} {
	return employee.NewEmployeeServiceGetEmployeeResult()
}

func listEmployeesHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*employee.EmployeeServiceListEmployeesArgs)
	realResult := result.(*employee.EmployeeServiceListEmployeesResult)
	success, err := handler.(employee.EmployeeService).ListEmployees(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newEmployeeServiceListEmployeesArgs() interface{} {
	return employee.NewEmployeeServiceListEmployeesArgs()
}

func newEmployeeServiceListEmployeesResult() interface{} {
	return employee.NewEmployeeServiceListEmployeesResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) CreateEmployee(ctx context.Context, req *employee.CreateEmployeeRequest) (r *employee.CreateEmployeeResponse, err error) {
	var _args employee.EmployeeServiceCreateEmployeeArgs
	_args.Req = req
	var _result employee.EmployeeServiceCreateEmployeeResult
	if err = p.c.Call(ctx, "CreateEmployee", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateEmployee(ctx context.Context, req *employee.UpdateEmployeeRequest) (r *employee.UpdateEmployeeResponse, err error) {
	var _args employee.EmployeeServiceUpdateEmployeeArgs
	_args.Req = req
	var _result employee.EmployeeServiceUpdateEmployeeResult
	if err = p.c.Call(ctx, "UpdateEmployee", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteEmployee(ctx context.Context, req *employee.DeleteEmployeeRequest) (r *employee.DeleteEmployeeResponse, err error) {
	var _args employee.EmployeeServiceDeleteEmployeeArgs
	_args.Req = req
	var _result employee.EmployeeServiceDeleteEmployeeResult
	if err = p.c.Call(ctx, "DeleteEmployee", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetEmployee(ctx context.Context, req *employee.GetEmployeeRequest) (r *employee.GetEmployeeResponse, err error) {
	var _args employee.EmployeeServiceGetEmployeeArgs
	_args.Req = req
	var _result employee.EmployeeServiceGetEmployeeResult
	if err = p.c.Call(ctx, "GetEmployee", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ListEmployees(ctx context.Context, req *employee.ListEmployeesRequest) (r *employee.ListEmployeesResponse, err error) {
	var _args employee.EmployeeServiceListEmployeesArgs
	_args.Req = req
	var _result employee.EmployeeServiceListEmployeesResult
	if err = p.c.Call(ctx, "ListEmployees", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
