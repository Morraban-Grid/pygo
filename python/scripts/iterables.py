from collections.abc import Iterable

texto = "hola mundo"
numero = 32
lista = [1,2,3,4,5,6,7]

print(isinstance(texto, Iterable))
print("Número es un objeto: ", isinstance(numero, int))
print(isinstance(lista, Iterable))

class Perro():
    def __init__(self, nombre, raza):
        self.nombre = nombre
        self.raza = raza

firulais = Perro("firulais", "chihuahua")

print("firulais es instancia de Perro: ", isinstance(firulais, Perro))

tupla01 = (1,2,3,4,5,6,7)
print("\nSuma de los elementos de una tupla: ", sum(tupla01))


lista_01 = ["1", "2", "3", 4, 5, "6", 7, "8", "9", "10"]
elementos_unidos = "-".join(map(str, lista_01))
print(elementos_unidos)