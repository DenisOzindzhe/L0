package config // пакет config

/*
	На девелоперской машине, где сразу запущено много проектов, хранить параметры
	в переменных окружениях не всегда удобно; логичнее будет разделить их между проектами с помощью env-файлов.
	Сделать это можно, например, с помощью godotenv — это портированная на Go Ruby-библиотека dotenv.
	Она позволяет устанавливать необходимые для приложения переменные окружения из файла .env. (источник habr.com)
*/

//ServerConfiguration структура которая содержит порт сервера в формате string
type ServerConfiguration struct {
	ServerPort string // Server port
}

//DatabaseConfiguration структура которая содержит поля для подключения к базе данных
type DatabaseConfiguration struct {
	PostgresUser     string // Postgres username
	PostgresPassword string // Postgres password
	DBHost           string // Database host
	DBHostPort       string // Database host port
	DBName           string // Database name
}

//NatsConfiguration структура которая содержит поля для NUTS-Straming Service
type NatsConfiguration struct {
	NatsClusterID  string // Nats cluster
	NatsClientID   string // Nats client
	NatsChannel    string // Channel
	NatsDurableID  string // Durable
	NatsQueueGroup string // Que
}

// RedisConfiguration структура которая содержит поля для Redis Server (планируется использовать для хранения данных)
type RedisConfiguration struct {
	Host     string
	HostPort string
}
