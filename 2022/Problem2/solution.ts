import { resolve } from "path";
import * as fs from "fs";

type Moves = "A" | "B" | "C";
type Result = "X" | "Y" | "Z";
type ParsedMoves = { opponent: Moves; player: Moves };
type ParsedResults = { opponent: Moves; result: Result };

const drawPoints = 3;
const winPoints = 6;

const pointsMap: Record<Moves, number> = {
  A: 1,
  B: 2,
  C: 3,
} as const;

const winnersMap: Record<Moves, Moves> = {
  A: "C",
  B: "A",
  C: "B",
};
const losersMap: Record<Moves, Moves> = {
  A: "B",
  B: "C",
  C: "A",
};
const readData = (file: string): Array<string> => {
  return fs
    .readFileSync(resolve(__dirname, file))
    .toString()
    .trim()
    .split(/\r?\n/);
};
const calculateScore = (file: string): number => {
return readData(file)
.map((item: string): ParsedMoves =>{
const parseItem = item
.replace('X','A')
.replace('Y','B')
.replace('Z','C')
.split(' ')
return{
    opponent: parseItem[0] as Moves,
    player: parseItem[1] as Moves
}
}
).reduce<number>((score, item) => {
    score += pointsMap[item.player]
    if (item.player === item.opponent){
        score += drawPoints
    }if (winnersMap[item.player] === item.opponent){
        score += winPoints
    }
    return score
}, 0
)
}
const calculateScore2 = (file: string): number => {
  return readData(file)
    .map((item: string): ParsedResults => {
      const parsedItem = item.split(' ')

      return {
        opponent: parsedItem[0] as Moves,
        result: parsedItem[1] as Result
      }
    })
    .reduce<number>((score, item) => {
      // Draw
      if (item.result === 'Y') {
        score += pointsMap[item.opponent]
        score += drawPoints
      }

      // Win
      if (item.result === 'Z') {
        score += pointsMap[losersMap[item.opponent]]
        score += winPoints
      }

      // Lose
      if (item.result === 'X') {
        score += pointsMap[winnersMap[item.opponent]]
      }

      return score
    }, 0)
}
console.log(calculateScore('input.txt'))
console.log(calculateScore2('input.txt'))