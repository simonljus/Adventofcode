export function problem1(text: string) {
  const multiplications = text.match(/mul\(\d+,\d+\)/gm) ?? [];
  return multiplications.reduce((sum, multiplication) => {
    return multiply(multiplication) + sum;
  }, 0);
}

export function multiply(text: string) {
  const factors = (text.match(/\d+/g) ?? []).map((n) => Number.parseInt(n));
  const [factorA, factorB] = factors;
  return factorA * factorB;
}

export function problem2(text: string) {
  const matches = Array.from(
    text.matchAll(/(?<mul>mul\(\d+,\d+\))|(?<do>do\(\))|(?<dont>don't\(\))/gm),
  );
  const { sum } = matches.reduce((acc, match) => {
    if (match.groups?.mul) {
      if (!acc.doMul) {
        return acc;
      }
      const product = multiply(match.groups?.mul);
      return { ...acc, sum: acc.sum + product };
    } else if (match.groups?.do) {
      return { ...acc, doMul: true };
    } else if (match.groups?.dont) {
      return { ...acc, doMul: false };
    }
    return acc;
  }, { doMul: true, sum: 0 });
  return sum;
}
