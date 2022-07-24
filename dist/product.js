"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const fs = require("fs");
function Main(input) {
    const [a_s, b_s] = input.split(" ");
    const a = parseInt(a_s);
    const b = parseInt(b_s);
    const product = a * b;
    const quotient = product % 2;
    const result = quotient === 0 ? "Even" : "Odd";
    console.log("%s", result);
}
Main(fs.readFileSync("/dev/stdin", "utf8"));
//# sourceMappingURL=product.js.map