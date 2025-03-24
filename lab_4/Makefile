run:
	kubectl apply -f mysql-configmap.yaml
	kubectl apply -f mysql-secret.yaml

	kubectl apply -f mysql-pv.yaml
	kubectl apply -f mysql-pvc.yaml

	kubectl apply -f mysql-deployment.yaml
	kubectl apply -f mysql-service.yaml

	kubectl apply -f wordpress-deployment.yaml
	kubectl apply -f wordpress-service.yaml

	@echo "All entities have been created."

check:
	kubectl get pods
	kubectl get configmap mysql-config -o yaml
	kubectl get secret mysql-secret -o yaml
	kubectl get pv
	kubectl get pvc
	kubectl get services
	kubectl logs $$(kubectl get pods -l app=mysql -o jsonpath='{.items[0].metadata.name}')
	kubectl logs $$(kubectl get pods -l app=wordpress -o jsonpath='{.items[0].metadata.name}')
	
	@echo "WordPress: http://$$(kubectl get nodes -o jsonpath='{.items[0].status.addresses[?(@.type=="InternalIP")].address}'):30000"

remove:
	kubectl delete -f wordpress-service.yaml
	kubectl delete -f wordpress-deployment.yaml
	kubectl delete -f mysql-service.yaml
	kubectl delete -f mysql-deployment.yaml
	kubectl delete -f mysql-pvc.yaml
	kubectl delete -f mysql-pv.yaml
	kubectl delete -f mysql-secret.yaml
	kubectl delete -f mysql-configmap.yaml
	
	@echo "All entities have been removed."