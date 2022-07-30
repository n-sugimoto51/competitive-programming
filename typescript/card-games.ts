import * as fs from "fs";

function Main(input: string) {
  const [n_s, a_s, b_s] = input.split(" ");
  const n: number = parseInt(n_s);
  const a: number = parseInt(a_s);
  const b: number = parseInt(b_s);

  let result: number = 0;

  for (let i = 1; i <= n; i++) {
    const digitArray = [...i.toString()].map((num_s) => parseInt(num_s));
    const digitSum = digitArray.reduce((sum, el) => {
      return sum + el;
    }, 0);

    if (a <= digitSum && digitSum <= b) {
      result += i;
    }
  }

  console.log("%s", result);
}

Main(fs.readFileSync("/dev/stdin", "utf8"));
