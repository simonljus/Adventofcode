def count_passwords():
    minpass = 147981
    maxpass = 691423
    p1 = 0
    p2 = 0
    for password in range(minpass,maxpass):
        p1_increment,p2_increment = isValid(password)
        p1 += p1_increment
        p2 += p2_increment
    print("problem 1",p1)
    print("problem 2",p2)
def isValid(password):
    strpass =str(password)
    p1_inc = 0
    p2_inc = 0
    foundPair = False
    foundLarge = False
    if (strpass != ''.join(sorted(strpass))):
        return 0,0
    for i in range(10):
        groupsize =strpass.count(str(i))
        if(groupsize ==2):
            foundPair = True
        elif(groupsize >2):
            foundLarge = True
    if(foundLarge or foundPair):
        p1_inc = 1
    if(foundPair):
        p2_inc = 1

    return p1_inc,p2_inc
count_passwords()
