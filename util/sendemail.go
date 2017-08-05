package util

type Email struct {

}

func SendMail(e Email) (error) {

}

func SendMasivemail(es []Email) {

}

func GetTemplateHTMLEmail()  {

}

func SendEmailError(message string) {
	e := Email{
		From:"uncorreo@gmail.com",
		To: "otrocorreo@hotmail.com",
		Subject: "Error en sistema",
		Content: "Error en: " + message,
	}
	e.From = "uncorreo@gmail.com"
	pass, _ := base64.StdEncoding.DecodeString("x123456v")
	e.To = "otrocorreo@hotmail.com"

	message := "From: " + e.From + "\n" +
		"To: " + e.To + "\n" +
		"Subject: " + e.Subject + "\n" +
		e.Content

	err := smtp.SendMail("smtp.gmail.com:487",
		smtp.PlainAuth("", e.From, string(pass), "smtp.gmail.com"),
		e.From, []string{e.To}, []byte(message))

	if err != nil {
		log.Fatal(err.Error())
	}

	return err
}
