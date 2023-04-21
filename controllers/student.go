package controllers

import (
	"echo-lab-go/models"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// CreateStudentScore ..
func CreateStudentScore(c echo.Context) error {
	studentModel := models.StudentScore{}
	if err := c.Bind(&studentModel); err != nil {
		log.Fatal(err)
	}

	fmt.Println("studentModel: ", studentModel)

	if err := models.Db.Create(&studentModel).Error; err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusOK, map[string]interface{}{"success": false, "message": "Create StudentScore error"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"success": true})
}

// GetStudentScore ..
func GetStudentScore(c echo.Context) error {
	name := c.Param("name")

	var students []models.StudentScore
	if err := models.Db.Where("name = ?", name).Find(&students).Error; err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusOK, map[string]interface{}{"success": false, "message": "Get Students by name error"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"success": true, "data": students})
}

// GetAllStudentScore ..
func GetAllStudentScore(c echo.Context) error {
	var studentModels []models.StudentScore

	if err := models.Db.Find(&studentModels).Error; err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusOK, map[string]interface{}{"success": false, "message": "Get All StudentScore error"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"success": true, "data": studentModels})
}

// UpdateStudentScore ..
func UpdateStudentScore(c echo.Context) error {
	studentModel := models.StudentScore{}
	if err := c.Bind(&studentModel); err != nil {
		log.Fatal(err)
	}

	fmt.Println("studentModel: ", studentModel)

	id := c.Param("id")

	if err := models.Db.Model(&studentModel).Where("id = ?", id).Updates(&studentModel).Error; err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusOK, map[string]interface{}{"success": false, "message": "Update StudentScore error"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"success": true, "message": "StudentScore updated successfully"})
}

// DeleteStudentScore ..
func DeleteStudentScore(c echo.Context) error {
	studentModel := models.StudentScore{}
	if err := c.Bind(&studentModel); err != nil {
		log.Fatal(err)
	}

	fmt.Println("studentModel: ", studentModel)

	id := c.Param("id")

	if err := models.Db.Model(&studentModel).Where("id = ?", id).Delete(&studentModel).Error; err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusOK, map[string]interface{}{"success": false, "message": "Delete StudentScore error"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"success": true})
}

// RecoverStudentScore ..
func RecoverStudentScore(c echo.Context) error {
	studentModel := models.StudentScore{}
	if err := c.Bind(&studentModel); err != nil {
		log.Fatal(err)
	}

	fmt.Println("studentModel: ", studentModel)

	id := c.Param("id")

	if err := models.Db.Unscoped().Model(&studentModel).Where("id = ?", id).Update("deleted_at", nil).Error; err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusOK, map[string]interface{}{"success": false, "message": "Recover StudentScore error"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"success": true})
}

// GetSubjectGrade ..
func GetSubjectGrade(c echo.Context) error {
	studentModels := []models.StudentScore{}
	if err := models.Db.Find(&studentModels).Error; err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusOK, map[string]interface{}{"success": false, "message": "Get Subject Grade error"})
	}

	groupedStudents := make(map[string][]models.StudentScore)
	for _, student := range studentModels {
		groupedStudents[student.Subject] = append(groupedStudents[student.Subject], student)
	}

	gradeMap := make(map[string]string)
	for subject, students := range groupedStudents {
		totalScore := 0
		for _, student := range students {
			totalScore += student.Score
		}
		averageScore := totalScore / len(students)

		var grade string
		switch {
		case averageScore >= 81:
			grade = "A"
		case averageScore >= 71:
			grade = "B"
		case averageScore >= 61:
			grade = "C"
		case averageScore >= 51:
			grade = "D"
		default:
			grade = "F"
		}

		gradeMap[subject] = grade
	}

	var response []map[string]string
	for subject, grade := range gradeMap {
		students := groupedStudents[subject]
		for _, student := range students {
			response = append(response, map[string]string{
				"name":    student.Name,
				"subject": subject,
				"grade":   grade,
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"success": true, "data": response})
}

// GetSubjectGradeByStudentName ..
func GetSubjectGradeByStudentName(c echo.Context) error {
	name := c.Param("name")

	var studentScores []models.StudentScore
	if err := models.Db.Where("name = ?", name).Find(&studentScores).Error; err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusOK, map[string]interface{}{"success": false, "message": "Get Subject Grade by Student Name error"})
	}

	subjectScores := make(map[string]int)
	for _, studentScore := range studentScores {
		subjectScores[studentScore.Subject] += studentScore.Score
	}
	for subject, score := range subjectScores {
		subjectScores[subject] = score / len(studentScores)
	}

	groupedStudents := make(map[string][]models.StudentScore)
	for _, student := range studentScores {
		groupedStudents[student.Subject] = append(groupedStudents[student.Subject], student)
	}

	gradeMap := make(map[string]string)
	for subject, students := range groupedStudents {
		totalScore := 0
		for _, student := range students {
			totalScore += student.Score
		}
		averageScore := totalScore / len(students)
		var grade string
		switch {
		case averageScore >= 81:
			grade = "A"
		case averageScore >= 71:
			grade = "B"
		case averageScore >= 61:
			grade = "C"
		case averageScore >= 51:
			grade = "D"
		default:
			grade = "F"
		}
		gradeMap[subject] = grade
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"success": true, "data": gradeMap})
}
