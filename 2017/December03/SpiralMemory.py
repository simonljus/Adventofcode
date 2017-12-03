
def part1_testCases(test_input):
    pass_count =0
    number_of_tests= len(test_input)
    for test_value in test_input:
        print( "testing string "+ test_value )
        test_output = SpiralMemory(int(test_value))
        expected_output = test_input[test_value]
        if test_output == expected_output:
            #print( " Part1: Test of " + test_string +  " passed" )
            pass_count +=1
        else:
             print( "Part1: Test of " + test_value + " failed, returned " + str(test_output)  )
    print "Part1: %d of %d passed" % (pass_count,number_of_tests)

def getLayerWidthFromSquare(square_number):
    layer_width =1
    layer_number=0
    layer_max= layer_width * layer_width
    while layer_max< square_number:
        layer_width +=2
        layer_number +=1
        layer_max = layer_width * layer_width
    return layer_number, layer_width, layer_max
    
def getCross(layer_width,layer_number,elements_in_inner_square):
    layer_max = layer_width * layer_width
    right = elements_in_inner_square +layer_number
    cross_arr =[right]
    for i in range(3):
        next_cross = (layer_width -1)*(i+1) + right
        cross_arr.append(next_cross)
    return cross_arr
    


def getDistanceFromClosest(number,my_arr):
    min_distance = number
    #print(number,my_arr)
    for el in my_arr:
        distance = abs(number -el)
        if distance < min_distance:
            min_distance = distance
    return min_distance

def SpiralMemory(input_square):
    steps =0
    number_of_elements =1
    elements_in_inner_square =0
    layer_number, layer_width, layer_max = getLayerWidthFromSquare(input_square)
    
    if(layer_width >1):
        elements_in_inner_square = (layer_width -2) *  (layer_width -2)
        number_of_elements = layer_max - elements_in_inner_square
        cross_list =getCross(layer_width,layer_number,elements_in_inner_square)
        distance = getDistanceFromClosest(input_square,cross_list)
        steps = distance + layer_number
    
    
    #print input_square,layer_number,layer_width
    return steps

def getNextSpiralCoords(prev_square,layer_index):
    prev_x =prev_square["x"]
    prev_y =prev_square["y"]
    prev_val =prev_square["val"]
    #print "Next coord"
    #print prev_x,prev_y, prev_val, layer_index
    if prev_x ==0 and prev_y == 0:
        return 1,0
    if prev_x >= layer_index:
        #print "right side"
        
        if prev_y >= layer_index:
            #print "prev was top corner"
            x = prev_x -1
            y = prev_y
            return x,y
        elif prev_y <= -layer_index:
            #print "prev was bottom right corner"
            x = prev_x +1
            y = prev_y
            return x,y
        else:
            #print "somewhere else on right side"
            x = prev_x
            y = prev_y +1
            return x,y
    
    
   
    elif prev_x <= -layer_index:
        #print "left side"
       
        if(prev_y <= -layer_index):
            x = prev_x +1
            y = prev_y
            return x,y
        else:
            x = prev_x 
            y = prev_y -1
            return x,y
    
    elif prev_y >= layer_index:
        #print "top side"
        
        if(prev_x <= -layer_index):
            x = prev_x
            y = prev_y -1
            return x,y
        else:
            x = prev_x  -1
            y = prev_y
            return x,y
    

    elif (prev_y <= -layer_index):
        #print "bottom side"
        
        if(prev_x >= layer_index):
            x = prev_x +1
            y = prev_y
            return x,y
        else:
            x = prev_x  +1
            y = prev_y
            return x,y
    else:
        return 0,0



def draw_spiral(square_number):
    spiral_matrix = []
    square = {"x":0, "y":0,"val":1}
    first_layer =[square]
    spiral_matrix.append(first_layer)
    biggest = 1
    layer_width = 1
    prev_layer_size =0
    layer_index =0
    while (biggest < square_number):
        prev_layer_size += len(spiral_matrix[-1])
        layer_index +=1
        layer =[]
        layer_width += 2
        square_size= layer_width * layer_width;
        #print "prev sq" + str(prev_layer_size)
       #print "sq" + str(square_size)
        layer_size = square_size - prev_layer_size;

        for i in range(layer_size):
            #print "loop " + str(i)
            if len(layer) >0:
                prev_square = layer[-1]
                x,y = getNextSpiralCoords(prev_square,layer_index)
                 
            else:
                prev_square = spiral_matrix[-1][-1]
                x,y = getNextSpiralCoords(prev_square,layer_index -1)
                 
            next_value = getNextSpiralValue(spiral_matrix[-1],layer,x,y)
            new_square = {"x":x, "y":y,"val":next_value}
            biggest = next_value
            layer.append(new_square)
        spiral_matrix.append(layer)



    return spiral_matrix



def sumNeighborsInLayer(layer,x,y):
    sum_val =0
    for el in layer:
        if isNeighbor(el["x"],el["y"],x,y):
            sum_val += el["val"]
    return sum_val

def isNeighbor(a_x,a_y,b_x,b_y):
    if abs(a_x -b_x) <= 1 and abs(a_y -b_y) <=1:
        return True
    else:
        return False

def getNextSpiralValue(prev_layer,this_layer,x,y):
    layer_sum = sumNeighborsInLayer(this_layer,x,y)
    prev_sum = sumNeighborsInLayer(prev_layer,x,y)
    return prev_sum + layer_sum

def part2(puzzle_input):
    part2_spiral= draw_spiral(puzzle_input)
    for layer in part2_spiral:
        #print "layer"
        for square in layer:
            #print str(square["val"]),str(square["x"]),str(square["y"])
            #print ""
            if square["val"] > puzzle_input:
                return square["val"]
if __name__ == "__main__":
    part1_input = 265149
    #part1_test_input = {"1": 0, "12": 3,"23":2,"1024":31}
    part1_output = SpiralMemory(part1_input)
    #part1_testCases(part1_test_input)
    print ("Result part1 " + str(part1_output))
    part2_output = part2(part1_input)
    print ("Result part2 " + str(part2_output))