import * as fs from 'fs'
import { doubleNewline } from '../../utils/DOUBLE_NEWLINE.js';
import { resolve } from 'path';
const numberMap: Record<string, string> = {
  one: "1",
  two: "2",
  three: "3",
  four: "4",
  five: "5",
  six: "6",
  seven: "7",
  eight:"8",
  nine: "9",
};
const numberMap2: Record<string,number> ={
    "one": 1,
    "two": 2,
    "three": 3,
    "four": 4,
    "five": 5,
    "six": 6,
    "seven": 7,
    "eight": 8,
    "nine": 9
}
const calculateSums = (file: string): Number =>{ 
return fs.readFileSync(resolve(__dirname, file))
.toString().split(/\r?\n/).map(item =>{
    let cleanedItem = item
      .replace(/\D/g,'')
      .replace(/(?!^|\d$)[^]/g, '');
      if (cleanedItem.length == 1){
        cleanedItem +=cleanedItem;
      }
      return cleanedItem
})
      .map(Number)
      .reduce<number>((acc, item) => acc + item, 0)
    }
// Calculate sums did not work :(
const calculateSums2 = (file: string): Array<Number> =>{ 
return fs.readFileSync(resolve(__dirname, file))
.toString().split(/\r?\n/).map(item =>{
    var cleanedItem = item
    let array1;
    let regex1 = /\d/;
    while ((array1 = regex1.exec(cleanedItem)) !==null){
        console.log(array1) 
        break
        let num = numberMap[array1[0]]
        let reg = new RegExp(array1[0])
        cleanedItem = cleanedItem.replace(reg, num)
        
    }
   cleanedItem = cleanedItem
    .replace(/\D/g,'')
    .replace(/(?!^|\d$)[^]/g, '')
    if (cleanedItem.length == 1){
        cleanedItem +=cleanedItem;
      }
      return cleanedItem})
      .filter(cleanedItem => cleanedItem !== '')
      .map(Number)
      //.reduce<number>((acc, item) => acc + item, 0)
    }
function findNum(line, i){
    if(Number.isInteger(+line[i])){
        return +line[i];
    }
    for(const num of Object.keys(numberMap2)){
        if(line.substring(i).indexOf(num) === 0){
            return numberMap2[num];
        }
    }
    return 0;
}
const ifInDoubtForLoopItOut = (file: string): Number => {
    let sum = 0;
    let arrayOut = [];
    var data =  fs.readFileSync(resolve(__dirname, file), 'utf-8')
    .split(/\r\n/)
    for(let line of data) {
    let num1 = 0;
    let num2 = 0;
    for(let i = 0; i < line.length; i++) {
        num1 = findNum(line, i);
        if(num1 !== 0) break;
    }
    for(let i = line.length - 1; i >= 0; i--) {
        num2 = findNum(line, i);
        if(num2 !== 0) break;
    }
    arrayOut.push(`${num1}${num2}`)
    sum += +`${num1}${num2}`;
}
   return sum
}
console.log(calculateSums('input.txt'))
console.log(ifInDoubtForLoopItOut('input.txt'))
console.log(calculateSums2('example.txt'))