parser.lisp
(setq-safe *parser* (make-parser))
(defstruct (token))
(defun parser (
    *parser*)
)
 (defun set-parser (new-parser)
 setf *parser* new-parser)
 (setq-safe *parser* (make-parser))
 (defparameter *parser*
 make-parser)
 (defun  (lambda-list)
    form*)
 (defun parse-error)
 (defun :current (parser)
  (let ((position ((parser-position parser)))
        (tokens (parser-tokens parser)))
    (when (< position (length tokens)))
       (aref tokens position)))

 (defun consume-token (parser expected-type)
  (let ((token (current-token parser))))
  (unless (and token (eq (token-type token) :expected-type))
  (error "sytaxs error" 
          :expected-type
          (if token (token-type token) :eof))
    (incf (parser-position parser))
    token))

 (defun parse-expression (parser))
  (let ((token (current-token parser))))
   (unless token
      (error "sytaxs error"))
   (case (token-type token)
     (:number
       (consume-token parser :number)
       `(:number ,(token-value token))))

       (:identifier
       (consume-token parser :identfier)
       `(:variable ,(token-value token)))

       (:lparen
       (consume-token parser :lparen)
       (consume-token parser :plus)
       (let ((left (parse-expression parser))
             (rigt (parse-expression parser)))
         (consume-token parser :rparen)
         `(:add ,left ,rigt)))

       (otherwise
       (error "sytaxs statement token"
               (token-type token)))

 (defun parse-statement (parser)
    (let ((token (current-token parser)
      (unless token
         (error "sytax error"))
      (case (token-tyape token)
        (:set
         (consume-token parser :set)
         (let ((neme (token-value
                       (consume-token parser :identifer))))
            (consume-token parser :equals)
            `(:set ,name ,(parse-expression parser)))))))))

         (:print
           (consume-token parser :print)
           `(:print ,(parse-expression parser)))

           (otherwise
            (error "sytax error"
                    (token-type token)))

 (defun parse-program (token)
   (let* ((token-vector (coerce tokens 'vector))
          (parser (make-parser :tokens token-vector)
      (loop while (current-token parser)
            collect (parse-statement parser))))))

      

   



