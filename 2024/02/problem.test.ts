import { assertEquals } from "@std/assert";
import { problem1, problem2 } from "./problem.ts";
const testInputP1 = await Deno.readTextFile("02/test-p1.txt");
const input = await Deno.readTextFile("02/input.txt");
Deno.test("p1 test", () => {
  assertEquals(problem1(testInputP1), 2);
});

Deno.test("p2 test", () => {
  assertEquals(problem2(testInputP1), 4);
});
Deno.test("p1 Solution", () => {
  assertEquals(problem1(input), 670);
});

Deno.test("p2 Solution", () => {
  assertEquals(problem2(input), 700);
});
