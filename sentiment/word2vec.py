import codecs

from gensim.models import word2vec

from corpus import HotelCorpus as hotel
from corpus import Waimai2Corpus as waimai2
from corpus import WaimaiCorpus as waimai


def corpus_to_csv():
    """
    Cleaning corpus files and write result to csv files
    :return:
    """
    # Save corpus to csv files
    # hotel corpus
    hotel_pos_csv = codecs.open("data/corpus_csv/hotel_pos.csv", 'w', 'utf-8')
    for line in hotel().pos_list:
        hotel_pos_csv.write("%s " % ' '.join(line))
    hotel_neg_csv = codecs.open("data/corpus_csv/hotel_neg.csv", 'w', 'utf-8')
    for line in hotel().neg_list:
        hotel_neg_csv.write("%s " % ' '.join(line))

    # waimai corpus
    waimai_pos_csv = codecs.open("data/corpus_csv/waimai_pos.csv", 'w', 'utf-8')
    for line in waimai().pos_list:
        waimai_pos_csv.write("%s " % ' '.join(line))
    waimai_neg_csv = codecs.open("data/corpus_csv/waimai_neg.csv", 'w', 'utf-8')
    for line in waimai().neg_list:
        waimai_neg_csv.write("%s " % ' '.join(line))

    # waimai2 corpus
    waimai2_pos_csv = codecs.open("data/corpus_csv/waimai2_pos.csv", 'w', 'utf-8')
    for line in waimai2().pos_list:
        waimai2_pos_csv.write("%s " % ' '.join(line))
    waimai2_neg_csv = codecs.open("data/corpus_csv/waimai2_neg.csv", 'w', 'utf-8')
    for line in waimai2().neg_list:
        waimai2_neg_csv.write("%s " % ' '.join(line))


def csv2vec():
    hotel_corpus = word2vec.Text8Corpus("corpus.csv")


if __name__ == "__main__":
    corpus_to_csv()
