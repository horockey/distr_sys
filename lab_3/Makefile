PYTHON_POD_YAML=python-time-pod.yaml
REDIS_DEPLOYMENT_YAML=redis-deployment.yaml
REDIS_SERVICE_YAML=redis-service.yaml

run:
	kubectl apply -f $(PYTHON_POD_YAML)
	kubectl apply -f $(REDIS_DEPLOYMENT_YAML)
	kubectl apply -f $(REDIS_SERVICE_YAML)
	@echo "All entities have been created."

check:
	kubectl logs python-time-pod | tail -n 1

	kubectl get deployments
	kubectl get pods -l app=redis

	kubectl get services redis-service

	@NODE_IP=$$(kubectl get nodes -o jsonpath='{.items[0].status.addresses[?(.type=="InternalIP")].address}'); \
	redis-cli -h $$NODE_IP -p 30000 ping

remove:
	kubectl delete -f $(PYTHON_POD_YAML)
	kubectl delete -f $(REDIS_DEPLOYMENT_YAML)
	kubectl delete -f $(REDIS_SERVICE_YAML)
	@echo "All entities have been removed."