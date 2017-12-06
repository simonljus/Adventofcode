def readPassPhrases(filename):
    input_phrases =[]
    counter =0
    with open(filename,'r') as infile: 
       
        for filerow in infile:
            row_strings = filerow.split()
            row_set = set(row_strings)
            if len(row_set) == len(row_strings):
                counter +=1
        return counter 

def passPhrasesAnagram(filename):
    input_phrases =[]
    counter =0
    with open(filename,'r') as infile: 
       
        for filerow in infile:
            row_strings = filerow.split()
            anagram =isAnagramInside(row_strings)
            if  not anagram:
                counter +=1
        return counter 
def isAnagramInside(row_strings):
    for i,el in enumerate(row_strings):
        for j,comp_el  in enumerate(row_strings):
            if  len(el) == len(comp_el) and i != j:
                for c in el:
                    if c in comp_el:
                        comp_el =comp_el.replace(c,"",1)
                        print c, comp_el
                if len(comp_el) == 0:
                    return True
    return False
def reverse_me(string):
    revstring =""
    for i in range(len(string)):
        revstring += string[-(i+1)]
    return revstring
        

        

if __name__ == "__main__":
    print "Result: "  + str(readPassPhrases("input.txt"))
    print "Result: "  + str(passPhrasesAnagram("input.txt"))

           