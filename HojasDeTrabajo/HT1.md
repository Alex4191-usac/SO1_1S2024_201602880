### HOJA DE TRABAJO 1
#### Módulos de Kernel y Wails

desea la inserción de un módulo del Kernel creado por usted mismo en el
lenguaje de programación C y las librerías de Linux para kernel, el cual debe
obtener información importante de la memoria RAM como el porcentaje usado y
libre; esta información debe ser extraída por un programa en lenguaje Golang
usando el framework Wails, el cual se conectará con una aplicación de React y
mostrará dinámicamente como cambia la información de la memoria RAM

### Como utilizarlo:
* Generacion de modulo RAM

- generar y montar el modulo
```
make all

sudo insmod modulo_ram.ko

```

* utilizar la aplicacion en modo desarrollo

```
 wails dev
```

* generar el ejecutable para produccion

```
wails build
```
pd: el ejecutable se genera regularmente en la carpeta dist / bin * ver mensaje de consola

### ANEXOS

* Video de ejecucion: [HT1_VIDEO](https://drive.google.com/file/d/1xK21aA43nP4U11KD3_mJcAPKopZ2HYIo/view?usp=drive_link)