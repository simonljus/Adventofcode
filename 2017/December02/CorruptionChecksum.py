


def sumDifferenceMinMax(input_matrix):
    sum =0
    for row in input_matrix:
        max_el =0 
        min_el = 100000000
        for el in row:
            if el > max_el:
                max_el = el
            if el < min_el:
                min_el = el
        sum += abs(max_el - min_el)
    return sum
def sumEvenDivision(input_matrix):
    sum =0
    for row in input_matrix:
        flag = False 
        for namn in sorted(row):
                for talj in row:
                    div_res = float(talj)/float(namn)
                    if div_res == float(talj)//float(namn) and div_res != 1.0:
                        sum += int(div_res)
                        break;
    
    return sum

    
    
    return my_res
def fileToIntMatrix(filename):
    with open(filename,'r') as infile: 
        input_matrix =[]
        for filerow in infile:
            row_strings = filerow.split()
            row_ints = map(int,row_strings)
            input_matrix.append(row_ints)
    return input_matrix


if __name__ == "__main__":
    input_matrix = fileToIntMatrix("input.txt")
    res_part1 =sumDifferenceMinMax(input_matrix)
    res_part2 =sumEvenDivision(input_matrix)


    print "Part1 result: %d" % (res_part1)
    print "Part2 result: %d" % (res_part2)