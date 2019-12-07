def get_instructions(filename):
    instructions=[]
    with open(filename) as file:
        for line in file:
            instruction_line=list(map(int, line.rstrip().split(",")))
            instructions += instruction_line
    return instructions

def parse_instruction(instruction):
    instruction_string = ("00000" + str(instruction))[-5:]
    mode_a=int(instruction_string[2])
    mode_b=int(instruction_string[1])
    mode_c=int(instruction_string[0])
    opcode = int(instruction_string[-2:])
    return mode_a,mode_b,mode_c,opcode

def get_operand(val,mode,instructions):
    position_mode=0
    immediate_mode=1
    if(mode == immediate_mode):
        return val
    elif(mode ==position_mode):
        return  instructions[val]
    
def problem_1(start,filename):
    ptr=0
    input =start
    terminate=99
    adder=1
    multiplier=2
    read =3
    write =4
    step=4
    jump_if_true=5
    jump_if_false=6
    less_than =7
    equals=8
    position_mode=0
    immediate_mode=1
    instructions=get_instructions(filename)
    while ptr < len(instructions):
        instruction=instructions[ptr]
        mode_a,mode_b,mode_c,opcode=parse_instruction(instruction)
        val_a=instructions[ptr+1]
        val_b=instructions[ptr+2]
        val_c=instructions[ptr+3]
        #print("pos",ptr,"instruction",instruction,"vals",val_a,val_b,val_c)
        #print("opcodes",mode_a,mode_b,mode_c,opcode)
        step=0
        if opcode == read :
            step=2
            #op1 = get_operand(val_a,mode_a,instructions)
            instructions[val_a]=input
        elif opcode == write:
            step=2
            op1 = get_operand(val_a,mode_a,instructions)
            #print("write",val_a,mode_a,op1)
            output =op1
            input =output
            if(output != 0):
                print("invalid output since output is",output)
                return output
        elif opcode==adder :
            step=4
            op1 = get_operand(val_a,mode_a,instructions)
            op2 = get_operand(val_b,mode_b,instructions)
            instructions[val_c]=op1+op2
        elif opcode==multiplier :
            step=4
            op1 = get_operand(val_a,mode_a,instructions)
            op2 = get_operand(val_b,mode_b,instructions)
            instructions[val_c]=op1*op2
        elif opcode == jump_if_true:
            step=3
            op1 = get_operand(val_a,mode_a,instructions)
            if(op1 !=0):
                op2 = get_operand(val_b,mode_b,instructions)
                ptr =op2
                continue
        elif opcode == jump_if_false :
            step=3
            op1 = get_operand(val_a,mode_a,instructions)
            if(op1 ==0):
                op2 = get_operand(val_b,mode_b,instructions)
                ptr =op2
                continue
        elif opcode == less_than :
            step=4
            op1 = get_operand(val_a,mode_a,instructions)
            op2 = get_operand(val_b,mode_b,instructions)
            opstore=0
            if(op1< op2):
                opstore =1
            instructions[val_c]=opstore
        elif opcode == equals :
            step =4
            op1 = get_operand(val_a,mode_a,instructions)
            op2 = get_operand(val_b,mode_b,instructions)
            opstore=0
            if(op1 == op2):
                opstore=1
            instructions[val_c]=opstore
        elif instruction ==terminate :
            return input
        else:
            print("instruction not understood",instruction)
            return -1
        ptr +=step
    return input
    
ans1=problem_1(1,"input.txt")
ans2=problem_1(5,"input.txt")
print("problem 1",ans1)
print("problem 2",ans2)