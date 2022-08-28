import copy

class Prototype:
    def clone(self):
        return copy.deepcopy(self)

if __name__ == "__main__":
    prototype = Prototype()
    prototype_copy = prototype.clone()
    print(prototype_copy)