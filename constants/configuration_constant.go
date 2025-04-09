// Package constants contains constants for the application
package constants

// Constants for Configuration Messages
const (
	MsgConfigLoadFail    = "Failed to load configuration"
	MsgConfigLoadSuccess = "Successfully loaded configuration"
	MsgServerStart       = "Server started successfully"
	MsgServerShutdown    = "Shutting down server"
	MsgServerShutdownErr = "Failed to Shutting down server"
	MsgServerError       = "Error starting server"
	MsgServerGraceful    = "Server graceful shutdown"
	MsgZerologInit       = "Zerolog initiated successfully"
	MsgDirrectoryError   = "Failed to create directory"
	MsgFileOpenError     = "Failed to open file"
	MsgFileCloseError    = "Failed to close file"
	MsgRandomNumberError = "Failed to generate secure random number"
	MsgCommandFail       = "Failed to execute terminal command"
)
