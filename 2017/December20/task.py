
def fileToPoints(filename):
    points =[]
    removables = "><avp=\n "
    with open(filename,"r") as fp:
        for filerow in fp:
            for c in removables:
                filerow = filerow.replace(c,"")
            file_list = filerow.split(",")
            print file_list
            if len(file_list) >=9:
                coords = map(int,file_list[0:3])
                velocity = map(int,file_list[3:6])
                acceleration = map(int,file_list[6:9])
                point = [coords,velocity,acceleration]
                points.append(point)
        return points
    
def movePoints(points,runs,remove_collisions):
    latest_closest_index = -1
    latest_closest = 2000000000
    for run in xrange(runs):
        closest =2000000000
        closest_index = -1
        for i in xrange(len(points)):
            for val_i,acceleration_val in enumerate(points[i][2]):
                points[i][1][val_i] += acceleration_val
            for val_i,velocity_val in enumerate(points[i][1]):
                points[i][0][val_i] += velocity_val
            m_dist =0
            for val_i,pos_val in enumerate(points[i][0]):
                m_dist += abs(pos_val)
            if m_dist < closest:
                closest = m_dist 
                closest_index = i
                
        print "closest index",closest_index, "left",len(points)
        latest_closest_index = closest_index
        latest_closest = closest
        if remove_collisions:
            new_points = []
            ok_points = []
            for i,el_i in enumerate(points):
                i_coords = el_i[0]
                collision = False
                for j,el_j in enumerate(points):
                    j_coords = el_j[0]
                    if i_coords == j_coords and i !=j:
                        collision = True
                        break
                if not collision and i not in ok_points:
                    new_points.append(el_i)
                    ok_points.append(i)
            points = new_points[:]



                    

                

        
    return latest_closest,latest_closest_index
            
if "__main__" == __name__:
    #note: will converge earlier than 10000 steps
    points = fileToPoints("input.txt")
    #print "result",movePoints(points[:],10000,False)
   
    print "result",movePoints(points[:],10000,True)
    
