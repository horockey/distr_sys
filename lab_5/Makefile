docker_build_all:
	docker build -t localhost:5000/frontend:latest ./frontend
	docker build -t localhost:5000/backend:latest ./backend

prepare_local_registry:
	docker run -d -p 5000:5000 --name registry registry:2
	curl http://localhost:5000/v2/_catalog

run:
	docker push localhost:5000/frontend:latest
	docker push localhost:5000/backend:latest

	kubectl apply -f clickhouse-deployment.yaml
	kubectl apply -f clickhouse-service.yaml

	kubectl apply -f backend-deployment.yaml
	kubectl apply -f backend-service.yaml

	kubectl apply -f frontend-deployment.yaml
	kubectl apply -f frontend-service.yaml
	
	@echo "All entities have been created."

check:
	@echo "TODO: check section"

remove:
	kubectl delete -f frontend-service.yaml
	kubectl delete -f frontend-deployment.yaml
	kubectl delete -f backend-service.yaml
	kubectl delete -f backend-deployment.yaml
	kubectl delete -f clickhouse-service.yaml
	kubectl delete -f clickhouse-deployment.yaml

	@echo "All entities have been removed."
