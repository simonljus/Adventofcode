import { expect, test } from "vitest";
import { p1, p2 } from "./problem";


test('p1', () => {
    expect(p1('./test_input.txt')).toBe(15)
    expect(p1('./input.txt')).toBe(14827)
} )

test('p2', () => {
    expect(p2('./test_input.txt')).toBe(12)
    expect(p2('./input.txt')).toBe(13889)
} )