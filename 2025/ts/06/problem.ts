type Operation = "*" | "+";
function parse(text: string) {
  const lines = text.split("\n").map((line) => line.trim().split(/\s+/));
  const operations =
    lines.pop()?.map((operation) => operation.trim() as Operation) ?? [];
  const numbersLines = lines.map((line) =>
    line.map((number) => Number(number.trim()))
  );
  return { numbersLines, operations };
}
export function problem1(text: string): number {
  const { numbersLines, operations } = parse(text);
  let count = 0;
  for (const [index, operation] of operations.entries()) {
    const terms = numbersLines.map((numberLine) => numberLine[index]);
    count += terms.reduce(
      (acc, term) => operation === "*" ? acc * term : acc + term,
      operation === "*" ? 1 : 0,
    );
  }

  return count;
}

export function problem2(text: string) {
  const lines = text.split("\n").map((line) =>
    line.split("").reverse().join("")
  );
  const operationMatches = lines.pop()?.matchAll(/(?<ADD>\+)|(?<MUL>\*)/gm);
  const indexes = Array.from(operationMatches ?? []).map((match) => {
    return { endsAt: match.index, operation: match[0] as Operation };
  });
  let sum = 0;
  for (const [i, { endsAt, operation }] of indexes.entries()) {
    const terms = lines.map((line) => {
      if (i === 0) {
        return line.slice(0, endsAt + 1);
      }
      return line.slice(indexes[i - 1].endsAt + 1, endsAt + 1);
    });
    const maxLength = Math.max(...terms.map((n) => n.length));
    const numbers = new Array(maxLength).keys().map((v) =>
      terms.map((term) => term[v] ?? "").join("").trim()
    ).filter(Boolean).map((v) => Number(v));
    sum += numbers.reduce(
      (acc, term) => operation === "*" ? acc * term : acc + term,
      operation === "*" ? 1 : 0,
    );
  }
  return sum;
}
