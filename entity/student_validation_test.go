package entity

import (
    "testing"

    "github.com/asaskevich/govalidator"
    . "github.com/onsi/gomega"
)


func TestValid(t *testing.T){
	g := NewGomegaWithT(t)

	t.Run(`student is valid`, func(t *testing.T){
		user := Student{
			FullName : "Bobby",
			Age : 22,
			Email : "Bobby@gmail.com",
			GPA : 4.00,
		}

		ok, err := govalidator.ValidateStruct(user)
		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})
}

func TestFullName(t *testing.T){
	g := NewGomegaWithT(t)

	t.Run(`FullName is valid`, func(t *testing.T){
		user := Student{
			FullName : "",
			Age : 22,
			Email : "Bobby@gmail.com",
			GPA : 4.00,
		}

		ok, err := govalidator.ValidateStruct(user)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("FullName is required"))
	})
}

func TestAge(t *testing.T){
	g := NewGomegaWithT(t)

	t.Run(`Age must be at least 18`, func(t *testing.T){
		user := Student{
			FullName : "Bobby",
			Age : 11,
			Email : "Bobby@gmail.com",
			GPA : 4.00,
		}

		ok, err := govalidator.ValidateStruct(user)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Age must be at least 18"))
	})
}

func TestEmail(t *testing.T){
	g := NewGomegaWithT(t)

	t.Run(`Email is required`, func(t *testing.T){
		user := Student{
			FullName : "Bobby",
			Age : 22,
			Email : "",
			GPA : 4.00,
		}

		ok, err := govalidator.ValidateStruct(user)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Email is required"))
	})
}

func TestGPA(t *testing.T){
	g := NewGomegaWithT(t)

	t.Run(`GPA must be between 0.00 and 4.00`, func(t *testing.T){
		user := Student{
			FullName : "Bobby",
			Age : 22,
			Email : "Bobby@gmail.com",
			GPA : 6.00,
		}

		ok, err := govalidator.ValidateStruct(user)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("GPA must be between 0.00 and 4.00"))
	})
}