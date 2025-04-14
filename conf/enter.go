package conf

type Config struct {
	System System `yaml:"system"`
	Log    Log    `yaml:"log"`
	DB     DB     `yaml:"db"`  //读
	DB1    DB     `yaml:"db1"` //写
	Jwt    Jwt    `yaml:"jwt"`
	Redis  Redis  `yaml:"redis"`
	Site   Site   `yaml:"site"`
	Email  Email  `yaml:"email"`
	QQ     QQ     `yaml:"qq"`
	QiNiu  QiNiu  `yaml:"qiniu"`
	Ai     Ai     `yaml:"ai"`
	Upload Upload `yaml:"upload"`
	ES     ES     `yaml:"es"`
}
