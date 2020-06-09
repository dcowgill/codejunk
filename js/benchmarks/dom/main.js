let rowSeq = 1;
let itemSeq = 1;

let elapsedTimes = [];

const meanStdev = (xs) => {
    const sum = xs.reduce((acc, cur) => acc + cur, 0);
    const mean = sum / xs.length;
    if (xs.length <= 1) return { mean, stdev: 0 };
    const squaredErrors = xs.map((x) => Math.pow(x - mean, 2));
    const errorSum = squaredErrors.reduce((acc, cur) => acc + cur, 0);
    const variance = errorSum / (xs.length - 1);
    const stdev = Math.sqrt(variance);
    return { mean, stdev };
};

const makeHandler = (fn) => {
    let renderStart;
    const reportElapsedTime = () => {
        elapsedTimes.push(window.performance.now() - renderStart);
        const sts = meanStdev(elapsedTimes);
        const msg =
            `${elapsedTimes.map((x) => x.toFixed(1)).join(", ")} ` +
            `(x̄ = ${sts.mean.toFixed(1)}, s = ${sts.stdev.toFixed(1)})`;
        document.getElementById("elapsed").textContent = msg;
    };
    return (evt) => {
        evt.preventDefault();
        window.requestAnimationFrame(() => {
            renderStart = window.performance.now();
            fn();
            window.requestAnimationFrame(reportElapsedTime);
            const numRows = getTableBody().getElementsByTagName("tr").length;
            document.getElementById("curNumRows").textContent = numRows;
        });
    };
};

const getN = () => Number(document.getElementById("N").value);
const getTableBody = () => document.getElementById("main-table").getElementsByTagName("tbody")[0];
const getList = () => document.getElementById("main-list");

const rowTemplate = document.createElement("tr");
rowTemplate.innerHTML =
    `<td class="col0"></td>` +
    `<td class="col1"></td>` +
    `<td class="col2"></td>` +
    `<td class="col3"></td>` +
    `<td class="col4"></td>`;

const createNewRow = () => {
    const tr = rowTemplate.cloneNode(true);
    const td0 = tr.firstChild;
    const td1 = td0.nextSibling;
    const td2 = td1.nextSibling;
    const td3 = td2.nextSibling;
    const td4 = td3.nextSibling;
    td0.textContent = rowSeq;
    td1.textContent = rowSeq * 2;
    td2.textContent = rowSeq * 3;
    td3.textContent = rowSeq * 4;
    td4.textContent = rowSeq * 5;
    rowSeq++;
    return tr;
};

const createNewItem = () => {
    const li = document.createElement("li");
    li.textContent = itemSeq;
    itemSeq++;
    return li;
};

// Appends a character to the first cell in every nth row.
const updateEveryNthRow = (n) => {
    if (n <= 0) return;
    const rows = getTableBody().getElementsByTagName("tr");
    const numRows = rows.length;
    for (let i = 0; i < numRows; i += n) {
        rows[i].firstChild.textContent += "!";
    }
};

// Appends a character to every nth list item.
const updateEveryNthItem = (n) => {
    if (n <= 0) return;
    const items = getList().getElementsByTagName("li");
    const numItems = items.length;
    for (let i = 0; i < numItems; i += n) {
        items[i].textContent += "!";
    }
};

const actions = {
    //
    // table actions
    //

    clearRows: () => {
        getTableBody().textContent = "";
    },
    appendRows: () => {
        const numRows = getN();
        const tbody = getTableBody();
        for (let i = 0; i < numRows; i++) {
            tbody.appendChild(createNewRow());
        }
    },
    replaceRowsInnerHTML: () => {
        const numRows = getN();
        const tbody = getTableBody();
        let rows = [];
        for (let i = 0; i < numRows; i++) {
            rows.push(
                "<tr>" +
                    `<td class="col0">${rowSeq}</td>` +
                    `<td class="col1">${rowSeq * 2}</td>` +
                    `<td class="col2">${rowSeq * 3}</td>` +
                    `<td class="col3">${rowSeq * 4}</td>` +
                    `<td class="col4">${rowSeq * 5}</td>` +
                    "</tr>",
            );
            rowSeq++;
        }
        tbody.innerHTML = rows.join("");
    },
    replaceRowsAppendChild: () => {
        const numRows = getN();
        const tbody = getTableBody();
        tbody.textContent = "";
        for (let i = 0; i < numRows; i++) {
            tbody.appendChild(createNewRow());
        }
    },
    updateFirstRow: () => (getTableBody().firstChild.firstChild.textContent += "!"),
    updateLastRow: () => (getTableBody().lastChild.firstChild.textContent += "!"),
    updateAllRows: () => updateEveryNthRow(1),
    updateEvery10thRow: () => updateEveryNthRow(10),
    insertRowAtTop: () => {
        const tbody = getTableBody();
        const tr = createNewRow();
        tbody.insertBefore(tr, tbody.firstChild);
    },

    //
    // list actions
    //

    clearItems: () => {
        getList().textContent = "";
    },
    appendItems: () => {
        const numItems = getN();
        const list = getList();
        for (let i = 0; i < numItems; i++) {
            list.appendChild(createNewItem());
        }
    },
    replaceItemsInnerHTML: () => {
        const numItems = getN();
        let items = [];
        for (let i = 0; i < numItems; i++) {
            items.push(`<li>${itemSeq}</li>`);
            itemSeq++;
        }
        getList().innerHTML = items.join("");
    },
    replaceItemsAppendChild: () => {
        const numItems = getN();
        const ul = getList();
        ul.textContent = "";
        for (let i = 0; i < numItems; i++) {
            ul.appendChild(createNewItem());
        }
    },
    updateFirstItem: () => (getList().firstChild.textContent += "!"),
    updateLastItem: () => (getList().lastChild.textContent += "!"),
    updateAllItems: () => updateEveryNthItem(1),
    updateEvery10thItem: () => updateEveryNthItem(10),
    insertItemAtTop: () => {
        const list = getList();
        const li = createNewItem();
        list.insertBefore(li, list.firstChild);
    },

    resetElapsed: () => {},
};

// Decorate links with the appropriate handlers.
const getClickHandler = (el) => {
    const action = el.dataset.action;
    const fn = actions[action];
    if (!fn) {
        throw new Error(`link action "${action}" not found`);
    }
    return fn;
};
document.querySelectorAll("a").forEach((el) => {
    if (el.id === "resetElapsed") {
        el.addEventListener("click", (evt) => {
            evt.preventDefault();
            document.getElementById("elapsed").textContent = "n/a";
            elapsedTimes = [];
        });
    } else {
        el.addEventListener("click", makeHandler(getClickHandler(el)));
    }
});
