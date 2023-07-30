import * as fs from "fs";

function Main(input: string) {
  const [n_s, cardsSt] = input.split("\n");
  const n: number = parseInt(n_s);
  const cards = cardsSt.split(" ");
  let result: number = 0;

  cards.sort((a, b) => parseInt(a) - parseInt(b));
  console.log(cards);

  console.log("%s", result);
}

Main("3\n2 7 4");
// Main(fs.readFileSync("/dev/stdin", "utf8"));
