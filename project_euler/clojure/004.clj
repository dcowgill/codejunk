(defn palindrome? [n]
  (let [s (str n)]
    (= (apply str (reverse s)) s)))

(def products-of-three-digit-numbers (for [x (range 100 1000) y (range x 1000)] (* x y)))

(def palindromes (filter palindrome? products-of-three-digit-numbers))

;; "Elapsed time: 867.239 msecs"
(time (apply max palindromes))
