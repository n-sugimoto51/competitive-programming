"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
function Main(input) {
    const [n_s, cardsSt] = input.split("\n");
    const n = parseInt(n_s);
    const cards = cardsSt.split(" ");
    let result = 0;
    cards.sort((a, b) => parseInt(a) - parseInt(b));
    console.log(cards);
    console.log("%s", result);
}
Main("3\n2 7 4");
// Main(fs.readFileSync("/dev/stdin", "utf8"));
//# sourceMappingURL=card-games.js.map