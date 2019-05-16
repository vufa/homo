#! /bin/sh

cd homo/ && \
   rm -f homo.vocab && \
   text2wfreq < homo.txt | wfreq2vocab > homo.vocab && \
   rm -f homo.idngram && \
   text2idngram -vocab homo.vocab -idngram homo.idngram < homo.txt && \
   rm -f homo.arpa && \
   idngram2lm -vocab_type 0 -idngram homo.idngram -vocab homo.vocab -arpa homo.arpa && \
   rm -f homo.lm.bin && \
   sphinx_lm_convert -i homo.arpa -o homo.lm.bin
