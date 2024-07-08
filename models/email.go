package models

type Email struct {
    Address string `json:"address"`
    IsValid bool   `json:"is_valid"`
    Error   string `json:"error"`
}
