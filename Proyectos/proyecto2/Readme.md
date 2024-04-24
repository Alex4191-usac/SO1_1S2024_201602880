#### Instalacion

##### Locust
para utilizar el generador de tráfico, necesitamos:

1. Crear el entorno virtual (venv):

```bash
python3 -m venv venv
```

2. Activar el entorno virtual:

```bash
source venv/bin/activate
```

3. Instalar las dependencias (requirements.txt)

```bash
pip install -r requirements.txt 
```

4. Compilar y Ejecutar el archivo .py

```bash
locust -f lcst.py
```

##### Estructura de K8S
para poder levantar la estructura debe ser ingresado los .yaml encontrados en la carpeta k8s, se sugiere seguir el siguiente orden:

1. Ingreso de Namespace

```bash
kubectl apply -f namespace.yaml
```

2. levantar Kafka

```bash
kubectl create -f 'https://strimzi.io/install/latest?namespace=so1p2' -n so1p2
```

```bash
kubectl apply -f https://strimzi.io/examples/latest/kafka/kafka-persistent-single.yaml -n so1p2
```

crear el topic

```bash
kubectl apply -f topic.yaml
```


3. levantar las bd
carpeta databases como origen

```bash
kubectl apply -f mongo.yaml
```


4. levantar el los proveedores

```bash
kubectl apply -f grpc-depoy.yaml
```

5. levantar el consumer

```bash
kubectl apply -f consumer.yml
```

6. habilitar el hpa al consumer

```bash
kubectl apply -f hpa-consumer.yaml
```

7. levantar ingress

```bash
kubectl apply -f ingress.yaml
```

##### Cloud Run
para levantar nuestra api y cliente en CloudRun debemos dockerizar nuestras
aplicaciones y ademas asegurarnos que docker tenga permiso para manipular
el Artifact Registry


* en caso de no tener ningun artefacto debemos darle autenticacion a gcloud
como en mi caso:

```bash
$ gcloud auth configure-docker us-east1-docker.pkg.dev
```

luego debemos construir las imagenes para linux/arm

api:
```bash
docker buildx build -t us-east1-docker.pkg.dev/sopes2024/mongotools/api --platform linux/amd64 .
```
cliente:
```bash
docker buildx build -t us-east1-docker.pkg.dev/sopes2024/mongotools/web --platform linux/amd64 .
```

y subirlas al Artifact Registry:
```bash
docker push  us-east1-docker.pkg.dev/sopes2024/mongotools/api 
```

```bash
docker push  us-east1-docker.pkg.dev/sopes2024/mongotools/web 
```

configurar manualmente en CloudRun seleccionando servicios por medio
de Artifact Registry y asignarle el respecivo puerto

api: *3000*
web: *80*


