def file_to_array(filename):
    input_instructions =[]
    with open(filename,'r') as infile: 
       
        for filerow in infile:
            if len(filerow) > 0:
                row_array = filerow.split()
                if len(row_array) >= 2:
                    input_instructions.append(row_array)    
    return input_instructions

def part1(instructions):
    registers ={}
    current_instruction =0
    #print instructions
    number_of_instructions = len(instructions)
    last_frequency_value =0
    while(current_instruction < number_of_instructions):
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
        if op_call == "add":
            registers[first_arg] = first_arg_val + second_arg_val

        elif op_call == "jgz":
            if first_arg_val >0:
                current_instruction += second_arg_val -1
        elif op_call == "mod":
            if second_arg_val != 0:
                registers[first_arg] = first_arg_val % second_arg_val

        elif op_call == "mul":
            registers[first_arg] = first_arg_val * second_arg_val
            
        elif op_call == "rcv":
            if first_arg_val != 0:
                break

        elif op_call == "set":
            registers[first_arg] = second_arg_val

        elif op_call == "snd": 
            #print "Sending"
            last_frequency_value = first_arg_val
        current_instruction +=1
    return last_frequency_value

def runprogram(registers,current_instruction,instruction,sending_count,receive_q,sending_q,waiting,waiting_other):
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
        if op_call == "add":
            registers[first_arg] = first_arg_val + second_arg_val

        elif op_call == "jgz":
            if first_arg_val >0:
                current_instruction += second_arg_val -1
        elif op_call == "mod":
            if second_arg_val != 0:
                registers[first_arg] = first_arg_val % second_arg_val

        elif op_call == "mul":
            registers[first_arg] = first_arg_val * second_arg_val
            
        elif op_call == "rcv":
            if len(receive_q) > 0:
                registers[first_arg] = receive_q[0]
                del receive_q[0]
            else:
                current_instruction -=1 #stayonthesame
                waiting = True

        elif op_call == "set":
            registers[first_arg] = second_arg_val

        elif op_call == "snd": 
            sending_q.append(first_arg_val)
            waiting_other = False
            sending_count +=1
        current_instruction +=1
        return current_instruction,sending_count,waiting,waiting_other
def part2(instructions):
    registers_0 ={"p":0}
    registers_1 ={"p":1}
    receive_q_0= []
    receive_q_1= []
    sending_count_1 =0 
    sending_count_0 =0
    current_instruction_0 =0
    current_instruction_1 =0
    number_of_instructions = len(instructions)
    waiting_0 = False
    waiting_1 = False
    while((waiting_0 == False  or waiting_1 ==False)  and (current_instruction_0 < number_of_instructions or current_instruction_1 < number_of_instructions)):
        while(current_instruction_0 < number_of_instructions and waiting_0 == False):
            instruction = instructions[current_instruction_0]
            #print "running 0", current_instruction_0
            current_instruction_0,sending_count_0,waiting_0,waiting_1  = runprogram(registers_0,current_instruction_0,instruction,sending_count_0,receive_q_0,receive_q_1,waiting_0,waiting_1)
        while(current_instruction_1 < number_of_instructions and waiting_1 == False):
            instruction = instructions[current_instruction_1]
            #print "running 1", current_instruction_1
            current_instruction_1,sending_count_1,waiting_1,waiting_0  = runprogram(registers_1,current_instruction_1,instruction,sending_count_1,receive_q_1,receive_q_0,waiting_1,waiting_0)
    
    return sending_count_1





if "__main__" == __name__:
    instruction_array=file_to_array("input.txt")
    print "part1", part1(instruction_array)
    print "part2", part2(instruction_array)
