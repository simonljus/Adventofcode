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

export function problem1(text: string) {
  /**
   * width
   * height
   * find numbers
   * check if number is a part number (col start, col end,row)
   *  check surrounding number (max(col start -1,0), min(col end +1,maxcol), max(row-1,0), min(row+1,maxrow) )
   *      find symbol (not dot or digit)
   *
   */
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
  let total = 0;
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
      const contains = containsSymbol(rows, {
        column: {
          min: Math.max(matchIndex - 1, grid.column.min),
          max: Math.min(matchIndex + numeric.length, grid.column.max),
        },
        row: {
          min: Math.max(rowIndex - 1, grid.column.min),
          max: Math.min(rowIndex + 1, grid.column.max),
        },
      });
      if (contains) {
        total += Number.parseInt(numeric);
      }
    }
  }
  return total;
}
const symbolPattern = new RegExp(/[^\d.]/);
export function containsSymbol(rows: string[], search: Grid) {
  return rows.slice(search.row.min, search.row.max + 1).some((row) => {
    const part = row.slice(search.column.min, search.column.max + 1);
    if (Array.from(symbolPattern.exec(part) ?? []).length) {
      return true;
    }
  });
}

export function problem2(text: string) {
  return 0;
}
