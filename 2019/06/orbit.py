def get_orbits(filename="input.txt"):
    orbit_dict = dict()
    with open(filename) as file:
            for line in file:
                a,b= line.rstrip().split(")")
                orbit_dict[b]=a
    return orbit_dict
def get_total_size(orbit_dict,nodes):
    count =0
    for node in nodes:
        chain_size = get_size(node,orbit_dict,0)
        count +=chain_size
    return count

def get_size(root,orbit_dict,count):
    if(root in orbit_dict):
        node = orbit_dict[root]
        node_count =get_size(node,orbit_dict,count)
        return node_count +1
    else:
        return count

def problem_1(filename="input.txt"):
    o_dict =get_orbits(filename)
    elements =set(o_dict.values() + o_dict.keys())
    size =get_total_size(o_dict,elements)
    return size
def get_transfers(filename="input.txt"):
    transfer_dict = dict()
    with open(filename) as file:
            for line in file:
                a,b= line.rstrip().split(")")
                if(a not in transfer_dict):
                    transfer_dict[a]=set()
                if(b not in transfer_dict):
                    transfer_dict[b]=set()
                transfer_dict[a].add(b)
                transfer_dict[b].add(a)
    return transfer_dict
def problem_2(filename="input.txt"):
    from_orbit ="YOU"
    to_orbit ="SAN"
    length_dict =dict()
    transfer_dict =get_transfers(filename)
    visited=set()
    return get_distance(transfer_dict,from_orbit,to_orbit,length_dict,visited) -2
def get_distance(transfer_dict,from_orbit="YOU",to_orbit="SAN",length_dict=dict(),visited=set()):
    visited.add(from_orbit)
    max_distance =len(transfer_dict.keys()) +1
    min_distance =max_distance
    transfers = transfer_dict[from_orbit]
    if(from_orbit == to_orbit):
        min_distance=0
    for transfer in transfers:
        distance = max_distance
        if transfer not in length_dict:
            if transfer not in visited:
                distance = get_distance(transfer_dict,transfer,to_orbit,length_dict,visited)
        else:
            distance =length_dict[transfer]
        min_distance= min(min_distance,distance+1)
    length_dict[from_orbit] =min_distance
    return min_distance


test =problem_1("test.txt")
print("Test problem 1, should be 42:",test)
ans_1= problem_1()
print("Problem 1:",ans_1)
test2= problem_2("test2.txt")
print("test Problem 2, should be 4",test2)
ans_2= problem_2("input.txt")
print("Problem2:",ans_2)

