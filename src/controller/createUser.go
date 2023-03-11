package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/heldelopess/crud-golang/src/configuration/rest_err"
)

func CreateUser(c *gin.Context) {
	// retornando o erro em JSON para a aplicação/solicit ante
	// passando o err para o rest  erro

	err := rest_err.NewBadRequestError("Voce chamou a rota de forma incorreta, corrija usando ")
	c.JSON(err.Code, err)
	// retorna biding de objeto
}
