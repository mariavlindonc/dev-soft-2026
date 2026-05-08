package util

// Using Person from structs.go

type Employee struct {
	name      string
	age       int
	position  string
	seniority int
}

// struct embedding allows Manager to inherit fields and methods from Employee
type Manager struct {
	Employee   // Embedding Employee struct
	department string
}

// transformToEmployee takes a Person and transforms it into an Employee
func transformToEmployee(p Person) Employee {
	return Employee{
		name:      p.name,
		age:       p.age,
		position:  "Unknown", // Default position, can be updated later
		seniority: 0,         // Default seniority, can be updated later
	}
}

func PrintFunctionsStruct() {
	person := Person{name: "Alice", age: 30}

	// Transforming Person to Employee
	employee := transformToEmployee(person)

	// Updating employee's position and seniority
	employee.position = "Software Engineer"
	employee.seniority = 5

	println("Employee Name:", employee.name)
	println("Employee Age:", employee.age)
	println("Employee Position:", employee.position)
	println("Employee Seniority:", employee.seniority)

	// Creating a Manager using the Employee struct
	Manager := Manager{
		Employee:   employee,
		department: "Engineering",
	}

	println("Manager Name:", Manager.name)
	println("Manager Age:", Manager.age)
	println("Manager Position:", Manager.position)
	println("Manager Seniority:", Manager.seniority)
	println("Manager Department:", Manager.department)
}
