export function problem1(text: string):number{
    const grid = text.split("\n").map(row => row.split(""))
    let count = 0
    const word ="XMAS"
    const firstLetter ="X"
    const rest =word.slice(1)
    for(const [y,row] of grid.entries()){
        for(const [x,letter] of row.entries() ){
            if(letter !== firstLetter){
               continue
            }
            const ydirs = [-1,0,1]
            const xdirs = [-1,0,1]
            for(const dy of ydirs){
                for(const dx of xdirs){
                    if(dx === 0 && dy === 0){
                            continue
                    }
                    if(findWord(grid,rest,{x,y},{dx,dy})) {
                        count+=1
                    }
                }
            }
        }
    }
    return count
}

export function findWord(grid: string[][],word:string,pos: {x: number, y: number}, dir: {dx: number, dy:number}):boolean{
    const xCount = grid[0].length
    const yCount = grid.length
    for(const [index,char] of word.split("").entries()){
        const distance = index +1
        const yPos = (pos.y + dir.dy * distance)
        const xPos = (pos.x + dir.dx * distance)
        if(xPos < 0 || xPos >= xCount){
            return false
        }
        if(yPos < 0 || yPos >= yCount){
            return false
        }
        const foundLetter = grid.at(yPos)?.at(xPos)
        if(!foundLetter){
            throw new Error(`indexing issue!!! ${xPos},${yPos}`)
        }
        if(foundLetter !== char){
            return false
        }
    }
    return true
}

export function problem2(text: string):number{
    const grid = text.split("\n").map(row => row.split(""))
    let count = 0
    for(const [y,row] of grid.entries()){
        for(const [x,letter] of row.entries() ){
            if(letter !== "A"){
               continue
            }
            const topLeft = {x:-1,y:-1}
            const bottomRight = {x:1,y:1}
            const topRight = {y:-1,x: 1}
            const bottomLeft = {y:1,x:-1}
            const pairs = [[topLeft,bottomRight],[topRight,bottomLeft]]
            if(pairs.every(pair => pair.map(coord => grid[coord.y +y]?.[coord.x +x] ?? "").sort().join("") === "MS")){
                count +=1
            }
        }
    }
    return count
}