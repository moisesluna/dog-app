# GOLANG REST API SIMPLE

# Package
* gorilla/mux: 'go get github.com/gorilla/mux'
* SQLite3      'go get github.com/mattn/go-sqlite3'


# Pasos de Ejecucion
*docker build -t go-dog-app .
este ejecutara el archivo Docker para crear el contenedor de la aplicacion
*docker images 
para verificar la correcta creacion de este contenedor
*docker run -d -p 8081 go-dog-app
con esto el servidor estará arriba escuchando desde el puerto 8081
*se puede verificar un listado de los contenedores activos con docker ps.

#Pasos para ejecutar pruebas
La aplicacion es un CRUD muy sencillo en el cual podemos administrar los datos de diferentes perros.
Al iniciar la aplicacion creara si no existe la tabla Dog en la base de datos moises.db
luego insertara un registro de prueba, alli ya podemos probar el CRUD

#Consultar todos los perros
GET http://localhost:8081/dogs

#Consultar un perro
GET http://localhost:8081/dogs/{id}

#Eliminar un perro
DELETE http://localhost:8081/dogs/{id}

#Crear un perro
POST http://localhost:8081/dogs
Ejemplo de Body:
    {
        "Name": "Desmond",
        "Breed": "Labrador"
    }

#Modificar datos de un perro
PUT http://localhost:8081/dogs/{id}
Ejemplo de Body:
    {
        "Name": "Desmond",
        "Breed": "Labrador"
    }