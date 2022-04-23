from attributes import attributes


class Object:
    def __init__(self, splits) -> None:
        self.attributes = []
        for index in range(len(attributes)):
            value = attributes[index].index(splits[index])
            self.attributes.append(value)
