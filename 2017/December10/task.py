

def fileToIntList(filename):
   row=  open(filename,'r').readline()
   row_strings = row.split(",")
   row_ints = map(int,row_strings)
   return row_ints

def reverseAndReplace(m_list,from_index,number_of_elements):
    i =0
    n = len(m_list)
    reverse_list = []
    while i < number_of_elements:
        reverse_list.append(m_list[(from_index + i) %n])
        i +=1
    i = number_of_elements -1
    j = 0
    while i >=0:
        m_list[(from_index + j) % n ] = reverse_list[i]
        i -=1
        j +=1



def knot_algorithm(length_list, m_list,current_position,skip_size):
    n = len(m_list)
    
    for task in length_list:
        if task < n:
            reverseAndReplace(m_list,current_position,task)
            current_position = (current_position + task + skip_size) % n
            skip_size += 1

    return current_position, skip_size

def part1(length_list, m_list):
    skip_size =0
    current_position=0
    current_position , skip_size = knot_algorithm(length_list,m_list,current_position,skip_size)
    return m_list[0] * m_list[1]
    

def denseHash(m_list):
    n = len(m_list)
    from_index =0
    hash_size = 16
    dense_hash = []
    while from_index <n:
        dense_hash.append( reduce(lambda i, j: i ^ j, m_list[from_index:from_index + hash_size]))
        from_index +=16
    return dense_hash
def hexaHash(dense_hash):
    hexa_string = ""
    for d in dense_hash:
        hexa_ans = hex(d)[2:]
        if len(hexa_ans)< 2:
            hexa_ans ="0" + hexa_ans
        hexa_string+= hexa_ans
    return hexa_string
def part2(m_list,input_string, ascii_input):
    length_list = []
    rounds = 64
    current_round = 0
    skip_size =0
    current_position=0
    for c in input_string:
        length_list.append(ord(c))
    for ascii_char in ascii_input:
        length_list.append(ascii_char)
    #print length_list
    while current_round < rounds:
        #print str(current_round)
        current_position, skip_size = knot_algorithm(length_list,m_list,current_position,skip_size)
        current_round +=1
    #print m_list
    dense_hash =  denseHash(m_list)
    hexa_hash = hexaHash(dense_hash)
    return hexa_hash
def testCases(part1_input,part2_input,part1_list,part2_ascii):
    part2_tests(part1_list[:], part2_input,part2_ascii[:])

def part2_tests(input_list,length_string,length_ascii):
    test_cases = {"":"a2582a3a0e66e6e86e3812dcb672a272","AoC 2017":"33efeb34ea91902bb2f59c9920caa6cd","1,2,3":"3efbe78a8d82f29979031a4aa0b16a9d","1,2,4":"63960835bcdc130f0b66d7ff4f6a5a8e"}
    for test in test_cases:
        print "testing",test
        test_answer = test_cases[test]
        if test_answer == part2(input_list[:],test,length_ascii[:]):
            #print "passed test: ", test 
            print ""
        else:
            print "failed test:",test

if __name__ == "__main__":
    task_input = fileToIntList("input.txt")
    part2_input = open("input.txt",'r').readline()
    part1_list = range(256)
    #part1_output = part1(task_input[:],part1_list[:])
    part2_ascii = [17, 31, 73, 47, 23]
    part2_output = part2(part1_list[:],part2_input,part2_ascii[:])
    #print str(part1_output)
    #testCases(task_input,part2_input,part1_list,part2_ascii)
    print part2_output
