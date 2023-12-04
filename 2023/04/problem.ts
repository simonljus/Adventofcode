function intersection<T extends string | number>(a: T[], b: Set<T>): T[] {
  return a.filter((item) => b.has(item));
}

function parseRow(row: string): [number[], number[]] {
  const [winning, bet] = row
    .split(":")[1]
    .split("|")
    .map((card) => parseCard(card));

  const winningSet = new Set(winning);
  const betSet = new Set(bet);
  if (winningSet.size !== winning.length) {
    throw Error("contains duplicates");
  }
  if (betSet.size !== bet.length) {
    throw Error("contains duplicates");
  }
  return [winning, bet];
}

const numberPattern = new RegExp(/\d+/g);
function parseCard(card: string) {
  const matches = card.matchAll(numberPattern);
  if (!matches) {
    throw new Error("Could not parse card");
  }
  return Array.from(matches).flatMap((v) =>
    Array.from(v.values()).map((v) => Number.parseInt(v))
  );
}

export function problem1(text: string) {
  const rows = text.split("\n");

  return rows.reduce((total, row) => {
    const [winning, bet] = parseRow(row);
    const winningSet = new Set(winning);
    const correctNumbers = intersection(bet, winningSet);
    return (
      total +
      (correctNumbers.length ? Math.pow(2, correctNumbers.length - 1) : 0)
    );
  }, 0);
}

export function problem2(text: string) {
  const rows = text.split("\n");
  const cardMap = new Map<number, number>();
  const solution = rows.reduce((total, row, rowIndex) => {
    const [winning, bet] = parseRow(row);
    const copies = (cardMap.get(rowIndex) ?? 0) + 1;
    const winningSet = new Set(winning);
    const correctNumbers = intersection(bet, winningSet);
    const score = correctNumbers.length;
    const maxIndex = score + rowIndex;
    if (score) {
      for (let i = rowIndex + 1; i <= maxIndex; i++) {
        const prevCount = cardMap.get(i) ?? 0;
        cardMap.set(i, prevCount + copies);
      }
    }
    return total + copies;
  }, 0);
  return solution;
}
