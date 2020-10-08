package validations
type QueryValidator struct {
	names string `validate:"required"`
	email string `validate:"required"`
	content string `validate:"required"`
}
