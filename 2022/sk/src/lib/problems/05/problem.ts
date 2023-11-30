import * as fs from 'fs'
import * as path from 'path'


type Instruction = {
    moveCount: number,
    fromIndex:number
    toIndex:number 
}
const getCrates = (str:string) => {
    const lines = str.split("\n")
    const crates : string[][] = []
    lines.forEach(line=> {
        for(let i=0; i < line.length; i+=4){
            const column = i/4
            if(!crates[column]){
                crates[column] = []
            }
            const l = line.slice(i,i+3).replace(/\s+/g,'').slice(1,2)
           
            if(l){
                crates[column].unshift(l)
            }
            
        }
        return line
    })
    return crates
}
const getMoves  = (str:string): Array<Instruction> => {
    return str.split("\n").map(line=> {
        const words = line.split(" ")
        return {
            moveCount: parseInt(words[1]),
            fromIndex: parseInt(words[3]) -1,
            toIndex: parseInt(words[5]) -1
        }
    })
}
const performMoves = (crates: string[][], moves: Instruction[],{keepOrder = false} : {keepOrder?:boolean} = {} ) => {
    return moves.reduce((state,move)=>{
       
        const {moveCount,fromIndex,toIndex} = move
        const fromArr = state[fromIndex].slice()
        const toArr = state[toIndex].slice()
        if(keepOrder){
            const l = fromArr.length
            toArr.push(...fromArr.splice(l-moveCount,l))
        }
        else {
            for(let i =0; i < moveCount; i++){
                const el = fromArr.pop()
                if(el){
                    toArr.push(el)
                }
               
            }
        }
        
        const copy = state.slice()
        copy[fromIndex]= fromArr
        copy[toIndex] = toArr
        return copy
    },crates)
}
const getTopCrates = (crates: string[][]) => {
    return crates.map(crate => crate[crate.length-1]|| '')
}
export const p1  = (filename = './test_input.txt') => {
    const {moves,crates} = parseInput(filename)
    const  crateState = performMoves(crates,moves)
    return getTopCrates(crateState).join("")
}
export const p2  = (filename = './test_input.txt') => {
    const {moves,crates} = parseInput(filename)
    const  crateState = performMoves(crates,moves,{keepOrder:true})
    return getTopCrates(crateState).join("")
}
export const parseInput  = (filename = './test_input.txt') => {
    const lines = readFile(filename)
    const sep = lines.split("\n\n")
    
    const [cratelines,movelines] = sep
    const crates = getCrates(cratelines)
   
    const moves = getMoves(movelines)
    return {moves,crates}
}

const readFile = (filename = './test_input.txt'):string => {
    const p = path.resolve(__dirname, filename)
    return fs.readFileSync(p).toString()
}