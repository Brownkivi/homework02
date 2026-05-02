package logic

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID int
}

func (e *Employee) PrintInfo() {
	println("Name:", e.Name)
	println("Age:", e.Age)
	println("Employee ID:", e.EmployeeID)
}
