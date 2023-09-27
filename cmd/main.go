package main

// "content-management-microservice/utils"
import (
	"context"
	"fmt"
	"log"
	"os"

	//	"github.com/joho/godotenv"

	"github.com/Spacio-app/content-management-microservice/routes"
	"github.com/Spacio-app/content-management-microservice/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	// Inicializar la base de datos antes de configurar el servidor
	utils.InitDatabase()
	// if err := godotenv.Load(); err != nil {
	//     fmt.Println("Error cargando el archivo .env")
	//     os.Exit(1)
	// }

	// // Acceder a las variables de entorno
	// dbHost := os.Getenv("DB_HOST")
	// dbPort := os.Getenv("DB_PORT")
	// dbUser := os.Getenv("DB_USER")
	// dbPassword := os.Getenv("DB_PASSWORD")
	// dbName := os.Getenv("DB_NAME")
	// Crear instancia de Fiber
	app := fiber.New()

	// Configurar middleware CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, PUT, DELETE",
	}))
	// app.Use(middleware.SessionValidationMiddleware())
	// Configurar rutas
	// Configurar el archivo de registro
	logFile, err := os.Create("app.log")
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	// Configurar el registro para que escriba en el archivo
	log.SetOutput(logFile)

	routes.SetupRoutes(app)

	// Obtener una referencia a la colección
	collection := utils.GetCollection("Content") // <colección>

	// Consulta
	// filter := bson.M{"campo": "valor"}
	filter := bson.M{} // <filtro>
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())

	// Iterar a través de los resultados
	for cursor.Next(context.Background()) {
		var result bson.M
		if err := cursor.Decode(&result); err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	// Iniciar el servidor
	port := ":3001"
	err = app.Listen(port)
	if err != nil {
		fmt.Printf("Error al iniciar el servidor: %s\n", err)
		return
	} else {
		fmt.Printf("Servidor en ejecución en el puerto %s\n", port)
	}
}
