# API_Falabella

install minikube

run kubectl apply -f deployment.yaml

copy last 5 digits generated in NAME using command "kubectl get pods" and replace in the next command:

expose port on machine kubectl port-forward api-service-cbc647987-"generated" 3000:3000
