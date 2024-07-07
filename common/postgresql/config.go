package postgresql

type Config struct {
	Host                  string
	Port                  string
	User                  string
	Password              string
	Database              string
	MaxConnections        string
	MaxConnectionIdleTime string
}
