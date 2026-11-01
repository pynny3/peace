#!/usr/bin/env sbcl --script

;;; Minimal Peace interpreter hosted in Common Lisp.

(defparameter *environment* (make-hash-table :test #'equal))

(defun trim (text)
  (string-trim '(#\Space #\Tab #\Return #\Newline) text))

(defun split-on-equals (text)
  (let ((position (position #\= text)))
    (and position
         (list (trim (subseq text 0 position))
               (trim (subseq text (1+ position)))))))

(defun replace-variables (text)
  (let ((result text))
    (maphash
     (lambda (name value)
       (setf result
             (with-output-to-string (output)
               (loop with start = 0
                     for position = (search name result :start2 start)
                     do (if position
                            (progn
                              (write-string result output :start start :end position)
                              (format output "~a" value)
                              (setf start (+ position (length name))))
                            (progn
                              (write-string result output :start start)
                              (return)))))))
     *environment*)
    result))

(defun evaluate-expression (text)
  (let ((expression (replace-variables (trim text))))
    (handler-case
        (eval (read-from-string expression))
      (error ()
        (error "Peace: invalid expression: ~a" text)))))

(defun set-variable (name value type)
  (when (and type (not (typep value (read-from-string type))))
    (error "Peace: ~a is not of type ~a" name type))
  (setf (gethash name *environment*) value)
  value)

(defun run-line (line)
  (let ((line (trim line)))
    (cond
      ((or (zerop (length line)) (char= (char line 0) #\;)) nil)
      ((uiop:string-prefix-p "print " line)
       (format t "~a~%" (evaluate-expression (subseq line 6))))
      ((uiop:string-prefix-p "set " line)
       (let* ((parts (split-on-equals (subseq line 4)))
              (name (first parts))
              (right (second parts)))
         (unless (and name right)
           (error "Peace: expected `set name = expression`"))
         (let* ((type-position (search ": type " right))
                (expression (if type-position
                               (subseq right 0 type-position)
                               right))
                (type (and type-position
                           (trim (subseq right (+ type-position 7))))))
           (set-variable name (evaluate-expression expression) type))))
      (t (error "Peace: unknown statement: ~a" line)))))

(defun run-file (path)
  (with-open-file (stream path)
    (loop for line = (read-line stream nil nil)
          while line
          do (run-line line))))

(when (= (length sb-ext:*posix-argv*) 2)
  (run-file (second sb-ext:*posix-argv*)))

(defun run-repl ()
  (format t "Peace REPL (:quit end)~%")
  (loop
    (format t "peace> ")
    (finish-output)
    (let ((line (read-line *STANDARD-INPUT* nil nil)))
      (unless line
        (return))
      (if (string= (trim line) ":quit")
          (return)
          (handler-case
              (run-file line)
            (error (condition)
               (format *ERROR-OUTPUT* "Error: ~a~%" condition)))))))

(load "lexer.lisp")
(load "parser.lisp")
(load "evaluator.lisp")
(load "go-generator.lisp")

(defun run-peace-file (path)
   (let* ((source (read-file-string path))
          (tokens (tokenize source))
          (ast (parse-program tokens)))
      (evaluate-program ast)))

(defun compile-peace-file (source-path go-path)
  (let* ((source (read-file-string source-path))
         (tokens (tokenize source))
         (ast (parse-program tokens)))
    (emit-go ast go-path)))