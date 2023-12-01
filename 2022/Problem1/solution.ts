import * as fs from 'fs'
import { doubleNewline } from '../../utils/DOUBLE_NEWLINE.js';
import { resolve } from 'path';
const calculateSums = (file: string): Array<number> =>{ 
return fs.readFileSync(resolve(__dirname, file))
.toString().split(doubleNewline).map(item =>{
    return item
      .split('\n')
      .map(item => parseInt(item, 10))
      .reduce<number>((acc, item) => acc + item, 0)
})}
const sumOfXLargest = (file: string, x: number): number => {
    return calculateSums(file).sort((lhs, rhs) => lhs - rhs).slice(-x).reduce<number>((acc, item) => acc + item, 0)
}
console.log(Math.max(...calculateSums('input.txt')))
console.log(Math.max(sumOfXLargest('input.txt', 3)))