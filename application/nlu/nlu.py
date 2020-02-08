#!/usr/bin/env python
# -*- coding:utf-8 -*-
"""
function to say hi in python
"""

import logging

import jieba
from rasa.nlu.model import Interpreter

logging.getLogger('tensorflow').setLevel(logging.ERROR)
jieba.setLogLevel(logging.INFO)


class NLU():
    def __init__(self, codedir):
        self.model = Interpreter.load("models/")
        self.codedir = codedir

    def set_codedir(self, path):
        self.codedir = path

    def get_codedir(self):
        return self.codedir

    def handler(self, event, context):
        res = {}
        if isinstance(event, dict):
            if "err" in event:
                raise TypeError(event['err'])
            res = event
        elif isinstance(event, bytes):
            res['bytes'] = event.decode("utf-8")
        if 'messageQOS' in context:
            res['messageQOS'] = context['messageQOS']
        if 'messageTopic' in context:
            res['messageTopic'] = context['messageTopic']
        if 'messageTimestamp' in context:
            res['messageTimestamp'] = context['messageTimestamp']
        if 'functionName' in context:
            res['functionName'] = context['functionName']
        if 'functionInvokeID' in context:
            res['functionInvokeID'] = context['functionInvokeID']
        res['Say'] = self.model.parse(u"你好")
        return res
