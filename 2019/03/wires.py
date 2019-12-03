

def getGrid():
    wires = []
    with open ("input.txt") as file:
        for wire in file:
            instructions=(wire.rstrip().split(","))
            wires.append(draw(instructions))
    wireA=wires[0]
    wireB=wires[1]
    common = wireA.keys() & wireB.keys()
    distances= []
    lengths= []
    for x,y in common:
        print(x,y)
        distances.append(abs(x)+abs(y))
        lengths.append(wireA[(x,y)] + wireB[(x,y)])
    problem1=min(distances)
    problem2=min(lengths)
    print("problem 1",problem1)
    print("problem 2",problem2)

def draw(instructions):
    x=0
    y=0
    length =0
    positions = dict()
    for instruction in instructions:
        direction = instruction[0]
        distance = int(instruction[1:])
        for i in range(distance):
            length +=1
            if(direction == "U"):
                y+=1
            elif(direction == "D"):
                y-=1
            elif(direction == "R"):
                x+=1
            elif(direction =="L"):
                x-=1
            if((x,y) not in positions):
                #will not add since shorter
                positions[(x,y)]=length
    return positions
getGrid()