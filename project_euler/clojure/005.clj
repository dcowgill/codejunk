;; "Elapsed time: 6289.962 msecs"
(time (letfn [(f [n] (every? #(= (rem n %) 0) [19 18 17 16 15 14 13 12 11]))]
        (first (filter f (iterate (partial + 20) 20)))))
