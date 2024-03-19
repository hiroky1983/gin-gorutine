package main

import (
	"fmt"
	"sync"
	"time"

	rand "github.com/Pallinder/go-randomdata"
	"github.com/gin-gonic/gin"
)
type User struct {
	ID uint64 
	Username string 
	Email string 
	Address string
	Phone string
}

type Teacher struct {
	ID uint64
	TeacherName string
	Commission  int
	GradeNumber uint64
	ClassNumber uint64
	Subject *int
}

type Class struct {
	User *User
	Teacher *Teacher
}

const (
	NoCommission = iota
	MainCommission
	SubCommission
)

const (
	MathSubject = iota
	ScienceSubject
	EnglishSubject
)

func main () {
	r := gin.Default()
	r.GET("/" , func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	r.GET("/user", handleUsers)
	r.GET("/teacher", handleTeachers)
	r.GET("/class", handleClass)
	r.GET("/classgo", handleClassGoRoutine)
	r.Run()
}

func handleUsers(c *gin.Context) {
	user := &User{}
	c.JSON(200, gin.H{
		"user": user.newUser(),
	})
}

func handleTeachers(c *gin.Context) {
	teacher := &Teacher{}
	c.JSON(200, gin.H{
		"teacher": teacher.newTeacher(),
	})
}

func handleClass(c *gin.Context) {
	s := time.Now()
	user := &User{}
	teacher := &Teacher{}
	class := &Class{
		User: user.newUser(),
		Teacher: teacher.newTeacher(),
	}
	fmt.Println("finish")
	fmt.Printf("process time as %d\n",time.Since(s).Milliseconds())
	c.JSON(200, gin.H{
		"class": class,
	})
}

func handleClassGoRoutine(c *gin.Context) {
	s := time.Now()
	user := &User{}
	teacher := &Teacher{}
	var wg sync.WaitGroup
	wg.Add(2)
	go func () {
		defer wg.Done()
		teacher = teacher.newTeacher()
	}()
	go func () {
		defer wg.Done()
		user = user.newUser()
	}()
  wg.Wait()
	fmt.Println("finish2")
	fmt.Printf("process time as %d\n",time.Since(s).Milliseconds())
	c.JSON(200, gin.H{
		"class": &Class{
			User: user,
			Teacher: teacher,
		},
	})
}

func (u *User) newUser() *User {
	return &User{
		ID: u.randUserID(),
		Username: u.randUserName(),
		Email: u.randEmail(),
		Address: u.randAddress(),
		Phone: u.randPhone(),
	}
}

func (u *User)randUserID() uint64 {
	return uint64(rand.Number(1, 10000))
}

func (u *User)randUserName() string {
	return rand.FullName(1)
}

func (u *User)randEmail() string {
	return rand.Email()
}

func (u *User)randAddress() string {
	return rand.Address()
}

func (u *User)randPhone() string {
	return rand.PhoneNumber()
}

func (t *Teacher) newTeacher() *Teacher {
	return &Teacher{
		ID: t.randTeacherID(),
		TeacherName: t.randTeacherName(),
		Commission: t.randCommission(),
		GradeNumber: t.randGradeNumber(),
		ClassNumber: t.randClassNumber(),
		Subject: t.randSubject(),
	}
}

func (t *Teacher)randTeacherID() uint64 {
	return uint64(rand.Number(1, 10000))
}

func (t *Teacher)randTeacherName() string {
	return rand.FullName(1)
}

func (t *Teacher)randCommission() int {
	return rand.Number(0, 2)
}

func (t *Teacher)randGradeNumber() uint64 {
	return uint64(rand.Number(1, 6))
}

func (t *Teacher)randClassNumber() uint64 {
	return uint64(rand.Number(1, 6))
}

func (t *Teacher)randSubject() *int {
	if t.Commission == NoCommission {
		return nil
	}
	subject := rand.Number(1, 2)
	return &subject
}
