import * as fs from 'fs'
import * as path from 'path'

export const add = (a: number, b: number) => {
    readFile()
    return a + b;
}
export const p1 = (filename='./test_input.txt') => {
    return sumHighestCalories({filename,top:1})
}
export const p2 = (filename='./test_input.txt') => {
    return sumHighestCalories({filename,top:3})
}

const sumHighestCalories = ({filename,top}: {filename:string,top: number}) => {
    const elves: Array<Array<number>> = readFile(filename).split("\n\n").map(elf => elf.split("\n").map(calorie=> parseInt(calorie)))
    const calories = elves.map(elf => countCalories(elf))
    return calories.sort((a,b) => b-a).slice(0,top).reduce((calorie,count)=> calorie+count,0)
}
const countCalories = ( calories: Array<number>) => {
    return calories.reduce((c,sum) =>  c+sum,0)
}

const readFile = (filename= './test_input.txt') => {
    const p = path.resolve(__dirname,filename)
    return fs.readFileSync(p).toString()
}
