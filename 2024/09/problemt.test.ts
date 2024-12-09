import { assertEquals } from "@std/assert";
import { problem1, problem2 } from "./problem.ts";
const testInputP1 = await Deno.readTextFile("09/test-p1.txt");
const testInputP2 = await Deno.readTextFile("09/test-p1.txt");
const input = await Deno.readTextFile("09/input.txt");
Deno.test("p1 test", () => {
  assertEquals(problem1(testInputP1), 1928);
});

Deno.test("p1 Solution", () => {
  assertEquals(problem1(input), 6332189866718);
});

Deno.test("p2 test", () => {
  assertEquals(problem2(testInputP2), 2858);
});

Deno.test("p2 Solution", () => {
  assertEquals(problem2(input), 6353648390778);
});
