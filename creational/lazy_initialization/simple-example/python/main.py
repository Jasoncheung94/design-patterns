class Fruit:
    def __init__(self, item: str) -> None:
        self.item = item

class Fruits:
    def __init__(self) -> None:
        self.items = {}

    def get_fruit(self, item: str) -> Fruit:
        if item not in self.items:
            self.items[item] = Fruit(item)

        return self.items[item]

    def __str__(self) -> str:
        return ', '.join([str(x) for x in self.items])

if __name__ == "__main__":
    fruits = Fruits()
    print("Has fruits?",len(fruits.items) > 0)
    print(fruits.get_fruit("Apple"))
    print("Has fruits?",len(fruits.items) > 0)
    print(fruits)
    print(fruits.get_fruit("Lime"))
    print(fruits)