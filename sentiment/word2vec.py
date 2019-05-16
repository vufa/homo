from corpus import HotelCorpus as hotel
# from corpus import WaimaiCorpus as waimai
# from corpus import Waimai2Corpus as waimai2
import codecs
from gensim.models import word2vec

def corpus_to_csv():
    # Save corpus to csv files
    hotel_pos_csv = codecs.open("data/corpus_csv/hotel_pos.csv", 'w', 'utf-8')
    for line in hotel().pos_list:
        hotel_pos_csv.write("%s " % ' '.join(line))
    hotel_neg_csv = codecs.open("data/corpus_csv/hotel_neg.csv", 'w', 'utf-8')
    hotel_neg_csv.write("%s" % ' '.join(str(hotel().neg_list)))

if __name__ == "__main__":
    corpus_to_csv()