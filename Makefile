# postgres
make_postgres:
	docker run --name my-pg \
	-e POSTGRES_USER=ROOT \
	-e POSTGRES_PASSWORD=whwhwhwhwhwh1231 \
	-e POSTGRES_DB=MAN \
	-p 5432:5432 \
	-d postgres

# redis
make_redis:
	docker run --name=my-rd \
	-p 6379:6379 \
	-d redis \
	-requirepass whwhwhwhwhwh1321



make_db: make_postgres make_redis

start_pg:
	docker start my-pg

start_rd:
	docker start my-rd

start_db: start_pg start_rd