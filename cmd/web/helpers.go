package main

func (app *config) sendEmail(msg Message) {
	app.wait.Add(1)
	app.mailer.MailerChan <- msg
}
