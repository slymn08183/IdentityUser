package constant

//<editor-fold desc="Mongo Error Codes">
const (
	MongoDuplicateKeyErrCode    = 11000
	MongoDuplicateKeyErrCodeStr = "11000"
)

//</editor-fold>

//<editor-fold desc="Mongo Error Keys">
const (
	MongoDuplicateMailErr  = "email_1 dup key"
	MongoMailValidationErr = "Error:Field validation for 'Email'"
)

//</editor-fold>

//<editor-fold desc="Error Strings">
const (
	DuplicateMailErr = "This user is already created!"

	UserDidNotCreatedErr = "An error accrued while creating the user!"

	InternalServerError = "Internal Server Error!"

	EmailValidationError = "Not a valid email!"

	PasswordOrMailValidationError = "Email or password is incorrect"

	InvalidToken = "The token is invalid!"
)

//</editor-fold>
