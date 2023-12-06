import { assertEquals } from "testing";
import { problem1, problem2 } from "./problem.ts";
const testInputP1 = await Deno.readTextFile("test-p1.txt");
const testInputP2 = await Deno.readTextFile("test-p2.txt");
const input = await Deno.readTextFile("input.txt");
Deno.test("p1 test", () => {
  assertEquals(problem1(testInputP1), 35);
});

Deno.test("p2 test", () => {
  assertEquals(problem2(testInputP2), 46);
});

Deno.test("p1 solution", () => {
  assertEquals(problem1(input), 600279879);
});
Deno.test("p2 solution", () => {
  assertEquals(problem2(input), 20191102);
});