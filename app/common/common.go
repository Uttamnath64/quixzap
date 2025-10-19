package common

type ctxKey string

const (

	// ✅ Success Codes
	StatusSuccess     = 3300 // General success
	StatusUserCreated = 3301 // User created successfully
	StatusUserFetched = 3302 // User fetched successfully
	StatusUpdated     = 3303 // Record updated successfully
	StatusDeleted     = 3304 // Record deleted successfully

	// ❌ Error Codes
	StatusBadRequest      = 3400 // Bad request
	StatusUnauthorized    = 3401 // Unauthorized access
	StatusForbidden       = 3403 // Forbidden
	StatusNotFound        = 3404 // Resource not found
	StatusConflict        = 3409 // Conflict (duplicate entry, etc.)
	StatusServerError     = 3500 // Internal server error
	StatusDatabaseError   = 3501 // Database error
	StatusValidationError = 3502 // Validation error

	CtxUserID    ctxKey = "user_id"
	CtxUserType  ctxKey = "user_type"
	CtxSessionID ctxKey = "session_id"
)
