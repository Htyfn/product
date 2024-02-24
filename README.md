# Сервис товаров для марктеплейса

для старта подключаем бд
psql -h localhost -U postgres -w -c "create database postgres;"

и накатываем миграции
Для этого необходимо установить migrate
go install github.com/golang-migrate/migrate

и запустить 

migrate -database DATABASE_URL -path db/migrations up

DATABASE_URL = postgres://user:password@host:port/dbname?sslmode=disable

или запустить команды из файла db/migrations/{version}_create_tables.up.sql