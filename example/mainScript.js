import myNumbers from './script1'
import mulSqrt from './script2'

const result = mulSqrt(myNumbers)

if (isNaN(result)) console.log('one of the numbers is negative')
else console.log('your result : ' + result)