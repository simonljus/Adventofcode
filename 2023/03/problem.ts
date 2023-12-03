type Grid = {
  row: {
    min: number;
    max: number;
  };
  column: {
    min: number;
    max: number;
  };
};

const symbolPattern = new RegExp(/[^\d.]/);
const gear = "*";
function getGearOrSymbol(
  rows: string[],
  search: Grid,
  numeric: number,
  gearMap: Map<string, number[]>
) {
  return rows
    .slice(search.row.min, search.row.max + 1)
    .some((row, rowOffset) => {
      const part = row.slice(search.column.min, search.column.max + 1);
      const execArray = symbolPattern.exec(part);
      if (!execArray || !execArray.length) {
        return false;
      }
      for (const symbol of execArray.values()) {
        if (symbol === gear) {
          const key = `${rowOffset + search.row.min}_${
            execArray.index + search.column.min
          }`;
          if (!gearMap.has(key)) {
            gearMap.set(key, []);
          }
          gearMap.get(key)?.push(numeric);
        }
      }
      return true;
    });
}

function solve(text: string) {
  const rows = text.split("\n");
  const maxRowIndex = rows.length - 1;
  const minRowIndex = 0;
  const minColIndex = 0;
  const maxColIndex = (rows.at(0)?.length ?? 0) - 1;
  const grid: Grid = {
    row: {
      min: minRowIndex,
      max: maxRowIndex,
    },
    column: {
      min: minColIndex,
      max: maxColIndex,
    },
  };
  const pattern = new RegExp(/\d+/g);
  let p1Total = 0;
  const gearMap = new Map<string, number[]>();
  for (const [rowIndex, row] of rows.entries()) {
    for (const matchArray of row.matchAll(pattern)) {
      const numeric = matchArray.at(0);
      if (!numeric) {
        console.error("no number found", rowIndex);
        throw new Error("No number found");
      }
      const matchIndex = matchArray.index ?? -1;
      if (matchIndex < 0) {
        console.error(rowIndex);
        throw new Error("Unknown index");
      }

      const hasSymbol = getGearOrSymbol(
        rows,
        {
          column: {
            min: Math.max(matchIndex - 1, grid.column.min),
            max: Math.min(matchIndex + numeric.length, grid.column.max),
          },
          row: {
            min: Math.max(rowIndex - 1, grid.column.min),
            max: Math.min(rowIndex + 1, grid.column.max),
          },
        },
        Number.parseInt(numeric),
        gearMap
      );
      if (hasSymbol) {
        p1Total += Number.parseInt(numeric);
      }
    }
  }
  let p2Total = 0;
  for (const numbers of gearMap.values()) {
    if (numbers.length === 2) {
      p2Total += numbers[0] * numbers[1];
    }
  }
  return { part1: p1Total, part2: p2Total };
}
export function problem1(text: string) {
  return solve(text).part1;
}
export function problem2(text: string) {
  return solve(text).part2;
}
