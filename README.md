# Customer-int

Este es el repositorio del proyecto Customer-int.



## Requisitos

- Go 1.20
- Git
- Postman o algun cliente donde pueda hacer peticiones REST

## Instalación y Ejecución

1. Clona el repositorio:


git clone https://github.com/stivenlaiton20/Customer-int.git

2. Correr el proyecto con 

go run main.go
y ya nuestra apliacion estara corriendo

## Descripción

Eh creado un aplicacion tipo rest-api 
donde interactuamos diferentes metodos tipo CRUD con https://jsonplaceholder.typicode.com/guide/
esta aplicacion hace peticiones a jsonplaceholder simunlando que es una base de datos
pero para interactuar con estos metodos primero toca logearnos con un usuario y contraseña 

## 1 paso para la apliacion

importar la configuracion de postman para hacer las diferentes peticiones que maneja la apliacion 
para ello deje un archivo llamado customer.postman_collection.json en el repositorio
teniendo instalado previamente postman selecionaos el boton de import para imporatar la coniguracion
![configuracion]([URL_de_la_imagen](https://raw.githubusercontent.com/stivenlaiton20/Customer-int/main/importar%20configuracion%20postman.png))

se nos abrira una ventan donde tenemos que seleccionar el archivo de configuracion llamado customer.postman_collection.json 
![configuracion2](https://raw.githubusercontent.com/stivenlaiton20/Customer-int/main/importar%20configuracion%20postman%202.png)

ya teniendo importado los difrentes metodos para comunicarnos con la aplicacion
procedemos a enviarle un Json con 
{
    "username": "usuario",
    "password": "contraseña"
} 
 para obetener un bearer token 
tal como se especifica en la imagen
![Login Con bearer Token](https://raw.githubusercontent.com/stivenlaiton20/Customer-int/main/login.PNG)

con este token ya podemos interactuar con los metodos para obtener todos los post, obtener un post, eliminar post, actualizar un post y crear un post
y para utilizarlo seleccionamos algun metodo y en la pestaña de authorizacion  seleccionamos Bearer tokent y en la casilla que nos aparece colocamos el token que nos dio previamente el login 
en caso que no lo colequemos o lo coloquemos mal nos arrojara un error 401 que es inautorizado
![utilizar token](https://raw.githubusercontent.com/stivenlaiton20/Customer-int/main/utilizar%20token%20bearer.PNG)
 y  ya con esto es replicarlo en cada uno de los metodos que aparecen disponibles en la carpeta de customers 

# crearPost
![crearPost](https://raw.githubusercontent.com/stivenlaiton20/Customer-int/main/crearPost.PNGn)
# eliminar Post 
![eliminar Post](https://raw.githubusercontent.com/stivenlaiton20/Customer-int/main/deletePost.PNG)
# actualizar  Post 
![actualizar post]([URL_de_la_imagen](https://raw.githubusercontent.com/stivenlaiton20/Customer-int/main/UpdatePost.PNG)https://raw.githubusercontent.com/stivenlaiton20/Customer-int/main/UpdatePost.PNG)







