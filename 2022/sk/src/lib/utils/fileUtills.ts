import * as fs from 'fs'
import * as path from 'path'

export const readFile = (dirname:string,filename= './test_input.txt') => {
    const p = path.resolve(dirname,filename)
    return fs.readFileSync(p).toString()
}