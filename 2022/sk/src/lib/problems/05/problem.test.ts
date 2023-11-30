import { expect, test } from "vitest";
import { p1, p2 } from "./problem";


test('p1', () => {
    expect(p1('./test_input.txt')).toBe('CMZ')
    expect(p1('./input.txt')).toBe('ZBDRNPMVH')
} )

test('p2', () => {
    expect(p2('./test_input.txt')).toBe('MCD')
    expect(p2('./input.txt')).toBe('WDLPFNNNB')
} )