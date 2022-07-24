"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const fs = require("fs");
function Main(input) {
    const s1 = parseInt(input.substring(0, 1));
    const s2 = parseInt(input.substring(1, 2));
    const s3 = parseInt(input.substring(2, 3));
    const result = s1 + s2 + s3;
    console.log("%d", result);
}
// Main("000");
Main(fs.readFileSync("/dev/stdin", "utf8"));
//# sourceMappingURL=placing.js.map