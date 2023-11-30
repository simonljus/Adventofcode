
import * as fs from 'fs'
import * as path from 'path'
export const p1 = (filename= '\test_input.txt') => {
    return readFile(filename).split("\n").reduce((count,str) => 
    {
        const start =  new Set(str.slice(0,str.length/2).split(''))
        const end = new Set(str.slice(str.length/2))
        let copy: string  = ''
        end.forEach(e=> {
            if(copy){
                return
            }
            if(start.has(e)){
                copy = e
            }
        })
        const v = copy.charCodeAt(0)
        if(v >= 'a'.charCodeAt(0) && v <= 'z'.charCodeAt(0)){
            const ascii =  (v - 'a'.charCodeAt(0)) +1
            return ascii +  count
        }
        if(v >= 'A'.charCodeAt(0) && v <= 'Z'.charCodeAt(0)){
            const ascii =  (v - 'A'.charCodeAt(0)) +27
            return ascii + count
            
        }
        return 0
              
        
    },0 )
}
const getChunks = (str: string,n=3)=> {
    const chunksize = str.length/n
    const chunks = []
    for(let i=0; i < str.length; i+chunksize){
        chunks.push( new Set(str.slice(i,i+chunksize).split("")))
    }
    return chunks
}
const getGroups = (item: Array<Set<string>>,chunksize=3) => {
    const groups: Array<Set<string>>[] = []
    const copy = item.slice()
    if(chunksize <=0){
        return groups
    }
    while(copy.length){
        groups.push(copy.splice(0,chunksize))
    }
    return groups
}
export const p2 = (filename= '\test_input.txt') => {
    const sets = readFile(filename).split("\n").map(str=> new Set(str.split("")))
    const groups = getGroups(sets)
    return  groups.reduce((count,groupn) => 
    {
        
        const group = groupn.slice()
        const first = group.splice(0,1)[0]
        const rest = group
        const match =  Array.from(first).find(char => rest.every(sack => sack.has(char)))
        if(!match){
            return 0;
        }
       
        const v = match.charCodeAt(0)
        if(v >= 'a'.charCodeAt(0) && v <= 'z'.charCodeAt(0)){
            const ascii =  (v - 'a'.charCodeAt(0)) +1
            return ascii +  count
        }
        if(v >= 'A'.charCodeAt(0) && v <= 'Z'.charCodeAt(0)){
            const ascii =  (v - 'A'.charCodeAt(0)) +27
            return ascii + count
            
        }
        return 0
              
        
    },0 )
}

const readFile = (filename = './test_input.txt'):string => {
    const p = path.resolve(__dirname, filename)
    return fs.readFileSync(p).toString()
}