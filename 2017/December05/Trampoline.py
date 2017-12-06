def file_to_int_array(filename):
    input_instructions =[]
    with open(filename,'r') as infile: 
       
        for filerow in infile:
            if len(filerow) > 0:
                input_instructions.append(int(filerow))
    return input_instructions

def part1(instructions):
    next_instruction =0
    current_pos =0
    instruction_counter =0
    while (next_instruction >=0 and next_instruction < len(instructions)):
        old_instruction = next_instruction
        next_instruction += instructions[next_instruction]
        instructions[old_instruction] += 1
        instruction_counter +=1


    return instruction_counter
def part2(instructions):
    next_instruction =0
    current_pos =0
    instruction_counter =0
    n = len(instructions)
    while (next_instruction >=0 and next_instruction < n):
        old_instruction = next_instruction
        next_instruction += instructions[next_instruction]
        if instructions[old_instruction] >= 3:
            instructions[old_instruction] -= 1
        else:
            instructions[old_instruction] += 1


        
        instruction_counter +=1


    return instruction_counter

if __name__ == "__main__":
    part1_input = file_to_int_array("input.txt")
    #part1_output = part1(part1_input)

    part2_output = part2(part1_input)
    #print str(part1_output)
    print str(part2_output)
