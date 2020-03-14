import os
import re


class Corpus:
    """
    Load corpus data from data/corpus/*
    pos_list: positive corpus's words list
    neg_list: negative corpus's words list
    train_num: corpus numbers used for training
    test_num: corpus numbers used for check(test)
    """

    def __init__(self, file_path):
        root_path = os.path.dirname(os.path.abspath(__file__))
        file_path = os.path.normpath(os.path.join(root_path, file_path))
        detected_file_path = os.path.normpath(os.path.join(root_path, "data/dict/detected_dict.txt"))

        re_split = re.compile("\s+")

        self.pos_list = []
        self.neg_list = []

        # Load detected file first
        print("Loading %s" % detected_file_path)
        with open(detected_file_path, 'r', encoding='utf-8', errors='ignore') as f:
            detected_list = f.readlines()

        # Delete the last'\n'of each line
        for i in range(0, len(detected_list)):
            detected_list[i] = detected_list[i][:-1]

        self.detected_dics = {}.fromkeys(detected_list, 1)
        print("Loading %s" % file_path)
        # Open corpus data file as utf-8
        with open(file_path, encoding="utf-8") as f:
            # get each line
            for line in f:
                # strip: remove space before and after a string
                line = line.strip()
                # remove all symbols
                # line = self.remove_symbol(line)
                # split: split string to list
                splits = re_split.split(line)
                # extract Chinese words
                chi_word_list = re.findall('[\u4e00-\u9fa5]+', line)
                if splits[0] == "pos":
                    # append list after 'pos' to pos_list
                    # Remove detected words
                    self.pos_list.append(self.clean_detected_words(chi_word_list))
                elif splits[0] == "neg":
                    self.neg_list.append(self.clean_detected_words(chi_word_list))
                else:
                    raise ValueError("Read corpus from %s failed, not have 'pos' or 'neg'\n" % file_path)
            f.close()
        self.pos_list_len = len(self.pos_list)
        self.neg_list_len = len(self.neg_list)

        # Check if post_list_len correctly
        # assert len(self.neg_list) == self.pos_list_len

        self.train_num = 0
        self.test_num = 0

        output_msg = "Using corpus: %s.\n" % file_path
        output_msg += "postive corpus: %d\tnegative corpus: %d" % \
                      (self.pos_list_len, self.neg_list_len)
        print(output_msg)

    def remove_symbol(self, sentence):
        """
        Remove all symbol from sentence
        :param sentence:
        :return:
        """
        result = re.sub("[\s+\.\!\/_,$%^*(+\"\']+|[+——！，。？、~@#￥%……&*（）；：．{}０１２３４８９]", " ", sentence)
        return result

    def list_clean(self, strings):
        """
        Remove all letters and numbers
        :param strings:
        :return: result
        """
        result = []
        for s in strings:
            cleaned = re.sub("[\s+|[A-Z]+|[a-z]+|[0-9]", "", s)
            # cleaned.replace("\n", "")
            if cleaned != "":
                result.append(cleaned)
        return result

    def clean_detected_words(self, cutted_list):
        """
        Remove detected words
        :param cutted_list:
        :return:
        """
        result = []
        for cutted in cutted_list:
            if cutted not in self.detected_dics:
                result.append(cutted)
        return result


class WaimaiCorpus(Corpus):
    def __init__(self):
        Corpus.__init__(self, "data/corpus/ch_waimai_corpus.txt")


class Waimai2Corpus(Corpus):
    def __init__(self):
        Corpus.__init__(self, "data/corpus/ch_waimai2_corpus.txt")


class HotelCorpus(Corpus):
    def __init__(self):
        Corpus.__init__(self, "data/corpus/ch_hotel_corpus.txt")
