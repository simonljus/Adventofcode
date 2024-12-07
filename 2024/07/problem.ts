export function problem1(text: string){
    return solve(text,['multiply','add'])
}

export function solve(text: string,operators:('multiply'| 'add' | 'concat')[]){
    let count =0
    for (const [lineIndex,line] of text.split("\n").entries()) {
        const [testValue,...numbers] = line.match(/\d+/g)?.map(n => Number.parseInt(n)) ?? []
        const orders: ('multiply' |'add' | 'concat')[][] = []
        const combinations = Math.pow(operators.length,numbers.length-1)
        for(let i = 0; i < combinations; i++){
           orders.push(toPaddedOperations(i,combinations-1,operators))
        }
        for(const order of orders){
            const calculatedValue = numbers.slice(1).reduce((acc,number,operatorIndex)=>{
                return doOperation(acc,number,order[operatorIndex])
            },numbers[0])
            if(calculatedValue === testValue){
                count = count +testValue
                break;
            }
        }
    }
    return count
}
function toPaddedOperations<T extends string>(n:number, max:number, operators: T[]): T[]{
    const bin= (n >>> 0).toString(operators.length);
    const maxBinLength = (max >>> 0).toString(operators.length).length;
    return bin.padStart(maxBinLength,"0").split("").map(v => operators[Number.parseInt(v)])
}
function doOperation(a: number,b: number,operation: 'multiply' | 'add' | 'concat'){
    if(operation === 'multiply'){
        return a * b
    }
    if(operation === 'add'){
        return a+b;
    }
    if(operation === 'concat'){
        return Number.parseInt(`${a}${b}`)
    }
    throw new Error("Unknown operation")
}

export function problem2(text:string){
    return solve(text,['multiply','add','concat'])
}