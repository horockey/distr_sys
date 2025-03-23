MYSQL_CONFIGMAP_YAML=mysql-configmap.yaml
MYSQL_SECRET_YAML=mysql-secret.yaml
WORDPRESS_DEPLOYMENT_YAML=wordpress-deployment.yaml
MYSQL_PV_YAML=mysql-pv.yaml
MYSQL_PVC_YAML=mysql-pvc.yaml
MYSQL_DEPLOYMENT_YAML=mysql-deployment.yaml
MYSQL_SERVICE_YAML=mysql-service.yaml
WORDPRESS_SERVICE_YAML=wordpress-service.yaml

run:
	kubectl apply -f $(MYSQL_CONFIGMAP_YAML)
	kubectl apply -f $(MYSQL_SECRET_YAML)

	kubectl apply -f $(MYSQL_PV_YAML)
	kubectl apply -f $(MYSQL_PVC_YAML)

	kubectl apply -f $(MYSQL_DEPLOYMENT_YAML)
	kubectl apply -f $(MYSQL_SERVICE_YAML)

	kubectl apply -f $(WORDPRESS_DEPLOYMENT_YAML)
	kubectl apply -f $(WORDPRESS_SERVICE_YAML)

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
	kubectl delete -f $(WORDPRESS_SERVICE_YAML)
	kubectl delete -f $(WORDPRESS_DEPLOYMENT_YAML)
	kubectl delete -f $(MYSQL_SERVICE_YAML)
	kubectl delete -f $(MYSQL_DEPLOYMENT_YAML)
	kubectl delete -f $(MYSQL_PVC_YAML)
	kubectl delete -f $(MYSQL_PV_YAML)
	kubectl delete -f $(MYSQL_SECRET_YAML)
	kubectl delete -f $(MYSQL_CONFIGMAP_YAML)
	
	@echo "All entities have been removed."