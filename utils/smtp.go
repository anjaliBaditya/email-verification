package utils

import (
    "net/smtp"
)

func IsValidSMTP(email string) bool {
    parts := strings.Split(email, "@")
    domain := parts[1]
    addr := fmt.Sprintf("mail.%s", domain)
    conn, err := smtp.Dial(addr + ":25")
    if err!= nil {
        return false
    }
    defer conn.Close()
    return true
}
