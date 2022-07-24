import * as fs from "fs";

function Main(input: string) {
  const inputArray: string[] = input.split("\n");
  const progression: string[] = inputArray[1].split(" ");
  let resultCount: number = 0;

  // const retryCheck: boolean = progression.every(evenCheck);

  let enableRetryCheck: boolean = true;
  while (enableRetryCheck) {
    for (let i = 0; i < progression.length; i++) {
      const num: number = parseInt(progression[i]);
      enableRetryCheck = num % 2 === 0 ? enableRetryCheck : false;
      if (!enableRetryCheck) {
        break;
      }
      progression[i] = (num / 2).toString();
    }

    if (enableRetryCheck) {
      resultCount++;
    }
  }

  console.log("%d", resultCount);
}

Main(fs.readFileSync("/dev/stdin", "utf8"));
