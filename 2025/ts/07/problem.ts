const BEAM = "|";
const START = "S";
const SPLITTER = "^";
const EMPTY = ".";
const markers = [BEAM, START, SPLITTER, EMPTY] as const;
type Marker = typeof markers[number];
function parse(text: string) {
  const lines = text.split("\n").map((line) => line.split("") as Marker[]);
  return { lines };
}
export function draw(text: string) {
  const { lines } = parse(text);
  const yMax = lines.length - 1;
  let count = 0;
  let count2 = 0;
  for (let y = 0; y <= yMax; y++) {
    const xMax = lines[y].length - 1;
    for (let x = 0; x <= xMax; x++) {
      const marker = lines[y][x];
      if (marker === START) {
        if (lines[y + 1]?.[x] === EMPTY) {
          lines[y + 1][x] = BEAM;
        }
        continue;
      }
      if (marker === BEAM) {
        if (lines[y + 1]?.[x] === EMPTY) {
          lines[y + 1][x] = BEAM;
        } else if (lines[y + 1]?.[x] === SPLITTER) {
          count += 1;
          count2 += 1;
          if (lines[y + 1]?.[x - 1] === EMPTY) {
            lines[y + 1][x - 1] = BEAM;
          }
          if (lines[y + 1]?.[x + 1] === EMPTY) {
            lines[y + 1][x + 1] = BEAM;
          }
        }
      }
    }
  }
  return { count, lines };
}
export function problem1(text: string): number {
  const { count } = draw(text);
  return count;
}
export function problem2(text: string): number {
  const { lines } = draw(text);
  const memory = new Map<string, number>();
  let timelines = 0;
  for (const [y, line] of lines.entries()) {
    for (const [x, mark] of line.entries()) {
      if (mark === START) {
        timelines += 1;
        timelines += bfs({ x, y: y + 1 }, lines, memory);
      }
    }
  }
  return timelines;
}
function bfs(
  pos: { x: number; y: number },
  lines: Marker[][],
  memory: Map<string, number>,
) {
  const yStart = pos.y;
  const xStart = pos.x;
  let count = 0;
  const key = `${xStart}_${yStart}`;
  if (memory.has(key)) {
    return memory.get(key) ?? 0;
  }
  const below = lines[pos.y + 1]?.[pos.x];
  if (below === SPLITTER) {
    count += 1;
    if (lines[pos.y + 1]?.[pos.x + 1] === BEAM) {
      count += bfs({ x: xStart + 1, y: yStart + 1 }, lines, memory);
    }
    if (lines[pos.y + 1]?.[pos.x - 1] === BEAM) {
      count += bfs({ x: xStart - 1, y: yStart + 1 }, lines, memory);
    }
  } else if (below === BEAM) {
    count += bfs({ x: xStart, y: yStart + 1 }, lines, memory);
  }
  memory.set(key, count);
  return count;
}
