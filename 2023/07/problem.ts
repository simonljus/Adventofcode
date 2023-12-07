type Play = {
  hand: string;
  bid: number;
};

const mapper = new Map<string, string>([
  ["2", "C"],
  ["3", "D"],
  ["4", "E"],
  ["5", "F"],
  ["6", "G"],
  ["7", "H"],
  ["8", "I"],
  ["9", "L"],
  ["T", "M"],
  ["J", "N"],
  ["Q", "O"],
  ["K", "P"],
  ["A", "R"],
]);
const mapperJoker = new Map<string, string>([
  ["2", "C"],
  ["3", "D"],
  ["4", "E"],
  ["5", "F"],
  ["6", "G"],
  ["7", "H"],
  ["8", "I"],
  ["9", "L"],
  ["T", "M"],
  ["J", "B"],
  ["Q", "O"],
  ["K", "P"],
  ["A", "R"],
]);
export function problem1(text: string) {
  const rows = text.split("\n").map((row) => row.split(" "));
  const sortedByRank = rows
    .map(([hand, bid]) => ({
      sorted: parseHand(hand, mapper),
      hand,
      strength: getStrength(hand),
      bid: Number.parseInt(bid),
    }))
    .sort((a, b) => {
      const strengthCompared = a.strength - b.strength;
      if (strengthCompared) {
        return strengthCompared;
      }
      if (a.sorted === b.sorted) {
        return 0;
      }
      return a.sorted < b.sorted ? -1 : 1;
    });
  console.log(sortedByRank.slice(900));
  const score = sortedByRank.reduce(
    (sum, play, index) => sum + play.bid * (index + 1),
    0
  );
  return score;
}
enum Strength {
  FIVE_OF_A_KIND = 7,
  FOUR_OF_A_KIND = 6,
  FULL_HOUSE = 5,
  THREE_OF_A_KIND = 4,
  TWO_PAIR = 3,
  ONE_PAIR = 2,
  HIGH_CARD = 1,
}
export function getStrength(hand: string): number {
  const counter = new Map<string, number>();
  for (const card of hand) {
    counter.set(card, (counter.get(card) ?? 0) + 1);
  }
  const counts = Array.from(counter.values());
  const maxCount = Math.max(...counts);
  switch (maxCount) {
    case 5:
      return Strength.FIVE_OF_A_KIND;
    case 4:
      return Strength.FOUR_OF_A_KIND;
    case 3:
      return counts.some((c) => c === 2)
        ? Strength.FULL_HOUSE
        : Strength.THREE_OF_A_KIND;
    case 2:
      return counts.filter((c) => c == 2).length == 2
        ? Strength.TWO_PAIR
        : Strength.ONE_PAIR;
    default:
      return Strength.HIGH_CARD;
  }
}
function getStrengthJoker(hand: string): number {
  const excludeJoker = hand.replaceAll("J", "");
  const jokerCount = hand.length - excludeJoker.length;
  if (jokerCount === 5) {
    return Strength.FIVE_OF_A_KIND;
  }
  if (jokerCount === 0) {
    return getStrength(hand);
  }
  const strength = getStrength(excludeJoker);
  switch (strength) {
    case Strength.FOUR_OF_A_KIND:
      return Strength.FIVE_OF_A_KIND;
    case Strength.THREE_OF_A_KIND:
      return jokerCount == 2
        ? Strength.FIVE_OF_A_KIND
        : Strength.FOUR_OF_A_KIND;
    case Strength.TWO_PAIR:
      return Strength.FULL_HOUSE;
    case Strength.FULL_HOUSE:
      return Strength.FULL_HOUSE;
    case Strength.ONE_PAIR:
      switch (jokerCount) {
        case 3:
          return Strength.FIVE_OF_A_KIND;
        case 2:
          return Strength.FOUR_OF_A_KIND;
        case 1:
          return Strength.THREE_OF_A_KIND;
        default:
          throw new Error("Unreachable one pair");
      }

    case Strength.HIGH_CARD: {
      switch (jokerCount) {
        case 4:
          return Strength.FIVE_OF_A_KIND;
        case 3:
          return Strength.FOUR_OF_A_KIND;
        case 2:
          return Strength.THREE_OF_A_KIND;
        case 1:
          return Strength.ONE_PAIR;
        default:
          throw new Error("Unreachable state");
      }
    }
  }
  throw new Error("Unreachable");
}

export function parseHand(hand: string, mapValue: Map<string, string>): string {
  return hand
    .split("")
    .map((card) => mapValue.get(card) ?? "X")
    .join("");
}

export function problem2(text: string) {
  const rows = text.split("\n").map((row) => row.split(" "));
  const sortedByRank = rows
    .map(([hand, bid]) => ({
      sorted: parseHand(hand, mapperJoker),
      hand,
      strength: getStrengthJoker(hand),
      bid: Number.parseInt(bid),
    }))
    .sort((a, b) => {
      const strengthCompared = a.strength - b.strength;
      if (strengthCompared) {
        return strengthCompared;
      }
      if (a.sorted === b.sorted) {
        return 0;
      }
      return a.sorted < b.sorted ? -1 : 1;
    });
  console.log(sortedByRank.slice(900));
  const score = sortedByRank.reduce(
    (sum, play, index) => sum + play.bid * (index + 1),
    0
  );
  return score;
}

function parseNumbers(str: string) {}
