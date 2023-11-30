import { expect, test } from "vitest";
import { hasSomeOverlap, p1, p2 } from "./problem";


test('p1', () => {
    expect(p1('./test_input.txt')).toBe(2)
    expect(p1('./input.txt')).toBe(651)
} )

test('hassomeoverlap', () => {
    expect(hasSomeOverlap({start:1, end: 2}, {start:2,end: 3})).toBe(true)
    expect(hasSomeOverlap({start:1, end: 2}, {start:2,end: 4})).toBe(true)
    expect(hasSomeOverlap({start:1, end: 2}, {start:3,end: 5})).toBe(false)
    expect(hasSomeOverlap({start:1, end: 3}, {start:2,end: 4})).toBe(true)
    expect(hasSomeOverlap({start:1, end: 1}, {start:2,end: 2})).toBe(false)
    expect(hasSomeOverlap({start:1, end: 1}, {start:2,end: 2})).toBe(false)
} )

test('p2', () => {
    expect(p2('./test_input.txt')).toBe(4)
    expect(p2('./input.txt')).toBe(956)
} )

