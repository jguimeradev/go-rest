# GO REST

RESTful API en Go para utilizar con otros proyectos enfocados a *DevOps* y *Cloud* 


## ¿QUÉ ES UNA REST API?

REST (Representational State Transfer) es un modelo que permite que aplicaciones se comuniquen mediante el uso de API sobre HTTP. REST no es un protocolo ( _como SOAP_ ), sino un conjunto de "indicaciones" lo que permite implementarlo de diferentes maneras.

## CARACTERÍSTICAS DE UNA REST API:

- Sin estado (Stateless): Las peticiones del cliente son independientes. El servidor no almacena información ni sesiones del cliente entre solicitudes, por lo que cada llamada debe incluir todos los datos necesarios para procesarse.
- Interfaz uniforme: Garantiza que las peticiones y respuestas sigan un formato coherente, predecible y estandarizado. Se utilizan métodos HTTP estándar para manipular datos:
  - `GET`: Para solicitar o leer datos.
  - `POST`: Para crear nuevos recursos.
  - `PUT / PATCH`: Para actualizar recursos existentes.
  - `DELETE`: Para eliminar un recurso.
- Cliente-Servidor: Ambos sistemas están desacoplados y son independientes. El servidor se encarga del almacenamiento y lógica de negocio, mientras que el cliente maneja la interfaz de usuario y la experiencia.
- Almacenamiento en caché (Cacheable): Las respuestas del servidor deben indicar explícitamente si son almacenables en caché o no. Esto mejora drásticamente el rendimiento y reduce la latencia al evitar procesar peticiones repetidas.
- Sistema en capas (Layered System): El cliente no necesita saber si se comunica directamente con el servidor o a través de intermediarios como proxies, balanceadores de carga o gateways. Esto permite escalar y asegurar la arquitectura de forma transparente.


## ¿POR QUÉ EN GO?
- Al ser un lenguaje compilado, aporta un mayor rendimiento.
- Poco _boilerplate_ lo que simplifica mucho el código.
- Concurrencia. ¿Que sería de Go sin concurrencia?.
  


  




