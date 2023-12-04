export function problem1(text: string) {
  const rows = text.split("\n");

  return rows.reduce((total, row) => {
    const [winning, bet] = row
      .split(":")[1]
      .split("|")
      .map((card) => parseCard(card));
    if (!winning || !bet) {
      throw new Error("Could not parse row");
    }
    const winningSet = new Set(winning);
    const betSet = new Set(bet);
    if (winningSet.size !== winning.length) {
      throw Error("contains duplicates");
    }
    if (betSet.size !== bet.length) {
      throw Error("contains duplicates");
    }

    const correctNumbers = intersection(bet, winningSet);

    return (
      total +
      (correctNumbers.length ? Math.pow(2, correctNumbers.length - 1) : 0)
    );
  }, 0);
}

function intersection<T extends string | number>(a: T[], b: Set<T>): T[] {
  return a.filter((item) => b.has(item));
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
export function problem2(text: string) {
  return 0;
}
