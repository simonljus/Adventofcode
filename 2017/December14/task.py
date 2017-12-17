def wordFromFile(filename):
    return "ugkiagan"

def hexaToBinary(hex_char):
    num_of_bits =4
    scale =16
    return str(bin(int(hex_char, scale))[2:].zfill(num_of_bits))

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

def knot_algorithm(length_list, m_list,current_position,skip_size):
    n = len(m_list)
    
    for task in length_list:
        if task < n:
            reverseAndReplace(m_list,current_position,task)
            current_position = (current_position + task + skip_size) % n
            skip_size += 1

    return current_position, skip_size

def knot_hash(input_string):
    ascii_input = [17, 31, 73, 47, 23]
    m_list = range(256)

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




def hashString(string_to_hash):
    asc_string = ""
    use_count =0
    binary_string =""
    hexa_string = knot_hash(string_to_hash)
    for h in hexa_string:
        binary_string += hexaToBinary(h)

    for b in binary_string:
        if b == "1":
            use_count +=1
    return binary_string,use_count
def part1(key_string):
    tot =0
    for n in range(128):
        hash_string =key_string + "-" + str(n)
        binary_string,number_squares = hashString(hash_string)
        tot += number_squares
    return tot
def part2(key_string):
    tot =0
    binary_matrix =[]
    for n in xrange(128):
        #print "run", n
        hash_string =key_string + "-" + str(n)
        binary_string,number_squares = hashString(hash_string)
        binary_string_array = []
        for c in binary_string:
            binary_string_array.append(int(c))
        binary_matrix.append(binary_string_array)
    print "calculate regions now \n"
    return regions(binary_matrix)

def groupMatrix(binary_matrix,region_queue,region_id):
    matrix_size = len(binary_matrix)
    while len(region_queue) >0:
        matrix_index = region_queue[0]
        del region_queue[0]
        el_index = matrix_index % matrix_size
        row_index = matrix_index / matrix_size
        binary_matrix[row_index][el_index] = region_id
        #check down
        if row_index  < 127:
            if binary_matrix[row_index +1][el_index] ==1:
                 region_queue.append(matrix_index + matrix_size)
        #check right
        if el_index < 127:
            if binary_matrix[row_index][el_index +1] == 1:
                region_queue.append(matrix_index +1)
        # check up
        if row_index  > 0:
            if binary_matrix[row_index -1][el_index] ==1:
                 region_queue.append(matrix_index - matrix_size)
        # check left
        if el_index > 0:
            if binary_matrix[row_index][el_index -1] ==1:
                 region_queue.append(matrix_index -1)
        
    
    

def regions(binary_matrix):
    regions =1
    row_index =0
    el_index =0
    matrix_size =len(binary_matrix)
    while row_index < matrix_size:
        #print "row", row_index
        el_index =0

        while el_index < matrix_size:
            el = binary_matrix[row_index][el_index]
            #print "row col", row_index,el_index
            if el == 1:
                regions +=1
                region_queue =[]
                region_queue.append(row_index * matrix_size + el_index)
                groupMatrix(binary_matrix,region_queue,regions)
            el_index +=1
        row_index +=1
            
                
    return regions -1


if __name__ == "__main__":
    task_input = wordFromFile("input.txt")
    part1_output = part1(task_input)
    print "part1 done", part1_output
    part2_output = part2(task_input)
    print "part1 ", part1_output
    print "part2", part2_output
    