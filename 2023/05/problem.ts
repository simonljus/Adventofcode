const numberPattern = new RegExp(/\d+/g);
export function problem1(text: string) {
  const [seedsRow, ...almanacStrings] = text.split("\n\n");
  const seeds = parseNumbers(seedsRow);
  const almanacs = almanacStrings.map((a) => parseAlmanac(a));
  almanacs.map((almanac) => {
    almanac.ranges;
  });
  const values: number[] = [];
  for (const seed of seeds) {
    let currentSeed = seed;
    for (const almanac of almanacs) {
      currentSeed = convert(currentSeed, almanac);
    }
    values.push(currentSeed);
  }
  return Math.min(...values);
}
type Range = {
  destination: number;
  source: number;
  range: number;
};
type Almanac = {
  ranges: Range[];
  seen: Set<number>;
};

function pairs<T>(arr: T[]) {
  const copy = arr.slice();
  const allPairs: { start: T; range: T }[] = [];
  while (copy.length) {
    const [start, range] = copy.splice(0, 2);
    allPairs.push({ start, range });
  }
  return allPairs;
}

function convert(n: number, almanac: Almanac) {
  for (const range of almanac.ranges) {
    const diff = n - range.source;
    if (diff < 0) {
      //("n too small", n, range);
      continue;
    }
    if (diff >= range.range) {
      //("n too large", n, range);
      continue;
    }
    return range.destination + diff;
  }
  return n;
}

function parseNumbers(str: string) {
  const numberPattern = new RegExp(/\d+/g);
  const numbers: number[] = [];
  for (const match of str.matchAll(numberPattern)) {
    for (const numeric of match.values()) {
      numbers.push(Number.parseInt(numeric));
    }
  }
  return numbers;
}

function parseAlmanac(almanac: string): Almanac {
  const mappers = almanac.split("\n").slice(1);
  const numbers = mappers
    .map((m) => parseNumbers(m))
    .map(
      ([destination, source, range]) =>
        ({ destination, source, range } as Range)
    );
  return { ranges: numbers, seen: new Set<number>() };
}

export function problem2(text: string) {
  const [seedsRow, ...almanacStrings] = text.split("\n\n");
  const seeds = pairs(parseNumbers(seedsRow));
  const almanacs = almanacStrings.map((a) => parseAlmanac(a));
  almanacs.map((almanac) => {
    almanac.ranges;
  });
  const values: Set<number> = new Set<number>();
  let minValue = Infinity;
  for (const [pairIndex, seedRange] of seeds.entries()) {
    console.log("Pair:", pairIndex);
    const maxSeed = seedRange.start + seedRange.range;
    loopSeed: for (let seed = seedRange.start; seed < maxSeed; seed++) {
      let currentSeed = seed;
      for (const almanac of almanacs) {
        currentSeed = convert(currentSeed, almanac);
      }

      if (currentSeed < minValue) {
        minValue = currentSeed;
      }
    }
  }
  return minValue;
}
