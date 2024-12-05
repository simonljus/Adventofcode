export function problem1(text: string) {
  const orderings = getRules(text);
  const updates = getUpdates(text);

  return updates.reduce((acc, update) => {
    if (!isSorted(orderings, update)) {
      return acc;
    }
    return acc + update[(update.length - 1) / 2];
  }, 0);
}

function isSorted(orderings: Map<number, Set<number>>, update: number[]) {
  return update.every((x, i) => {
    return update.slice(i + 1).every((y) => {
      return !orderings.get(y)?.has(x);
    });
  });
}

function getRules(text: string) {
  const orderings = new Map<number, Set<number>>();
  const [rulesText, updatesText] = text.split("\n\n");
  for (const rule of rulesText.split("\n").filter((r) => r)) {
    const [x, y] = rule.match(/\d+/g)?.map((n) => Number.parseInt(n)) ?? [];
    const s = orderings.get(x) ?? new Set<number>();
    s.add(y);
    orderings.set(x, s);
  }
  return orderings;
}

function getUpdates(text: string) {
  const [rulesText, updatesText] = text.split("\n\n");
  const updates = updatesText.split("\n")?.filter((l) => l).map((update) =>
    update.split(",").map((n) => Number.parseInt(n))
  );
  return updates;
}

export function problem2(text: string) {
  const orderings = getRules(text);
  const updates = getUpdates(text);
  return updates.reduce((acc, update) => {
    if (isSorted(orderings, update)) {
      return acc;
    }
    const copy = update.slice();
    copy.sort((a, b) => {
      return orderings.get(b)?.has(a) ? 1 : -1;
    });

    return acc + copy[(copy.length - 1) / 2];
  }, 0);
}
