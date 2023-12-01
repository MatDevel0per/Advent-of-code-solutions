import * as fs from 'fs'
import { doubleNewline } from '../../utils/DOUBLE_NEWLINE.js';
import { resolve } from 'path';
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
const calculatesums = (file: string): Array<Number> =>{ 
return fs.readFileSync(resolve(__dirname, file))
.toString().split(/\r?\n/).map(item =>{
    return item
      .replace(/\D/g,'')
      .replace(/(?!^|\d$)[^]/g, '')})
      .map(Number)
    }     
console.log(calculateSums('input.txt'))
console.log(calculateSums('example.txt'))