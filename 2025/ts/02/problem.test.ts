import { assertEquals } from "@std/assert";
import { problem1, problem2 } from "./problem.ts";
const testInputP1 = await Deno.readTextFile("./data/02/test-p1.txt");
const input = await Deno.readTextFile("./data/02/input.txt");
Deno.test("p1 test", () => {
  assertEquals(problem1(testInputP1), 1227775554);
});

Deno.test("p1 Solution", () => {
  assertEquals(problem1(input), 31210613313);
});

Deno.test("p2 test", () => {
  assertEquals(problem2(testInputP1), 4174379265);
});

Deno.test("p2 Solution", () => {
  assertEquals(problem2(input), 41823587546);
});
