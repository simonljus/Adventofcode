export function problem1(text: string) {
  return sumMultiplications(text);
}
export function sumMultiplications(text: string) {
  const multiplications = text.match(/mul\(\d+,\d+\)/gm) ?? [];
  return multiplications.reduce((sum, multiplication) => {
    const factors = (multiplication.match(/\d+/g) ?? []).map((n) =>
      Number.parseInt(n)
    );
    const [factorA, factorB] = factors;
    const product = factorA * factorB;
    return product + sum;
  }, 0);
}
export function problem2(text: string) {
  const matches = Array.from(
    `do()${text}do()`.matchAll(/(?<do>do\(\))|(?<dont>don't\(\))/gm),
  );
  const { sum } = matches.reduce((obj, match) => {
    const subtext = text.slice(obj.prevIndex, match.index);
    const product = obj.op === "dont" ? 0 : problem1(subtext);
    return {
      op: match.groups?.do ? "do" : "dont",
      sum: obj.sum + product,
      prevIndex: match.index,
    };
  }, { op: "do", sum: 0, prevIndex: 0 });
  return sum;
}
