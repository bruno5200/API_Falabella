#la imagen de golang en Docker
FROM golang

#  creamos el directorio
RUN mkdir /app
# copiamos el volumen al directorio
ADD . /app
#selecionamos app como directorio de trabajo
WORKDIR /app
# ejecutamos el comando build del Makefile
RUN make build
# ejecutamos el binario
CMD make run-bin
# exponemos el puerto 3000
EXPOSE 3000