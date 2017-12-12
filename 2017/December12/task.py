def fileToDict(filename):
    link_dict={}
    with open(filename,"r") as fp:
        for row in fp:
            if len(row) >0:
                row =row.replace(",","")
                row_arr =row.split()
                if len(row_arr) >=3:
                    link_dict[row_arr[0]]= row_arr[2:] 
    return link_dict

def findGroup(link_dict,parent,visited,groups):
    to_visit =[]
    group = set()
    if parent not in visited:
        to_visit.append(parent)
    while len(to_visit) >0:
        current_visit = to_visit[0]
        if current_visit not in visited:
            visited.add(current_visit)
            group.add(current_visit)
            to_visit += link_dict[current_visit]
        del to_visit[0]
    if len(group) >0:
        groups.append(group)

def part1(link_dict):
    visited = set()
    groups =[]
    findGroup(link_dict,"0",visited,groups)
    return len(visited)

def part2(link_dict):
    visited = set()
    groups = []

    for parent in link_dict:
        findGroup(link_dict,parent,visited,groups)
    return len(groups)



if __name__ =="__main__":
    part1_input = fileToDict("input.txt")
    part1_output = part1(part1_input)
    part2_output = part2(part1_input)
    print "Part 1: ",part1_output
    print "Part 2: ",part2_output
