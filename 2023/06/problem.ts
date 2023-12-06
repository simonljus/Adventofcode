export function problem1(text: string) {
  const [times, distances] = text.split("\n").map((row) => parseNumbers(row));
  return times.reduce((product, time, i) => {
    const distanceToBeat = distances.at(i);
    if (distanceToBeat === undefined) {
      throw Error("distance not found");
    }
    return (product || 1) * pq({ time, distance: distanceToBeat });
  }, 0);
}

function pq({ time, distance }: { time: number; distance: number }) {
  const delta = Math.sqrt(time * time - 4 * distance);
  const timeMax = (time + delta) / 2;
  const maxOffset = Number.isInteger(timeMax) ? -1 : 0;
  const timeMin = (time - delta) / 2;
  const minOffset = Number.isInteger(timeMin) ? 1 : 0;
  const minInt = Math.ceil(timeMin + minOffset);
  const maxInt = Math.floor(timeMax + maxOffset);

  const counts = maxInt - minInt + 1;
  return counts;
}

export function problem2(text: string) {
  const [time, distance] = text
    .split("\n")
    .map((row) => Number.parseInt(parseNumbers(row).join("")));
  return pq({ time, distance });
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
