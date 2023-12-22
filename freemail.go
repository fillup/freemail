package freemail

import (
	"errors"
	"fmt"
	"strings"
)

var ErrInvalidEmail = errors.New("invalid email address")

// IsFreeDomain determines if the given domain name is believed to be a free/public
// email service provider. Additional domains can be provided if they are not already
// in the default list, DomainList
func IsFreeDomain(domain string, additional ...string) bool {
	d := strings.ToLower(strings.TrimSpace(domain))
	if _, ok := DomainList[d]; ok {
		return true
	}
	for _, v := range additional {
		if v == d {
			return true
		}
	}
	return false
}

// IsFreeEmail determines if the given email address is believed to be from a
// free/public email service provider. Additional domains can be provided if
// they are not already in the default list, DomainList. An error, ErrInvalidEmail,
// is returned if the given email address does not look like a valid email address.
func IsFreeEmail(email string, additional ...string) (bool, error) {
	parts := strings.SplitN(email, "@", 2)
	if len(parts) != 2 {
		return false, fmt.Errorf("%w: %s", ErrInvalidEmail, email)
	}
	return IsFreeDomain(parts[1], additional...), nil
}
