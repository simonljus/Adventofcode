def fileToIntVector(filename):
   row=  open(filename,'r').readline()
   row_strings = row.split()
   row_ints = map(int,row_strings)
   return row_ints

def stateSeenBefore(bank,states):
    return bank in states
def findLoopSize(bank,states):
    for i,state in enumerate(states):
        if bank == state:
            return len(states) -i
def part1(bank):
    states =[]
    current_bank =0
    distrubute =0
    n = len(bank)
    while( not stateSeenBefore(bank,states)):
        states.append(bank[:])
        print len(states)
        distribute = max(bank)
        current_bank = bank.index(distribute)
        bank[current_bank] =0
        current_bank = (current_bank +1)%n
       
        while distribute >0: 
            bank[current_bank] += 1
            current_bank = (current_bank +1)%n
            distribute -=1
    loopsize = findLoopSize(bank,states)
    return len(states),loopsize

        





if __name__ == "__main__":
    part1_input = fileToIntVector("input.txt")
    print part1_input

    part1_result,loop_size = part1(part1_input)
    #part1_result = part1([0,2,7,0])
    print str(part1_result), str(loop_size)
