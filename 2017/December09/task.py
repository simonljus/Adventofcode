
def testPart1(test_cases):
    print "hello world"
    #print test_cases
    for test in test_cases:
        print "testing",test
        test_result = part1(test)
        if test_result ==  test_cases[test]:
            #print test, "passed"
            pass
        else:
            print test, "not passed returned", test_result
def calculateLayers(curly_list,cost):
    
    print "new layer"
    print curly_list
    sum =0
    n = len(curly_list)
    index = 0
    layer_score =cost
    found_start = False
    start_index = n
    children_count =0
    while len(curly_list) > 0:
        if curly_list[index] == "{":
            if not found_start:
                #print "start_layer at index", str(index)
                found_start = True
                start_index =index
                index +=1
            else:
                #print "increase children count "
                children_count +=1
                index +=1

        elif curly_list[index] == "}":
            if children_count >0:
                #print "decrease children count "
                children_count -=1
                index +=1
            else:
                #print "recursive, current layerscore:", str(layer_score)
                #print "statrt index ",  str(start_index), "end index ", str(index)
                
                sum += layer_score
                if index -start_index >1:

                    layer_score +=1
                del curly_list[index]
                del curly_list[start_index]
            
                #print curly_list
                start_index = len(curly_list)
                found_start = False
                index = 0
                children_count =0
        else:
            index +=1

    return sum


        


        
        


def part1(input_string):
    sum =0
    n = len(input_string)
    input_list = list(input_string)
    layer_score =0
    garbage_from = n
    start_group =[]
    end_group =[]  
    comma_sign = []
    start_garbage =[]
    end_garbage=[]
    curlies =[]
    curly_group =""
    curly_weight =0
    curlies =[]
    i =0 
    while (i <n):
        c = input_list[i]
        if i > garbage_from:
            if c == "!":
                input_list[i] = "a"
                i +=1
                if i <n:
                    input_list[i] = "a"
            elif c == ">":
                garbage_from = n
            else: 
                input_list[i] = "a"



        else:    
            if c == "!":
                input_list[i] = "a"
                i +=1
                if i <n:
                    input_list[i] = "a"
            elif c == "<":
                garbage_from = i
            elif c == ">":
                input_list[i] = "a"
            elif c == "{":
                curly_weight +=1
                start_group.append(i)
                curly_group+= "{"
            elif c == "}":
                curly_weight -=1
                end_group.append(i)
                curly_group+="}"
            elif c == ",":
                if curly_weight ==0:
                    curlies.append(list(curly_group))
                    curly_group = ""
                    garbage_from = n
            else:
                pass
        i +=1 
    curlies.append(list(curly_group))
    if len(start_group) == len(end_group):
        #print("Seems right")
        for curly_list in curlies:
            #print curly_list
            #print len(curlies)
            sum += calculateLayers(curly_list[:],1)
    else:
        print("Seems wrong")
        print(stat_group,end_group)
        


    return sum 
def readStringFromFile(filename):
    return open(filename,'r').readline()
def testCases():
    part1_test_input = {"{}": 1, "{{{}}}": 6,"{{},{}}":5,"{{{},{},{{}}}}":16,"{<a>,<a>,<a>,<a>}":1,"{{<ab>},{<ab>},{<ab>},{<ab>}}":9,"{{<!!>},{<!!>},{<!!>},{<!!>}}":9,"{{<a!>},{<a!>},{<a!>},{<ab>}}":3}
    working = {"{}": 1,"{{{}}}": 6}
    partly = {"{{},{}}":5}
    part1_complete_sum = 1 +6+5+16+1+9+9+3
    part1_complete_test_input = {"{},{{{}}},{{},{}},{{{},{},{{}}}},{<a>,<a>,<a>,<a>},{{<ab>},{<ab>},{<ab>},{<ab>}},{{<!!>},{<!!>},{<!!>},{<!!>}},{{<a!>},{<a!>},{<a!>},{<ab>}},":part1_complete_sum}
    #testPart1(partly)
    testPart1(part1_test_input)
    testPart1(part1_complete_test_input)

if __name__ == "__main__":
    testCases()
    input_string = readStringFromFile("input.txt")
    print str(part1(input_string))
