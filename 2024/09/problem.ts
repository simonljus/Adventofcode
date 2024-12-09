export function problem1(text: string) {
  const memory: (number | ".")[] = [];
  for (
    const [index, [fileSize, freeSize]]
      of (text.match(/\d{1,2}/g)?.map((m) =>
        m.split("").map((n) => Number.parseInt(n))
      ) ?? []).entries()
  ) {
    memory.push(
      ...Array(fileSize).fill(index),
      ...Array(freeSize ?? 0).fill("."),
    );
  }
  const freeIndexes: number[] = [];
  for (const [index, m] of memory.entries()) {
    if (m === ".") {
      freeIndexes.push(index);
    }
  }
  const lastIndex = memory.length - 1;
  const spaceIndexes: number[] = [];
  for (const [index, m] of memory.slice().reverse().entries()) {
    if (spaceIndexes.length === freeIndexes.length) {
      break;
    }
    if (m === ".") {
      continue;
    }
    spaceIndexes.push(lastIndex - index);
  }
  const swapped = memory.slice();
  for (const [index, freeIndex] of freeIndexes.entries()) {
    const spaceIndex = spaceIndexes[index];
    if (freeIndex > spaceIndex) {
      break;
    }
    swapped[freeIndex] = swapped[spaceIndex];
    swapped[spaceIndex] = ".";
  }
  return swapped.reduce((acc: number, curr, index) => {
    if (curr === ".") {
      return acc;
    }
    return acc + (curr * index);
  }, 0);
}

export function problem2(text: string) {
  const memory: (number | ".")[] = [];
  const freeSpaces: { startsAt: number; size: number }[] = [];
  const fileSpaces: { startsAt: number; size: number; id: number }[] = [];
  for (
    const [index, [fileSize, freeSize]]
      of (text.match(/\d{1,2}/g)?.map((m) =>
        m.split("").map((n) => Number.parseInt(n))
      ) ?? []).entries()
  ) {
    const file = new Array(fileSize).fill(index);
    const free = new Array(freeSize ?? 0).fill(".");
    freeSpaces.push({ startsAt: memory.length + file.length, size: freeSize });
    fileSpaces.push({ startsAt: memory.length, size: fileSize, id: index });
    memory.push(...file, ...free);
  }
  fileSpaces.reverse();

  for (const file of fileSpaces) {
    for (const freeSpace of freeSpaces) {
      if (file.startsAt < freeSpace.startsAt) {
        break;
      }
      if (freeSpace.size === 0) {
        continue;
      }

      if (file.size > freeSpace.size) {
        continue;
      }
      freeSpace.size -= file.size;
      file.startsAt = freeSpace.startsAt;
      freeSpace.startsAt += file.size;
      break;
    }
  }
  const checksum = fileSpaces.reduce((acc, curr) => {
    for (const [i, v] of new Array(curr.size).fill(curr.id).entries()) {
      acc += v * (i + curr.startsAt);
    }
    return acc;
  }, 0);
  return checksum;
}
