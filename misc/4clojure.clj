;; Last Element
;;
;; Write a function which returns the last element in a
;; sequence.

(defn my-last [s]
  (if (empty? (next s))
    (first s)
    (recur (next s))))

;; Penultimate Element
;;
;; Write a function which returns the second to last element
;; from a sequence.

(fn [s]
  (if (empty? (next (next s)))
    (first s)
    (recur (next s))))

;; Sum It All Up
;;
;; Write a function which returns the sum of a sequence of
;; numbers.

(partial reduce +)

;; Find the odd numbers
;;
;; Write a function which returns only the odd numbers from
;; a sequence.

(partial filter odd?)

;; http://4clojure.com/problem/21
;; Nth Element
;;
;; Write a function which returns the Nth element from a
;; sequence.

(fn [s n]
  (cond (empty? s) nil
        (<= n 0) (first s)
        :default (recur (next s) (dec n))))

;; http://4clojure.com/problem/22
;; Count a Sequence
;;
;; Write a function which returns the total number of
;; elements in a sequence.

#(reduce + (map (constantly 1) %))

;; http://4clojure.com/problem/23
;; Reverse a Sequence
;;
;; Write a function which reverses a sequence.

#(into () %)

;; http://4clojure.com/problem/26
;; Fibonacci Sequence
;;
;; Write a function which returns the first X fibonacci
;; numbers.

(fn [n]
  (take n (map first (iterate (fn [[a b]] [b (+ a b)]) [1 1]))))

;; http://4clojure.com/problem/27
;; Palindrome Detector
;;
;; Write a function which returns true if the given sequence
;; is a palindrome.

(fn [s]
  (if (string? s)
    (= (apply str (reverse s)) s)
    (= (reverse s) s)))

;; http://4clojure.com/problem/28
;; Flatten a Sequence
;;
;; Write a function which flattens a sequence.

(fn [s0]
  (let [helper (fn flt [coll acc]
                 (if (empty? coll)
                   acc
                   (let [left (first coll) right (next coll)]
                     (recur right
                            (if (sequential? left)
                              (into acc (flt left []))
                              (conj acc left))))))]
    (helper s0 [])))

;; http://4clojure.com/problem/29
;; Get the Caps
;;
;; Write a function which takes a string and returns a new
;; string containing only the capital letters.

