package main

func (app *Config) sendEmail(msg Message) {
	app.Wait.Add(1)
	app.Mailer.MailerChan <- msg
}
