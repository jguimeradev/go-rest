# GO REST

RESTful API en Go para utilizar con otros proyectos enfocados a *DevOps* y *Cloud* 


## ¿QUÉ ES UNA REST API?

REST, **Representational State Transfer**, es un conjunto de reglas sobre como se comunican servidores y clientes sobre el protocolo *http*


## CARACTERÍSTICAS DE UNA REST API:

### 1. Sin estado (Stateless): 
Las peticiones del cliente son independientes. El servidor no almacena los datos de sesión, por lo que cada petición deben incluir todos los datos necesarios para procesarse:
```
PUT /users/42 HTTP/1.1
Host: localhost:6969
Content-Type: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiJ9...

{"user": "John", "email": "john@example.com"}

```

El servidor, una vez procesada la petición, no necesita recordar nada. La siguiente llevará su propio conjunto de datos. El cliente **sí** que guarda los datos que necesite para construir la siguiente petición.

Al mismo tiempo, se utilizan cachés que guardan copias de respuestas concretas (por ejemplo, el resultado de `GET /users/1`) para optimizar el rendimiento de la aplicación sin comprometer el principio de *statelessness*


### 2. Interfaz uniforme: 

La API utiliza una formato coherente, predecible y estandarizado, del mismo modo que lo hacen los recursos. Aquí es dónde el diseño de las URL y los métodos HTTP aparecen.

#### ¿Qué es un recurso?

Un recurso es aquello sobre lo que trata la API. Siguiendo el ejemplo anterior, el recurso es `user`.

Un recurso tiene:
- Una **URL** que lo identifica
- Una **representación** (normalmente, JSON) que lo describe:

```json
{
  "id": "1",
  "username": "John",
  "email": "john@example.com"
}
```

Este objeto JSON es la representación del usuario con id `1`. El servidor contiene los datos reales, este JSON es una copia, (una *representación*), que se envía al cliente.

#### ¿Qué es un método HTTP?

A través del los **métodos HTTP** el cliente le comunica al servidor qué hacer con el recurso solicitado:
  - `GET`: Para solicitar o leer datos.
  - `POST`: Para crear nuevos recursos.
  - `PUT / PATCH`: Para actualizar recursos existentes.
  - `DELETE`: Para eliminar un recurso.
  
En resumen, la URL identifica el *qué* y el método indica *qué hacer*


### 3. Separación Cliente - Servidor: 

Son sistemas independientes y uno no sabe nada de como funciona internamente el otro. El cliente no sabe como el servidor almacena los datos, el servidor no sabe como funciona el cliente. Solo se comunican mediante los contratos establecidos en las API


## ¿POR QUÉ EN GO?
- Al ser un lenguaje compilado, aporta un mayor rendimiento.
- Poco _boilerplate_ lo que simplifica mucho el código.
- Concurrencia. ¿Que sería de Go sin concurrencia?.


**Este repositorio es solo para pruebas, no para un entorno de producción real**
  


  




