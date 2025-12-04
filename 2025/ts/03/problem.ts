function parse(text: string) {
  return text.split("\n").map((t) => {
    const bank = t.split("");
    return bank.map((b) => Number(b));
  });
}
export function problem1(text: string): number {
  const banks = parse(text);
  let count = 0;
  for (const bank of banks) {
    const rest = bank.slice();
    rest.pop();
    const firstValue = Math.max(...rest);
    const firstIndex = bank.indexOf(firstValue);
    const lastValue = Math.max(...bank.slice(firstIndex + 1));
    const bankJoltage = Number(`${firstValue}${lastValue}`);
    count += bankJoltage;
  }
  return count;
}
export function problem2(text: string): number {
  const banks = parse(text);
  const batterySize = 12;
  let count = 0;
  for (const bank of banks) {
    let sliced = bank.slice();
    const values: number[] = [];
    for (let i = batterySize - 1; i >= 0; i--) {
      const value = Math.max(...sliced.slice(0, sliced.length - i));
      const startIndex = sliced.indexOf(value);
      sliced = sliced.slice(startIndex + 1);
      values.push(value);
    }
    count += Number(values.join(""));
  }
  return count;
}
