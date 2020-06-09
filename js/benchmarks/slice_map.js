const slice = (xs, upto, fn) => xs.slice(0, upto).map(fn);

const loop = (xs, upto, fn) => {
    if (xs.length <= upto) {
        return xs.map(fn);
    }
    const ys = [];
    for (let i = 0; i < upto; i++) {
        ys.push(fn(xs[i]));
    }
    return ys;
};

const makeData = (n) => {
    const xs = [];
    for (let i = 0; i < n; i++) {
        xs.push(Math.random());
    }
    return xs;
};

const xs = makeData(2000);
const f1 = (x) => x * x;

slice(xs, 1000, f1);

loop(xs, 1000, f1);
