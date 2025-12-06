import { assertEquals } from "@std/assert";
import { problem1, problem2 } from "./problem.ts";
const testInputP1 = await Deno.readTextFile("./data/06/test-p1.txt");
const input = await Deno.readTextFile("./data/06/input.txt");
Deno.test("p1 test", () => {
  assertEquals(problem1(testInputP1), 4277556);
});

Deno.test("p1 Solution", () => {
  assertEquals(problem1(input), 4878670269096);
});

Deno.test("p2 test", () => {
  assertEquals(problem2(testInputP1), 3263827);
});

Deno.test("p2 Solution", () => {
  assertEquals(problem2(input), 8674740488592);
});
