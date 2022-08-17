package errors

import "fmt"

// required env variable not defined
type EnvVarNotDefined struct {
	Name string
}

func (e *EnvVarNotDefined) Error() string {
	return fmt.Sprintf("Environment variable [ %s ] not defined.", e.Name)
}

// Notification type not implemented
type NotificationNotImplemented struct {
	Name string
}

func (e *NotificationNotImplemented) Error() string {
	return fmt.Sprintf("Notificaton type [ %s ] not implemented.", e.Name)
}
