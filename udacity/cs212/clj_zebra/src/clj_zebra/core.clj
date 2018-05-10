(ns clj_zebra.core
  (:use [clojure.math.combinatorics :only [permutations]]))

(defn right-of [a b]
  (= (- a b) 1))

(defn next-to [a b]
  (or (right-of a b) (right-of b a)))

(defn solve []
  (let [[first _ middle :as houses] (range 1 6)
        orderings (permutations houses)]
    (for [[englishman, japanese, norwegian, spaniard, ukrainian] orderings
          :when (= norwegian first) ; 10
          [blue, green, ivory, red, yellow] orderings
          :when (and (= englishman red) ; 2
                     (right-of green ivory) ; 6
                     (next-to norwegian blue)) ; 15
          [coffee, milk, OJ, tea, water] orderings
          :when (and (= green coffee) ; 4
                     (= ukrainian tea) ; 5
                     (= milk middle)) ; 9
          [chesterfield, kool, luckystrike, oldgold, parliament] orderings
          :when (and (= kool yellow) ; 8
                     (= luckystrike OJ) ; 13
                     (= japanese parliament)) ; 14
          [dog, fox, horse, snails, zebra] orderings
          :when (and (= spaniard dog) ; 3
                     (= oldgold snails) ; 7
                     (next-to kool horse) ; 12
                     (next-to chesterfield fox) ; 11
                     (next-to kool horse))] ; 12
      [water zebra])))

;; (def counts (atom {}))
;; (defn c [k]
;;   (fn [whatever]
;;     (swap! counts assoc k (inc (@counts k 0)))
;;     whatever))
