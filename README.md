# curso_golang_platzi
Curso básico de Go/Golang en Platzi, con muchas cosas interesantes, para poder el el historial hay que cer los commits en la rama main (si, la idea es que tienes que navegar entre los distintos commits), cada commit es una clase con ejemplos tratados en el curso.

## Descarga de packages

1. Para descargar librerías de Go/Golang es usando el siguiente comando.

```bash
go get golang.org/x/website/tour 
```

2. Para descargar la librería y ver la versión, hay que usar el flag '-v'.

```bash
go get -v golang.org/x/website/tour 
```

3. Para obtener más información y volver a descargar la librería.

```bash
go get -v -u golang.org/x/website/tour 
```

## Modificar modulos en Go/Golang

Para modificar módulos, es necesario modificar el go.mod usando las funcionalidades del lenguaje (el módulo tiene que estar en la carpeta raíz del proyecto actual, es decir en 'curso_golang_platzi').

```bash
git clone git@github.com:labstack/echo.git // hacer los cambio en el módulo
go mod edit -replace github.com/labstack/echo=./echo
```

Para verificar el go.mod, se usa el comando 
```bash
go mod verify
```

Para volver el módulo a su estado original, para el cambio hay que usar las funcionalidades del lenguaje
```bash
go mod edit -dropreplace github.com/labstack/echo
```

## Empaquetar código

Para empaquetar el código, se usa el siguiente comando.
```bash
go mod vendor
```