(defn square [n] (* n n))

(let [ns (take 100 (iterate inc 1))]
  (Math/abs
   (- (square (apply + ns))
      (apply + (map square ns)))))
