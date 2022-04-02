package config

type (
	Configuration struct {
		RedisSettings
		DiscountDbSettings
	}
	RedisSettings struct {
		Address string
		Pass    string
		Db      int
	}
	DiscountDbSettings struct {
		Host   string
		DbName string
		Port   string
		User   string
		Pass   string
	}
)
