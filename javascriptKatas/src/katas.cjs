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

function toCamelCase(word){
    const myArray = word.split(/_|-/) 
    let concatenatedText = "";
    if (myArray.length < 1)
    {
        return "";
    }
    concatenatedText = myArray[0];
    for (i=1; i <myArray.length; i++)
    {
        const str = myArray[i]
        const str2 = str.charAt(0).toUpperCase() + str.slice(1);
        concatenatedText += str2;
    }
    return concatenatedText;
}

module.exports =
{
    isPrime,
    toCamelCase
} 
//let x = isPrime(5); 
//console.log(x)