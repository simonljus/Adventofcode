export function problem1(text:string){
    const memory: (number | '.')[] = []
    for(const [index,[fileSize,freeSize]] of (text.match(/\d{1,2}/g)?.map(m => m.split("").map(n => Number.parseInt(n))) ?? []).entries()){
        memory.push(...Array(fileSize).fill(index),...Array(freeSize ?? 0).fill('.'))
    }
    const freeIndexes: number[] = []
    for(const [index,m] of memory.entries()){
        if (m === '.'){
            freeIndexes.push(index)
        }
    }
    const lastIndex = memory.length -1
    const spaceIndexes: number[] = []
    for(const [index,m] of memory.slice().reverse().entries()){
        if(spaceIndexes.length === freeIndexes.length){
            break
        }
        if(m === '.'){
            continue;
        } 
        spaceIndexes.push(lastIndex -index)
    }
    console.log(freeIndexes)
    console.log(spaceIndexes)
    const swapped = memory.slice()
    for(const [index,freeIndex] of freeIndexes.entries()){
        const spaceIndex = spaceIndexes[index]
        if(freeIndex > spaceIndex){
            break
        }
        swapped[freeIndex] = swapped[spaceIndex]
        swapped[spaceIndex]='.'
    }
    return swapped.reduce((acc: number,curr,index) => {
        if(curr === '.'){
            return acc
        }
        return acc  + (curr * index)
    },0)
}

export function problem2(text:string){
    return 0
}