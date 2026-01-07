package entity

import(

	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func EmployeesNegativeTestValid(t *testing.T){
	g := NewGomegaWithT(t)

	t.Run("Salary no pass", func (t *testing.T) {
		e := Employees{
			Name: "qwer",
			Salary: ,
			EmployeeCode: "qw-1412",
		}
	})
	
	ok, err := govalidator.ValidateStruct(e)
	Expect(ok).To(Equal(""))
	Expect(err).To(Equal("Salary must be between 15000 and 200000"))

	
}

