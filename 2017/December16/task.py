
def fileToArray(filename):
    input_string = open(filename,"r").readline()
    input_array = input_string.split(",")
    return input_array

def dance_with_instructions(input_array,instruction_array):
    for instruction in instruction_array:
        #print input_array,instruction
        if instruction[0] =="s":
            i =int(instruction[1:])
            input_array = input_array[-i:] + (input_array[:-i])
            
        elif instruction[0] == "x":
            from_to = instruction[1:].split("/")
            first_index = int(from_to[0])
            second_index = int(from_to[1])
            temp =  input_array[first_index] 
            input_array[first_index] = input_array[second_index]
            input_array[second_index] = temp
        elif instruction[0] == "p":
            from_to = instruction[1:].split("/")
            first_val = from_to[0]
            second_val = from_to[1]
            first_index =input_array.index(first_val)
            second_index = input_array.index(second_val)
            temp =  input_array[first_index] 
            input_array[first_index] = input_array[second_index]
            input_array[second_index] = temp
    return input_array
    

def part1(input_array,instruction_array,number_of_runs):
    formation_string =""
    formation_list = []
    cycle_found = False
    i =0
    for c in input_array:
       formation_string += c
    formation_list.append(formation_string)
    while i < number_of_runs:
        input_array = dance_with_instructions(input_array[:],instruction_array)
        formation_string =""
        for c in input_array:
            formation_string += c
        if formation_string in formation_list:
            cycle_found = True
            break
        else:

            formation_list.append(formation_string)
        i+=1
    output_formation = formation_string
    if cycle_found:
        print "Cycle found", len(formation_list)
        formation_index = number_of_runs % len(formation_list)
        output_formation = formation_list[formation_index] 
    
    return output_formation




def part2(input_array,instruction_array,number_of_runs):
    compare_start = input_array[:]
    output_string = ""
    start_sequence =""
    sequence_found = False
    index_array = []
    dance_sequences= []
    for c in compare_start:
        start_sequence += c
    dance_sequences.append(start_sequence)
    print "input array ",input_array
    #print "instruction array", instruction_array
    print "number runs", number_of_runs
    print "comapre start", compare_start

    
    input_array = dance_with_instructions(input_array,instruction_array)
    for c in compare_start:
        index_array.append(input_array.index(c))
    
    dance_string =""
    for c in input_array:
        dance_string += c
    if dance_string in dance_sequences:
        sequence_found = True
    else: 
        dance_sequences.append(dance_string)
    run_i =1

    while( run_i < number_of_runs):
        temp_array = input_array[:]
        print temp_array
        print input
        for from_index,to_index in enumerate(index_array):
            input_array[to_index] = temp_array[from_index]
        dance_string =""
        for c in input_array:
            dance_string += c
        if dance_string in dance_sequences:
            sequence_found = True
            break;
        else: 
            dance_sequences.append(dance_string)
        run_i +=1
    if sequence_found:
       print "sequence length found", len(dance_sequences) 
       print dance_sequences
       sequence_index = number_of_runs % len(dance_sequences)
       return dance_sequences[sequence_index]
    

            
    for c in input_array:
        output_string += c

    return output_string
if "__main__" == __name__:
    instruction_array = fileToArray("input.txt")
    input_string = "abcdefghijklmnop"
    input_array =[]
    part1_runs =1000000000
    part2_test = 10
    part2_runs = 1000000000
    for c in input_string:
        input_array.append(c)
    part1_res =  part1(input_array[:],instruction_array[:],part1_runs)
    
    #part2_output = part2(input_array[:],instruction_array[:],part2_test)
    print "part1 res", part1_res
    #print "part2 res:", part2_output
