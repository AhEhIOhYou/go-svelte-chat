package constants

const (
	//Database proccess
	DatabaseConnectionStart = "Connecting to database..."
	DatabaseConnectionError = "Database connection error: %v"
	DatabaseConnectionSuccess = "Connected to database"
	DatabaseError = "Database error: %v"

	//Error messages
	Failed = "request failed"

	//User errors
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

	//Message errors
	MessageCantBeEmpty = "message can't be empty"
	FromUserIDCantBeEmpty = "from user id can't be empty"
	ChatIDCantBeEmpty = "chat id can't be empty"
	MessageDoesNotExist = "message does not exist"

	//Contact errors
	UserIDCantBeEmpty = "user id can't be empty"
	ContactIDCantBeEmpty = "contact id can't be empty"
	ContactDoesNotExist = "contact does not exist"


	//Success messages
	Successful = "Request successful"
	UsernameIsAvailable = "Username is available"
	RegistrationSuccessful = "Registration successful"
	LoginSuccessful = "Login successful"
	YouAreLoggedIn = "You are logged in"
	ContactAdded = "Contact added"
)