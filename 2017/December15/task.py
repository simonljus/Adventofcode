#Generator A starts with 699
#factor 16807
#Generator B starts with 124
#factor 48271

def part1(a_input,b_input):
    i =0
    count =0 
    a_prev_generated = a_input
    a_factor = 16807
    b_factor = 48271
    b_prev_generated = b_input
    iterations = 40000000
    divider = 2147483647

    while (i <iterations):
        print(i)
        a = a_factor * a_prev_generated
        b = b_factor * b_prev_generated
        a = a % divider
        b = b% divider
        a_prev_generated =a
        b_prev_generated =b
        bin_a = format(a,'#034b')
        bin_b = format(b,'#034b')
        if bin_a[-16:] == bin_b[-16:]:
            count +=1
        i +=1
    return count

def part2(a_input,b_input):
    i =0
    count =0 
    a_prev_generated = a_input
    a_accepted_multiple =4
    b_accepted_multiple =8 
    a_factor = 16807
    b_factor = 48271
    b_prev_generated = b_input
    iterations = 5000000
    divider = 2147483647

    while (i <iterations):
        print(i)
        a = a_factor * a_prev_generated
        a = a % divider
        b = b_factor * b_prev_generated
        b = b% divider
        while(a % a_accepted_multiple != 0):

            a *= a_factor
            a = a % divider
        while(b % b_accepted_multiple != 0):

            b *= b_factor
            b = b % divider

        a_prev_generated =a
        b_prev_generated =b
        bin_a = format(a,'#034b')
        bin_b = format(b,'#034b')
        if bin_a[-16:] == bin_b[-16:]:
            count +=1
        i +=1
    return count


#print("result:",part1(699,124))
print("result:",part2(699,124))
