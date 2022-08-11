const isPrime = require('./katas.cjs')

test('1 is not a prime', () => {
  expect(isPrime(1)).toBe(false);
});

test('1 is aprime', () => {
    expect(isPrime(2)).toBe(true);
  });

  test('3 is a prime', () => {
    expect(isPrime(3)).toBe(true);
  });

  test('9 is not a prime', () => {
    expect(isPrime(9)).toBe(false);
  });

  test('13 is a prime', () => {
    expect(isPrime(13)).toBe(true);
  });

  test('4 is not a prime', () => {
    expect(isPrime(4)).toBe(false);
  });