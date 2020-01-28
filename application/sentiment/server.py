# from sklearn.svm import SVC
import json
from http.server import BaseHTTPRequestHandler, HTTPServer
from time import process_time
from urllib.parse import urlparse

import jieba
import numpy as np
from gensim.models import KeyedVectors
# from sklearn.decomposition import PCA
from sklearn.externals import joblib

from word2vec import getWordVecs


def msg2Vec(message, model):
    wordList = list(jieba.cut(message))
    # print(wordList)
    resultList = getWordVecs(wordList, model)
    input = []
    if len(resultList) != 0:
        resultArray = sum(np.array(resultList)) / len(resultList)
        input.append(resultArray)
    return np.array(input)


def checkSVM(message):
    # time
    start = process_time()
    # clf = SVC(C=2, probability=True)
    # Load svm clf
    svm_model = joblib.load("models/SVC_origin.pkl")

    e1 = process_time()
    print('joblib.load cost: %s Seconds' % (e1 - start))

    clf = svm_model

    inp = 'models/wiki.zh.text.vector'
    # Need most time
    model = KeyedVectors.load_word2vec_format(inp, binary=False)

    e2 = process_time()
    print('load_word2vec_format cost: %s Seconds' % (e2 - e1))

    vector = msg2Vec(message, model)

    e3 = process_time()
    print('msg2Vec cost: %s Seconds' % (e3 - e2))

    # vector = PCA(n_components=100).fit_transform(vector)
    prediction = clf.predict(vector)

    e4 = process_time()
    print('clf.predict cost: %s Seconds' % (e4 - e3))

    print(prediction)

    if prediction[0]:
        print("正面")
    else:
        print("负面")

    print('total cost: %s Seconds' % (e4 - start))


class RequestHandler(BaseHTTPRequestHandler):
    def __init__(self, clf, model, *args):
        self.clf = clf
        self.model = model
        BaseHTTPRequestHandler.__init__(self, *args)

    def do_GET(self):
        parsed_path = urlparse(self.path)
        self.send_response(200)
        self.end_headers()
        self.wfile.write(json.dumps({
            'method': self.command,
            'path': self.path,
            'real_path': parsed_path.query,
            'query': parsed_path.query,
            'request_version': self.request_version,
            'protocol_version': self.protocol_version
        }).encode())
        return

    def do_POST(self):
        content_len = int(self.headers.get('Content-Length'))
        post_body = self.rfile.read(content_len)
        # extract request body and convert to array from json
        data = json.loads(post_body)

        # Print request
        parsed_path = urlparse(self.path)
        """
        print(json.dumps({
            'method': self.command,
            'path': self.path,
            'real_path': parsed_path.query,
            'query': parsed_path.query,
            'request_version': self.request_version,
            'protocol_version': self.protocol_version,
            'body': data
        }).encode())
        """
        # get text
        message = data['query']
        if message != "":
            # print(data['query'])
            # convert to vector
            vector = msg2Vec(message, self.model)
            prediction = self.clf.predict(vector)
            # print(prediction)

            # give response
            self.send_response(200)
            self.end_headers()
            self.wfile.write(json.dumps({
                'code': 200,
                'prediction': prediction.tolist(),
            }).encode())
        else:
            self.send_response(200)
            self.end_headers()
            self.wfile.write(json.dumps({
                'code': 500,
                'message': 'Invalid request parameter'
            }).encode())
        return


class SentimentServer:
    def __init__(self, svm_model_path, inp):
        # Load svm clf
        print("Loading svm clf...")
        e1 = process_time()
        clf = joblib.load(svm_model_path)
        e2 = process_time()
        print("Loading svm clf cost %.5f seconds" % (e2 - e1))

        print("Loading word2vec model...")
        # Need most time
        e2 = process_time()
        model = KeyedVectors.load_word2vec_format(inp, binary=False)
        e3 = process_time()
        print("Loading word2vec model cost %.5f seconds" % (e3 - e2))

        def handler(*args):
            RequestHandler(clf, model, *args)

        server = HTTPServer(('', 8000), handler)
        print('Starting server at http://localhost:8000')
        server.serve_forever()


if __name__ == "__main__":
    # checkSVM("不错，在同等档次酒店中应该是值得推荐的！")
    # checkSVM("餐厅打包人员少打包一个套餐里的薯条，让配送员跑了两趟才吃上，这顿饭用了一个半小时才全部送达到我手里。")
    # checkSVM("芝士肉酱薯难吃至极")
    SentimentServer("models/SVC_origin.pkl", 'models/wiki.zh.text.vector')
