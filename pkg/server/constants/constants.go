package constants

const (
	//Database proccess
	DatabaseConnectionStart = "connecting to database..."
	DatabaseConnectionError = "database connection error: %v"
	DatabaseConnectionSuccess = "connected to database"
	DatabaseError = "database error: %v"

	//Error messages
	Failed = "request failed: %v"

	//User errors
	UserIDInvalid = "user id invalid"
	UsernameCantBeEmpty = "username can't be empty"
	UsernameInvalid = "username is invalid"
	UsernameIsAlreadyTaken = "username is already taken"
	UsernameLenError = "username should be between 3-20 characters"
	PasswordCantBeEmpty = "password can't be empty"
	PasswordIncorrect = "password is incorrect"
	UsernameAndPasswordCantBeEmpty = "username and password can't be empty"
	NotLoggedIn = "not logged in"
	AccountDoesNotExist = "account does not exist"

	//Chat erros
	ChatNameCantBeEmpty = "chat name can't be empty"
	ChatMustHaveAtLeastTwoParticipants = "chat must have at least 2 participants"
	ChatDoesNotExist = "chat does not exist"
	ChatAlreadyExists = "chat already exists"

	//Message errors
	MessageCantBeEmpty = "message can't be empty"
	FromUserIDCantBeEmpty = "from user id can't be empty"
	ChatIDCantBeEmpty = "chat id can't be empty"
	MessageDoesNotExist = "message does not exist"
	MessageCannotBeDeleted = "message cannot be deleted"
	MessageCannotBeEdited = "message cannot be edited"

	//Contact errors
	UserIDCantBeEmpty = "user id can't be empty"
	ContactIDCantBeEmpty = "contact id can't be empty"
	ContactUsernameCantBeEmpty = "contact username can't be empty"
	ContactDoesNotExist = "contact does not exist"
	ContactAlreadyExists = "contact already exists"

	//Security errors
	PasswordHashError = "error hashing password: %v"
	PasswordVerifyError = "error verifying password: %v"


	//Success messages
	Successful = "request successful"

	UsernameIsAvailable = "Username is available"
	RegistrationSuccessful = "Registration successful"
	LoginSuccessful = "Login successful"
	YouAreLoggedIn = "You are logged in"
	ContactAdded = "Contact added"

	//User Statuses
	UserOffline = 0
	UserOnline = 1
	
	//Chat Types
	ChatTypeDialog = 0
	ChatTypeGroup = 1
)