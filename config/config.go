package config

type ServerConfig struct {
	Name        string        `mapstructure:"name"`
	LogsAddress string        `mapstructure:"logsAddress"`
	MysqlInfo   MysqlConfig   `mapstructure:"mysql"`
	RedisInfo   RedisConfig   `mapstructure:"redis"`
	UserInfo    UserConfig    `mapstructure:"user"`
	WeightInfo  UserConfig    `mapstructure:"weight"`
	MongodbInfo MongodbConfig `mapstructure:"mongodb"`
	JwtInfo     JwtConfig     `mapstructure:"jwt"`
}

type UserConfig struct {
	Port int `mapstructure:"port"`
}

type WeightConfig struct {
	Port int `mapstructure:"port"`
}

type MysqlConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Name     string `mapstructure:"name"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbName"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Db       int    `mapstructure:"db"`
}

type MongodbConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Name     string `mapstructure:"name"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbName"`
}

type JwtConfig struct {
	Realm         string `mapstructure:"realm"`
	Key           string `mapstructure:"key" json:"key"`
	TokenLookup   string `mapstructure:"tokenLookup"`
	TokenHeadName string `mapstructure:"tokenHeadName"`
}
