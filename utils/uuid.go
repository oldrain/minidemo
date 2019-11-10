package utils

import (
	"github.com/google/uuid"
	"strings"
)

func UUID() string {
	uu, _ := uuid.NewUUID()
	return uu.String()
}

func UUIDPure() string {
	uu, _ := uuid.NewUUID()
	return strings.ReplaceAll(uu.String(), "-", "")
}
