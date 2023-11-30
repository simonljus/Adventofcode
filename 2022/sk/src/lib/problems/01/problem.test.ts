import { expect, test } from "vitest";
import { add, p1, p2 } from "./problem";

test('add', () => {
    expect(add(4,6)).toBe(10)
})
test('p1_test', () => {
    expect(p1()).toBe(24000)
} )

test('p1', () => {
    expect(p1('./input.txt')).toBe(71300)
} )
test('p2', () => {
    expect(p2('./test_input.txt')).toBe(45000)
    expect(p2('./input.txt')).toBe(209691)
} )