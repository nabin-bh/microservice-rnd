package microservicernd

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
)

func main() {
	r := gin.Default()

	r.POST("/send-email", func(c *gin.Context) {
		// Parse request and extract email details
		var request struct {
			To      string `json:"to"`
			Subject string `json:"subject"`
			Body    string `json:"body"`
		}
		if err := c.BindJSON(&request); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request"})
			return
		}

		// Create and send the email
		m := gomail.NewMessage()
		m.SetHeader("From", "your-email@example.com")
		m.SetHeader("To", request.To)
		m.SetHeader("Subject", request.Subject)
		m.SetBody("text/plain", request.Body)

		d := gomail.NewDialer("smtp.example.com", 587, "your-smtp-username", "your-smtp-password")

		if err := d.DialAndSend(m); err != nil {
			c.JSON(500, gin.H{"error": "Failed to send email"})
			return
		}

		c.JSON(200, gin.H{"message": "Email sent successfully"})
	})

	r.Run(":8080")
}
