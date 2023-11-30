import { expect, test } from "vitest";
import { p1, p2 } from "./problem";


test('p1', () => {
    expect(p1('./test_input.txt')).toBe(157)
    expect(p1('./input.txt')).toBe(7997)
} )

test('p2', () => {
    expect(p2('./test_input.txt')).toBe(70)
    expect(p2('./input.txt')).toBe(2545)
} )