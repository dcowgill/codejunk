function _getRequireWildcardCache() {
    if (typeof WeakMap !== "function") return null;
    var cache = new WeakMap();
    _getRequireWildcardCache = function() {
        return cache;
    };
    return cache;
}

function _interopRequireWildcard(obj) {
    if (obj && obj.__esModule) {
        return obj;
    }
    if (obj === null || (typeof obj !== "object" && typeof obj !== "function")) {
        return { default: obj };
    }
    var cache = _getRequireWildcardCache();
    if (cache && cache.has(obj)) {
        return cache.get(obj);
    }
    var newObj = {};
    var hasPropertyDescriptor = Object.defineProperty && Object.getOwnPropertyDescriptor;
    for (var key in obj) {
        if (Object.prototype.hasOwnProperty.call(obj, key)) {
            var desc = hasPropertyDescriptor ? Object.getOwnPropertyDescriptor(obj, key) : null;
            if (desc && (desc.get || desc.set)) {
                Object.defineProperty(newObj, key, desc);
            } else {
                newObj[key] = obj[key];
            }
        }
    }
    newObj.default = obj;
    if (cache) {
        cache.set(obj, newObj);
    }
    return newObj;
}

function _createForOfIteratorHelper(o) {
    if (typeof Symbol === "undefined" || o[Symbol.iterator] == null) {
        if (Array.isArray(o) || (o = _unsupportedIterableToArray(o))) {
            var i = 0;
            var F = function F() {};
            return {
                s: F,
                n: function n() {
                    if (i >= o.length) return { done: true };
                    return { done: false, value: o[i++] };
                },
                e: function e(_e2) {
                    throw _e2;
                },
                f: F,
            };
        }
        throw new TypeError(
            "Invalid attempt to iterate non-iterable instance.\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method.",
        );
    }
    var it,
        normalCompletion = true,
        didErr = false,
        err;
    return {
        s: function s() {
            it = o[Symbol.iterator]();
        },
        n: function n() {
            var step = it.next();
            normalCompletion = step.done;
            return step;
        },
        e: function e(_e3) {
            didErr = true;
            err = _e3;
        },
        f: function f() {
            try {
                if (!normalCompletion && it.return != null) it.return();
            } finally {
                if (didErr) throw err;
            }
        },
    };
}

function _slicedToArray(arr, i) {
    return (
        _arrayWithHoles(arr) ||
        _iterableToArrayLimit(arr, i) ||
        _unsupportedIterableToArray(arr, i) ||
        _nonIterableRest()
    );
}

function _nonIterableRest() {
    throw new TypeError(
        "Invalid attempt to destructure non-iterable instance.\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method.",
    );
}

function _unsupportedIterableToArray(o, minLen) {
    if (!o) return;
    if (typeof o === "string") return _arrayLikeToArray(o, minLen);
    var n = Object.prototype.toString.call(o).slice(8, -1);
    if (n === "Object" && o.constructor) n = o.constructor.name;
    if (n === "Map" || n === "Set") return Array.from(o);
    if (n === "Arguments" || /^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(n))
        return _arrayLikeToArray(o, minLen);
}

function _arrayLikeToArray(arr, len) {
    if (len == null || len > arr.length) len = arr.length;
    for (var i = 0, arr2 = new Array(len); i < len; i++) {
        arr2[i] = arr[i];
    }
    return arr2;
}

function _iterableToArrayLimit(arr, i) {
    if (typeof Symbol === "undefined" || !(Symbol.iterator in Object(arr))) return;
    var _arr = [];
    var _n = true;
    var _d = false;
    var _e = undefined;
    try {
        for (var _i = arr[Symbol.iterator](), _s; !(_n = (_s = _i.next()).done); _n = true) {
            _arr.push(_s.value);
            if (i && _arr.length === i) break;
        }
    } catch (err) {
        _d = true;
        _e = err;
    } finally {
        try {
            if (!_n && _i["return"] != null) _i["return"]();
        } finally {
            if (_d) throw _e;
        }
    }
    return _arr;
}

function _arrayWithHoles(arr) {
    if (Array.isArray(arr)) return arr;
}

