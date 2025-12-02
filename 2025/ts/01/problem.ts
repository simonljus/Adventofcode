function parse(text: string) {
  const pattern = /(?<L>((?<=L)\d+))|(?<R>((?<=R)\d+))/gm;
  const moves: Array<{ direction: "L" | "R"; steps: number }> = [];
  for (const m of text.matchAll(pattern)) {
    const move = m.groups as Record<"L" | "R", `${number}`> | undefined;
    if (!move) {
      continue;
    }
    const direction = move.L ? "L" as const : "R" as const;
    const steps = Number(move.L ?? move.R);
    moves.push({ direction, steps });
  }
  return moves;
}
export function problem1(text: string): number {
  const moves = parse(text);
  const MAX_INDEX = 99;
  const VALUE_COUNT = MAX_INDEX + 1;
  const START = 50;
  const FIND_POS = 0;
  let pos = START;
  let count = 0;
  for (const move of moves) {
    const direction = move.direction === "L" ? -1 : 1;
    pos = (pos + (direction * (move.steps % VALUE_COUNT))) % VALUE_COUNT;

    if (pos < 0) {
      pos += VALUE_COUNT;
    }
    if (pos === FIND_POS) {
      count += 1;
    }
  }
  return count;
}
export function problem2(text: string): number {
  const moves = parse(text);
  const MAX_INDEX = 99;
  const VALUE_COUNT = MAX_INDEX + 1;
  const START = 50;
  let pos = START;
  let count = 0;
  for (const move of moves) {
    const prev = pos;
    const direction = move.direction === "L" ? -1 : 1;
    const stepsToZero = direction === -1 ? prev : VALUE_COUNT - prev;

    const offset = move.steps >= stepsToZero ? 1 : 0;
    const remainingSteps = offset ? move.steps - stepsToZero : move.steps;
    const lapses = Math.floor(remainingSteps / VALUE_COUNT);
    const initialLapse = stepsToZero === 0 ? 0 : offset;
    count += lapses + initialLapse;
    const startPos = offset ? 0 : pos;
    const adjustSteps = remainingSteps % VALUE_COUNT;
    if (startPos === 0 && adjustSteps > 0) {
      pos = direction === -1 ? VALUE_COUNT - adjustSteps : adjustSteps;
    } else {
      pos = direction === -1 ? startPos - adjustSteps : startPos + adjustSteps;
    }
  }
  return count;
}
