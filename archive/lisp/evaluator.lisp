evaluator.lisp

(defparameter *enviroment* (make-hash-table :test #'equal))

(defun evaluate-expression (ast)
  (case (first ast)
     (:number
      (second ast))))

      (:variable
       (multiple-value-bind (values found-p)
           (gethash (second ast) *enviroment*)
        (unless found-p
          (error "Undefined variable: ~a" (second ast)))
          values))
          
       (:add 
        (+ (evaluate-expression(second ast))
           (evaluate-expression (third ast))))

       (otherwise
        (error "Unsupported expression AST: ~a" ast))