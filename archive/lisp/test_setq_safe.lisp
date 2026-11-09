;;; tests/test_setq_safe.lisp
;;; Simple test script. Runs with SBCL or CLISP.

(defun run-setq-safe-tests ()
  (format t "Running setq-safe tests...~%")
  ;; Clean up.
  (when (boundp 'x) (makunbound 'x))
  (when (boundp 'y) (makunbound 'y))
  (when (boundp 'z) (makunbound 'z))

  ;; Successful cases.
  (setq-safe x 10 :type integer :doc "counter" :debug nil)
  (unless (= x 10) (error "test failed: x not 10"))
  (unless (string= (get 'x 'documentation) "counter") (error "test failed: doc not stored"))

  (setq-safe y "hello" :type string)
  (unless (string= y "hello") (error "test failed: y value"))

  ;; Failure case: a type mismatch should signal an error.
  (let ((caught nil))
    (handler-case
        (progn (setq-safe z 3.14 :type integer) (setf caught nil))
      (error (e) (setf caught t)))
    (unless caught (error "test failed: expected type mismatch error")))

  (format t "All tests passed.~%"))

;; Run the tests.
(run-setq-safe-tests)
