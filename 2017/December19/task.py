def readMap(filename):
    map_list=[]
    with open(filename,"r") as fp:
        for file_row in fp:
            map_list.append(list(file_row.replace(" ","#")))
        return map_list

def printNearby(map_matrix,row,col):
    row_length = len(map_matrix)
    top_left = "?"
    above = "?"
    top_right = "?"
    left = "?"
    right = "?"
    center = "?"
    bottom_left ="?"
    bottom = "?"
    bottom_right = "?"
    top_line = [top_left,above,right]
    center_line = [left,center,right]
    bottom_line = [bottom_left,bottom,bottom_right]
    
    number_of_rows =  len(map_matrix)
    #top_line
    if row -1 >= 0 and row -1 < row_length:
         #top_left
        if col -1>= 0 and col -1 < len(map_matrix[row -1]):
            top_line[0] = map_matrix[row -1][col -1]
        if col >= 0 and col < len(map_matrix[row -1]):
            top_line[1] = map_matrix[row -1][col]

        if col +1 >=0 and col +1 < len(map_matrix[row -1]):
            top_line[2] = map_matrix[row -1][col +1]
    #centerline
    if row  >= 0 and row < row_length:
         
        if col -1>= 0 and col -1 < len(map_matrix[row]):
            center_line[0] = map_matrix[row][col -1]
        if col >= 0 and  col < len(map_matrix[row]) :
            center_line[1] = map_matrix[row][col]

        if col +1 >=0 and col +1 < len(map_matrix[row]):
            center_line[2] = map_matrix[row][col +1]

    if row +1 >= 0 and row +1 < row_length:
         #top_left
        if col -1>= 0 and col -1 < len(map_matrix[row +1]):
            bottom_line[0] = map_matrix[row +1][col -1]
        if col >= 0 and col < len(map_matrix[row +1]):
            bottom_line[1] = map_matrix[row +1][col]

        if col +1 >=0 and col +1 < len(map_matrix[row +1]):
            bottom_line[2] = map_matrix[row +1 ][col +1]
    
    nearby_matrix = [top_line,center_line,bottom_line]
    
    print top_line
    print center_line
    print bottom_line
    


def verticalMove(map_list,sentence,current_row,current_col,direction,move_count):
    possible_moves ="|-+"
    move_found = True 
    new_direction = 1
    while(move_found):
        move_found = False
        if (current_row + direction < len(map_list)) and  current_row + direction >= 0:

            pos_char = map_list[current_row +direction][current_col]
            if pos_char in possible_moves:
                current_row +=direction
                move_found = True
                move_count +=1
            elif pos_char.isalpha():
                move_found = True
                sentence.append(pos_char)
                current_row +=direction
                move_count +=1
                print "vertical found_letter", pos_char,move_count
        if move_found == False:
            if current_col -1 > 0:
                 pos_char = map_list[current_row][current_col -1]
                 if pos_char in possible_moves or pos_char.isalpha():
                    new_direction = -1
    return current_row, new_direction,move_count

def horizontalMove(map_list,sentence,current_row,current_col,direction,move_count):
    possible_moves = "|-+"
    move_found = True
    new_direction = 1
    while(move_found):
        move_found = False
        if (current_col + direction < len(map_list[current_row])) and  current_col + direction > 0:
            pos_char = map_list[current_row][current_col + direction]
            if pos_char in possible_moves:
                current_col+=direction
                move_found = True
                move_count+=1
            elif pos_char.isalpha():
                move_found = True
                sentence.append(pos_char)
                current_col +=direction
                move_count +=1
                print " horizontal found letter", pos_char, move_count
        if move_found == False:
            if current_row -1 > 0:
                 pos_char = map_list[current_row -1][current_col]
                 if pos_char in possible_moves or pos_char.isalpha():
                    new_direction = -1

            

    return current_col,new_direction,move_count

    




def part1(map_list):
    number_of_rows = len(map_list)
    print "number of rows",number_of_rows
    pissbile_moves = "|-+"
    sentence = []
    move_found = True
    current_row =0
    current_col = map_list[0].index("|")
    print "start", current_row,current_col
    current_direction = "vertical"
    direction =1
    horizontal_Lock = False 
    vertical_lock = False
    move_count =0
    while(vertical_lock == False or horizontal_Lock== False):
        #print current_row, current_col
        if current_direction == "vertical" and vertical_lock == False:
            #print "new_row"
            new_row, new_direction,move_count = verticalMove(map_list,sentence,current_row,current_col,direction,move_count)
            if new_row != current_row:
                current_row = new_row
                horizontal_Lock = False
                # do something here so it can check vertical
            else:
                vertical_Lock = True
            direction = new_direction
            current_direction = "horizontal"
            continue
        elif current_direction == "horizontal" and horizontal_Lock ==False:
            #print "new_col"
            new_col, new_direction,move_count = horizontalMove(map_list,sentence,current_row,current_col,direction,move_count)
            #printNearby(map_list,current_row,new_col)
            #print "new_col_res"
            if new_col != current_col:
                current_col = new_col
                vertical_Lock = False    
            else:
                vertical_lock = True
            direction = new_direction
            current_direction = "vertical"
            continue
        else:
            break
    result = ""
    for c in sentence:
        result += c
    return result,move_count

        






    return sentence

if "__main__" == __name__:
    map_list = readMap("input.txt")
    #Can find AEGMPQRVY
    print "part1, part2", part1(map_list)
    print "no idea why palindrome, result is the middlechar and its steps +1"
            
