include "common.thrift"

namespace go salary

struct Salary {
    1: i64 id
    2: string created_at
    3: string updated_at
    4: i64 employee_id
    5: common.EnterpriseCode enterprise_code
    6: string month
    7: i32 salary
}

struct CreateSalaryRequest {
    1: i64 employee_id
    2: common.EnterpriseCode enterprise_code
    3: string month
    4: i32 salary
}

struct CreateSalaryResponse {
    1: Salary salary
}

struct UpdateSalaryRequest {
    1: i64 id
    2: i64 employee_id
    3: common.EnterpriseCode enterprise_code
    4: string month
    5: i32 salary
}

struct UpdateSalaryResponse {
    1: Salary salary
}

struct DeleteSalaryRequest {
    1: i64 id
}

struct DeleteSalaryResponse {
    1: bool success
}

struct GetSalaryRequest {
    1: i64 id
}

struct GetSalaryResponse {
    1: Salary salary
}

struct ListSalariesRequest {
    1: i64 employee_id
    2: common.EnterpriseCode enterprise_code
    3: string month
    4: i64 page
    5: i64 page_size
}

struct ListSalariesResponse {
    1: list<Salary> salaries
    2: i64 total
}

service SalaryService {
    CreateSalaryResponse CreateSalary(1: CreateSalaryRequest req)
    UpdateSalaryResponse UpdateSalary(1: UpdateSalaryRequest req)
    DeleteSalaryResponse DeleteSalary(1: DeleteSalaryRequest req)
    GetSalaryResponse GetSalary(1: GetSalaryRequest req)
    ListSalariesResponse ListSalaries(1: ListSalariesRequest req)
}