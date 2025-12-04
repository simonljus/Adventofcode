function parse(text: string) {
  return text.split(",").map((t) => {
    const [min, max] = t.split("-");
    return { min: Number(min), max: Number(max) };
  });
}
export function problem1(text: string): number {
  const allRanges = parse(text);
  const ranges = allRanges.map((range) => {
    const minLength = range.min.toString().length;
    const maxLength = range.max.toString().length;
    const min = minLength % 2 === 0 ? range.min : Math.pow(10, minLength);
    const max = maxLength % 2 === 0
      ? range.max
      : Math.pow(10, maxLength - 1) - 1;
    return { min, max };
  }).filter((range) => range.min <= range.max);
  let count = 0;
  for (const range of ranges) {
    for (let min = range.min; min <= range.max; min++) {
      const minString = min.toString();
      const minLength = minString.length;
      if (minLength % 2 === 1) {
        continue;
      }
      const first = minString.slice(0, minLength / 2);
      const second = minString.slice(minLength / 2);
      if (first === second) {
        count += min;
      }
    }
  }
  return count;
}
export function problem2(text: string): number {
  const allRanges = parse(text);
  const ranges = allRanges.filter((range) => range.min <= range.max);
  let count = 0;
  for (const range of ranges) {
    for (let min = range.min; min <= range.max; min++) {
      const minString = min.toString();
      const halfMinLength = Math.floor(minString.length / 2);
      for (let startIndex = 1; startIndex <= halfMinLength; startIndex++) {
        minString.slice(0, startIndex);
        if (minString.length % startIndex !== 0) {
          continue;
        }
        if (
          minString.replaceAll(minString.slice(0, startIndex), "").length === 0
        ) {
          count += min;
          break;
        }
      }
    }
  }
  return count;
}
