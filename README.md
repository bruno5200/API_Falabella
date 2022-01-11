# API-Falabella 🖥

## Tutorial de Kubernetes con Minikube

### ejecuta los siguientes comandos con minikube instalado en la terminal de tu preferencia:

kubectl apply -f deployment.yaml

### copia las ultimas 5 letras generadas en NAME al ejecutar el siguiente comando

kubectl get pods

### los reemplazas en el siguiente comando y ya tienes ejecutandose una imagen docker en tu computadora con la ruta de acceso en https://Localhost:3000

kubectl port-forward api-service-cbc647987-reemplaza 3000:3000


### te recomiendo usar las siguientes rutas:

con cualquier navegador:

GET http://localhost:3000/beers

GET http://localhost:3000/beers/1

usando postman:

POST http://localhost:3000/beers

Content-Type application/json

{
    "beer": {
        "id": 2,
        "name": "Paceña",
        "brewery": "CBN",
        "country": "Bolivia",
        "price": 15,
        "currency": "BOB"
    },
    "description": "Cerveza creada.",
    "success": true
}

Resultado:

{
    "beer": {
        "id": 2,
        "name": "Paceña",
        "brewery": "CBN",
        "country": "Bolivia",
        "price": 15,
        "currency": "BOB"
    },
    "description": "Cerveza creada.",
    "success": true
}


GET http://localhost:3000/beers/1/boxprice

Content-Type application/json

{
    "currency": "EUR",
    "quantity": 1
}

resultado:

{
    "currency": "EUR",
    "description": "Operación Exitosa.",
    "priceTotal:": 63,
    "success": true
}
