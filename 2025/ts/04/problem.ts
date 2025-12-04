type Mark = "." | "@" | "x";
type Grid = Mark[][];
const PAPER = "@";
const REMOVED = "x";
const ADJACENT_LIMIT = 4;

function parse(text: string) {
  return text.split("\n").map((t) => {
    const row = t.split("") as Mark[];
    return row;
  });
}

function countAdjacent(
  pos: { x: number; y: number },
  rows: Grid,
  mark: Mark,
): number {
  const dirs = [-1, 0, 1];
  let count = 0;
  for (const yDir of dirs) {
    for (const xDir of dirs) {
      if (xDir === 0 && yDir === 0) {
        continue;
      }
      const y = pos.y + yDir;
      const x = pos.x + xDir;
      const currentMark = rows[y]?.[x];
      if (currentMark === mark) {
        count += 1;
      }
    }
  }
  return count;
}

export function problem1(text: string): number {
  const rows = parse(text);
  let count = 0;
  for (const [y, row] of rows.entries()) {
    for (const [x, point] of row.entries()) {
      if (point !== PAPER) {
        continue;
      }
      const adjacent = countAdjacent({ x, y }, rows, PAPER);
      if (adjacent < ADJACENT_LIMIT) {
        count += 1;
      }
    }
  }
  return count;
}

function getPositions(rows: Grid, mark: Mark) {
  const positions: { x: number; y: number }[] = [];
  for (const [y, row] of rows.entries()) {
    for (const [x, point] of row.entries()) {
      if (point === mark) {
        positions.push({ x, y });
      }
    }
  }
  return positions;
}

function moveMarkerPositions(
  rows: Grid,
  positions: { x: number; y: number }[],
  mark: Mark,
) {
  return positions.filter((position) => {
    const adjacent = countAdjacent(position, rows, mark);
    if (adjacent < ADJACENT_LIMIT) {
      rows[position.y][position.x] = REMOVED;
      return false;
    }
    return true;
  });
}

export function problem2(text: string): number {
  const rows = parse(text);
  let positions = getPositions(rows, PAPER);
  const startCount = positions.length;
  while (true) {
    const oldLength = positions.length;
    positions = moveMarkerPositions(rows, positions, PAPER);
    if (positions.length === oldLength) {
      break;
    }
  }

  return startCount - positions.length;
}
