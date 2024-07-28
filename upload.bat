#FERNANDODEV
# Paso 1: Agregar y cometer cambios en Git
git add .
git commit -m "Ultimo Commit"
git push

# Paso 2: Configurar variables de entorno para la compilación cruzada de Go
set GOARCH=amd64
set GOOS=linux

# Paso 3: Construir el binario de Go y asignarle el nombre
# personalizado
go build -o bootstrap main.go

# Paso 4: Eliminar el archivo ZIP previo si existe
del deployment.zip

# Paso 5: Crear un nuevo archivo ZIP con el binario
tar.exe -a -cf deployment.zip bootstrap