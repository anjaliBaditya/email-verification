package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "strings"

    "email-verifier/models"
    "email-verifier/utils"
)

func main() {
    http.HandleFunc("/verify", verifyEmail)
    http.ListenAndServe(":8080", nil)
}

func verifyEmail(w http.ResponseWriter, r *http.Request) {
    email := r.URL.Query().Get("email")
    if email == "" {
        http.Error(w, "Email address is required", http.StatusBadRequest)
        return
    }

    emailModel := models.Email{Address: email}

    // Syntax check
    if!isValidSyntax(email) {
        emailModel.IsValid = false
        emailModel.Error = "Invalid email syntax"
        json.NewEncoder(w).Encode(emailModel)
        return
    }

    // DNS check
    if!utils.IsValidDNS(strings.Split(email, "@")[1]) {
        emailModel.IsValid = false
        emailModel.Error = "Invalid DNS record"
        json.NewEncoder(w).Encode(emailModel)
        return
    }

    // SMTP check
    if!utils.IsValidSMTP(email) {
        emailModel.IsValid = false
        emailModel.Error = "Email address does not exist"
        json.NewEncoder(w).Encode(emailModel)
        return
    }

    emailModel.IsValid = true
    json.NewEncoder(w).Encode(emailModel)
}

func isValidSyntax(email string) bool {
    // Implement email syntax validation using a regex or a library like go-validate
    return true
}
