type RGB = Map<string, number>;

const bag: RGB = new Map<string, number>([
  ["red", 12],
  ["green", 13],
  ["blue", 14],
]);

export function problem1(text: string) {
  const games = text.split("\n");
  return games.reduce((count, game, index) => {
    const gameId = index + 1;
    const hands = game.split(":").at(1)?.split(";") ?? [];
    if (hands.every((hand) => isHandValid(getRGBFromHand(hand)))) {
      return count + gameId;
    }

    return count;
  }, 0);
}

export function getPower(rgb: RGB) {
  rgb.entries;
}

export function getRGBFromHand(hand: string): RGB {
  const map = new Map<string, number>();
  const pattern = new RegExp(/(\d+)(\s+)(?=([a-z]+))/g);

  for (const match of hand.matchAll(pattern)) {
    const count = match?.at(1);
    const color = match?.at(3);
    if (!color || !count) {
      continue;
    }
    const prev = map.get(color) ?? 0;
    map.set(color, Number.parseInt(count) + prev);
  }

  return map;
}
export function isHandValid(hand: RGB) {
  return Array.from(hand.entries()).every(
    ([color, count]) => (bag.get(color) ?? 0) >= count
  );
}

export function problem2(text: string) {
  const games = text.split("\n");

  return games.reduce((powerSum, game) => {
    const hands = game.split(":").at(1)?.split(";") ?? [];
    const smallestBag = hands.reduce((minBag, hand) => {
      for (const [color, colorCount] of getRGBFromHand(hand)) {
        if ((minBag.get(color) ?? 0) < colorCount) {
          minBag.set(color, colorCount);
        }
      }
      return minBag;
    }, new Map<string, number>() as RGB);
    const power = Array.from(smallestBag.values()).reduce(
      (product, count) => (product || 1) * count,
      0
    );
    return power + powerSum;
  }, 0);
}
