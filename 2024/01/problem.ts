export function problem1(text: string): number{
const numbers = text.split("\n").map(line => (line.match(/[0-9]+/g) ?? []).map(n => Number.parseInt(n)))
const lefts: number[] = []
const rights: number[] = []
for(const [left,right] of numbers){
    lefts.push(left)
    rights.push(right)
}
lefts.sort((a,b) => a-b)
rights.sort((a,b) => (a-b))

return lefts.reduce((sum,v,index) => {
    const diff = Math.abs(rights[index]- v)
    return sum + diff
}, 0 )
}


export function problem2(text: string): number{
    const numbers = text.split("\n").map(line => (line.match(/[0-9]+/g) ?? []).map(n => Number.parseInt(n)))
    const lefts: number[] = []
    const rights: number[] = []
    const rightMap= new Map<number,number>()
    for(const [left,right] of numbers){
        lefts.push(left)
        rights.push(right)
        rightMap.set(right,(rightMap.get(right) ?? 0) +1)
    }
    
    return lefts.reduce((sum,v) => {
        const count =  rightMap.get(v) ?? 0
        return sum + (count * v)
    }, 0 )
    }