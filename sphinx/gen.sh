#! /bin/sh

cd aiicy/ && \
   rm -f aiicy.vocab && \
   text2wfreq < aiicy.txt | wfreq2vocab > aiicy.vocab && \
   rm -f aiicy.idngram && \
   text2idngram -vocab aiicy.vocab -idngram aiicy.idngram < aiicy.txt && \
   rm -f aiicy.arpa && \
   idngram2lm -vocab_type 0 -idngram aiicy.idngram -vocab aiicy.vocab -arpa aiicy.arpa && \
   rm -f aiicy.lm.bin && \
   sphinx_lm_convert -i aiicy.arpa -o aiicy.lm.bin
