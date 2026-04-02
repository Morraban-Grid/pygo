lista1 = ["rodolfo", "eva", "alejandro", "ana", "karen", "max"]
lista2 = [30, 14, 10, 19, 21, 25]

dict1 = {i:j for i,j in zip(lista1, lista2)}

for a, b in dict1.items():
    print(f"{a} tiene {b} años")