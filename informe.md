# Modelo MVC y patrón repositorio

## Introducción

MVC o Modelo Vista Controlador es un patrón en el diseño de software comúnmente utilizado para implementar interfaces de usuario, datos y lógica de control. Enfatiza una separación entre la lógica de negocio y su visualización.

El patrón repositorio por otro lado está relacionado también con el concepto de Unidad de Trabajo o Unit of Work los cuales están enfocados en crear una capa de abstracción entre la capa de acceso a datos y la capa de lógica de negocio.

## Modelo MVC

El MVC es un modelo de arquitectura de software el cual está compuesto de tres partes, ell modelo, la vista y el controlador.

El modelo define qué datos debe contener la aplicación. Si el estado de estos datos cambia, el modelo generalmente notificará a la vista (para que la pantalla pueda cambiar según sea necesario) y, a veces, el controlador (si se necesita una lógica diferente para controlar la vista actualizada).

El controlador contiene una lógica que actualiza el modelo y / o vista en respuesta a las entradas de los usuarios de la aplicación.

La vista es la interfaz de usuario, define cómo se deben mostrar los datos de la aplicación.

las vistas suelen recibir los tipos de datos que dan los modelos y se 
representan al usuario.