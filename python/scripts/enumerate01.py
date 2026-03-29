"""
Example of enumerate in Python.
"""

print("Example of enumerate in Python.\n\n")

# Create a list of items
sueldos = [1000, 2000, 3000, 4000, 5000, 10000, 20000, 30000, 40000, 50000]
# Creamos un diccionario con el índice y el sueldo utilizando enumerate
index_sueldos = {index : sueldo for index, sueldo in enumerate(sueldos, start=1)}

print("Index and sueldos:")
for index, sueldo in index_sueldos.items():
    print(f"Index: {index}, Sueldo: {sueldo}")

# Filtrar elementos por su posisición
frutas = ["manzana", "banana", "naranja", "pera", "uva", "kiwi", "melon", "sandia", "cereza", "durazno"]
print("\nFrutas con índice múltiplo de 3:")
print([fruta for index, fruta in enumerate(frutas, start=1) if index % 3 == 0])