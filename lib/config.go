package lib

type Config struct {
	Server struct {
		Host string
		Port int
	}

	Database struct {
		MySQL struct {
			Dsn string
		}
	}

	Log struct {
		Level string
	}
}

func LoadConfig() Config {
	return Config{
		Server: struct {
			Host string
			Port int
		}{
			Host: "localhost",
			Port: 8080,
		},
		Database: struct{ MySQL struct{ Dsn string } }{
			MySQL: struct{ Dsn string }{
				Dsn: "root:root@tcp(localhost:3306)/db",
			},
		},
		Log: struct{ Level string }{
			Level: "info",
		},
	}
}
