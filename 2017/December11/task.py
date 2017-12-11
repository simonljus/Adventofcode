#solved with information from https://www.redblobgames.com/grids/hexagons/
def fileToList(filename):
    direction_dict = {"sw":0,"se":0,"s":0,"nw":0,"ne":0,"n":0}
    row=  open(filename,'r').readline()
    direction_list = row.split(",")
    print len(direction_list)
    return direction_list


def shortestPath(direction_list):
    x =0
    y =0
    z=0 
    distance_list = []
    for direction in direction_list:
        if direction == "n":
            x +=1 
            y -=1
        if direction == "s":
            x -=1
            y +=1
        if direction == "nw":
            x+=1
            z -=1
        if direction == "se":
            x-=1
            z +=1
        if direction == "ne":
            y-=1
            z +=1
        if direction == "sw":
            y +=1
            z -=1
        distance_list.append((abs(x) + abs(y) + abs(z))/2)
    return distance_list 
    

def hexatask(direction_list):
    distance_list =shortestPath(direction_list)


    return distance_list[-1], max(distance_list)
    

if __name__ == "__main__":
    dir_list = fileToList("input.txt")
    part1_output, part2_output = hexatask(dir_list)
    print str(part1_output),str(part2_output)
