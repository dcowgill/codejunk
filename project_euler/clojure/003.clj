(defn primes []
  (letfn [(p? [n] (not (some #(= 0 (rem n %)) (range 3 (inc (Math/sqrt n)) 2))))]
    (cons 2 (filter p? (iterate (partial + 2) 3)))))

(defn prime-factors [n]
  (loop [factors [], ps (primes), x n]
    (if (<= x 1)
      factors
      (let [ps1 (drop-while #(and (not (divisible? x %)) (<= % x)) ps)
            p (first ps1)]
        (if (and (<= p x) (divisible? x p))
          (recur (conj factors p) (next ps1) (/ x p))
          factors)))))

(last (prime-factors 600851475143))
