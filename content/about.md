+++
date = '2024-12-05T23:55:20-06:00'
draft = true
title = 'Acerca de'
+++

Esta aplicación surge como un proyecto personal para automatizar el **proceso de cambio de nombres**
a las imagenes tomadas en los cursos de [CS50 y Web50 Nicaragua](https://code-fu.net.ni/).


Originalmente, este programa era un script de [Python](https://www.python.org/) 
al cual se le debía pasar la ruta de una carpeta con todas las imagenes y el nombre que se repetirá
durante todas las fotos:


```python
import os

def rename(path, name, extension):
    for root, dirs, files in os.walk(path):
        for count, file in enumerate(files):
            os.rename(os.path.join(root, file), os.path.join(root, f'{name}{count+1}.{extension}'))

def main():
    path = input('Ingrese el path donde se encuentran los archivos: ')
    name = input('Ingrese cómo desea que comience el nombre de los archivos: ')
    extension = input("Ingresa la extensión de los archivos (ej: png, jpg, txt, etc): ")
    rename(path, name, extension)

if __name__ == '__main__':
    main()
```

Sin embargo, decidí que si quería compartir esto con otra gente era más apropiado hacerlo un programa propio,
un poco más *formal*. 

Mi primera opción fue hacerlo en [winforms con C#](https://github.com/Auntjoestar/CS50-Name-Changer), que tenía los **programas/ciclos/semanas/grupos** estático. Más, pensando en la practicidad, pensé que sería apropiado agregarle una base de datos para poder agregar dinámicamente contenido.

Trabajando en esto, encontré algo frustrante el **manejo de interfaces en winforms** y no me gustaba la cantidad de **boilerplate** que había. Entonces me dispuse a buscar alternativas, como tengo algo de experiencia en el **desarrollo web** consideré a [Electron](https://www.electronjs.org/) como una opción, aunque **no quería empaquetar una versión de chromium** en una app de este estilo.

Finalmente, di con [Tauri](https://v2.tauri.app/) que me llevó a [Wails](https://wails.io/), entre las dos opciones preferí utilizar **Wails** porque se me hizo más sensato aprender algo de [Go](https://go.dev/) para terminar estre proyecto relativamente rápido que [Rust](https://www.rust-lang.org/) con su escala de complejidad. Este es mi pequeño aporte al mundo, mientras tanto creo que estaré aprendiendo más de **Go**.