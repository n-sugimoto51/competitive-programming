import * as fs from "fs";

function Main(input: string) {
  const s1: number = parseInt(input.substring(0, 1));
  const s2: number = parseInt(input.substring(1, 2));
  const s3: number = parseInt(input.substring(2, 3));

  const result = s1 + s2 + s3;

  console.log("%d", result);
}

// Main("000");
Main(fs.readFileSync("/dev/stdin", "utf8"));
