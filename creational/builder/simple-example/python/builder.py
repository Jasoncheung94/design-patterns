class Building:
    def __init__(self) -> None:
        self.build_door()
        self.build_window()
        self.build_floor()
        self.build_size()

    def build_window(self):
        raise NotImplementedError

    def build_door(self):
        raise NotImplementedError

    def build_floor(self):
        raise NotImplementedError

    def build_size(self):
        raise NotImplementedError

    def __repr__(self) -> str:
        return "Door: {0.door} | Window {0.window} | Floor: {0.floor} | Size: {0.size}".format(self)


class House(Building):
    def build_floor(self) -> None:
        self.floor = 2

    def build_size(self) -> None:
        self.size = "Big"

    def build_window(self) -> None:
        self.window = "Fixed"

    def build_door(self) -> None:
        self.door = "Wooden"


class Igloo(Building):
    def build_floor(self) -> None:
        self.floor = 1

    def build_window(self) -> None:
        self.window = "Ice"

    def build_door(self) -> None:
        self.door = "Ice"

    def build_size(self) -> None:
        self.size = "Small"

class ComplexBuilding:
    def __repr__(self) -> str:
        return "Door: {0.door} | Window {0.window} | Floor: {0.floor} | Size: {0.size}".format(self)


class ComplexHouse(ComplexBuilding):
    def build_floor(self) -> None:
        self.floor = "One"

    def build_size(self) -> None:
        self.size = "Big and fancy"

    def build_window(self) -> None:
        self.window = "Fixed"

    def build_door(self) -> None:
        self.door = "Metal Gates"


def construct(buildingType) -> Building:
    b = buildingType()
    b.build_floor()
    b.build_size()
    b.build_window()
    b.build_door()
    return b


def main():
    house = construct(House)
    flat = construct(Igloo)
    complex_house = construct(ComplexHouse)
    print(house)
    print(flat)
    print(complex_house)


if __name__ == "__main__":
    main()