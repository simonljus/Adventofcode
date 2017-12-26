import math
def fileToArray(filename):
    input_instructions =[]
    with open(filename,'r') as infile: 
       
        for filerow in infile:
            if len(filerow) > 0:
                row_array = filerow.split()
                if len(row_array) >= 2:
                    input_instructions.append(row_array)    
    return input_instructions

def calculate_increment(register_state,registers,register_arg):
    diff_register ={}
    if register_arg in registers:
        from_val = register_state.get(register_arg,0)
        to_val = registers.get(register_arg,0)
        step_length =  to_val -from_val
        number_of_steps = abs(to_val/step_length)
        for register in registers:
            diff_register[register] = registers[register] - register_state.get(register,0)
            registers[register] = registers[register] + diff_register[register]*number_of_steps
        print "diff", step_length
        print diff_register
        print registers
        print "function end"
    return registers
def isPrime(n):
    n_sqrt = int(math.sqrt(n))
    for i in xrange(2,n_sqrt):
        if n %i ==0:
            return 0
    return 1
    
def part2():
    a =1
    b = 107900 
    c = 107900 + 17000
    d =0
    e =0
    f =0
    g =0
    h =0
    while True: 
        f =1
        d =2
        f = isPrime(b)
        g = 0
        if f ==0:
            h +=1
            g =b
        g -=c 
        if g == 0:
            break
        b += 17
    return h



def part1(instructions):
    registers ={}
    current_instruction =0
    #print instructions
    number_of_instructions = len(instructions)
    last_frequency_value =0
    mul_count =0
    while(current_instruction < number_of_instructions):
        print loop_count
        instruction = instructions[current_instruction]
        op_call = instruction[0]
        #print "op_call", op_call
        first_arg = "" 
        second_arg = ""
        first_arg_val = 0
        second_arg_val =0
        if len(instruction) >=2:
            first_arg = instruction[1]
            if first_arg.isalpha():
                first_arg_val = registers.get(first_arg,0)
            else: 
                first_arg_val = int(first_arg)
        if len(instruction) >=3:
            second_arg = instruction[2]
            if second_arg.isalpha():
                second_arg_val = registers.get(second_arg,0)
            else:
                second_arg_val = int(second_arg)
        if op_call == "jnz":
            if first_arg_val !=0:
                current_instruction += second_arg_val -1
        elif op_call == "mul":
            registers[first_arg] = first_arg_val * second_arg_val
            mul_count +=1
        elif op_call == "set":
            registers[first_arg] = second_arg_val

        elif op_call == "sub": 
            #print "Sending"
            registers[first_arg] = first_arg_val - second_arg_val
        current_instruction +=1
    return mul_count

if "__main__" == __name__:
    instructions =fileToArray("input.txt")
    #print "part1",part1(instructions)
    registers ={'a': 1}
    current_instruction =0
    print part2()
    #print "part2",part2(instructions,1,registers,current_instruction) 
    #registers ={'a': 1, 'c': 124900, 'b': 107900, 'e': 107900, 'd': 107900, 'g': 0, 'f': 0}  
    #current_instruction =23
    #part2(instructions,registers,current_instruction) 
    #print calculate_increment({"a":2,"b":3},{"a":5,"b":2},"b")