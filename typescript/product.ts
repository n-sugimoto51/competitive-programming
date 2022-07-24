import * as fs from "fs";

function Main(input: string) {
  const [a_s, b_s] = input.split(" ");
  const a: number = parseInt(a_s);
  const b: number = parseInt(b_s);
  const product: number = a * b;
  const quotient: number = product % 2;

  const result: string = quotient === 0 ? "Even" : "Odd";

  console.log("%s", result);
}

Main(fs.readFileSync("/dev/stdin", "utf8"));
