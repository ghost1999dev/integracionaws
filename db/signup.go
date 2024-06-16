package db

import (
	"fmt"
	"github.com/integracionaws/models"
	"github.com/integracionaws/tools"
)

func SigunUp(sig models.SignUp) error {
	fmt.Println("Comienza Registro")
	err:= DbConnect()
	if err != nil{
		return err
	}
	defer Db.Close()
	sentencia := "INSERT INTO users (User_Email, User_UUID, User_DateAdd) VALUES ('"+sig.UserEmail+"','"+sig.UserUUID+"','"+tools.MySQLDate()+"')"
	fmt.Println(sentencia)
	_,err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("SignUp > Succesfull")
	return nil
}