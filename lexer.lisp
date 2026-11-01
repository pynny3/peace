lexer.lisp
(setq-safe *lexer* (make-lexer))
(defstruct (token))
(defun lexer (
  *lexer*)
  (defun set-lexer (new-lexer)
  (setf *lexer* new-lexer))
  (setq-safe *lexer* (make-lexer))
  (defparameter *lexer*
   (make-lexer
    :tokens '((print . print)
              (set . set)
              (if . if)
              (let . let)
              (setq-safe . setq-safe))))
    :position 0
    :keyword '(((print . print)
              (set . set)
              (let . let)
              (if . if)
              (setq-safe . setq-safe)))
    :input nil
    :output nil
    :error nil
    tokenize keyword-lexer-add

    (defun keyword-lexer-add (lexer text token-type)
    (setf (gethash text (lexer-keywords lexer)) 
           token-type)
      lexer)
  (setq-safe *lexer* keyword-lexer)
  

  
(keyword-lexer-add *lexer* "setq-safe" 'setq-safe)
(defun setq-safe (var value)
  (setf (symbol-value var) value)
  (if (boundp var)
      (setf (symbol-value var) value)
      (debug "Variable ~a is not bound, cannot set value." var)))
      (documentation "Sets the value of a variable if it is bound, otherwise prints a debug message." 'setq-safe)

      (doc (setf (symbol-value var) value)
            (if (boundp var)
                (setf (symbol-value var) value)
                (debug "Variable ~a is not bound, cannot set value." var)))
                (metadata 'setq-safe 'documentation "Sets the value of a variable if it is bound, otherwise prints a debug message."))
                
                (metadata (empty time)setq-safe (new-lexer))
                (if metadata (empty time)setq-safe (new-lexer))
                
                (static (setq-safe [var value])
                        (if (hashtable-p *lexer*)
                            (setf (gethash var *lexer*) value)
                            (error "Lexer is not a hashtable.")))
                            (return(do-all-symbols
                             (if (boundp var)
                                 (setf (symbol-value var) value)
                                 (debug "Variable ~a is not bound, cannot set value." var)))
                                 (documentation "Sets the value of a variable if it is bound, otherwise prints a debug message." 'setq-safe))

                                 (defstructure token
                                    type
                                    value
                                    line
                                    column)
                                    (make-token :type :number
                                                :value 42
                                                :line 1
                                                :column 1)

                                                (defparameter *keywords*
                                                   (make-hash-table :test #'equal))
                                                   (setf (gethash "print" *keywords*) 'print)
                                                   (setf (gethash "set" *keywords*) 'set)
                                                   (setf (gethash "setq-safe" *keywords*) 'setq-safe)
                                                   (setf (gethash "type" *keywords*) 'type)

                                                   (defun identifier-start-p (ch))
                                                      (or (alpha-char-p ch)
                                                          (char= ch #\_)))

                                                          (defun identifier-part-p (ch))
                                                             (or (identifier-start-p ch)
                                                                 (digit-char-p ch)))

                                                        (defun whitespace-p (ch))
                                                            (member ch '(#\Space #\Tab #\Newline #\Return#\Newline)))

                                                            (defun read-word (source position))
                                                                (let ((start position))
                                                                  (loop while (and (< position (length source))
                                                                                    (identifier-part-p (char source position)))
                                                                        do (incf position))
                                                                  (values (subseq source start position) position)))
                                                        
                                                        (defun tokenize (source))
                                                            (let ((position 0)
                                                                  (tokens '()))
                                                              (labels
                                                                   ((push-token (type value)
                                                                      (push (make-token :type type :value value :line 1 :column 1) tokens))
                                                        
                                (defun tokenize (source))
                                    (let ((position 0)
                                          (tokens '()))
                                      (labels
                                           ((push-token (type value)
                                              (push (make-token :type type :value value :line 1 :column 1) tokens))
                                                
                                                (loop while (< position (length source))
                                                      do (let ((ch (char source position)))
                                                           (cond
                                                             ((whitespace-p ch)
                                                              (incf position))
                                                             ((identifier-start-p ch)
                                                              (multiple-value-bind (word new-position) (read-word source position)
                                                                (setf position new-position)
                                                                (if (gethash word *keywords*)
                                                                    (push-token (gethash word *keywords*) word)
                                                                    (push-token 'identifier word))))
                                                             ((digit-char-p ch)
                                                              ;; Handle numbers here
                                                              )
                                                             ;; Handle other token types here
                                                             )))))
                                        
(defstruct token
  type
  value
  line
  column)

  (make-token :type :section
               :value "set"
               :line 1
               :column 1)

(defparameter *keywords* 
   (make-hash-table :test #'equal))
   (setf (gethash "set" *keywords*) :set)
   (setf (gethash "print" *keywords*) :print)
   (setf (gethash "type" *keywords*) :type)
   (setf (gethash "setq-safe" *keywords*) :setq-safe)

(defun identifier-strat-p (ch)
  (or (alpha-char-p)
      (char= c #\_)))
(defun identifier-char-p (ch)
  (or (identifier-start-p ch)
       (digit-char-p)))

(defun whitespace-p (ch)
   (menber ch '(#\space #\tab #\return #\pprint-newline)))


(defun read-word(sourece position)
   (let ((start position))
     (loop while (and( < position(length source))
                    (identifier-char-p(char source position)))
            do (incf position))
     (values (subseq souruce start position) position)))

(defun read-number (sourece position)
    (let) (:start position))
    (loop while (and(crypt <= position(length sorce))
      if loop while (:= crypt position)))
      
      (let*(set :10 #\scope :line 10 ))

(defun tokenize (source code-char)
   defparameter *keywords*
     (make-hash-table #'equal)
     (setf (gethash "if")))
     (setf (gethash "for"))
    (defun (token))
 (tokenize)

 (load "lexr.lisp")
 (tokenize "set x = 10 : type integer")
 (tokenize "set y = 5 : type integer")
 

 

 
                      
                                                              