const traitCodeToName = {
    0: "battlecry", // has Battlecry keyword
    1: "charge", // has Charge keyword
    2: "chooseone", // has Choose One keyword
    3: "collectible", // is a collectible card
    4: "combo", // has Combo keyword
    5: "deathrattle", // has Deathrattle keyword
    6: "discover", // has Discover keyword
    7: "divineshield", // has Divine Shield keyword
    8: "doublespelldamage", // has the "Arcane Blast" effect
    9: "echo", // has Echo keyword
    10: "elite", // has an elite/dragon border; i.e. Legendary
    11: "invoke", // invoke Galakrond
    12: "enrage", // has "enrage" effect (defunct keyword)
    13: "forgetful", // has ogre 50% chance to miss effect
    14: "freeze", // has Freeze keyword
    15: "galakrond", // is a Galakrond hero card
    16: "heropowerdamage", // has the "Fallen Hero" effect
    17: "immune", // has Immune keyword
    18: "inspire", // has Inspire keyword
    19: "lifesteal", // has Lifesteal keyword
    20: "lackey", // is a lackey
    21: "magnetic", // has Magnetic keyword
    22: "notarget", // can't be targetted
    25: "overkill", // has Overkill keyword
    26: "poisonous", // has Poisonous keyword
    27: "realquest", // is a Quest card
    28: "reborn", // has Reborn keyword
    29: "buffscthun", // buffs C'Thun
    30: "rush", // has Rush keyword
    31: "secret", // has Secret keyword
    32: "sidequest", // is a Side Quest card
    33: "silence", // has Silence keyword
    34: "sparepart", // is a Spare Part card
    35: "startofgame", // has a start-of-game trigger
    36: "stealth", // has Stealth keyword
    37: "taunt", // has Taunt keyword
    38: "twinspell", // has Twinspell keyword
};

const traits = Object.values(traitCodeToName);

var keywordTypes = (function() {
    var types = {};

    traits.forEach(function(kw) {
        types[kw] = Boolean;
    });

    types["quest"] = Boolean; // special case

    ["armor", "attack", "durability", "health", "mana", "overload", "spelldamage"].forEach(function(
        kw,
    ) {
        types[kw] = Number;
    });
    ["class", "description", "name", "rarity", "text", "type", "usable"].forEach(function(kw) {
        types[kw] = String;
    });
    return types;
})(); // Set of all recognized comparison keywords.

var validKeywords = Object.keys(keywordTypes); // Converts an abbreviated or prefix keyword to its canonical name.

var canonicalKeyword = function canonicalKeyword(kw) {
    // const foo = (kw) => {
    //     // Handle single-letter abbreviations.
    //     switch (kw) {
    //         case "a":
    //             return "attack";
    //         case "c":
    //             return "class";
    //         case "h":
    //             return "health";
    //         case "m":
    //             return "mana";
    //         case "t":
    //             return "type";
    //         case "u":
    //             return "usable";
    //     }
    //     if (keywordTypes.hasOwnProperty(kw)) {
    //         return kw; // exact match
    //     }
    //     for (const k of validKeywords) {
    //         if (k.startsWith(kw)) {
    //             return k; // prefix match
    //         }
    //     }
    //     return kw; // unrecognized
    // };
    // const answer = foo(kw);
    // console.log(`DEBUG: canonicalKeyword("${kw}") => "${answer}"`);
    // return answer;
    // Handle single-letter abbreviations.
    console.log('DEBUG: canonicalKeyword("'.concat(kw, '")'));

    switch (kw) {
        case "a":
            return "attack";

        case "c":
            return "class";

        case "h":
            return "health";

        case "m":
            return "mana";

        case "t":
            if (kw !== "t") {
                throw new Error('*** ERROR: kw = "'.concat(kw, '", matched "type" in switch'));
            }

            return "type";

        case "u":
            return "usable";
    } // switch (true) {
    //     case kw === "a":
    //         return "attack";
    //     case kw === "c":
    //         return "class";
    //     case kw === "h":
    //         return "health";
    //     case kw === "m":
    //         return "mana";
    //     case kw === "t":
    //         console.log(`DEBUG: kw = "${kw}", matched "type" in switch`);
    //         return "type";
    //     case kw === "u":
    //         return "usable";
    // }
    // if (kw === "a") {
    //     return "attack";
    // } else if (kw === "c") {
    //     return "class";
    // } else if (kw === "h") {
    //     return "health";
    // } else if (kw === "m") {
    //     return "mana";
    // } else if (kw === "t") {
    //     console.log(`DEBUG: kw = "${kw}", matched "type" in switch`);
    //     return "type";
    // } else if (kw === "u") {
    //     return "usable";
    // }

    if (keywordTypes.hasOwnProperty(kw)) {
        console.log('DEBUG: kw = "'.concat(kw, '", returning as-is'));
        return kw; // exact match
    }

    var _iterator = _createForOfIteratorHelper(validKeywords),
        _step;

    try {
        for (_iterator.s(); !(_step = _iterator.n()).done; ) {
            var k = _step.value;

            if (k.startsWith(kw)) {
                console.log('DEBUG: kw = "'.concat(kw, '", returning "').concat(k, '"'));
                return k; // prefix match
            }
        }
    } catch (err) {
        _iterator.e(err);
    } finally {
        _iterator.f();
    }

    console.log('DEBUG: kw = "'.concat(kw, '", not recognized'));
    return kw; // unrecognized
};

let tries = 0;
const numTries = 1000;
const bug = () => {
    const answer = canonicalKeyword(tries % 2 === 0 ? "ta" : "t");
    document.getElementById("target").innerHTML = answer;
    tries++;
    console.log(`DEBUG: tries = ${tries}`);
    if (tries < numTries) {
        setTimeout(bug, 10);
    }
};
setTimeout(bug, 1000);
