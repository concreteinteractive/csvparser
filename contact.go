package parser

import "time"

//ContactGetter is a wrapper for the different types of contact structs we want to test
type ContactGetter interface {
	GetFirstName() string
	GetLastName() string
	GetWorking() bool
	GetAge() int
	GetSalary32() float32
	GetSalary64() float64
	GetVacationDays() uint
	GetTitle() string
	GetEmail() string
	GetBirthdate() time.Time
	GetDescription() string
}

//ExampleContact1 specifies all the csv struct tag fields
type ExampleContact1 struct {
	FirstName    string    `csv:"0"`
	LastName     string    `csv:"1"`
	Working      bool      `csv:"2"`
	Age          int       `csv:"3"`
	Salary32     float32   `csv:"4"`
	Salary64     float64   `csv:"5"`
	VacationDays uint      `csv:"6"`
	Title        string    `csv:"7"`
	Email        string    `csv:"8"`
	Birthdate    time.Time `csv:"9" csvDate:"2006-01-02"`
	Description  string    `csv:"10"`
}

//ExampleContact2 specifies some of the csv struct tag fields
type ExampleContact2 struct {
	FirstName    string
	LastName     string `csv:"1"`
	Working      bool
	Age          int
	Salary32     float32
	Salary64     float64
	VacationDays uint
	Title        string `csv:"7"`
	Email        string
	Birthdate    time.Time `csv:"9" csvDate:"2006-01-02"`
	Description  string    `csv:"10"`
}

//ExampleContact3 specifies any of the csv struct tag fields
type ExampleContact3 struct {
	FirstName    string
	LastName     string
	Working      bool
	Age          int
	Salary32     float32
	Salary64     float64
	VacationDays uint
	Title        string
	Email        string
	Birthdate    time.Time `csvDate:"2006-01-02"`
	Description  string
}

//ExampleContact4 specifies csv columns names
type ExampleContact4 struct {
	FirstName    string    `csvColumn:"FirstName"`
	LastName     string    `csvColumn:"LastName"`
	Working      bool      `csvColumn:"Working"`
	Age          int       `csvColumn:"Age"`
	Salary32     float32   `csvColumn:"Salary32"`
	Salary64     float64   `csvColumn:"Salary64"`
	VacationDays uint      `csvColumn:"VacationDays"`
	Title        string    `csvColumn:"Title"`
	Email        string    `csvColumn:"Email"`
	Birthdate    time.Time `csvColumn:"Birthdate" csvDate:"2006-01-02"`
	Description  string    `csvColumn:"Description"`
}

//ExampleContactInvalidBoolean is made for testing bool wrong field
type ExampleContactInvalidBoolean struct {
	FirstName    string
	LastName     bool //this is not a bool!
	Working      bool
	Age          int
	Salary32     float32
	Salary64     float64
	VacationDays uint
	Title        string
	Email        string
	Birthdate    time.Time
	Description  string
}

//ExampleContactInvalidUint is made for testing uint wrong field
type ExampleContactInvalidUint struct {
	FirstName    string
	LastName     uint //this is not an uint!
	Working      bool
	Age          int
	Salary32     float32
	Salary64     float64
	VacationDays uint
	Title        string
	Email        string
	Birthdate    time.Time
	Description  string
}

//ExampleContactInvalidInt is made for testing uint wrong field
type ExampleContactInvalidInt struct {
	FirstName    string
	LastName     int //this is not an int!
	Working      bool
	Age          int
	Salary32     float32
	Salary64     float64
	VacationDays uint
	Title        string
	Email        string
	Birthdate    time.Time
	Description  string
}

//ExampleContactInvalidFloat32 is made for testing uint wrong field
type ExampleContactInvalidFloat32 struct {
	FirstName    string
	LastName     float32 //this is not an int!
	Working      bool
	Age          int
	Salary32     float32
	Salary64     float64
	VacationDays uint
	Title        string
	Email        string
	Birthdate    time.Time
	Description  string
}

//ExampleContactInvalidFloat64 is made for testing uint wrong field
type ExampleContactInvalidFloat64 struct {
	FirstName    string
	LastName     float64 //this is not an int!
	Working      bool
	Age          int
	Salary32     float32
	Salary64     float64
	VacationDays uint
	Title        string
	Email        string
	Birthdate    time.Time
	Description  string
}

