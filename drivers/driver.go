package drivers

// Driver represents methods to work with data source
type Driver interface {
	Open(filename string) error
	HasTabs() bool
	UseMainTab()
	UseTab(name string) error
	UseTabID(id int) error
	GetHeaders() []string
}
