import math

def calculate_square_root(a, b):
    sqrt_a = math.sqrt(a)
    if sqrt_a.is_integer():
        print(sqrt_a)
    else:
        print("The square root is not an integer.")

calculate_square_root(16, 25)
