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
const calculateSums2 = (file: string): Array<Number> =>{ 
return fs.readFileSync(resolve(__dirname, file))
.toString().split(/\r?\n/).map(item =>{
    var cleanedItem = item
    let array1;
    let regex1 = RegExp('one|two|three|four|five|six|seven|eight|nine', 'g');
    while ((array1 = regex1.exec(cleanedItem)) !==null){
        let num = numberMap[array1[0]]
        let reg = new RegExp('^'+ array1[0], "g")
        cleanedItem = cleanedItem.replace(reg, num)
        console.log(array1)
    }
    
   cleanedItem 
    .replace(/\D/g,'')
    .replace(/(?!^|\d$)[^]/g, '')
    if (cleanedItem.length == 1){
        cleanedItem +=cleanedItem;
      }
      return cleanedItem})
      .map(Number)
      //.reduce<number>((acc, item) => acc + item, 0)
    }
console.log(calculateSums('input.txt'))
//console.log(calculateSums2('input.txt'))
console.log(calculateSums2('example.txt'))