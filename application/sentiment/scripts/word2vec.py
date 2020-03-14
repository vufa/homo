import codecs

import matplotlib.pyplot as plt
import numpy as np
import pandas as pd
from gensim.models import KeyedVectors
from sklearn.decomposition import PCA
from sklearn.externals import joblib
from sklearn.linear_model import LogisticRegression
from sklearn.metrics import roc_curve, auc, accuracy_score, classification_report, roc_auc_score
from sklearn.model_selection import GridSearchCV
from sklearn.model_selection import train_test_split
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
            # print(model[word])
        except KeyError:
            continue
    return np.array(vecs, dtype='float')


def buildVec(csv_file, model):
    fileVecs = []
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
                fileVecs.append(resultArray)
    return fileVecs


def drawX(X, dimension):
    """
    Plot the PCA spectrum
    :param pca:
    :return:
    """
    pca = PCA(n_components=dimension)
    pca.fit(X)
    # figsize: w,h tuple in inches
    plt.figure(1, figsize=(6, 4.5))
    plt.clf()
    plt.axes([.2, .2, .7, .7])
    plt.plot(pca.explained_variance_, linewidth=2)
    plt.axis('tight')
    plt.xlabel('n_components')
    plt.ylabel('explained_variance_')
    plt.show()


def Vecs2CSV():
    inp = 'models/wiki.zh.text.vector'
    model = KeyedVectors.load_word2vec_format(inp, binary=False)
    hotel_pos = buildVec("data/corpus_csv/hotel_pos.csv", model)
    hotel_neg = buildVec("data/corpus_csv/hotel_neg.csv", model)

    waimai_pos = buildVec("data/corpus_csv/waimai_pos.csv", model)
    waimai_neg = buildVec("data/corpus_csv/waimai_neg.csv", model)

    waimai2_pos = buildVec("data/corpus_csv/waimai2_pos.csv", model)
    waimai2_neg = buildVec("data/corpus_csv/waimai2_neg.csv", model)

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
    plt.figure(figsize=(6, 4.5))
    plt.plot(fpr, tpr, label='area = %.2f' % roc_auc)
    plt.plot([0, 1], [0, 1], 'k--')
    plt.xlim([0.0, 1.0])
    plt.ylim([0.0, 1.05])
    plt.legend(loc='lower right')
    plt.show()


def LR(x_pca, Y):
    """
    Using LogisticRegression
    :param X_train:
    :param X_test:
    :param y_train:
    :param y_test:
    :return:
    """
    # cut dataset
    X_train, X_test, y_train, y_test = train_test_split(x_pca, Y, test_size=0.25, random_state=0)

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


def cutTest(x_pca, Y):
    """
    test_size = 0.25
    :param clf:
    :param x_pca:
    :param Y:
    :return:
    """
    X_train, X_test, y_train, y_test = train_test_split(x_pca, Y, test_size=0.25, random_state=0)
    clf = SVC(C=2, probability=True)
    clf.fit(X_train, y_train)
    print("Train num = %s" % len(X_train))
    print("Test num = %s" % len(X_test))
    print('Test Accuracy: %.2f' % clf.score(X_test, y_test))
    displayRoc(clf, X_test, y_test)


def buildModel():
    df = pd.read_csv('models/all_vector.csv')
    # first column is index, second column is label
    Y = df.iloc[:, 1]
    # 3th and all latter columns is word vectors
    X = df.iloc[:, 2:]
    # standardization
    X = scale(X)

    # Plot the PCA spectrum
    drawX(X, 400)

    #
    # Original clf
    #
    # clf_orig = SVC(C=2, probability=True)
    # clf_orig.fit(X, Y)
    # joblib.dump(clf_orig, "models/SVC_origin.pkl")

    # Reduce to 100 dimension
    # len(x_pca) = 21978 == len(X)
    x_pca = PCA(n_components=100).fit_transform(X)
    drawX(x_pca, 100)

    # SVM (RBF)
    # using training data with 100 dimensions
    clf = SVC(C=2, probability=True)
    clf.fit(x_pca, Y)
    print('Test Accuracy: %.2f' % clf.score(x_pca, Y))
    # Save SVMClassifier to file
    joblib.dump(clf, "models/SVC.pkl")
    # Load use:
    # clf = joblib.load("models/SVC.pkl")
    displayRoc(clf, x_pca, Y)

    #
    # cut dataset to train and test
    #
    # cutTest(x_pca, Y)

    # with LR
    LR(x_pca, Y)


if __name__ == "__main__":
    # Extract data from raw text
    # Corpus2CSVs()

    # Vecs2CSV()
    buildModel()