(fn [s] (apply str (filter #(Character/isUpperCase %) s)))

;; http://4clojure.com/problem/30
;; Compress a Sequence
;;
;; Write a function which removes consecutive duplicates
;; from a sequence.

(fn compress [coll]
  (when-let [[f & r] coll]
    (if (= f (first r))
      (compress r)
      (cons f (compress r)))))

;; http://4clojure.com/problem/31
;; Pack a Sequence
;;
;; Write a function which packs consecutive duplicates into
;; sub-lists.

#(partition-by identity %)

;; http://4clojure.com/problem/32
;; Duplicate a Sequence
;;
;; Write a function which duplicates each element of a
;; sequence.

#(mapcat (fn [x] (list x x)) %)

;; http://4clojure.com/problem/33
;; Replicate a Sequence
;;
;; Write a function which replicates each element of a
;; sequence a variable number of times.

#(mapcat (partial repeat %2) %1)

;; http://4clojure.com/problem/34
;; Implement range
;;
;; Write a function which creates a list of all integers in
;; a given range.

(fn r [a b]
  (let [n (inc a)]
    (if (< a b)
      (cons a (r n b))
      ())))

;; http://4clojure.com/problem/38
;; Maximum Value
;;
;; Write a function which takes a variable number of
;; parameters and returns the maximum value.

(fn [& args] (reduce (fn [a b] (if (> a b) a b)) args))

;; http://4clojure.com/problem/39
;; Interleave Two Seqs
;;
;; Write a function which takes two sequences and returns
;; the first item from each, then the second item from each,
;; then the third, etc.

#(mapcat vector %1 %2)

;; http://4clojure.com/problem/40
;; Interpose a Seq
;;
;; Write a function which separates the items of a sequence
;; by an arbitrary value.

(fn [sep coll] (drop-last (mapcat vector coll (repeat sep))))

;; http://4clojure.com/problem/41
;; Drop Every Nth Item
;;
;; Write a function which drops every Nth item from a
;; sequence.

(fn [coll n]
  (map first (filter #(not= (second %) n)
                     (map vector coll (cycle (range 1 (inc n)))))))

;; http://4clojure.com/problem/42
;; Factorial Fun
;;
;; Write a function which calculates factorials.

(fn [n] (reduce * (range 2 (inc n))))

;; http://4clojure.com/problem/43
;; Reverse Interleave
;;
;; Write a function which reverses the interleave process
;; into x number of subsequences.

(fn [coll n]
  (apply map list (partition n coll)))

;; http://4clojure.com/problem/44
;; Rotate Sequence
;;
;; Write a function which can rotate a sequence in either
;; direction.

(fn [n coll]
  (let [offset (mod n (count coll))]
    (concat (drop offset coll) (take offset coll))))

;; http://4clojure.com/problem/46
;; Flipping out
;;
;; Write a higher-order function which flips the order of
;; the arguments of an input function.

(fn [f]
  (fn [& args]
    (apply f (reverse args))))

;; http://4clojure.com/problem/49
;; Split a Sequence
;;
;; Write a function which will split a sequence into two parts.

(fn [n coll] (vector (take n coll) (drop n coll)))

;; http://4clojure.com/problem/50
;; Split by Type
;;
;; Write a function which takes a sequence consisting of
;; items with different types and splits them up into a set
;; of homogeneous sub-sequences. The internal order of each
;; sub-sequence should be maintained, but the sub-sequences
;; themselves can be returned in any order (this is why
;; 'set' is used in the test cases).

#(vals (group-by type %))

;; http://4clojure.com/problem/53
;; Longest Increasing Sub-Seq
;;
;; Given a vector of integers, find the longest consecutive
;; sub-sequence of increasing numbers. If two sub-sequences
;; have the same length, use the one that occurs first. An
;; increasing sub-sequence must have a length of 2 or
;; greater to qualify.

;; Too ugly
(fn [v]
  (if (empty? v)
    []
    (let [n (count v)
          f (fn [start i previous [best-start best-len :as best]]
              (let [new-best (fn []
                               (let [new-len (- i start)]
                                 (if (and (> new-len 1) (< best-len (- i start)))
                                   [start (- i start)]
                                   [best-start best-len])))]
                (if (= i n)
                  (let [[start length] (new-best)]
                    (subvec v start (+ start length)))
                  (let [current (nth v i)]
                    (if (> current previous)
                      (recur start (inc i) current best)
                      (recur i (inc i) current (new-best)))))))]
      (f 0 0 (nth v 0) [0 0]))))

;; Cleaner, but inefficient
(fn [coll]
  (reduce
   (fn [a b]
     (let [x (count a) y (count b)]
       (cond (<= x y 1) []
             (>= x y) a
             :default b)))
   (next (reductions
          (fn [a b]
            (if (sequential? a)
              (if (> b (last a)) (conj a b) [b])
              (if (> b a) [a b] [b])))
          coll))))


;; http://4clojure.com/problem/54
;; Partition a Sequence
;;
;; Write a function which returns a sequence of lists of x
;; items each. Lists of less than x items should not be
;; returned.

(fn my-partition [n coll]
  (let [head (take n coll)]
    (if (= (count head) n)
      (cons head (my-partition n (drop n coll)))
      ())))

;; http://4clojure.com/problem/55
;; Count Occurrences
;;
;; Write a function which returns a map containing the
;; number of occurences of each distinct item in a sequence.

(fn [coll0]
  (let [helper (fn [acc coll]
                 (if (empty? coll)
                   acc
                   (let [hd (first coll)]
                     (recur (assoc acc hd (inc (get acc hd 0)))
                            (next coll)))))]
    (helper {} coll0)))

;; http://4clojure.com/problem/56
;; Find Distinct Items
;;
;; Write a function which removes the duplicates from a
;; sequence. Order of the items must be maintained.

(fn [coll]
  (first
   (reduce (fn [[acc seen] a]
             (if (seen a)
               [acc seen]
               [(conj acc a) (conj seen a)]))
           [[] #{}] coll)))

;; http://4clojure.com/problem/58
;; Function Composition
;;
;; Write a function which allows you to create function
;; compositions. The parameter list should take a variable
;; number of functions, and create a function applies them
;; from right-to-left.

(fn my-comp [f & fns]
  (if (not (empty? fns))
    (fn [& xs] (f (apply (apply my-comp fns) xs)))
    (fn [& xs] (apply f xs))))

;; http://4clojure.com/problem/59
;; Juxtaposition
;;
;; Take a set of functions and return a new function that
;; takes a variable number of arguments and returns a
;; sequence containing the result of applying each function
;; left-to-right to the argument list.

(fn [& fs]
  (fn [& args]
    (reduce #(conj %1 (apply %2 args)) [] fs)))

;; http://4clojure.com/problem/60
;; Sequence Reductions
;;
;; Write a function which behaves like reduce, but returns
;; each intermediate value of the reduction. Your function
;; must accept either two or three arguments, and the return
;; sequence must be lazy.

(fn my-reductions
  ([f coll]
     (if-let [s (seq coll)]
       (my-reductions f (first s) (rest s))
       (list (f))))
  ([f init coll]
     (cons init
           (lazy-seq
            (when-let [s (seq coll)]
              (my-reductions f (f init (first s)) (rest s)))))))

;; http://4clojure.com/problem/61
;; Map Construction
;;
;; Write a function which takes a vector of keys and a
;; vector of values and constructs a map from them.

(fn my-zipmap [keys vals]
  (loop [map {}
         ks (seq keys)
         vs (seq vals)]
    (if (and ks vs)
      (recur (assoc map (first ks) (first vs))
             (next ks)
             (next vs))
      map)))

;; or: #(into {} (map vector %1 %2))

;; http://4clojure.com/problem/62
;; Re-implement Iterate
;;
;; Given a side-effect free function f and an initial value
;; x write a function which returns an infinite lazy
;; sequence of x, (f x), (f (f x)), (f (f (f x))), etc.

(fn my-iterate [f x]
  (cons x (lazy-seq (my-iterate f (f x)))))

;; http://4clojure.com/problem/63
;; Group a Sequence
;;
;; Given a function f and a sequence s, write a function
;; which returns a map. The keys should be the values of f
;; applied to each item in s. The value at each key should
;; be a vector of corresponding items in the order they
;; appear in s.

(fn [f coll]
  (loop [map {} s (seq coll)]
    (if s
      (let [v (first s) k (f v)]
        (recur (assoc map k (conj (get map k []) v)) (next s)))
      map)))

;; http://4clojure.com/problem/65
;; Black Box Testing
;;
;; Clojure has many sequence types, which act in subtly
;; different ways. The core functions typically convert them
;; into a uniform "sequence" type and work with them that
;; way, but it can be important to understand the behavioral
;; and performance differences so that you know which kind
;; is appropriate for your application.
;;
;; Write a function which takes a collection and returns one
;; of :map, :set, :list, or :vector - describing the type of
;; collection it was given.
;;
;; You won't be allowed to inspect their class or use the
;; built-in predicates like list? - the point is to poke at
;; them and understand their behavior.

;; http://4clojure.com/problem/66
;; Greatest Common Divisor
;;
;; Given two integers, write a function which returns the
;; greatest common divisor.

(fn gcd [a b] (if (= b 0) a (gcd b (mod a b))))

;; http://4clojure.com/problem/67
;; Prime Numbers
;;
;; Write a function which returns the first x number of
;; prime numbers.

(fn [n]
  (take n (letfn [(sieve [s]
                    (let [x (first s)]
                      (cons x (lazy-seq (filter #(not (zero? (mod % x))) (sieve (next s)))))))]
            (cons 2 (sieve (iterate (partial + 2) 3))))))

;; http://4clojure.com/problem/69
;; Merge with a Function
;;
;; Write a function which takes a function f and a variable
;; number of maps. Your function should return a map that
;; consists of the rest of the maps conj-ed onto the first.
;; If a key occurs in more than one map, the mapping(s) from
;; the latter (left-to-right) should be combined with the
;; mapping in the result by calling (f val-in-result
;; val-in-latter)

(fn [f & maps]
  (let [merge (fn [m e]
                (let [k (key e) v (val e)]
                  (if (contains? m k)
                    (assoc m k (f (m k) v))
                    (assoc m k v))))]
    (reduce (fn [m1 m2] (reduce merge m1 (seq m2))) maps)))

;; http://4clojure.com/problem/70
;; Word Sorting
;;
;; Write a function which splits a sentence up into a sorted
;; list of words. Capitalization should not affect sort
;; order and punctuation should be ignored.

(fn [s]
  (sort #(compare (.toLowerCase %1) (.toLowerCase %2))
        (re-seq #"[A-Za-z]+" s)))

;; http://4clojure.com/problem/73
;; Analyze a Tic-Tac-Toe Board
;;
;; A tic-tac-toe board is represented by a two dimensional
;; vector. X is represented by :x, O is represented by :o,
;; and empty is represented by :e. A player wins by placing
;; three Xs or three Os in a horizontal, vertical, or
;; diagonal row. Write a function which analyzes a
;; tic-tac-toe board and returns :x if X has won, :o if O
;; has won, and nil if neither player has won.

(fn [rows]
  (let [columns (apply map list rows)
        two-d (fn [i j] (nth (nth rows i) j))
        diagonals [[(two-d 0 0) (two-d 1 1) (two-d 2 2)]
                   [(two-d 0 2) (two-d 1 1) (two-d 2 0)]]
        lines (concat rows columns diagonals)
        won? (fn [player]
               (some (fn [line] (every? #(= player %) line)) lines))]
    (cond (won? :x) :x
          (won? :o) :o)))

;; http://4clojure.com/problem/74
;; Filter Perfect Squares
;;
;; Given a string of comma separated integers, write a
;; function which returns a new comma separated string that
;; only contains the numbers which are perfect squares.

(fn [s]
  (letfn [(square? [n] (= n (int (Math/pow (int (Math/sqrt n)) 2))))
          (parse [s] (map #(Integer/parseInt %) (re-seq #"\d+" s)))]
    (->> s parse (filter square?) (interpose ",") (apply str))))

;; http://4clojure.com/problem/75
;; Euler's Totient Function
;;
;; Two numbers are coprime if their greatest common divisor
;; equals 1. Euler's totient function f(x) is defined as the
;; number of positive integers less than x which are coprime
;; to x. The special case f(1) equals 1. Write a function
;; which calculates Euler's totient function.

(fn [x]
  (letfn [(gcd [a b] (if (zero? b) a (gcd b (mod a b))))
          (coprime? [a b] (= 1 (gcd a b)))]
    (count (filter (partial coprime? x) (range 1 (inc x))))))

;; http://4clojure.com/problem/77
;; Anagram Finder
;;
;; Write a function which finds all the anagrams in a vector
;; of words. A word x is an anagram of word y if all the
;; letters in x can be rearranged in a different order to
;; form y. Your function should return a set of sets, where
;; each sub-set is a group of words which are anagrams of
;; each other. Each sub-set should have at least two words.
;; Words without any anagrams should not be included in the
;; result.

;; ["meat" "mat" "team" "mate" "eat"]
;; ["veer" "lake" "item" "kale" "mite" "ever"]

(fn [words]
  (loop [anagrams {} more words]
    (if-let [s (seq more)]
      (let [word (first s) key (sort (seq word))]
        (recur (assoc anagrams key (conj (get anagrams key #{}) word))
               (next s)))
      (set (filter #(> (count %) 1) (vals anagrams))))))

;; http://4clojure.com/problem/78
;; Reimplement Trampoline
;;
;; Reimplement the function described in "Intro to
;; Trampoline".

(fn my-trampoline
  ([f]
     (let [ret (f)]
       (if (fn? ret)
         (recur ret)
         ret)))
  ([f & args]
     (my-trampoline #(apply f args))))

;; http://4clojure.com/problem/79
;; Triangle Minimal Path
;;
;; Write a function which calculates the sum of the minimal
;; path through a triangle. The triangle is represented as a
;; vector of vectors. The path should start at the top of
;; the triangle and move to an adjacent number on the next
;; row until the bottom of the triangle is reached.

(fn [triangle]
  (let [max-row (dec (count triangle))
        get (fn [i j] (nth (nth triangle i) j))]
    (with-local-vars
        [search (memoize
                 (fn [i j]
                   (if (= i max-row)
                     (get i j)
                     (+ (get i j)
                        (min (search (inc i) j)
                             (search (inc i) (inc j)))))))]
      (search 0 0))))

;; http://4clojure.com/problem/80
;; Perfect Numbers
;;
;; A number is "perfect" if the sum of its divisors equal
;; the number itself. 6 is a perfect number because 1+2+3=6.
;; Write a function which returns true for perfect numbers
;; and false otherwise.

(fn [n]
  (let [divisors (filter #(zero? (rem n %)) (range 1 (inc (/ n 2))))]
    (= (reduce + divisors) n)))

;; http://4clojure.com/problem/81
;; Set Intersection
;;
;; Write a function which returns the intersection of two
;; sets. The intersection is the sub-set of items that each
;; set has in common.

(fn [a b] (set (filter #(a %) (seq b))))

;; http://4clojure.com/problem/82
;; Word Chains
;;
;; A word chain consists of a set of words ordered so that
;; each word differs by only one letter from the words
;; directly before and after it. The one letter difference
;; can be either an insertion, a deletion, or a
;; substitution. Here is an example word chain:
;;
;; cat -> cot -> coat -> oat -> hat -> hot -> hog -> dog
;;
;; Write a function which takes a sequence of words, and
;; returns true if they can be arranged into one continous
;; word chain, and false if they cannot.

(defn one-letter-difference? [w1 w2]
  (let [s (seq w1) s-len (count s) t (seq w2) t-len (count t)
        changes (fn [n s t]
                  (if (empty? s)
                    true
                    (if (= (first s) (first t))
                      (recur n (next s) (next t))
                      (if (zero? n)
                        false
                        (recur (dec n) (next s) (next t))))))
        insertions (fn [n s t]
                     (if (empty? s)
                       true
                       (if (= (first s) (first t))
                         (recur n (next s) (next t))
                         (if (zero? n)
                           false
                           (recur (dec n) s (next t))))))]
    (cond (= s-len t-len)       (changes    1 s t)
          (= s-len (inc t-len)) (insertions 1 t s)
          (= s-len (dec t-len)) (insertions 1 s t))))

(defn continuous-word-chain? [words]
  (let [s (seq words)]
