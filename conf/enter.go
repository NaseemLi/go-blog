package conf

type Config struct {
	System System `yaml:"system"`
	Log    Log    `yaml:"log"`
	DB     DB     `yaml:"db"`  //读
	DB1    DB     `yaml:"db1"` //写
	Jwt    Jwt    `yaml:"jwt"`
}
