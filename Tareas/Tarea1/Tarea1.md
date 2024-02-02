### TAREA 1
#### Docker y Contenedores


La siguiente actividad contiene la realizacion de una aplicación web utilizando el framework React que se conecta a una API
utilizando el lenguaje de programación Go, la cual consume un endpoint especifico el cual es "/data" de tipo GET, el cual retorna datos pertenecientes a mi registro estudiantil, los cuales son visibles a su vez mediante el cliente por medio de un boton. Ambos cliente y API-REST fueron dockerizadas para facilitar su despliegue.


#### Entregables graficos de la actividad

* Comandos utilizados para crear las imagenes y comandos para ejecutar los contenedores:




```
#Comando para crear la imagen del cliente

docker build -t client-app .


#Comando para crear la imagen del API-REST

docker build -t go-api-image

```

* Comando para mostrar las imagenes creadas

```
docker images
```

* sentencias para levantar los contenedores


```
docker run -p 8000:80 client-app
docker run -p 8080:8080 go-api-image

```

* sentencia para vizualizar los contenedores activos


```
docker ps
```


### Anexos

* Captura de construccion de Imagenes:
![Captura de construccion de Imagenes](https://github.com/Alex4191-usac/SO1_1S2024_201602880/blob/tarea1/Tareas/Tarea1/SCREENS/Screenshot%20from%202024-02-01%2019-12-20.png)

* Captura de imagenes construidas:

![Imagenes](https://github.com/Alex4191-usac/SO1_1S2024_201602880/blob/tarea1/Tareas/Tarea1/SCREENS/Screenshot%20from%202024-02-01%2019-13-59.png)


* Captura de construccion de Contenedores:

![Captura de construccion de Contenedores](https://github.com/Alex4191-usac/SO1_1S2024_201602880/blob/tarea1/Tareas/Tarea1/SCREENS/Screenshot%20from%202024-02-01%2019-13-46.png)



*Captura de contenedores activos:
![Activos Containers](https://github.com/Alex4191-usac/SO1_1S2024_201602880/blob/tarea1/Tareas/Tarea1/SCREENS/Screenshot%20from%202024-02-01%2019-26-54.png)


* Captura de interaccion cliente - servidor
![Inicial Carga Cliente](https://github.com/Alex4191-usac/SO1_1S2024_201602880/blob/tarea1/Tareas/Tarea1/SCREENS/Screenshot%20from%202024-02-01%2019-16-03.png)

![Fin](https://github.com/Alex4191-usac/SO1_1S2024_201602880/blob/tarea1/Tareas/Tarea1/SCREENS/Screenshot%20from%202024-02-01%2019-16-08.png)






* enlace a Grabacion (ADV: volumen bajo):
[GrabacionTarea1](https://drive.google.com/file/d/1v2xaHeyX8mBXnX9CxVXrYX0t3XrgOKph/view?usp=sharing)