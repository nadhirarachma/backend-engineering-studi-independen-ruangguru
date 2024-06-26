package main

type Employee interface {
	GetBonus() float64
}

type Junior struct {
	Name string
	BaseSalary int
	WorkingMonth int
}

type Senior struct {
	Name string
	BaseSalary int
	WorkingMonth int
	PerformanceRate float64
}

type Manager struct {
	Name string
	BaseSalary int
	WorkingMonth int
	PerformanceRate float64
	BonusManagerRate float64
}

func (j Junior) GetBonus() float64 {
	return float64(j.BaseSalary) * (float64(j.WorkingMonth) / 12.0)
}

func (s Senior) GetBonus() float64 {
	return 2 * float64(s.BaseSalary) * (float64(s.WorkingMonth) / 12.0) + (s.PerformanceRate * float64(s.BaseSalary))
}

func (m Manager) GetBonus() float64 {
	return 2 * float64(m.BaseSalary) * (float64(m.WorkingMonth) / 12.0) + (m.PerformanceRate * float64(m.BaseSalary)) + (m.BonusManagerRate * float64(m.BaseSalary))
}

func EmployeeBonus(employee Employee) float64 {
	return employee.GetBonus()
}

func TotalEmployeeBonus(employees []Employee) float64 {
	total := 0.0
	for _, employee := range employees {
		total += employee.GetBonus()
	}
	return total
}
