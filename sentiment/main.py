# from sklearn.svm import SVC
import jieba
import numpy as np
from gensim.models import KeyedVectors
from sklearn.externals import joblib

from word2vec import getWordVecs


def msg2Vec(message, model):
    wordList = list(jieba.cut(message))
    print(wordList)
    resultList = getWordVecs(wordList, model)
    input = []
    if len(resultList) != 0:
        resultArray = sum(np.array(resultList)) / len(resultList)
        input.append(resultArray)
    return np.array(input)


def checkSVM(message):
    # clf = SVC(C=2, probability=True)
    # Load svm clf
    svm_model = joblib.load("models/SVC_origin.pkl")
    clf = svm_model

    inp = 'models/wiki.zh.text.vector'
    model = KeyedVectors.load_word2vec_format(inp, binary=False)
    vector = msg2Vec(message, model)
    # vector = PCA(n_components=100).fit_transform(vector)
    prediction = clf.predict(vector)

    print(prediction)

    if prediction[0]:
        print("正面")
    else:
        print("负面")


if __name__ == "__main__":
    checkSVM("不错，在同等档次酒店中应该是值得推荐的！")
