# coding=utf-8

from matrix import Matrix

class Item:
    def __init__(self, value=0, weight=0):
        self.weight = weight
        self.value = value
    
    def __str__(self):
        return "{}Kg / {}$".format(self.weight, self.value)

def make_matrix(m=Matrix(), items=[], max_weight=0):
    items = sorted(items, key=lambda i: i.weight)
    for idx in range(0, len(items)):
        iw = items[idx].weight
        for weight in range(0, max_weight+1):
            if iw > weight:
                # se o item atual é mais pesado do que o peso máximo da mochila
                # pega o valor da mochila considerando o peso anterior
                m.put(idx, weight, m.get(idx, weight-1))
                continue
            value_with_cur_idx = items[idx].value + m.get(idx - 1, weight - iw)
            value_without = m.get(idx-1, weight)
            if value_without > value_with_cur_idx:
                m.put(idx, weight, value_without)
            else:
                m.put(idx, weight, value_with_cur_idx)
        print "For item ({}) the best value is {}".format(items[idx], m.get(idx, max_weight))

if __name__=='__main__':
    items = [
        Item(weight=100,value=40),
        Item(weight=50,value=35),
        Item(weight=45,value=18),
        Item(weight=20,value=4),
        Item(weight=10,value=10),
        Item(weight=5,value=2)]
    make_matrix(items=items,max_weight=100)
