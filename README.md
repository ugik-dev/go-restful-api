# go-restful-api


depedency :
go get github.com/go-sql-driver/mysql 
go get github.com/julienschmidt/httprouter 
go get github.com/go-playground/validator 

# Library Depedency Injection
https://github.com/google/wire
go get github.com/google/wire

go install github.com/google/wire/cmd/wire@latest


# db Migration sql
https://github.com/golang-migrate/migrate
go install -tags "mysql" github.com/golang-migrate/migrate/v4/cmd/migrate@latest
go install -tags "postgres,mysql" github.com/golang-migrate/migrate //untuk multi

-> masuk terminal, key"migrate" harusnya di database driver di bagian bawah ada ket mysql

migrate create -ext sql -dir db/migrations nama_file_migration //jika mongo db -ext json (ext itu extension)

migrate create -ext sql -dir db/migrations create_table_skills

running syntax
migrate -database "mysql://root:@tcp(localhost:3306)/go_restful2" -path db/migrations up

running rolback syntax
migrate -database "mysql://root:@tcp(localhost:3306)/go_restful2" -path db/migrations down

pada up bisa menambahkan up 2 untuk maju 2 step dan down 2 untuk rollback 2step migrate

DirtyState kesalahan saat penulisan sql harus dipperbaiki manual pada phpmyadmin
cek versi sekarang / bisa cek langsung di db tbl schema_migration
migrate -database "mysql://root:@tcp(localhost:3306)/go_restful2" -path db/migrations version

lalu eksekusi
migrate -database "mysql://root:@tcp(localhost:3306)/go_restful2" -path db/migrations force <versisekarang> 

# redis (inmemory database)
doc https://github.com/redis/go-redis
go get github.com/redis/go-redis/v9

