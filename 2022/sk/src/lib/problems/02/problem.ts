import * as fs from 'fs'
import * as path from 'path'

enum RPS {
    ROCK,
    PAPER,
    SCISSOR
}
type Opponent = 'A' | 'B' | 'C'
type Strategy = 'X' | 'Y' | 'Z'

const getOpponentMove = (move: Opponent): RPS => {
    switch (move) {
        case ('A'):
            return RPS.ROCK
        case ('B'):
            return RPS.PAPER
        case ('C'):
            return RPS.SCISSOR
    }

}

const getWinningStrategyMove = (move: Strategy): RPS => {
    switch (move) {
        case ('X'):
            return RPS.ROCK
        case ('Y'):
            return RPS.PAPER
        case ('Z'):
            return RPS.SCISSOR
    }

}

const getResultStrategyMove = (opponent: RPS,move: Strategy): RPS => {
    switch (move) {
        case ('X'):
            return getLosingMove(opponent)
        case ('Y'):
            return opponent
        case ('Z'):
            return getWinningMove(opponent)
    }

}

const countP1Score = (opponent:Opponent,strategy:Strategy) =>{
    const op = getOpponentMove(opponent)
    const st = getWinningStrategyMove(strategy)
    return countScore(op,st)

}
const countP2Score = (opponent:Opponent,strategy:Strategy) =>{
    const op = getOpponentMove(opponent)
    const st = getResultStrategyMove(op,strategy)
    return countScore(op,st)

}
const countScore = (opponent:RPS,strategy:RPS) =>{
    const outcome = getOutcome(opponent,strategy)
    const shapeScore = getShapeScore(strategy)
    return outcome + shapeScore
}

const getOutcome = (opponent: RPS,strategy:RPS) => {
    if(opponent === strategy){
        return 3
    }
    const winning =  getWinningMove(opponent)
    return winning === strategy ? 6: 0
}

const getShapeScore = (rps:RPS) => {
    switch(rps){
        case(RPS.ROCK): 
        return 1
        case(RPS.PAPER): 
        return 2
        case(RPS.SCISSOR): 
        return 3
    }
}

const getWinningMove = (rps: RPS): RPS =>{
    switch(rps){
        case(RPS.ROCK):
            return RPS.PAPER
        case(RPS.PAPER):
            return RPS.SCISSOR
        case(RPS.SCISSOR): 
            return RPS.ROCK
    }
}
const getLosingMove = (rps: RPS): RPS =>{
    switch(rps){
        case(RPS.ROCK):
            return RPS.SCISSOR
        case(RPS.PAPER):
            return RPS.ROCK
        case(RPS.SCISSOR): 
            return RPS.PAPER
    }
}
export const p1 = (filename: string) => {

    const data = readFile(filename)
    const scores = data.split("\n").map(round=> {
        const [opponent,strategy] =  round.split(" ")
        return countP1Score(opponent as Opponent,strategy as Strategy)
   
})
    return  scores.reduce((r,total) => r+total,0)
}

export const p2 = (filename: string) => {

    const data = readFile(filename)
    const scores = data.split("\n").map(round=> {
        const [opponent,strategy] =  round.split(" ")
        return countP2Score(opponent as Opponent,strategy as Strategy,)
   
})
    return  scores.reduce((r,total) => r+total,0)
}

const readFile = (filename = './test_input.txt') => {
    const p = path.resolve(__dirname, filename)
    return fs.readFileSync(p).toString()
}