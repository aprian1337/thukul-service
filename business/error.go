package businesses

import "errors"

var (
	ErrInternalServer           = errors.New("something gone wrong, contact administrator")
	ErrNotFound                 = errors.New("data not found")
	ErrIDNotFound               = errors.New("id not found")
	ErrNewsIDResource           = errors.New("(NewsID) not found or empty")
	ErrNewsTitleResource        = errors.New("(NewsTitle) not found or empty")
	ErrCategoryNotFound         = errors.New("category not found")
	ErrDuplicateData            = errors.New("duplicate data")
	ErrInvalidId                = errors.New("invalid id, id not numeric")
	ErrUserIdNotFound           = errors.New("user id not found")
	ErrInvalidDate              = errors.New("invalid date, date must be formed : yyyy-mm-dd")
	ErrUsernamePasswordNotFound = errors.New("username or password empty")
	ErrInvalidAuthentication    = errors.New("authentication failed: invalid user credentials")
	ErrInvalidTokenCredential   = errors.New("token not found or expired")
	ErrBadRequest               = errors.New("bad requests")
	ErrNothingDestroy           = errors.New("no data found to delete")
	ErrTypeActivity             = errors.New("type must be only 'income' and 'expense'")
	ErrInsufficientPermission   = errors.New("Insufficient Permission")
)
