package entity

import(

	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func EmployeesTestValid(t *testing.T){
	g := NewGomegaWithT(t)

	t.Run("pass", func (t *testing.T) {
		e := Employees{
			Name: "qwer",
			Salary: 150000,
			EmployeeCode: "qw-1412",
		}
	})
	
	ok, err := govalidator.ValidateStruct(e)
	Expect(ok).To(Equal(""))
	Expect(err).To(Equal(""))

	
}

