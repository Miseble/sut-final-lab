package entity

import(

	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func EmployeesNegativesTestValid(t *testing.T){
	g := NewGomegaWithT(t)

	t.Run("pass", func (t *testing.T) {
		e := Employees{
			Name: "qwer",
			Salary: 150000,
			EmployeeCode: "",
		}
	})
	
	ok, err := govalidator.ValidateStruct(e)
	Expect(ok).To(Equal(""))
	Expect(err).To(Equal("EmployeeCode must be 2 uppercase English letters (A-Z) followed by ‘-’ and 4 digits (0-9)"))

	
}

