import codecs

import matplotlib.pyplot as plt
import numpy as np
import pandas as pd
from gensim.models import KeyedVectors
from gensim.models import word2vec
from sklearn.decomposition import PCA
from sklearn.linear_model import LogisticRegression
from sklearn.metrics import roc_curve, auc, accuracy_score, classification_report, roc_auc_score
from sklearn.model_selection import GridSearchCV
from sklearn.preprocessing import scale
from sklearn.svm import SVC

from corpus import HotelCorpus as hotel
from corpus import Waimai2Corpus as waimai2
from corpus import WaimaiCorpus as waimai


def corpus2csv(data_list, csv_file):
    file = codecs.open(csv_file, 'w', 'utf-8')
    sum = 1
    for line in data_list:
        # the last line
        if sum == len(data_list):
            file.write("%s" % ' '.join(line))
        else:
            file.write("%s\n" % ' '.join(line))
        sum += 1


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
    # Load corpus from csv file
    sentences = word2vec.Text8Corpus(csv_file)
    """
    f = codecs.open(csv_file, 'r', 'utf-8')
    lines = f.read().split('\n')
    sentences = []
    for line in lines:
        # remove space line
        res = line.replace(' ', '')
        if len(res) != 0:
            sentences.append(line.split())
    """
    # 训练skip - gram模型
    model = word2vec.Word2Vec(sentences, size=400)
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
    """
    Get vector from a sentence's words
    :param wordList:
    :param model:
    :return:
    """
    vecs = []
    for word in wordList:
        word = word.replace('\n', '')
        try:
            vecs.append(model[word])
        except KeyError:
            continue
    result = np.array(vecs, dtype='float')
    return result


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
            resultList = getWordVecs(line.split(' '), model)
            # for each sentence, the mean vector of all its vectors is used to represent this sentence
            if len(resultList) != 0:
                resultArray = sum(np.array(resultList)) / len(resultList)
                input.append(resultArray)
    return input


def drawX(X, dimension):
    """
    Plot the PCA spectrum
    :param pca:
    :return:
    """
    pca = PCA(n_components=dimension)
    pca.fit(X)
    # figsize: w,h tuple in inches
    plt.figure(1, figsize=(4, 3))
    plt.clf()
    plt.axes([.2, .2, .7, .7])
    plt.plot(pca.explained_variance_, linewidth=2)
    plt.axis('tight')
    plt.xlabel('n_components')
    plt.ylabel('explained_variance_')
    plt.show()


def Vecs2CSV():
    hotel_pos = buildVec("data/corpus_csv/hotel_pos.csv", "models/hotel_pos.vec")
    hotel_neg = buildVec("data/corpus_csv/hotel_neg.csv", "models/hotel_neg.vec")

    waimai_pos = buildVec("data/corpus_csv/waimai_pos.csv", "models/waimai_pos.vec")
    waimai_neg = buildVec("data/corpus_csv/waimai_neg.csv", "models/waimai_neg.vec")

    waimai2_pos = buildVec("data/corpus_csv/waimai2_pos.csv", "models/waimai2_pos.vec")
    waimai2_neg = buildVec("data/corpus_csv/waimai2_neg.csv", "models/waimai2_neg.vec")

    # use 1 for positive sentiment, 0 for negative
    Y = np.concatenate((np.ones(len(hotel_pos) + len(waimai_pos) + len(waimai2_pos)),
                        np.zeros(len(hotel_neg) + len(waimai_neg) + len(waimai2_neg))))
    # Merge all data
    X = np.concatenate((hotel_pos, hotel_neg))
    X = np.concatenate((X, waimai_pos, waimai_neg))
    X = np.concatenate((X, waimai2_pos, waimai2_neg))

    X = np.array(X)

    # Save to file
    df_x = pd.DataFrame(X)
    df_y = pd.DataFrame(Y)
    data = pd.concat([df_y, df_x], axis=1)
    data.to_csv('models/all_vector.csv')


def displayRoc(clf_fited, X, y):
    # Create ROC curve
    pred_probas = clf_fited.predict_proba(X)[:, 1]  # score

    fpr, tpr, _ = roc_curve(y, pred_probas)
    roc_auc = auc(fpr, tpr)
    plt.plot(fpr, tpr, label='area = %.2f' % roc_auc)
    plt.plot([0, 1], [0, 1], 'k--')
    plt.xlim([0.0, 1.0])
    plt.ylim([0.0, 1.05])
    plt.legend(loc='lower right')
    plt.show()


def LR(X_train, X_test, y_train, y_test):
    """
    Using LogisticRegression
    :param X_train:
    :param X_test:
    :param y_train:
    :param y_test:
    :return:
    """
    param_grid = {'C': [0.01, 0.1, 1, 10, 100, 1000, ], 'penalty': ['l1', 'l2']}
    # Logistic Regression
    grid_search = GridSearchCV(LogisticRegression(), param_grid, cv=10)
    grid_search.fit(X_train, y_train)
    print(grid_search.best_params_, grid_search.best_score_)

    # 预测拆分的test
    LR = LogisticRegression(C=grid_search.best_params_['C'], penalty=grid_search.best_params_['penalty'])
    LR.fit(X_train, y_train)
    lr_y_predict = LR.predict(X_test)
    print(accuracy_score(y_test, lr_y_predict))
    print('使用LR进行分类的报告结果：')
    print(classification_report(y_test, lr_y_predict))
    print("AUC值:", roc_auc_score(y_test, lr_y_predict))
    displayRoc(LR, X_test, y_test)


def buildModel():
    df = pd.read_csv('models/all_vector.csv')
    # first column is index, second column is label
    Y = df.iloc[:, 1]
    # 3th and all latter columns is word vectors
    X = df.iloc[:, 2:]
    # standardization
    X = scale(X)

    # Plot the PCA spectrum
    # drawX(X, 400)

    # Reduce to 50 dimension
    # len(X_reduced) = 21978 == len(X)
    x_pca = PCA(n_components=50).fit_transform(X)
    # drawX(X_reduced, 50)

    # SVM (RBF)
    # using training data with 100 dimensions
    clf = SVC(C=2, probability=True)
    clf.fit(x_pca, Y)
    print('Test Accuracy: %.2f' % clf.score(x_pca, Y))
    displayRoc(clf, x_pca, Y)

    """
    #
    # test_size = 0.25
    #
    X_train, X_test, y_train, y_test = train_test_split(x_pca, Y, test_size=0.25, random_state=0)
    clf.fit(X_train, y_train)
    print('Test Accuracy: %.2f' % clf.score(X_test, y_test))
    displayRoc(clf, X_test, y_test)

    # with LR
    # LR(X_train, X_test, y_train, y_test)
    """


if __name__ == "__main__":
    # Corpus2CSVs()
    # CSV2Vecs()
    # Vecs2CSV(
    buildModel()
