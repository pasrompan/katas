const katas = require('../src/katas.cjs')

test('1 is not a prime', () => {
  expect(katas.isPrime(1)).toBe(false);
});

test('1 is aprime', () => {
    expect(katas.isPrime(2)).toBe(true);
  });

  test('3 is a prime', () => {
    expect(katas.isPrime(3)).toBe(true);
  });

  test('9 is not a prime', () => {
    expect(katas.isPrime(9)).toBe(false);
  });

  test('13 is a prime', () => {
    expect(katas.isPrime(13)).toBe(true);
  });

  test('4 is not a prime', () => {
    expect(katas.isPrime(4)).toBe(false);
  });


  test("An empty string was provided but not returned", () => {
    expect(katas.toCamelCase('')).toBe('');
  });

  test("toCamelCase('the_stealth_warrior') did not return correct value", () => {
    expect(katas.toCamelCase("the_stealth_warrior")).toBe("theStealthWarrior");
  });

  test("toCamelCase('the_stealth-warrior') did not return correct value", () => {
    expect(katas.toCamelCase("the_stealth-warrior")).toBe("theStealthWarrior");
  });

  test("toCamelCase('the_stealth-warrior') did not return correct value", () => {
    expect(katas.toCamelCase("_--")).toBe("");
  });