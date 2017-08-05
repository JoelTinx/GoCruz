package models

import (
	"log"
)

func WhenError(err error) {
	if err != nil {
		log.Println(err.Error())
		_ := SendEmailError(err.Error())
		return
	}
}

func RegErrorInDB(err error) {

}

func ErrorReg()  {
	
}