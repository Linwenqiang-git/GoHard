// 编写SQL查询以查找每个部门中薪资最高的员工。

select d.name as Department ,e.name as Employee ,e.Salary as Salary 
from Employee e
left join Department d on d.id = e.departmentId 
left join (
    select departmentId,max(salary) as Salary 
    from Employee 
    group by departmentId 
) s on s.departmentId = d.id
where e.salary >= s.salary

SELECT
    Department.name AS 'Department',
    Employee.name AS 'Employee',
    Salary
FROM
    Employee
        JOIN
    Department ON Employee.DepartmentId = Department.Id
WHERE
    (Employee.DepartmentId , Salary) IN
    (   SELECT
            DepartmentId, MAX(Salary)
        FROM
            Employee
        GROUP BY DepartmentId
	)
;
