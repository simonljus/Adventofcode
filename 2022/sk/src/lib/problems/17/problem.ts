import * as fs from 'fs'
import * as path from 'path'
import { identity, svg_element } from 'svelte/internal'

const parseRocks = () => {
    return readFile("./rocks.txt").split("\n\n").map(rock=> rock.split("\n").map(segment => segment.trim().split("")))
}

enum Move {
    Right = '>',
    Left = '<'
}
interface Position {
    col:number,
    row:number
}
type Tower = string[][]
type Rock = string[][]
const parseMoves= (lines: string): Array<Move> =>{
    return lines.split("") as Array<Move>
}

const createSpace = (n: number) => {
    const row: string[]  = Array(7).fill(".")
    const rows:  string[][] = []
    for (let i = 0; i<n; i++){
        rows.push(row.slice())
    }
    return rows
}

const willCollide =({pos,tower,rock}: {pos:Position,tower:Tower,rock:Rock}) => {
    const copy = tower.map(r=> r.slice())
    placeRock({position:pos,tower:copy,rock, ignoreError: true})
    //console.log("________START_____")
    //console.log(stringifyTower(copy))
    //console.log("________END_____")
    const rockBottom =rock.length -1
    for(let r = 0; r < rock.length; r++){
        const rockSegment = rock[rockBottom -r]
        const towerSegment = tower[pos.row+rockBottom -r]
        // console.log({rockSegment,towerSegment})
        for(let c =0; c < rockSegment.length; c++){
            if(rockSegment[c] === "."){
                continue;
            }
            if(towerSegment[c+ pos.col] !== "."){
                return true
            }
        }
    }
    return false;
}

const placeRock = ({rock,position,tower, ignoreError}: {rock:Rock,position:Position,tower:Tower, ignoreError?:boolean}) => {
    for(let r= 0; r < rock.length; r++){
        const towerSegment = tower[r + position.row]
        const rockSegment = rock[r]
        for(let c = 0; c < rockSegment.length; c++){
            const col = c+position.col
            if(rockSegment[c] === '.'){
                continue;
            }
            if(towerSegment[col] !== "."){
                if(ignoreError){
                    towerSegment[col] ="x"
                    continue
                }
                throw new Error(`Expected emptyness found something else ${towerSegment[col]}`)
                
            }
            towerSegment[col] ="#"
        }
    }
}

const stringifyTower = (tower:Tower) => {
    return tower.map(t=> t.join("")).join("\n")
}


const play = ({rock,tower, moves, moveCount}: {rock: string[][], tower: string[][], moves: Move[], moveCount:number }) => {
    const rockWidth = rock[0].length
    const rockHeight = rock.length
    const pos = {
        col: 2,
        row: 0
    }
    let currentMoveCount = moveCount
    let moveSize = moves.length
    const towerHeight = tower.length;
    const maxCol = tower[0].length -1;
    const maxLevel =  towerHeight-rockHeight
    console.log("NEW ROCK_________")
    for(let level =0; level <= maxLevel; level++){
        const rockBottom = level + rockHeight -1
        const checkCollision = tower[rockBottom].some(c=> c !== '.')
        //console.log({level,maxLevel,checkCollision})
        const move = moves[currentMoveCount % moveSize]
        if(move === Move.Left){
           //console.log("move left")
            if(pos.col === 0){
            }
            else if(checkCollision && willCollide({pos: {row:level,col:pos.col -1},tower,rock})){
                //console.log("willcollide")
            }
            else{
                pos.col = pos.col -1
            }
           
        }
        else if(move === Move.Right){
            //console.log("move right",rockWidth,rockHeight)
            if((pos.col + rockWidth) > maxCol ){
            }
            else if(checkCollision && willCollide({pos: {row:level,col:pos.col +1},tower,rock})){
                //console.log("willcollide right")
            }
            else {
                //console.log("MOVING")
                pos.col = pos.col +1
            }
           
        }
        currentMoveCount +=1
        if(level === maxLevel){
            //console.log("reached maxlevel")
            return {col: pos.col,row: level,currentMoveCount}
            
        }
        if(willCollide({pos: {row:level +1,col:pos.col},tower,rock})){
            //console.log("will collide if moving down")
            return {col: pos.col,row: level,currentMoveCount}
        }
    }
    //console.log("outside loop")
    return {col:pos.col,row:maxLevel,currentMoveCount}
}


export const p1 = (filename: string) => {
    const moves = parseMoves(readFile(filename))
    const rocks = parseRocks()
    const nRocks = rocks.length
    let moveCount =0;
    let tower: string[][] = []
    for (let i = 0; i <2022; i++){
        const rock = rocks[i %  nRocks]
        //console.log({rock})
        const space = createSpace(rock.length +3)
        tower = space.map(s=>s.slice()).concat(tower)
        const position = play({rock,tower,moves,moveCount});
        moveCount = position.currentMoveCount
        //console.log({position})
        placeRock({rock,position,tower})
        tower  = tower.filter(r=> r.some(c=> c !== '.'))
    }
    console.log(tower.map(t=> t.join("")).join("\n"))
    readFile(filename)


    return tower.length
}

export const p2 = (filename: string) => {
    const moves = parseMoves(readFile(filename))
    const rocks = parseRocks()
    const nRocks = rocks.length
    let moveCount =0;
    let tower: string[][] = []
    const occurrences = new Map<number,number[]>()
    const inserts = new Map<string,number[]>()
    let maxRow = 0
    for (let i = 0; i <500; i++){
        const rockIndex= i %  nRocks
        const rock = rocks[rockIndex]
        const space = createSpace(rock.length +3)
        tower.unshift(...space.map(s=>s.slice()))
        
        const position = play({rock,tower: tower.slice(0,20),moves,moveCount});
        moveCount = position.currentMoveCount
        if(position.row > maxRow){
            maxRow = position.row
        }
        const key = `${rockIndex}_${position.row}_${position.col}`
        if(!inserts.has(key)){
            inserts.set(key,[])
        }
        inserts.get(key)?.push(i)
        //console.log({position})
        placeRock({rock,position,tower})
        tower  = tower.filter(r=> r.some(c=> c !== '.'))
        const states = tower.map(r=> parseInt(r.map(c=> c === '.' ? "0": "1").join(""),2))
        console.log({i,states})
    }
    //console.log(tower.map(t=> t.join("")).join("\n"))
    const states = tower.map(r=> parseInt(r.map(c=> c === '.' ? "0": "1").join(""),2))
    states.slice().reverse().forEach((state,stateIndex) => {
        if(!occurrences.has(state)){
            occurrences.set(state,[])
        }
        occurrences.get(state)?.push(stateIndex)
    })
    console.log(states)
    console.log(occurrences)
    console.log(maxRow)
    console.log({inserts})



    return tower.length
}
const readFile = (filename = './test_input.txt'):string => {
    const p = path.resolve(__dirname, filename)
    return fs.readFileSync(p).toString()
}