(let [divisible? (fn [a b] (zero? (rem a b)))]
  (reduce + (filter #(or (divisible? % 3) (divisible? % 5)) (range 1000))))
