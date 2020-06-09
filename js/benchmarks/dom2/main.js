const randomValue = () => Math.floor(Math.random() * Math.floor(100000));

//====================
// table operations
//====================

const getTableBody = () => document.getElementById("main-table").getElementsByTagName("tbody")[0];

const rowTemplate = document.createElement("tr");
rowTemplate.innerHTML =
    `<td class="col0"></td>` +
    `<td class="col1"></td>` +
    `<td class="col2"></td>` +
    `<td class="col3"></td>` +
    `<td class="col4"></td>`;

const createRow = () => {
    const tr = rowTemplate.cloneNode(true);
    const td0 = tr.firstChild;
    const td1 = td0.nextSibling;
    const td2 = td1.nextSibling;
    const td3 = td2.nextSibling;
    const td4 = td3.nextSibling;
    td0.textContent = randomValue();
    td1.textContent = randomValue();
    td2.textContent = randomValue();
    td3.textContent = randomValue();
    td4.textContent = randomValue();
    return tr;
};

const clearRows = () => (getTableBody().textContent = "");
const replaceRowsInnerHTML = (numRows) => {
    let rows = [];
    for (let i = 0; i < numRows; i++) {
        rows.push(
            "<tr>" +
                `<td class="col0">${randomValue()}</td>` +
                `<td class="col1">${randomValue()}</td>` +
                `<td class="col2">${randomValue()}</td>` +
                `<td class="col3">${randomValue()}</td>` +
                `<td class="col4">${randomValue()}</td>` +
                "</tr>",
        );
    }
    getTableBody().innerHTML = rows.join("");
};
const replaceRowsAppendChild = (numRows) => {
    const body = getTableBody();
    body.textContent = "";
    for (let i = 0; i < numRows; i++) {
        body.appendChild(createRow());
    }
};
const updateEveryNthRow = (n) => {
    if (n <= 0) return;
    const rows = getTableBody().getElementsByTagName("tr");
    const numRows = rows.length;
    for (let i = 0; i < numRows; i += n) {
        rows[i].firstChild.textContent += "!";
    }
};
const updateFirstRow = () => (getTableBody().firstChild.firstChild.textContent += "!");
const updateLastRow = () => (getTableBody().lastChild.firstChild.textContent += "!");
const insertRowAtTop = () => {
    const body = getTableBody();
    body.insertBefore(createRow(), body.firstChild);
};

//====================
// list operations
//====================

const getList = () => document.getElementById("main-list");

const createItem = () => {
    const li = document.createElement("li");
    li.textContent = randomValue();
    return li;
};

const clearItems = () => {
    getList().textContent = "";
};
const appendItems = (numItems) => {
    const list = getList();
    for (let i = 0; i < numItems; i++) {
        list.appendChild(createItem());
    }
};
const replaceItemsInnerHTML = (numItems) => {
    let items = [];
    for (let i = 0; i < numItems; i++) {
        items.push(`<li>${randomValue()}</li>`);
    }
    getList().innerHTML = items.join("");
};
const replaceItemsAppendChild = (numItems) => {
    const ul = getList();
    ul.textContent = "";
    for (let i = 0; i < numItems; i++) {
        ul.appendChild(createItem());
    }
};
const updateEveryNthItem = (n) => {
    if (n <= 0) return;
    const items = getList().getElementsByTagName("li");
    const numItems = items.length;
    for (let i = 0; i < numItems; i += n) {
        items[i].textContent += "!";
    }
};
const updateFirstItem = () => (getList().firstChild.textContent += "!");
const updateLastItem = () => (getList().lastChild.textContent += "!");
const insertItemAtTop = () => {
    const list = getList();
    list.insertBefore(createItem(), list.firstChild);
};

//====================
// benchmarks
//====================

const stats = (xs) => {
    if (xs.length === 0) return [0, 0];
    if (xs.length === 1) return [xs[0], 0];
    const sum = xs.reduce((acc, cur) => acc + cur, 0);
    const mean = sum / xs.length;
    const squaredErr = xs.map((x) => Math.pow(x - mean, 2));
    const errSum = squaredErr.reduce((acc, cur) => acc + cur, 0);
    const variance = errSum / (xs.length - 1);
    const stdev = Math.sqrt(variance);
    return [mean, stdev];
};

const logEl = document.getElementById("log");
const log = (msg) => (logEl.textContent += msg);

const bench = (tests) => {
    if (tests.length === 0) {
        log("--- benchmarks complete ---\n");
        return;
    }

    const runsPerTest = 5;
    let numRuns = 0;
    let startTime;
    let times = [];

    const test = tests[0];
    log(test.name + ":");
    if (test.setup) test.setup();

    const start = () => {
        window.requestAnimationFrame(() => {
            window.requestAnimationFrame(() => {
                startTime = window.performance.now();
                test.run();
                window.requestAnimationFrame(finish);
            });
        });
    };

    const finish = () => {
        const elapsed = Math.round(window.performance.now() - startTime);
        log(" " + elapsed);
        times.push(elapsed);
        if (++numRuns < runsPerTest) {
            window.requestAnimationFrame(start);
            return;
        }
        const [mean, stdev] = stats(times);
        log(` (x̄ = ${mean.toFixed(1)}, s = ${stdev.toFixed(1)})\n`);
        setTimeout(() => bench(tests.slice(1)), 250);
    };

    start();
};

const runBenchmarks = () => {
    const factors = [0, 1, 2, 3, 4].map((n) => Math.pow(2, n));
    const base = 1000;
    bench(
        Array.prototype.concat(
            factors.map((n) => {
                return { name: `replaceRows_${n}K`, run: () => replaceRowsInnerHTML(n * base) };
            }),
            factors.map((n) => {
                return {
                    name: `updateRows_${n}K`,
                    setup: () => replaceRowsInnerHTML(n * base),
                    run: () => updateEveryNthRow(1),
                };
            }),
            factors.map((n) => {
                return {
                    name: `updateFirstRowOf_${n}K`,
                    setup: () => replaceRowsInnerHTML(n * base),
                    run: updateFirstRow,
                };
            }),
            factors.map((n) => {
                return {
                    name: `updateLastRowOf_${n}K`,
                    setup: () => replaceRowsInnerHTML(n * base),
                    run: updateLastRow,
                };
            }),
            factors.map((n) => {
                return {
                    name: `insertRowBefore_${n}K`,
                    setup: () => replaceRowsInnerHTML(n * base),
                    run: insertRowAtTop,
                };
            }),
            factors.map((n) => {
                return {
                    name: `replaceItems_${n}K`,
                    setup: clearRows, // clear table so list is visible
                    run: () => replaceItemsInnerHTML(n * base),
                };
            }),
            factors.map((n) => {
                return {
                    name: `updateItems_${n}K`,
                    setup: () => replaceItemsInnerHTML(n * base),
                    run: () => updateEveryNthItem(1),
                };
            }),
            factors.map((n) => {
                return {
                    name: `updateFirstItemOf_${n}K`,
                    setup: () => replaceItemsInnerHTML(n * base),
                    run: updateFirstItem,
                };
            }),
            factors.map((n) => {
                return {
                    name: `updateLastItemOf_${n}K`,
                    setup: () => replaceItemsInnerHTML(n * base),
                    run: updateLastItem,
                };
            }),
            factors.map((n) => {
                return {
                    name: `insertItemBefore_${n}K`,
                    setup: () => replaceItemsInnerHTML(n * base),
                    run: insertItemAtTop,
                };
            }),
        ),
    );
};

// Wire up the button.
document.getElementById("run").addEventListener("click", (evt) => {
    evt.preventDefault();
    runBenchmarks();
});
