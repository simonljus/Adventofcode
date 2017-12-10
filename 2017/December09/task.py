
def testPart1(test_cases):
    #print test_cases
    for test in test_cases:
        print "testing",test
        test_result = part1(test)
        if test_result ==  test_cases[test]:
            #print test, "passed"
            pass
        else:
            print test, "not passed returned", test_result


def findSiblings(sibling_brackets,curly_list,index,layer_cost):
    n = len(curly_list)
    weight =0
    start_index = n
    end_index = n
    
    while (index < n):
        if curly_list[index] == "{":
            weight +=1
            if weight ==1:
                start_index = index
                sibling_brackets.append(start_index)
        elif curly_list[index] =="}":
            weight -=1
            if weight ==0:
                end_index =index
                sibling_brackets.append(end_index)
        index +=1
    if len(sibling_brackets) % 2 ==0:
        i = len(sibling_brackets) -1
        while i >= 0:
            del curly_list[sibling_brackets[i]]
            i-=1
        return layer_cost * (len(sibling_brackets)/2)
    else: 
        print "Something went wrong"
    return 0


def findPair(curly_list,index):
    weight =1
    n = len(curly_list)
    while (index < n):
        if curly_list[index] == "{":
            weight +=1
        elif curly_list[index] =="}":
            weight -=1
            if weight ==0:
                return index
        index +=1


def calculateSiblings(curly_list,cost):
    index =0
    msum =0
    while len(curly_list) > 0:
        #print curly_list
        if curly_list[index] == "{":
            end_index =findPair(curly_list,index +1)
            #print "end done"
            sibling_arr= [index,end_index]
            msum += findSiblings(sibling_arr,curly_list,end_index +1,cost)
            #print "sibling done"
            cost +=1
            index =0
    return msum


def part1(input_string):
    msum =0
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
    garbage_count =0
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
                garbage_count +=1



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
            msum += calculateSiblings(curly_list[:],1)
    else:
        print("Seems wrong")
        print(stat_group,end_group)

    return msum,garbage_count
def readStringFromFile(filename):
    return open(filename,'r').readline()
def testCases():
    part1_test_input = {"{}": 1, "{{{}}}": 6,"{{},{}}":5,"{{{},{},{{}}}}":16,"{<a>,<a>,<a>,<a>}":1,"{{<ab>},{<ab>},{<ab>},{<ab>}}":9,"{{<!!>},{<!!>},{<!!>},{<!!>}}":9,"{{<a!>},{<a!>},{<a!>},{<ab>}}":3, "{ { { {} } },{} }":12,"{ { { {} {} } },{} }":16  }
    working = {"{}": 1,"{{{}}}": 6}
    partly = {"{{},{}}":5}
    part1_complete_sum = 1 +6+5+16+1+9+9+3
    part1_complete_test_input = {"{},{{{}}},{{},{}},{{{},{},{{}}}},{<a>,<a>,<a>,<a>},{{<ab>},{<ab>},{<ab>},{<ab>}},{{<!!>},{<!!>},{<!!>},{<!!>}},{{<a!>},{<a!>},{<a!>},{<ab>}},":part1_complete_sum}
    #testPart1(partly)
    testPart1(part1_test_input)
    testPart1(part1_complete_test_input)

if __name__ == "__main__":
    #testCases()
    input_string = readStringFromFile("input.txt")
    msum, garbage_count = part1(input_string)
    print str(msum), str(garbage_count)
