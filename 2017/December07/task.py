def fileToIntVector(filename):
   row=  open(filename,'r').readline()
   row_strings = row.split()
   row_ints = map(int,row_strings)
   return row_ints

def arrayToDict(part1_input):
    my_dict = {}
    for tower_item in part1_input:
        child_dict = {}
        my_dict[tower_item[0]] = child_dict
        my_dict[tower_item[0]]["weight"] = int(tower_item[1])
        if len(tower_item) >= 4:
                    my_dict[tower_item[0]]["children"] =tower_item[3:]
    return my_dict


def getWeights(part1_input,parent):
    for row in part1_input:
        if row[0] == parent:
            if len(row) >=4:
                child_weights=[]
                single_weights =[]
                for child in row[3:]:
                    single_weight,child_weight, flag = getWeights(part1_input,child)
                    if flag != -1:
                        return single_weight,child_weight,flag
                    else:
                        child_weights.append(child_weight)
                        single_weights.append(single_weight)
                sorted_child_weights = sorted(child_weights)
                if sorted_child_weights[0] != sorted_child_weights[-1]:
                    print sorted_child_weights, "fix the parenthesus suff below, should be correct thoug"
                    if sorted_child_weights[0] == sorted_child_weights[1]:
                        wrong_val = sorted_child_weights[-1]
                        wrong_index =child_weights.index(wrong_val)
                        return single_weights[wrong_index],sum(child_weights),sorted_child_weights[1] - wrong_val
                    else:
                        wrong_val = sorted_child_weights[0]
                        wrong_index =child_weights.index(wrong_val)
                        return single_weights[0],sum(child_weights),sorted_child_weights[1] - wrong_val

                else:
                    return int(row[1]),int(row[1]) + sum(child_weights), -1
                    
                    
            else:
                return int(row[1]),int(row[1]), -1



def part1(part1_input):
    for row in part1_input:
        seen = False
        for row_comp in part1_input:
            if len(row_comp) >=4:
                if row[0] in row_comp[3:]:
                    seen = True 
                    break;
        if not seen:
            return row[0]


            
def file_to_array(filename):
    input_instructions =[]
    discholders =[]
    bottomnames =[]
    with open(filename,'r') as infile: 
       
        for filerow in infile:
            if len(filerow) > 0:
                filerow =filerow.replace(",","")
                row_array = filerow.split()
                input_instructions.append(row_array)
                if len(row_array) >= 4:
                    discholders.append(row_array[:])
                    bottomnames.append(row_array[0])

    return input_instructions, discholders,bottomnames

if __name__ == "__main__":
    part1_input,discholders,bottomnames = file_to_array("input.txt")
    #my_dict = arrayToDict(part1_input)
    part1_result = part1(part1_input)
    single_weight,child_weight,flag= getWeights(part1_input,part1_result)
    #part1_result = part1([0,2,7,0])
    print str(part1_result)
    print str(abs(single_weight+flag))