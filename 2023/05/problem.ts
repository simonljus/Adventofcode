const numberPattern = new RegExp(/\d+/g);
export function problem1(text: string) {
  const [seedsRow, ...almanacStrings] = text.split("\n\n");
  console.log(seedsRow, almanacStrings);
  const seeds = parseNumbers(seedsRow);
  const almanacs = almanacStrings.map((a) => parseAlmanac(a));
  console.log(almanacs);
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
  console.log(values);
  return Math.min(...values);
}
type Range = {
  destination: number;
  source: number;
  range: number;
};
type Almanac = {
  ranges: Range[];
};

function convert(n: number, almanac: Almanac) {
  for (const range of almanac.ranges) {
    const diff = n - range.source;
    if (diff < 0) {
      console.log("n too small", n, range);
      continue;
    }
    if (diff >= range.range) {
      console.log("n too large", n, range);
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
  return { ranges: numbers };
}

export function problem2(text: string) {
  return 0;
}
