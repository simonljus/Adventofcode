export function problem1(text: string): number {
  let count = 0;
  for (const line of text.split("\n")) {
    const levels = line.split(" ").map((level) =>
      Number.parseInt(level.trim())
    );
    if (isSafe(levels)) {
      count += 1;
    }
  }

  return count;
}

function isSafe(levels: number[]) {
  if (
    levels.toSorted((a, b) => a - b).join() !== levels.join() &&
    levels.toSorted((a, b) => b - a).join() !== levels.join()
  ) {
    return false;
  }
  for (const [index, level] of levels.entries()) {
    if (index === 0) {
      continue;
    }
    const diff = Math.abs(levels[index - 1] - level);
    if (diff < 1 || diff > 3) {
      return false;
    }
  }
  return true;
}

export function problem2(text: string): number {
  let count = 0;
  for (const line of text.split("\n")) {
    const levels = line.split(" ").map((level) =>
      Number.parseInt(level.trim())
    );
    if (isSafe(levels)) {
      count += 1;
      continue;
    }
    for (let i = 0; i < levels.length; i++) {
      const copy = levels.slice();
      copy.splice(i, 1);
      if (isSafe(copy)) {
        count += 1;
        break;
      }
    }
  }

  return count;
}
