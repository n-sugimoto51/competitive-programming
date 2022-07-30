import * as fs from "fs";

function Main(input: string) {
  const inputArray: string[] = input.split("\n");
  const aCount: number = parseInt(inputArray[0]);
  const bCount: number = parseInt(inputArray[1]);
  const cCount: number = parseInt(inputArray[2]);
  const targetPrice: number = parseInt(inputArray[3]);

  let result: number = 0;

  for (let aIndex = 0; aIndex <= aCount; aIndex++) {
    if (aIndex * 500 > targetPrice) {
      break;
    }
    if (aIndex * 500 === targetPrice) {
      result++;
      continue;
    }

    for (let bIndex = 0; bIndex <= bCount; bIndex++) {
      if (aIndex * 500 + bIndex * 100 > targetPrice) {
        break;
      }
      if (aIndex * 500 + bIndex * 100 === targetPrice) {
        result++;
        break;
      }

      for (let cIndex = 1; cIndex <= cCount; cIndex++) {
        if (aIndex * 500 + bIndex * 100 + cIndex * 50 > targetPrice) {
          break;
        }
        if (aIndex * 500 + bIndex * 100 + cIndex * 50 === targetPrice) {
          result++;
          break;
        }
      }
    }
  }

  console.log("%d", result);
}

// Main("2\n2\n2\n100");
Main(fs.readFileSync("/dev/stdin", "utf8"));
