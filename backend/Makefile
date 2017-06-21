.PHONY: glide deps initdb inittest

glide:
	mkdir ${GOPATH}/bin
	curl https://glide.sh/get | sh

deps:
	glide up
	glide install

initdb:
	@echo "initializing database..."
	@mysql -u root -e "DROP DATABASE IF EXISTS faultinfo_db;"
	@mysql -u root -e "CREATE DATABASE faultinfo_db;"
	@mysql -u root faultinfo_db < sql/faultinfo.sql

inittest:
	@mysql -u root faultinfo_db < sql/test_data.sql
