import * as fs from 'fs'
import * as path from 'path'

interface Interval {
    start:number,
    end:number
}
export const hasFullOverlap=(a:Interval,b:Interval) => {
    return hasFullOverlapSingle(a,b) || hasFullOverlapSingle(b,a)
}
const hasFullOverlapSingle=(a:Interval,b:Interval) => {
    return (a.start <= b.start && a.end >=b.end)
}
export const hasSomeOverlap=(a:Interval,b:Interval) => {
    return hasSomeOverlapSingle(a,b) || hasSomeOverlapSingle(b,a)
}

const hasSomeOverlapSingle=(a:Interval,b:Interval) => {
    return (a.start <= b.end && a.end >=b.start)
}
export const p1 = (filename ='./test_input.txt') => {
   return readFile(filename).split("\n").reduce((c,line)=> {
    const [intervalA,intervalB] = line.split(",").map(worker => {
        const [start,end] = worker.split("-").map(v=> parseInt(v))
        return {start,end}
    })
    const overlaps = hasFullOverlap(intervalA,intervalB)
    return overlaps ? 1 + c : c
},0)
}

export const p2 = (filename ='./test_input.txt') => {
    return readFile(filename).split("\n").reduce((c,line)=> {
     const [intervalA,intervalB] = line.split(",").map(worker => {
         const [start,end] = worker.split("-").map(v=> parseInt(v))
         return {start,end}
     })
     const overlaps = hasSomeOverlap(intervalA,intervalB)
     return overlaps ? 1 + c : c
 },0)
 }

const readFile = (filename = './test_input.txt'):string => {
    const p = path.resolve(__dirname, filename)
    return fs.readFileSync(p).toString()
}