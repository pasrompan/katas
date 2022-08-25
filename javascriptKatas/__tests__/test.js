const kata = require('../src/kata.cjs')

test('1 is not a prime', () => {
  expect(kata.isPrime(1)).toBe(false);
});

test('1 is aprime', () => {
    expect(kata.isPrime(2)).toBe(true);
  });

  test('3 is a prime', () => {
    expect(kata.isPrime(3)).toBe(true);
  });

  test('9 is not a prime', () => {
    expect(kata.isPrime(9)).toBe(false);
  });

  test('13 is a prime', () => {
    expect(kata.isPrime(13)).toBe(true);
  });

  test('4 is not a prime', () => {
    expect(kata.isPrime(4)).toBe(false);
  });


  test("An empty string was provided but not returned", () => {
    expect(kata.toCamelCase('')).toBe('');
  });

  test("toCamelCase('the_stealth_warrior') did not return correct value", () => {
    expect(kata.toCamelCase("the_stealth_warrior")).toBe("theStealthWarrior");
  });

  test("toCamelCase('the_stealth-warrior') did not return correct value", () => {
    expect(kata.toCamelCase("the_stealth-warrior")).toBe("theStealthWarrior");
  });

  test("toCamelCase('the_stealth-warrior') did not return correct value", () => {
    expect(kata.toCamelCase("_--")).toBe("");
  });

  test("Testing for fixed tests", () =>{
    expect(kata.countBits(4)).toBe(1);
  }) 

  test("failing test", () => {
    expect(kata.countBits(6015773801)).toBe(16)
  })