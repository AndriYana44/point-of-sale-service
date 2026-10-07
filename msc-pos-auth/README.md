# point-of-sale-service

run db migration => goose -dir migrations mysql "root:root@tcp(127.0.0.1:3306)/point_of_sale?parseTime=true" up
run db seeds => go run ./cmd/seed
