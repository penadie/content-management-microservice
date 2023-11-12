// services/content_service.go
package services

import (
	"errors"
	"fmt"
	"log"

	"github.com/Spacio-app/content-management-microservice/domain"
	"github.com/Spacio-app/content-management-microservice/domain/models"
	"github.com/Spacio-app/content-management-microservice/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// llamar repos, clientes, apis externas, para guardar, actualizar, eliminar.
func CreateCourse(content domain.CourseReq) error {
	fmt.Println("contentasds", content)
	log.Println("Guardando curso en la base de datos...")

	return repositories.CreateCourse(content)
}
func CreatePost(content domain.PostReq) error {
	return repositories.CreatePost(content)
}
func CreateFile(content domain.FileReq) error {
	return repositories.CreateFile(content)
}
func CreateTest(content domain.TestReq) error {
	return repositories.CreateTest(content)
}
func CreateFeed(content domain.FeedReq) error {
	return repositories.CreateFeed(content)
}
func SaveTestResult(content domain.TestResultReq) error {
	return repositories.SaveTestResult(content)
}
func CreateAnnouncement(announcement domain.FeedReq) error {
	err := repositories.CreateFeed(announcement)
	return err
}
func UpdatePostComments(id string, comment domain.FeedCommentsReq) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	log.Println("objectID", objectID)
	log.Println("updatedComments service", comment)
	if err != nil {
		return errors.New("ID inválido")
	}
	return repositories.UpdatePostComments(objectID, comment)

}
func GetAllFeeds() ([]models.Feed, error) {
	feeds, err := repositories.GetAllFeeds()
	if err != nil {
		return nil, errors.New("error al obtener el contenido")
	}
	return feeds, nil
}
func GetContentByID(id primitive.ObjectID) (models.GenericContent, error) {
	return repositories.GetContentByID(id)
}
func GetAllContent() ([]models.GenericContent, error) {
	content, err := repositories.GetAllContent()
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("error al obtener el contenido")
	}
	return content, nil
}
func GetContentFeedOrderByDate(skip int, limitInt int) ([]models.Feed, error) {
	//
	return repositories.GetContentFeedOrderByDate(skip, limitInt)
}
func GetPostsByAuthor(authorID string) ([]models.Feed, error) {
	// Llamar al repositorio para obtener los posts del feed por autor\

	posts, err := repositories.GetFeedByAuthor(authorID)
	if err != nil {
		return nil, errors.New("error al obtener el contenido")
	}
	return posts, nil

}
func GetContentByAuthor(author string) ([]models.GenericContent, error) {
	return repositories.GetContentByAuthor(author)
}
func GetAllCourses() ([]models.Courses, error) {
	return repositories.GetAllCourses()
}
func GetAllPosts() ([]models.Posts, error) {
	return repositories.GetAllPosts()
}
func GetAllFiles() ([]models.Files, error) {
	return repositories.GetAllFiles()
}
func GetAllTests() ([]models.Tests, error) {
	return repositories.GetAllTests()
}

// updates
func UpdateCourse(id string, content domain.CourseReq) error {
	// Convertir la cadena de texto a ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("ID inválido")
	}

	// Verificar si existe el contenido
	_, err = repositories.GetContentByID(objectID)
	if err != nil {
		return errors.New("el contenido no existe")
	}
	return repositories.UpdateCourse(objectID, content)
}

func UpdatePost(id string, content domain.PostReq) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("ID inválido")
	}
	_, err = repositories.GetContentByID(objectID)
	if err != nil {
		return errors.New("el contenido no existe")
	}
	return repositories.UpdatePost(objectID, content)
}
func UpdateFile(id string, content domain.FileReq) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("ID inválido")
	}
	_, err = repositories.GetContentByID(objectID)
	if err != nil {
		return errors.New("el contenido no existe")
	}
	return repositories.UpdateFile(objectID, content)
}
func UpdateTest(id string, content domain.TestReq) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("ID inválido")
	}
	_, err = repositories.GetContentByID(objectID)
	if err != nil {
		return errors.New("el contenido no existe")
	}
	return repositories.UpdateTest(objectID, content)
}

// DeleteContentByID elimina un registro de contenido por ID del repositorio
func DeleteContentByID(id primitive.ObjectID) error {
	//verificar si existe el contenido
	_, err := repositories.GetContentByID(id)
	if err != nil {
		return errors.New("el contenido no existe")
	}
	return repositories.DeleteContentByID(id)
}

func DeleteFeedComment(feedID primitive.ObjectID, commentID primitive.ObjectID) error {
	// Verificar si existe el contenido
	_, err := repositories.GetContentByID(feedID)
	if err != nil {
		return errors.New("el contenido no existe")
	}
	return repositories.DeleteFeedComment(feedID, commentID)
}

// Continuar con las funciones para actualizar y eliminar contenido...
