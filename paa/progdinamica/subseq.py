# coding=utf-8

from matrix import Matrix

class Slice:
    def __init__(self, start=0,end=0):
        self.start = start
        if end < start:
            end = start
        self.end = end
    
    def len(self):
        return self.end - self.start
    
    def expand(self, ):
        return Slice(self.start, self.end+1)

def subseq(m=Matrix(zero=lambda: 0), a="",b=""):
    max_seq = (-1, -1)
    for ia in range(0, len(a)):
        for ib in range(0, len(b)):
            if a[ia] == b[ib]:
                seq = m.get(ia-1, ib-1) + 1
                m.put(ia, ib, seq)
                if seq > m.get(max_seq[0], max_seq[1]):
                    max_seq = (ia, ib)
            else:
                m.put(ia, ib, m.get(ia-1, ib-1))
    seq = m.get(max_seq[0], max_seq[1])
    if seq == 0:
        print "{} and {} don't share a sub-sequence".format(a, b)
        return 0
    print "{} and {} share the sequence of {} items".format(a,
        b,
        a[max_seq[0]-seq+1:max_seq[0]+1])

if __name__ == '__main__':
    subseq(a="abc1234561234", b="123456789")