package domain

import (
	"time"

	"github.com/efrag/blog-posts/go_generate/utils"
)

type Person struct {
	name        string
	dateOfBirth time.Time
	phone       utils.Phone
}
