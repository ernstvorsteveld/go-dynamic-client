package oauth2client

type ClientType int

const (
	Confidential ClientType = iota
	Public
)

var clientTypes = map[string]ClientType{
	"confidential": Confidential,
	"public":       Public,
}

func (d ClientType) String() string {
	return [...]string{
		"confidential",
		"public"}[d]
}

func ClientTypeValueOf(val string) ClientType {
	clientType := clientTypes[val]
	return clientType
}
