#! /bin/sh

gortana --hmm "cmusphinx-zh-cn-5.2/zh_cn.cd_cont_5000" \
        --dict "cmusphinx-zh-cn-5.2/zh_cn.dic" \
        --lm "cmusphinx-zh-cn-5.2/zh_cn.lm.bin"
