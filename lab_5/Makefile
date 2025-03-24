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

	kubectl apply -f frontend-ingress.yaml

	kubectl apply -f https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml
	kubectl apply -f backend-hpa.yaml

	@echo "All entities have been created."

update_etc_hosts:
	@INGRESS_IP=$$(kubectl get ingress frontend-ingress -o jsonpath='{.status.loadBalancer.ingress[0].ip}'); \
	if [ -z "$$INGRESS_IP" ]; then \
		echo "Ingress IP not found. Please check if the Ingress is properly configured."; \
		exit 1;\
	fi; \
	echo "$$INGRESS_IP frontend.example.com" | sudo tee -a /etc/hosts

check:
	kubectl get pods -o wide
	kubectl get services
	kubectl get ingress
	kubectl get deployment metrics-server -n kube-system
	kubectl get hpa

remove:
	kubectl delete -f frontend-service.yaml
	kubectl delete -f frontend-deployment.yaml
	kubectl delete -f backend-service.yaml
	kubectl delete -f backend-deployment.yaml
	kubectl delete -f clickhouse-service.yaml
	kubectl delete -f clickhouse-deployment.yaml
	kubectl delete -f backend-hpa.yaml

	@echo "All entities have been removed."
