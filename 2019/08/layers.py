def get_layers(filename="input.txt",height=6,width=25):
    layers= []
    pixels=""
    with open(filename) as file:
            for line in file:
                pixels +=line.rstrip()
    layers =list(chunks(pixels, height*width))
    return layers
def get_counts(layers):
    dictionaries=[]
    min_count =999999
    min_dict = dict()
    for layer in layers:
        d = dict()
        for i in range(10):
            i_count =layer.count(str(i))
            d[i]=layer.count(str(i))
        dictionaries.append(d)
    for counts in dictionaries:
        if counts[0] < min_count:
            min_count=counts[0]
            min_dict =counts
    #print(dictionaries)
    return min_dict[1] * min_dict[2]

def chunks(lst, n):
    for i in range(0, len(lst), n):
        yield lst[i:i + n]

def decode_image(layers,width=25,height=6):
    image=[]
    transparent ="2"
    for i in range(width*height):
        pixel =transparent
        layer_str=""
        for layer_i,layer in enumerate(layers):
            layer_pixel = layer[i]
            layer_str+=layer_pixel
            #print("pixel_index",i,"layer",layer_i,"pixel value",layer_pixel)
            if(pixel == transparent):
                pixel = layer_pixel
        image.append(pixel)
    message = list(chunks(image,width))
    image_str =""
    for line in message:
        m_binary ="".join(line).replace("0"," ")
        print(m_binary)
    return message


def problem_1(filename="input.txt",height=6,width=25):
    layers =get_layers(filename,width,height)
    return get_counts(layers)
def problem_2(filename="input.txt",height=6,width=25):
    layers =get_layers(filename,width,height)
    return decode_image(layers,width,height)
ans_1 =problem_1()
print("Problem 1",ans_1)
print("Problem 2:")
ans_2=problem_2()

