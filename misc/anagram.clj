(def dictionary-filename "/usr/share/dict/web2")

;; read the dictionary and exclude proper names
(def all-words
  (filter #(not (Character/isUpperCase (first %)))
          (clojure.contrib.duck-streams/read-lines dictionary-filename)))

(def dictionary (group-by sort all-words))

(defn subwords [coll]
  (let [f (fn f [word]
            (if (<= (count word) 3)
              (list word)
              (let [drop-nth (fn [n coll] (concat (take (dec n) coll) (drop n coll)))]
                (cons word (apply concat (map f (map #(drop-nth % word)
                                                     (range 1 (inc (count word))))))))))]
    (sort (fn [a b] (compare (count b) (count a)))
          (distinct (f (sort coll))))))

(defn anagrams [word]
  (distinct (filter (complement nil?) (flatten (map dictionary (subwords word))))))
