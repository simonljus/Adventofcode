




def part1(number_of_steps,iterations,value_after):
    my_arr =[0]
    i =1 
    current_pos =0

    while(i <= iterations):
        current_pos = (current_pos + number_of_steps) % (i)
        current_pos +=1
        my_arr.insert(current_pos,i)
        i +=1
    return my_arr[ (my_arr.index(value_after) +1) %i]

def part2(number_of_steps,iterations):
    my_arr =[0]
    i =1 
    current_pos =0
    value_after =0

    while(i <= iterations):
        current_pos = (current_pos + number_of_steps) % (i)
        current_pos +=1
        if current_pos == 1:
            value_after = i
        i +=1
    return value_after

if "__main__" == __name__:
    task_input = 356
    print "part1",part1(task_input,2017,2017)
    print "part2",part2(task_input,50000000)
