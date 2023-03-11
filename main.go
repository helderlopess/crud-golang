package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/heldelopess/crud-golang/src/controller/routes"
	"github.com/joho/godotenv"
)

//carreegando a bib godotenv para ler as variaveis do arquivo .env
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println(os.Getenv ("TEST"))
	
//gin.default inicia o roteador com logger e middlewares de recovery
//ja o gin.new inicia o roteador novo sem nenhum middleware (panic recover)
	router :=gin.Default()

	routes.InitRoutes(&router.RouterGroup)
	//se estiver rodando outra aplicação nesta porta ele ira registrar log fatal
	if err := router.Run(":8080"); err != nil{
		log.Fatal(err)
	}
}