//ExampleContactInvalidTime is made for testing an invalid date format
type ExampleContactInvalidTime struct {
	FirstName    string    `csv:"0"`
	LastName     string    `csv:"1"`
	Working      bool      `csv:"2"`
	Age          int       `csv:"3"`
	Salary32     float32   `csv:"4"`
	Salary64     float64   `csv:"5"`
	VacationDays uint      `csv:"6"`
	Title        string    `csv:"7"`
	Email        string    `csv:"8"`
	Birthdate    time.Time `csv:"9" csvDate:"invalidDateFormat"`
	Description  string    `csv:"10"`
}

//ExampleContactWithCsvColumnFieldTooHigh has a csv column tag that exceed the number of csv columns in a row
type ExampleContactWithCsvColumnFieldTooHigh struct {
	LastName string `csv:"1000"`
}

//ExampleContactWithCsvTagLessThanZero has a csv column tag that is negative
type ExampleContactWithCsvTagLessThanZero struct {
	LastName string `csv:"-2"`
}

//ExampleContactWithCsvTagNotAnInteger has a csv column tag that is not an integer
type ExampleContactWithCsvTagNotAnInteger struct {
	LastName string `csv:"notAnInteger"`
}

func (c1 ExampleContact1) GetFirstName() string    { return c1.FirstName }
func (c1 ExampleContact1) GetLastName() string     { return c1.LastName }
func (c1 ExampleContact1) GetWorking() bool        { return c1.Working }
func (c1 ExampleContact1) GetAge() int             { return c1.Age }
func (c1 ExampleContact1) GetSalary32() float32    { return c1.Salary32 }
func (c1 ExampleContact1) GetSalary64() float64    { return c1.Salary64 }
func (c1 ExampleContact1) GetVacationDays() uint   { return c1.VacationDays }
func (c1 ExampleContact1) GetTitle() string        { return c1.Title }
func (c1 ExampleContact1) GetEmail() string        { return c1.Email }
func (c1 ExampleContact1) GetBirthdate() time.Time { return c1.Birthdate }
func (c1 ExampleContact1) GetDescription() string  { return c1.Description }

func (c2 ExampleContact2) GetFirstName() string    { return c2.FirstName }
func (c2 ExampleContact2) GetLastName() string     { return c2.LastName }
func (c2 ExampleContact2) GetWorking() bool        { return c2.Working }
func (c2 ExampleContact2) GetAge() int             { return c2.Age }
func (c2 ExampleContact2) GetSalary32() float32    { return c2.Salary32 }
func (c2 ExampleContact2) GetSalary64() float64    { return c2.Salary64 }
func (c2 ExampleContact2) GetVacationDays() uint   { return c2.VacationDays }
func (c2 ExampleContact2) GetTitle() string        { return c2.Title }
func (c2 ExampleContact2) GetEmail() string        { return c2.Email }
func (c2 ExampleContact2) GetBirthdate() time.Time { return c2.Birthdate }
func (c2 ExampleContact2) GetDescription() string  { return c2.Description }

func (c3 ExampleContact3) GetFirstName() string    { return c3.FirstName }
func (c3 ExampleContact3) GetLastName() string     { return c3.LastName }
func (c3 ExampleContact3) GetWorking() bool        { return c3.Working }
func (c3 ExampleContact3) GetAge() int             { return c3.Age }
func (c3 ExampleContact3) GetSalary32() float32    { return c3.Salary32 }
func (c3 ExampleContact3) GetSalary64() float64    { return c3.Salary64 }
func (c3 ExampleContact3) GetVacationDays() uint   { return c3.VacationDays }
func (c3 ExampleContact3) GetTitle() string        { return c3.Title }
func (c3 ExampleContact3) GetEmail() string        { return c3.Email }
func (c3 ExampleContact3) GetBirthdate() time.Time { return c3.Birthdate }
func (c3 ExampleContact3) GetDescription() string  { return c3.Description }

func (c4 ExampleContact4) GetFirstName() string    { return c4.FirstName }
func (c4 ExampleContact4) GetLastName() string     { return c4.LastName }
func (c4 ExampleContact4) GetWorking() bool        { return c4.Working }
func (c4 ExampleContact4) GetAge() int             { return c4.Age }
func (c4 ExampleContact4) GetSalary32() float32    { return c4.Salary32 }
func (c4 ExampleContact4) GetSalary64() float64    { return c4.Salary64 }
func (c4 ExampleContact4) GetVacationDays() uint   { return c4.VacationDays }
func (c4 ExampleContact4) GetTitle() string        { return c4.Title }
func (c4 ExampleContact4) GetEmail() string        { return c4.Email }
func (c4 ExampleContact4) GetBirthdate() time.Time { return c4.Birthdate }
func (c4 ExampleContact4) GetDescription() string  { return c4.Description }
