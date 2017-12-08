def file_to_array(filename):
    input_instructions =[]
    with open(filename,'r') as infile: 
       
        for filerow in infile:
            if len(filerow) > 0:
                row_array = filerow.split()
                if len(row_array) >= 7:
                    input_instructions.append(row_array)
    return input_instructions

def part1(instructions):
    # g dec 231 if bfx > -10
    registers ={}
    highest_value_ever = 0
    for instruction in instructions:
        write_register = instruction[0]
        operation = instruction[1]
        term_operation = int(instruction[2])
        read_register = instruction[4]
        conditional = instruction[5]
        term_compare = int(instruction[6])
        conditional_value = False
        read_register_value = registers.get(read_register,0);

        if conditional == ">":
            if read_register_value > term_compare:
                conditional_value = True
            
        elif conditional == ">=":
            if read_register_value >= term_compare:
                conditional_value = True

        elif conditional == "<":
            if read_register_value < term_compare:
                conditional_value = True
        elif conditional == "<=":
            if read_register_value <= term_compare:
                conditional_value = True
        elif conditional == "==":
            if read_register_value == term_compare:
                conditional_value = True
        elif conditional == "!=":
            if read_register_value != term_compare:
                conditional_value = True
        

        if conditional_value:
            if operation == "dec":
                registers[write_register] = registers.get(write_register,0) - term_operation
            elif operation == "inc":
                registers[write_register] = registers.get(write_register,0) + term_operation
        highest_value_yet = registers.get(max(registers,key=registers.get),0)
        if highest_value_yet > highest_value_ever:
            highest_value_ever = highest_value_yet

    return registers.get(max(registers,key=registers.get),0), highest_value_ever



       

if __name__ == "__main__":
    part1_input = file_to_array("input.txt")
    part1_output,part2_output = part1(part1_input)
    print str(part1_output)
    print str(part2_output)