include "common.thrift"

namespace go employee

struct Employee {
    1: i64 id
    2: string created_at
    3: string updated_at
    4: common.EnterpriseCode enterprise_code
    5: string name
    6: common.Gender gender
    7: i64 age
    8: string avatar
    9: string mobile
    10: string introduce
    11: string id_card_num
    12: string bank_card_num
}

struct CreateEmployeeRequest {
    1: common.EnterpriseCode enterprise_code
    2: string name
    3: common.Gender gender
    4: i64 age
    5: string avatar
    6: string mobile
    7: string introduce
    8: string id_card_num
    9: string bank_card_num
}

struct CreateEmployeeResponse {
    1: Employee employee
}

struct UpdateEmployeeRequest {
    1: i64 id
    2: common.EnterpriseCode enterprise_code
    3: string name
    4: common.Gender gender
    5: i64 age
    6: string avatar
    7: string mobile
    8: string introduce
    9: string id_card_num
    10: string bank_card_num
}

struct UpdateEmployeeResponse {
    1: Employee employee
}

struct DeleteEmployeeRequest {
    1: i64 id
}

struct DeleteEmployeeResponse {
    1: bool success
}

struct GetEmployeeRequest {
    1: i64 id
}

struct GetEmployeeResponse {
    1: Employee employee
}

struct ListEmployeesRequest {
    1: common.EnterpriseCode enterprise_code
    2: i64 page
    3: i64 page_size
}

struct ListEmployeesResponse {
    1: list<Employee> employees
    2: i64 total
}

service EmployeeService {
    CreateEmployeeResponse CreateEmployee(1: CreateEmployeeRequest req)
    UpdateEmployeeResponse UpdateEmployee(1: UpdateEmployeeRequest req)
    DeleteEmployeeResponse DeleteEmployee(1: DeleteEmployeeRequest req)
    GetEmployeeResponse GetEmployee(1: GetEmployeeRequest req)
    ListEmployeesResponse ListEmployees(1: ListEmployeesRequest req)
}