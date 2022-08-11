function isPrime(num) {
    if (num <=1)
        return false;
    if (num <= 3)
        return true;
    let iterationStep = 1;
    for (let i = 2; i*i <= num; i=i+iterationStep) {
            let remainder = num % i;
            if (remainder == 0)
            {
                return false;
            }
            if (i==3)
            {
                iterationStep = 2;
            }
        }
    return true;
}

module.exports = isPrime;
//let x = isPrime(5); 
//console.log(x)