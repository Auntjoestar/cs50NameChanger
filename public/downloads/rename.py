import os

# PROCEDURAL APPROACH
def rename(path, name, extension):
    # C:\Users\sebal\Downloads\semana 1
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
