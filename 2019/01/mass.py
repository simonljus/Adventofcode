
def calculate_fuel():
    with open("input.txt") as file:
        prob_2 = 0
        prob_1 = 0
        for number in file:
            mass =int(number.rstrip())
            fuel = mass_to_fuel(mass)
            prob_1 += fuel
            extra = calculate_extra(fuel)
            prob_2 += fuel + extra
            
    print('Problem 1:',prob_1,"units of fuel")
    print('Problem 2:',prob_2,"units of fuel")

def mass_to_fuel(mass):
    fuel = int(mass/3) -2
    return fuel

def calculate_extra(mass):
    extra_fuel = 0
    while True:
        mass = mass_to_fuel(mass)
        if mass <=0:
            return extra_fuel
        extra_fuel += mass  

calculate_fuel()     
print('Jag och Alex proggar för första gången ihop<3')








