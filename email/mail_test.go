package mail

import "testing"

func TestSendMail(t *testing.T) {
	t.Run("TestSendMail", func(t *testing.T) {
		SendMail()
	})
}

/* func TestSendMailOldVer(t *testing.T) {
	t.Run("TestSendMail2", func(t *testing.T) {
		SendMailOldVer()
	})
} */
