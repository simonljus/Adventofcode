import time

def fileToArrays(filename):
    firewall_dict = {}
    with open(filename,"r") as fp:
        for row in fp:
            mrow = row.replace(":","")
            row_arr =mrow.split()
            if len(row_arr) == 2:
                firewall_dict[int(row_arr[0])] = {"array_length": int(row_arr[1]),"scanner_pos":0,"velocity": 1}
    return firewall_dict.copy()

def copyDictInDict(dict_to_copy):
    new_dict ={}
    for dict_el in dict_to_copy:
        new_dict[dict_el] = dict_to_copy[dict_el].copy()
    return new_dict

def posAtTime(array_length,curr_time):
    current_pos =0
    if array_length >1:
        max_index = array_length -1
        mod_ans = curr_time % (max_index)
        if (curr_time / max_index) % 2 == 0:
            current_pos = mod_ans
        else:
            current_pos = (max_index) - mod_ans
    return current_pos

def part2(firewall_inp):
    max_firewall =max(firewall_inp.iterkeys())
    delay =0
    current_time =0
    current_layer =0
    caught = True
    deepest =0
    while caught:
        #print "delay and deepest",delay,deepest,"\n"
        current_layer =0
        caught = False
        current_time = delay
        while(current_layer <= max_firewall):
            if current_layer in firewall_inp:
                if posAtTime(firewall_inp[current_layer]["array_length"],current_time) ==0:
                    caught= True
                    delay +=1
                    break
                else:
                    current_layer +=1
                    current_time +=1
            else:
                current_layer +=1
                current_time +=1
            if current_layer > deepest:
                deepest = current_layer
                print "deepest: ", deepest,delay,"\n"

    return delay

def part1(firewall_dict):
    max_firewall =max(firewall_dict.iterkeys())
    current_layer =0
    severity =0
    caught =True
    while(current_layer <= max_firewall):
        #print current_layer,firewall_dict
        print "\n"
        if current_layer in firewall_dict:
            layer_info = firewall_dict[current_layer]
            if layer_info["scanner_pos"] == 0:
                severity += (current_layer * (layer_info["array_length"]))
        timePasses(firewall_dict)
        current_layer +=1          
    print "debug this pls"
    return severity


def isScannerApproaching(layer_info):
    next_step = layer_info["scanner_pos"] + layer_info["velocity"]
    return next_step ==0 or next_step == layer_info["array_length"] -1

def timePasses(firewall_dict):
    for firewall in sorted(firewall_dict):
        if firewall_dict[firewall]["array_length"] >2:
            if firewall_dict[firewall]["array_length"] -1 == firewall_dict[firewall]["scanner_pos"]:
               firewall_dict[firewall]["velocity"] = -1
            elif firewall_dict[firewall]["scanner_pos"] == 0:
                firewall_dict[firewall]["velocity"] =1
            #print firewall_dict[firewall]["scanner_pos"], firewall_dict[firewall]["max_pos"],"before\n" 
            firewall_dict[firewall]["scanner_pos"] += firewall_dict[firewall]["velocity"]
            #print firewall_dict[firewall]["scanner_pos"], "after\n"
        else:
            firewall_dict[firewall]["scanner_pos"] = (firewall_dict[firewall]["scanner_pos"] +1) %  2
    return firewall_dict.copy()

if "__main__" == __name__:      
    my_dict = fileToArrays("input.txt")
    #print part1(my_dict)
    start_time = time.time()
    print part2(my_dict.copy()),"\n"
    print str(time.time() - start_time)

            
        