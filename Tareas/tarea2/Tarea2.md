### TAREA 2
#### Docker Compose y Docker Volumes

Se realiza una aplicacion multicontenedores compuesta exactamente por 3 de ellos los cuales estaran conformados por:

* API (base en Nodejs)
* CLIENTE (base en la libreria React)
* BASE DE DATOS ( mongoDB)

el funcionamiento en concreto forma parte de una aplicación web para ingreso de fotografías tomadas desde el navegador web en una base de datos no relacional, así mismo que se puedan visualizar las fotos ingresadas. Estas deben guardarse en una base de datos MongoDB y se pide guardar de forma persistente.

el manejo de los contenedores esta proporcionado mediante un archivo de docker-compose el cual lo unico que requiere para hacer funcionar el sistema es el siguiente comando:

```
 docker compose up -d

```

para poder visualizar e interactuar con la aplicacion, al momento de levantar los contenedores, se debe ingresar la siguente direccion con puerto en: 

```
 localhost:8000
```

### Anexos

* enlace del Video (DISCLAIMER: volumen Bajo):
[VideoTarea2](https://drive.google.com/file/d/1HJ4CJxYH8t_3Up6_IBh1A_JBWdnRFkI0/view?usp=sharing)