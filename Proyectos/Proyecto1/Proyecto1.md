# Proyecto1

En este proyecto, se tiene como objetivo principal implementar un sistema de monitoreo de
recursos del sistema y gestión de procesos, empleando varias tecnologías y lenguajes de
programación. El sistema resultante permitirá obtener información clave sobre el rendimiento del
computador, procesos en ejecución y su administración a través de una interfaz amigable

## Componentes:

- Se realiza una aplicacion multicontenedores compuesta exactamente por 3 de ellos los cuales estaran conformados por:

    API 
    CLIENTE 
    BASE DE DATOS 

* API: 

Se encargará de realizar los llamados a los módulos ubicados en la carpeta proc y almacenar los datos en la BD de MYSQL. Realiza el envío de datos hacia el monitoreo histórico y monitoreo a tiempo real, el envío de datos necesarios para el árbol de procesos en tiempo real y realizar las operaciones del Simulador de Procesos y Graficarlo de manera inmediata.

el framework utilizado para el manejo de la comunicacion http fue:  [Gin](https://gin-gonic.com/)


* CLIENTE:

para la interaccion con el usuario final se utilizo la libreria React
con complementos para poder mostrar graficamente la manipulacion de los datos que se describieron en la seccion api, los detalles del como se utiliza se explican en secciones posteriores.


* BD:

Es el encargado de guardar los datos de la informacion de memoria ram y cpu. esta fue implementada mediante un contenedor de Docker
ademas la base de datos cuenta con persistencia por lo que se implemento un volumen para llevar acabo el requerimiento y que al momento que el contenedor se reinicia o equipo los datos persisten.

* Modulos:

para la obtencion de los datos requeridos en esta aplicacion se necesitan la utilizacion de dos modulos los cuales se encuentran disponibles en el Folder MODULES, cada uno de los modulos estan desarrollados en el lenguaje de C.

- Modulo RAM: este modulo tiene como finalidad retornar la informacion correspondiente con la capacidad con la que trabaja la memoria RAM, los datos que retorna este modulo son: Memoria total, Memoria Disponible, Memoria usada y el porcentaje de utilizacion.

- Modulo Cpu: el modulo Cpu tiene cierta similitud con respecto al modulo de Ram ya que puede retornar, los mismos datos con la adicion que tambien retorna un arreglo con los procesos que tiene actualmente asignados, cada uno de estos procesos: estan identificados por un PID, nombre y si tienen procesos "hijos" ligados a ellos.


## Instalacion:

1. Instalacion de Modulos

para hacer uso de esta aplicacion se requiere la instalacion de los modulos ya que estos proporcionaran parte los datos principales para manipular e visualar la data, para la instalacion de los mismos debemos ir a la carpeta MODULES, alli se encontraran dos carpetas CPU  Y RAM, cada uno posee su archivo.c y su archivo Makefile,


para poder obtener nuestros modulos abrimos nuestra terminal y en la carpeta raiz ya sea de de CPU O ram ejecutamos el comando 

```bash
Make all
```
esto nos generara varios archivos, en nuestro caso los que nos interesan son los de extension .ko

para la integracion de esos modulos ejecutamos la siguiente instruccion en la terminal

```bash
sudo insmod nombre_modulo.ko

#para cpu y ram seran con los nombres: modulo_cpu, modulo_ram
```

y con ello podemos verificar en nuestro SO, redirigiendonos a la carpeta /proc y visualizar los modulos anteriores.



2. Ejecucion de aplicacion

Para hacer uso de la aplicacion tenemos dos caminos ya sea obtener las imagenes de docker hub o utilizar este repositorio, en caso de ser la segunda opcion, seria el siguiente caso:

clonar el actual repositorio y ubicarnos en la carpeta raiz donde tenemos el archivo `docker-compose.yml`

este archivo ahorrara el tiempo de ejecutar cada uno de los archivos `Dockerfile` y generar uno por uno para cada uno de los servicios descritos anteriores.

para hacer uso de ello ejecutamos el siguiente comando:

```bash
Docker compose up

#puede agregar la etiqueta -d , para evitar visualizar los registros de inicializacion
```


y con ello los servicios serian iniciados exitosamente y lo que queda es utilizar el portal, para hacer uso de el solo se debe utilizar ya sea la ip de una maquina virtual o fisica 
seguida del puerto por defecto que en este caso es 80 en y colocar, esta informacion en el navegador web de preferencia.

ej: `xxxx.xxx.xxx.xxx:80` or `localhost:80`

3. Detener uso & Remover:

para parar completamente la aplicacion basta con solamente parar la ejecucion en la terminal donde se inicio, para ello podemos hacer uso de la combinacion `Ctrl + c` para parar los contenedores.

si deseamos removerlos por completo utilizamos el comando

```bash
 Docker compose down
```

nota: debemos estar localizados en la carpeta contenedora del archivo `docker-compose.yml`

## Uso:

Al iniciar nuestra aplicacion mostrara por defecto la siguiente pantalla tiene una barra de navegacion en la parte superior donde se encuentran las siguientes opciones:

- System Monitor
- History Ram & Cpu
- Tree Process
- State Managment

PD: al seleccionar cualquiera de las opciones sera redirigido a la pagina correspondiente a la opcion seleccionada.

#### System Monitor

En esta opcion podremos visualizar

- Gráfica en Tiempo Real del porcentaje de utilización de la memoria RAM.

- Gráfica en Tiempo Real del porcentaje de utilización del CPU


![pagina Principal](https://github.com/Alex4191-usac/SO1_1S2024_201602880/blob/proyecto1/Proyectos/Proyecto1/assets/pagina%20principal.png)


#### History Ram & Cpu

en base a los datos que se obtienen a lo largo del tiempo tanto de nuestro modulo ram y nuestro modulo cpu, en esta opcion tiene la caracteristica de visualizar como se comportan esos datos a lo largo del tiempo tanto de RAM como de CPU.

Ademas podemos ver la opcion de poder actualizar las graficas manualmente mediante el boton de actualizar que se muestra en la pagina.



![Historial](https://github.com/Alex4191-usac/SO1_1S2024_201602880/blob/proyecto1/Proyectos/Proyecto1/assets/Historial.png)


#### Tree Process

en esta pagina se mostra el arbol de procesos que es administrado por el modulo_cpu que se cargo anteriormente, estos procesos se podran visualizar al momento de seleccionar el SelectBox


![TreeProcess](https://github.com/Alex4191-usac/SO1_1S2024_201602880/blob/proyecto1/Proyectos/Proyecto1/assets/Tree%20Process.png)


ademas tiene la capacidad que cuando este es seleccionado, y al presionar el boton de View que se encuentra en la parte izquierda del Select box, mostrara de manera grafica como se compone desde la raiz hasta sus hijos (si es que los tuviera)

cada nodo mostrara el nombre del proceso y su PID correspondiente:

* caso en el que el proceso no tiene hijos

![TreeProcess](https://github.com/Alex4191-usac/SO1_1S2024_201602880/blob/proyecto1/Proyectos/Proyecto1/assets/tree%20process2.png)

* caso en el que el padre tiene varios hijos:

![TreeProcess2](https://github.com/Alex4191-usac/SO1_1S2024_201602880/blob/proyecto1/Proyectos/Proyecto1/assets/tree%20process3.png)

***NOTA: para poder visualizar diferentes arboles de procesos es recomendado utilizar el boton de limpiar que se encuentra en la parte derecha de color amarillo para generar un nuevo arbol***


#### State Managment

En esta pagina se podra  simular los estados de los procesos a través de botones, a manera de interaccion podra visualizar de manera grafica como es que se va comportando el estado de los procesos, las opciones se detallan a continuacion:

**Botones:**


-  New: Este botón creará un nuevo proceso, debe de retornar el PID de dicho proceso y
generará los siguientes estados: New, Ready y Running.
- Kill: Debe de Terminar el Proceso definitivamente. Estado Terminated.
- Stop: este botón debe de cambiar el estado de Running a Ready.
- Resumen: Este botón debe de cambiar el estado de Ready a Running.

![StateManagment](https://github.com/Alex4191-usac/SO1_1S2024_201602880/blob/proyecto1/Proyectos/Proyecto1/assets/stateManagment.png)



##### Opcion New & estados

Al ser seleccionado un boton como se indico anteriormente se mostrara de manera grafica la generacion de los estados, cabe decir que cada estado tendra un colo diferente para indicar cual es el estado actual y ademas se mostrara en pantalla en la parte superior izquierda el estado actual de manera textual.

para la opcion de new el color representante sera `verde`


-  New: El proceso se está creando
- Running: Se están ejecutando instrucciones.
- Waiting: El proceso está esperando que ocurra algún evento.

![StateManagment2](https://github.com/Alex4191-usac/SO1_1S2024_201602880/blob/proyecto1/Proyectos/Proyecto1/assets/state%20Managment2.png)


##### Opcion Stop & estados

al seleccionar stop el proceso pasara a un estado de Ready

- Ready: El proceso está esperando ser asignado a un procesador.

el color que identifica este estado es el color `amarillo`

[StateManagment3](https://github.com/Alex4191-usac/SO1_1S2024_201602880/blob/proyecto1/Proyectos/Proyecto1/assets/state%20Managment3.png)

##### Opcion Resume & estados

Al seleccionar Resume el proceso pasara a un estado de Running

para la opcion de new el color representante sera `verde`


##### Opcion Kill & estados
al seleccionar la opcion Kill el proceso pasa a al estado Terminated.

* Terminated: El proceso ha finalizado su ejecución.


para la opcion de new el color representante sera `rojo`


![StateManagment4](https://github.com/Alex4191-usac/SO1_1S2024_201602880/blob/proyecto1/Proyectos/Proyecto1/assets/state%20Managment4.png)





