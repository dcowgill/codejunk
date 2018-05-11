(defun main ()
  (write-line (hello-world)))

(defun hello-world ()
  (let ((username (getenv "USER")))
    (format nil "Hello, ~A!" username)))

(save-application "helloworld.bin"
                  :toplevel-function #'main
                  :prepend-kernel t
                  :error-handler :quit)
