package custom_error

const errorCodeBase = 0

const (
	// original codes
	UnknownError                 int32 = errorCodeBase + 1
	InvalidJSONString            int32 = errorCodeBase + 2
	InputValidationError         int32 = errorCodeBase + 3
	UnauthorizedError            int32 = errorCodeBase + 4
	InvalidCredentialsError      int32 = errorCodeBase + 5
	IntervalServerError          int32 = errorCodeBase + 6
	PermissionDenied             int32 = errorCodeBase + 7
	DocumentVersionItemsNotExist int32 = errorCodeBase + 10

	// proprietary codes
	DatabaseError      int32 = errorCodeBase + 8
	AlreadyExistsError int32 = errorCodeBase + 9
	TenantNotExist     int32 = errorCodeBase + 11
)
