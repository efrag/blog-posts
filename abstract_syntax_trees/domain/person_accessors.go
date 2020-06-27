// DO NOT EDIT: file has been automatically generated
package domain

import "time"
import "github.com/efrag/blog-posts/abstract_syntax_trees/utils"

func (t *Person) GetName() string {
	return t.name
}

func (t *Person) SetName(f string) {
	t.name = f
}

func (t *Person) GetDateOfBirth() time.Time {
	return t.dateOfBirth
}

func (t *Person) SetDateOfBirth(f time.Time) {
	t.dateOfBirth = f
}

func (t *Person) GetPhone() utils.Phone {
	return t.phone
}

func (t *Person) SetPhone(f utils.Phone) {
	t.phone = f
}
