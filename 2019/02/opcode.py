def get_instructions(first,second):
    instructions=[]
    with open("input.txt") as file:
        for line in file:
            instruction_line=list(map(int, line.rstrip().split(",")))
            instructions += instruction_line
    instructions[1]=first
    instructions[2]=second
    return instructions

def problem_1():
    ptr=0
    terminate=99
    adder=1
    multiplier=2
    step=4
    instructions=get_instructions(12,2)
    opcode=instructions[ptr]
    while opcode !=terminate:
        pos1 =instructions[ptr+1]
        pos2=instructions[ptr+2]
        pos3=instructions[ptr+3]
        op1 = instructions[pos1]
        op2 = instructions[pos2]
        store = pos3
        if(opcode==adder):
            instructions[store]=op1+op2
        elif(opcode==multiplier):
            instructions[store]=op1*op2
        else:
            continue
        ptr +=step
        opcode =instructions[ptr]
    return instructions[0]

def problem_2():
    ptr=0
    terminate=99
    adder=1
    multiplier=2
    step=4
    code = 19690720
    for first in range(100):
        for second in range(100):
            ptr=0
            print(first,second)
            instructions=get_instructions(first,second)
            opcode=instructions[ptr]
            while opcode !=terminate:
                op1 = instructions[instructions[ptr+1]]
                op2 = instructions[instructions[ptr+2]]
                store = instructions[ptr+3]
                if(opcode==adder):
                    instructions[store]=op1+op2
                elif(opcode==multiplier):
                    instructions[store]=op1*op2
                else:
                    continue
                ptr +=step
                opcode =instructions[ptr]
            if(instructions[0] == code):
                return 100*first +second
ans1=problem_1()
ans2=problem_2()
print("problem 1",ans1)
print("problem 2",ans2)