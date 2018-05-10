(ns poker
  ;;:(:require [foo :as bar])
  (:use [clojure.string :only [split]]))

(import '(java.util ArrayList Collections))

(defn rank-to-num [r]
  (cond (= r "A") 14
        (= r "K") 13
        (= r "Q") 12
        (= r "J") 11
        (= r "T") 10
        true (Integer/parseInt r)))

(defn card-ranks [hand]
  (let [ranks (map (comp rank-to-num str first) hand)
        ranks (sort > ranks)]
    (if (= ranks [14 5 4 3 2])
      [5 4 3 2 1]
      ranks)))

(defn card-suits [hand]
  (map (comp str second) hand))

(defn straight [ranks]
  (and (= 4 (- (apply max ranks) (apply min ranks)))
       (= 5 (count (set ranks)))))

(defn flush [hand]
  (= 1 (count (set (card-suits hand)))))

(defn kind [n ranks]
  (let [freqs (frequencies ranks)
        answer (map first (filter (fn [[r cnt]] (= cnt n)) freqs))]
    (when (not (empty? answer))
      answer)))

(defn hand-rank [hand]
  (let [ranks (card-ranks hand)]
    (cond (and (straight ranks) (flush hand)) [8, ranks]
          (kind 4 ranks)                      [7, (kind 4 ranks), (kind 1 ranks)]
          (and (kind 3 ranks) (kind 2 ranks)) [6, (kind 3 ranks), (kind 2 ranks)]
          (flush hand)                        [5, ranks]
          (straight ranks)                    [4, ranks]
          (kind 3 ranks)                      [3, (kind 3 ranks), (kind 1 ranks)]
          (= 2 (count (kind 2 ranks)))        [2, (kind 2 ranks), (kind 1 ranks)]
          (kind 2 ranks)                      [1, (kind 2 ranks), (kind 1 ranks)]
          true                                [0, ranks])))

;; FIXME: max-key expects nums.
;; FIXME: doesn't respect ties.
(defn poker [hands]
  (apply max-key hand-rank hands))

(defn deal [numhands &{:keys [handsize deck] :or {handsize 5}}]
  (let [deck (if (nil? deck) (for [r "23456789TJQKA" s "SHDC"] (str r s)) deck)
        deck (vec (shuffle deck))]
    (for [i (range numhands)]
      (let [n (* i handsize)]
        (subvec deck n (+ n handsize))))))

(defn hand-frequencies [n]
  (let [hands (for [_ (range n)] (first (deal 1)))
        ranks (map hand-rank hands)
        types (map first ranks)
        freqs (frequencies types)]
    (doseq [i (reverse (range 9))]
      (printf "%d: %6.3f\n" i (/ (* 100.0 (freqs i 0)) n)))))

(defn test []
  (let [sf (parse-hand "7H 8H 9H TH JH")
        fk (parse-hand "6S 6H 6D 6C KS")
        fh (parse-hand "8H 8D 8S QH QC")]
    (assert (straight (card-ranks sf)))
    (assert (flush sf))
    (assert (= (kind 1 (card-ranks sf)) [11 10 9 8 7]))
    (assert (= (kind 4 (card-ranks fk)) [6]))
    (assert (= (kind 1 (card-ranks fk)) [13]))
    (assert (= (kind 3 (card-ranks fh)) [8]))
    (assert (= (kind 2 (card-ranks fh)) [12]))
    (assert (nil? (kind 1 (card-ranks fh))))
    (assert (nil? (kind 2 (card-ranks fk))))
    (assert (nil? (kind 3 (card-ranks sf))))
    (assert (= sf (poker [sf fk fh])))
    (assert (= fk (poker [fk fh])))
    (assert (= fh (poker [fh]))))
  "All tests passed.")
