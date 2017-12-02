
def part1_testCases(test_input):
    pass_count =0
    number_of_tests= len(test_input)
    for test_string in test_input:
        #print( "testing string "+ test_string )
        test_result = sumOfMatchingDistance(test_string,1)
        test_output = test_input[test_string]
        if test_result == test_output:
            #print( " Part1: Test of " + test_string +  " passed" )
            pass_count +=1
        else:
             print( "Part1: Test of " + test_string + " failed, returned " + str(test_result)  )
    print "Part1: %d of %d passed" % (pass_count,number_of_tests)


def part2_testCases(test_input):
    pass_count =0
    number_of_tests= len(test_input)
    for test_string in test_input:
        #print( "testing string "+ test_string )
        test_result = sumOfMatchingDistance(test_string,len(test_string)/2)
        test_output = test_input[test_string]
        if test_result == test_output:
            #print( "Part2: Test of " + test_string +  " passed" )
            pass_count +=1
        else:
             print( "Part2: Test of " + test_string + " failed, returned " + str(test_result)  )
    print "Part2: %d of %d passed" % (pass_count,number_of_tests)


def testCases():
    part1_test_input = {"1122": 3, "1111": 4,"1234":0,"91212129":9,"":0,"55":10}
    part2_test_input = {"1212": 6, "1221": 0,"123425":4,"123123":12,"12131415":4}
    part1_testCases(part1_test_input)
    part2_testCases(part2_test_input)


def sumOfMatchingDistance(input_string,jump):
    sum =0
    n = len(input_string)
    for i,digit in enumerate(input_string) :
        if  digit == input_string[(i + jump)%n]:
            sum += int(digit)
    return sum 


def fileToString(filename):
    textfile = open(filename,"r")
    return textfile.read()


if __name__ == "__main__":
    testCases()
    captcha = fileToString("input.txt")
    sum_part1 = sumOfMatchingDistance(captcha,1)
    print "Part1 result: %d" % (sum_part1)
    sum_part2= sumOfMatchingDistance(captcha,len(captcha)/2)
    print "Part2 result: %d" % (sum_part2)
