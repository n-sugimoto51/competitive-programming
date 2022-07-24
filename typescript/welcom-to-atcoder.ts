import * as fs from "fs";

// inputに入力データ全体が入る
function Main(input: string) {
  // 1行目がinput[0], 2行目がinput[1], …に入る
  const input_array: string[] = input.split("\n");
  const a: number = parseInt(input_array[0]);
  const [b_s, c_s] = input_array[1].split(" ");
  const b = parseInt(b_s);
  const c = parseInt(c_s);
  const s = input_array[2];
  const sum = a + b + c;

  //出力
  console.log("%d %s", sum, s);
}

Main(fs.readFileSync("/dev/stdin", "utf8"));
