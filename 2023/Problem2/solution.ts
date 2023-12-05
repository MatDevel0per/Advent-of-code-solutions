import * as fs from 'fs'
import { resolve } from 'path';


type CubeCount = { red: number; green: number; blue: number };
type GameRecord = { id: number; rounds: CubeCount[] };
// Input data
const input = fs.readFileSync(resolve(__dirname, 'input.txt'))
.toString().split(/\r?\n/)

const powerUpTheCubes = (cubeCounts: CubeCount): number => {
  return cubeCounts.red * cubeCounts.green * cubeCounts.blue;
}

const minimumCubes = (game: CubeCount[]): CubeCount =>{
  const minCubes: CubeCount = { red: 0, green: 0, blue: 0  }
  for( let round of game){
    minCubes.red = Math.max(minCubes.red, round.red);
    minCubes.green = Math.max(minCubes.green, round.green);
    minCubes.blue = Math.max(minCubes.blue, round.blue);
  }
  return minCubes

}

const isPossibleGame = (cubeCounts: CubeCount, game: CubeCount[]): boolean => {
  for (const round of game) {
    if (
      round.red > cubeCounts.red ||
      round.green > cubeCounts.green ||
      round.blue > cubeCounts.blue
    ) {
      return false;
    }
  }
  return true;
}

const possibleGames = (cubeCounts: CubeCount, games: GameRecord[]): number[] => {
  let possibleGameIds: number[] = [];

  for (const game of games) {
    if (isPossibleGame(cubeCounts, game.rounds)) {
      possibleGameIds.push(game.id);
    }
  }

  return possibleGameIds;
}


// Parse input data
const games: GameRecord[] = input.map((line) => {
  const [idPart, roundsPart] = line.split(": ");
  const id = parseInt(idPart.split(" ")[1]);
  const rounds = roundsPart.split("; ").map((round) => {
    const cubes: CubeCount = { red: 0, green: 0, blue: 0 };
    round.split(", ").forEach((item) => {
      const [count, color] = item.split(" ");
      cubes[color as keyof CubeCount] = parseInt(count);
    });
    return cubes;
  });
  return { id, rounds };
});

// Target cube configuration
const targetCubeCounts: CubeCount = { red: 12, green: 13, blue: 14 };

// Find possible games
const possibleGameIds = possibleGames(targetCubeCounts, games);

// Calculate the sum of possible game IDs
const part1 = possibleGameIds.reduce((sum, id) => sum + id, 0);
// Calc the power of possible games
const part2 = games.reduce((sum, game) => {
  let minCubes =  minimumCubes(game.rounds);
  return sum + powerUpTheCubes(minCubes);
},0)
console.log("Sum of possible game IDs:", part1);
console.log("Sum of game power:", part2);

