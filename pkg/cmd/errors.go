package cmd

type UsageError struct {
	helpFunc func() error
	message  string
}

func NewUsageError(helpFunc func() error, message string) UsageError {
	return UsageError{
		helpFunc: helpFunc,
		message:  message,
	}
}

func (e UsageError) Error() string {
	return e.message
}

func (e UsageError) Help() error {
	return e.helpFunc()
}
