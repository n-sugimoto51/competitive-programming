"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const evenCheck = (num_s) => parseInt(num_s) % 2 === 0;
function Main(input) {
    const inputArray = input.split("\n");
    const progression = inputArray[1].split(" ");
    let resultCount = 0;
    // const retryCheck: boolean = progression.every(evenCheck);
    let enableRetryCheck = true;
    while (enableRetryCheck) {
        for (let i = 0; i < progression.length; i++) {
            const num = parseInt(progression[i]);
            enableRetryCheck = num % 2 === 0 ? enableRetryCheck : false;
            if (!enableRetryCheck) {
                break;
            }
            progression[i] = (num / 2).toString();
        }
        console.log(progression);
        if (enableRetryCheck) {
            resultCount++;
        }
    }
    console.log("%d", resultCount);
}
Main("3\n8 12 40");
// Main(fs.readFileSync("/dev/stdin", "utf8"));
//# sourceMappingURL=shift-only.js.map