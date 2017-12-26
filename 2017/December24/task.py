def fileToBlocks(filename):
    block_list = [] 

    with open(filename,"r") as fp:
        for filerow in fp:
            row_arr= filerow.rstrip().split("/")
            if len(row_arr) >= 2:
                block_arr = [int(row_arr[0]),int(row_arr[1])]
                block_list.append(block_arr)
    return block_list 


def best_bridge(bridge,blocks,last_connection, curr_strength):
    number_of_blocks = len(blocks)
    queue_children = []
    max_strength =0
    i =0
    max_bridge = []
    while i < number_of_blocks:
        if (i not in bridge) and (last_connection in blocks[i]):
            queue_children.append(i)
        i+=1
    if len(queue_children) ==0:
        return curr_strength,bridge
    for child in queue_children:
        childvals = blocks[child]
        temp_last= childvals[(childvals.index(last_connection) + 1) %2]
        temp = bridge[:]
        temp.append(child)
        strength,n_bridge = best_bridge(temp,blocks,temp_last ,curr_strength + childvals[0] + childvals[1])
        if strength > max_strength:
            max_strength = strength
            max_bridge = n_bridge 
    return max_strength,max_bridge
def part1(blocks):
    max_strength =0
    number_of_blocks =len(blocks)
    strength =0
    bridge = [0]
    max_bridge = []
    for i,block in enumerate(blocks):
        #print bridge, strength
        if 0 in block: 
            bridge = [i]
            strength = block[0] + block[1]
            last_connection = max(block)
            #print i, max_strength
            strength,new_bridge = best_bridge(bridge,blocks[:],last_connection,strength)
            if strength > max_strength:
                max_strength = strength
                max_bridge =new_bridge 
    return max_strength,max_bridge

def eval_longest_bridge(max_strength,max_bridge,strength,bridge):
    max_length = len(max_bridge)
    length = len(bridge)
    if length >= max_length:
        if length == max_length:
            if strength > max_strength:
                max_strength =strength
                max_bridge = bridge
        else:
            max_strength =strength
            max_bridge = bridge
    return max_strength,max_bridge

def eval_strongest_bridge(max_strength,max_bridge,strength,bridge):
    if strength > max_strength:
        max_strength = strength
        max_bridge = bridge 
    return max_strength,max_bridge
    

def part2(blocks,want_longest):
    max_strength =0
    number_of_blocks =len(blocks)
    strength =0
    bridge = [0]
    max_bridge = []
    for i,block in enumerate(blocks):
        #print bridge, strength
        if 0 in block: 
            bridge = [i]
            strength = block[0] + block[1]
            last_connection = max(block)
            #print i, max_strength
            strength,new_bridge = longest_bridge(bridge,blocks[:],last_connection,strength,want_longest)
            length = len(new_bridge)
            max_length = len(max_bridge)
            if want_longest:
                print "want longest big function"
                max_strength, max_bridge = eval_longest_bridge(max_strength,max_bridge,strength,new_bridge) 
                
            else:
                max_strength, max_bridge = eval_strongest_bridge(max_strength,max_bridge,strength,new_bridge) 

                
    return max_strength,len(max_bridge)

def longest_bridge(bridge,blocks,last_connection, curr_strength,want_longest):
    number_of_blocks = len(blocks)
    queue_children = []
    max_strength =0
    i =0
    max_bridge = []
    while i < number_of_blocks:
        if (i not in bridge) and (last_connection in blocks[i]):
            queue_children.append(i)
        i+=1
    if len(queue_children) ==0:
        return curr_strength,bridge
    for child in queue_children:
        childvals = blocks[child]
        temp_last= childvals[(childvals.index(last_connection) + 1) %2]
        temp = bridge[:]
        temp.append(child)
        strength,new_bridge = longest_bridge(temp,blocks,temp_last ,curr_strength + childvals[0] + childvals[1],want_longest)
        if want_longest:
            max_strength, max_bridge = eval_longest_bridge(max_strength,max_bridge,strength,new_bridge) 
        else:
            max_strength, max_bridge = eval_strongest_bridge(max_strength,max_bridge,strength,new_bridge)
    return max_strength,max_bridge
    
        
if __name__ == "__main__":
    block_input = fileToBlocks("input.txt")
    #print "result",part1(block_input)
    print "result", part2(block_input,True)