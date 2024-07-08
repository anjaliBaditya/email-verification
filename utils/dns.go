package utils

import (
    "net"
)

func IsValidDNS(domain string) bool {
    _, err := net.LookupNS(domain)
    return err == nil
}
