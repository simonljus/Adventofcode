function parse(text: string) {
  const [rangesText, ingredientsText] = text.split("\n\n");
  const ranges = rangesText.split("\n").map((range) => {
    const [min, max] = range.trim().split("-");
    return { min: Number(min), max: Number(max) };
  });
  const ingredients = ingredientsText.split("\n").map((ingredient) =>
    Number(ingredient.trim())
  );
  return { ranges, ingredients };
}
export function problem1(text: string): number {
  const { ranges, ingredients } = parse(text);
  return ingredients.filter((ingredient) => {
    return ranges.some((range) =>
      range.min <= ingredient && range.max >= ingredient
    );
  }).length;
}
type Range = { min: number; max: number };
function isSame(a: Range, b: Range): boolean {
  return a.min === b.min && a.max === b.max;
}
function isWithin(small: Range, large: Range): boolean {
  return large.max >= small.max && large.min <= small.min;
}

function canCombine(a: Range, b: Range): boolean {
  return a.max <= b.max && a.max >= b.min;
}
function merge(range1: Range, range2: Range): Range | null {
  if (isSame(range1, range2)) {
    return range1;
  }

  if (isWithin(range1, range2)) {
    return range2;
  }
  if (isWithin(range2, range1)) {
    return range1;
  }
  if (canCombine(range1, range2)) {
    return { min: range1.min, max: range2.max };
  }
  if (canCombine(range2, range1)) {
    return { min: range2.min, max: range1.max };
  }
  return null;
}
export function problem2(text: string): number {
  const { ranges } = parse(text);
  let iterations = 0;
  while (true) {
    let changed = false;
    for (const [i, range] of ranges.entries()) {
      let improvedRange: { range: Range; index: number } | null = null;
      iterations += 1;
      for (const [j, range2] of ranges.entries()) {
        if (i >= j) {
          continue;
        }

        const result = merge(range, range2);
        if (result) {
          improvedRange = { range: result, index: j };
          break;
        }
      }
      if (improvedRange) {
        ranges.splice(improvedRange.index, 1);
        ranges[i] = improvedRange.range;
        changed = true;
        break;
      }
    }
    if (!changed) {
      break;
    }
  }
  let sum = 0;
  for (const range of ranges) {
    sum += range.max - range.min + 1;
  }
  return sum;
}
