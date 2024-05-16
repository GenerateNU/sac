package calendar

type CalendarClient struct {
	Outlook 	OutlookClient
	Google 		GoogleClient
}

type CalendarClientInterface interface {
	Authorize() string
	Token(grantType string, code *string) error
}
