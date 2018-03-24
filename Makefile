#!make
include .env

install:
	@grep "DJANGO_CHANGE_ME" .env && \
		sed -i "s/DJANGO_CHANGE_ME/\"$(shell python secret_generator.py)\"/" .env || \
		echo "Skipping DJANGO_PASSWORD generation"
	@grep "PG_CHANGE_ME" .env && \
		sed -i "s/PG_CHANGE_ME/\"$(shell python secret_generator.py)\"/" .env || \
		echo "Skipping POSTGRES_PASSWORD generation"
	@if [ ! -f ./ssl/nginx.crt ]; then\
		echo "ERROR: Could not find symlink for ./ssl/nginx.crt";\
		exit 1;\
	fi
	@if [ ! -f ./ssl/nginx.key ]; then\
		echo "ERROR: Could not find symlink for ./ssl/nginx.key";\
		exit 1;\
	fi
	sudo mkdir -p $(PGDATA)
	sudo mkdir -p $(CT_STATIC)
	docker-compose build
	docker-compose up --no-start
	docker-compose start
	@echo "Waiting for DB..."
	@docker-compose run --rm app sh -c "while ! nc -z db 5432 </dev/null; do sleep 5; done"
	docker-compose run --rm app sh -c "python ./manage.py makemigrations cloudtablesui && \
									python ./manage.py migrate"
	docker-compose run --rm app python ./manage.py collectstatic --no-input
	docker-compose run --rm app curl -k -I https://nginx/sync
	
update:
	docker-compose down
	docker-compose build
	docker-compose up --no-start
	docker-compose start

self_signed:
	wget https://github.com/corysm1th/cloudtables-python/releases/download/v1.0/cfssl
	@cfssl selfsign cloudtables.selfsigned.smdh ./ssl/www.json \
	| cfssljson -bare ./ssl/cloudtables.selfsigned.smdh
	@if [ ! -f ./ssl/nginx.crt ]; then\
		ln -s cloudtables.selfsigned.smdh.pem ./ssl/nginx.crt;\
	fi
	@if [ ! -f ./ssl/nginx.key ]; then\
		ln -s cloudtables.selfsigned.smdh-key.pem ./ssl/nginx.key;\
	fi

reset_secrets:
	@cp .env.default .env

clean_certs:
	rm -f ./ssl/cloudtables.selfsigned.smdh*
	rm -f ./ssl/nginx.crt
	rm -f ./ssl/nginx.key

clean:
	docker-compose down
	sudo rm -Rf $(PGDATA)
	sudo rm -Rf $(CT_STATIC)

test:
	@echo "No tests yet."
