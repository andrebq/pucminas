class Matrix:
    def __init__(self, zero=lambda:0):
        self.items = {}
        self.zero = zero

    def put(self, x, y, value):
        self.items[(x,y)] = value

    def get(self, x, y):
        try:
            return self.items[(x,y)]
        except KeyError:
            return self.zero()