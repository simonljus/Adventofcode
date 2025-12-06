import { assertEquals } from "@std/assert";
import { problem1, problem2 } from "./problem.ts";
const testInputP1 = await Deno.readTextFile("./data/05/test-p1.txt");
const input = await Deno.readTextFile("./data/05/input.txt");
Deno.test("p1 test", () => {
  assertEquals(problem1(testInputP1), 3);
});

Deno.test("p1 Solution", () => {
  assertEquals(problem1(input), 744);
});

Deno.test("p2 test", () => {
  assertEquals(problem2(testInputP1), 14);
});

Deno.test("p2 Solution", () => {
  assertEquals(problem2(input), 347468726696961);
});
