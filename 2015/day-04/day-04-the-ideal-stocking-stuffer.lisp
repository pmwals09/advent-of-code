(require :asdf)
(require :sb-md5)

(defun hash-string (str)
  (let ((result ""))
    (map ()
       (lambda (byte) (setf result (format nil "~A~A" result (format nil "~2,'0x" byte))))
       (sb-md5:md5sum-string str :external-format :utf-8))
    result))

(defun make-hash-string (seed suffix)
  (format nil "~A~D" seed suffix))

(with-open-file (f "./day-04-data.txt" :direction :input)
  (let* ((seed (read-line f))
         (suffix 0)
         (part-one-found nil)
         (hashed-val (hash-string (make-hash-string seed suffix))))
    (loop while (not (string= (subseq hashed-val 0 6) "000000"))
          do (when (and (null part-one-found) 
                        (string= (subseq hashed-val 0 5) "00000"))
               (setf part-one-found t)
               (format t "Part one: ~A~%" suffix))
          do (incf suffix)
          do (setf hashed-val (hash-string (make-hash-string seed suffix))))
    (format t "Part two: ~A~%" suffix)))
