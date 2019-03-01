(defn prime?
  [n]
  (not (some #(= 0 (rem n %)) (range 3 (inc (Math/sqrt n)) 2))))

(let [n 19] (take 10 (range 3 (inc (Math/sqrt n)) 2)))

(nth (filter prime? (iterate inc 1)) 6)
