export function problem1(text:string): number{
    let count =0
    for(const line of text.split("\n")){
        let safe = true;
        const levels = line.split(" ").map(level => Number.parseInt(level.trim()))
        if(levels.toSorted((a,b) =>a - b).join() !== levels.join() && levels.toSorted((a,b) =>b - a).join() !== levels.join() ){
            continue
        }
        for(const [index,level] of levels.entries()){
           if(index === 0){
            continue
           }
           const diff = Math.abs(levels[index-1] - level)
           if(diff <1 || diff > 3){
                safe = false;
                break
           }
        }
        if(safe){
            count +=1
        }
    }
    
    return count
}