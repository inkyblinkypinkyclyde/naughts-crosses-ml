setup:
	docker run --name naughts-and-crosses-db -e POSTGRES_PASSWORD=postgres -p 5434:5432 -d postgres

create-db:
	docker exec -it naughts-and-crosses-db psql -U postgres -c "CREATE DATABASE naughts_and_crosses;"

drop-db:
	docker exec -it naughts-and-crosses-db psql -U postgres -c "DROP DATABASE naughts_and_crosses;"

teardown:
	docker stop naughts-and-crosses-db
	docker rm naughts-and-crosses-db

connect-db:
	psql postgresql://postgres:postgres@localhost:5434/naughts_and_crosses
