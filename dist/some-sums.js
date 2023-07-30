"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const fs = require("fs");
function Main(input) {
    const [n_s, a_s, b_s] = input.split(" ");
    const n = parseInt(n_s);
    const a = parseInt(a_s);
    const b = parseInt(b_s);
    let result = 0;
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
//# sourceMappingURL=some-sums.js.map