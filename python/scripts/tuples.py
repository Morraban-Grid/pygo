"""

Example of tuples in Python.


"""

print("\nTuples are immutable sequences of values.\n")

nombres = ("alice", "bob", "charlie", "dave", "eve", "frank", "grace", "heidi", "ivan", "judy")
cargos = ("devops", "developer", "designer", "manager", "analyst", "tester", "support", "sales", "marketing", "hr")
salarios = (50000, 60000, 70000, 80000, 90000, 100000, 110000, 120000, 130000, 140000)
nacionalidades = ("usa", "canada", "mexico", "brazil", "argentina", "chile", "peru", "colombia", "venezuela", "ecuador")

estructura = zip(nombres, cargos, salarios, nacionalidades)

for nombre, cargo, salario, nacionalidad in estructura:
    print(f"{nombre} es un {cargo} que gana {salario} y es de {nacionalidad}.")

"""lista1 = [1, 2, 3, 4]
lista2 = ["A", "B", "C", "D", "E"]
lista3 = list(zip(lista1, lista2, strict = True))

print("Lista de tuplas:", lista3)"""