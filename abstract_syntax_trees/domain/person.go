package domain

import (
	"time"

	"github.com/efrag/blog-posts/abstract_syntax_trees/utils"
)

type Person struct {
	name        string
	dateOfBirth time.Time
	phone       utils.Phone
}
