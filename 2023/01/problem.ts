export function problem1(text: string) {
  const rows = text.split("\n");
  return rows.reduce((count, row) => {
    const digits = row.match(/\d/g);
    const first = digits?.at(0);
    const last = digits?.at(-1);
    if (!first || !last) {
      throw Error(`could not find digits in ${row}`);
    }
    const combined = Number.parseInt(first + last);
    return combined + count;
  }, 0);
}

export function problem2(text: string) {
  const names = [
    "one",
    "two",
    "three",
    "four",
    "five",
    "six",
    "seven",
    "eight",
    "nine",
  ];
  const rows = text.split("\n");
  const pattern = `(?=(${names.join("|")}|[1-9]))`;
  const regex = new RegExp(pattern, "g");
  return rows.reduce((count, row) => {
    const matches = Array.from(row.matchAll(regex));
    const first = matches.at(0)?.at(-1);
    const last = matches.at(-1)?.at(-1);
    if (!first || !last) {
      throw new Error("Could not find values in row");
    }
    const combined = Number.parseInt(getDigit(first) + getDigit(last));
    return combined + count;
  }, 0);

  function getDigit(word: string) {
    const index = names.indexOf(word);
    if (index === -1) {
      return word;
    }
    return `${index + 1}`;
  }
}
