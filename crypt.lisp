crypt.lisp

(defun genrete-crypt-go (&optional (path "crypt.go"))
  (with-open-file (out path
                       :direction :output
                       :if-exists :supersede
                       :if-does-not-exist :create)
      (format out "package crypt~%~
 import (~%~
   \"crypto/aes\"~%~
   \"crypto/cipher\"~%~
   \"crypto/rand\"~%~
   \"io\"~%~
 )~%~
 func Encrypto(key, plaintext []byte) ([]byte,error) {~%~
     block, err := aes.NewCipher(key)~%~
     if err != nil { return nil, err }~%~
     gcm, err := cipher.NewGCM(block)~%~
     if err != nil { return nil, err }~%~
     nonce := make([]byte, gcm.NonceSize())~%~
     if _, err := io.ReadFull(rand.Reader, nonce); err != nil { return nil, err}~%~
     return gcm.Seal(nonce, nonce, plaintext, nil), nil~%~
 }~%~")))

(genrete-crypt-go)