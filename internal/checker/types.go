package checker

type CheckOptions struct {
	CheckType string
	FileNames []string
}

type CheckResult struct {
	CheckNotRequired bool // true if f.e.".ssh" not exists etc
	IsModyfied       bool
	Msg              string // TEMP
	Err              error
}

type Checker interface {
	Check() CheckResult
}
