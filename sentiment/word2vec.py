import codecs

import matplotlib.pyplot as plt
import numpy as np
from gensim.models import KeyedVectors
from gensim.models import word2vec
from sklearn.decomposition import PCA
from sklearn.preprocessing import scale

from corpus import HotelCorpus as hotel
from corpus import Waimai2Corpus as waimai2
from corpus import WaimaiCorpus as waimai


def corpus2csv(data_list, csv_file):
    file = codecs.open(csv_file, 'w', 'utf-8')
    for line in data_list:
        if line == len(data_list):
            file.write("%s" % ' '.join(line))
        else:
            file.write("%s\n" % ' '.join(line))


def Corpus2CSVs():
    """
    Cleaning corpus files and write result to csv files
    :return:
    """
    # Save corpus to csv files
    # hotel corpus
    corpus2csv(hotel().pos_list, "data/corpus_csv/hotel_pos.csv")
    corpus2csv(hotel().neg_list, "data/corpus_csv/hotel_neg.csv")

    # waimai corpus
    corpus2csv(waimai().pos_list, "data/corpus_csv/waimai_pos.csv")
    corpus2csv(waimai().neg_list, "data/corpus_csv/waimai_neg.csv")

    # waimai2 corpus
    corpus2csv(waimai2().pos_list, "data/corpus_csv/waimai2_pos.csv")
    corpus2csv(waimai2().neg_list, "data/corpus_csv/waimai2_neg.csv")


def csv2vec(csv_file, model_file):
    # 加载语料
    corpus = word2vec.Text8Corpus(csv_file)
    # 训练skip - gram模型
    model = word2vec.Word2Vec(corpus, size=400)
    # Save models to file
    model.save(model_file)


def CSV2Vecs():
    """
    Build word vector model
    :return:
    """
    print("Traing word vector model...")
    csv2vec("data/corpus_csv/hotel_pos.csv", "models/hotel_pos.vec")
    print("15%...")
    csv2vec("data/corpus_csv/hotel_neg.csv", "models/hotel_neg.vec")
    print("30%...")
    csv2vec("data/corpus_csv/waimai_pos.csv", "models/waimai_pos.vec")
    print("45%...")
    csv2vec("data/corpus_csv/waimai_neg.csv", "models/waimai_neg.vec")
    print("60%...")
    csv2vec("data/corpus_csv/waimai2_pos.csv", "models/waimai2_pos.vec")
    print("75%...")
    csv2vec("data/corpus_csv/waimai2_neg.csv", "models/waimai2_neg.vec")
    print("done.")


def getWordVecs(wordList, model):
    vecs = []
    for word in wordList:
        word = word.replace('\n', '')
        try:
            vecs.append(model[word])
        except KeyError:
            continue
    # vecs = np.concatenate(vecs)
    return np.array(vecs, dtype='float')


def buildVec(csv_file, vec_model):
    # Load word2vec model
    # model = word2vec.Word2Vec.load_word2vec_format(vec_model, binary = True)
    model = KeyedVectors.load(vec_model)

    input = []
    # Load csv file
    f = codecs.open(csv_file, 'r', 'utf-8')
    lines = f.read().split('\n')
    for line in lines:
        # remove space line
        res = line.replace(' ', '')
        if len(res) != 0:
            resultList = getWordVecs(list(line), model)
            # for each sentence, the mean vector of all its vectors is used to represent this sentence
            if len(resultList) != 0:
                resultArray = sum(np.array(resultList)) / len(resultList)
                input.append(resultArray)
    return input


def buildVecs():
    waimai_pos = buildVec("data/corpus_csv/waimai_pos.csv", "models/waimai_pos.vec")
    waimai_neg = buildVec("data/corpus_csv/waimai_neg.csv", "models/waimai_neg.vec")

    # use 1 for positive sentiment, 0 for negative
    y = np.concatenate((np.ones(len(waimai_pos)), np.zeros(len(waimai_neg))))

    X = waimai_pos[:]
    for neg in waimai_neg:
        X.append(neg)
    X = np.array(X)

    # standardization
    X = scale(X)

    # 无监督使用PCA训练X
    pca = PCA(n_components=400)
    pca.fit(X)
    # 创建图表并指定图表大小
    # figsize: w,h tuple in inches
    plt.figure(1, figsize=(4, 3))
    plt.clf()
    plt.axes([.2, .2, .7, .7])
    plt.plot(pca.explained_variance_, linewidth=2)
    plt.axis('tight')
    plt.xlabel('n_components')
    plt.ylabel('explained_variance_')
    plt.show()


if __name__ == "__main__":
    # Corpus2CSVs()
    # CSV2Vecs()
    buildVecs()
