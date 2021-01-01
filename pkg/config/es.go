package config

var ElasticSearch es

type es struct {
	User     string
	Password string
	Host     string
	Port     string
}
