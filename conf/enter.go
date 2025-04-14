package conf

type Config struct {
	System System `yaml:"system"`
	Log    Log    `yaml:"log"`
	DB     []DB   `yaml:"db"` //数据库连接列表
	Jwt    Jwt    `yaml:"jwt"`
	Redis  Redis  `yaml:"redis"`
	Site   Site   `yaml:"site"`
	Email  Email  `yaml:"email"`
	QQ     QQ     `yaml:"qq"`
	QiNiu  QiNiu  `yaml:"qiniu"`
	Ai     Ai     `yaml:"ai"`
	Upload Upload `yaml:"upload"`
	ES     ES     `yaml:"es"`
	River  River  `yaml:"river"`
}
