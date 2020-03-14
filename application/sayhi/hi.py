#!/usr/bin/env python
# -*- coding:utf-8 -*-
"""
nlu module
"""


class HI():
    def __init__(self, codedir):
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
        res['Say'] = 'Hello Aiicy'
        return res
