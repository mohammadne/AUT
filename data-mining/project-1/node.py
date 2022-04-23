from math import log

from objects import Object
from attributes import attributes


class Node:
    def __init__(self,  objects: list[Object], max_depth: int) -> None:
        self.objects = objects
        self.max_depth = max_depth
        self.label = -1

    def split(self, available_attributes: list[int]) -> None:
        maj = self.majority(self.objects)
        we_reach = True

        for obj in self.objects:
            if (obj.attributes[-1] != maj):
                we_reach = False
                break

        depth = (len(attributes) - 1) - len(available_attributes)
        if (we_reach or self.max_depth == depth):
            self.label = maj
            return

        splits = []  # contain all splits of nodes

        for attribute in available_attributes:
            nodes = []  # contain nodes with this attribute

            for attribute_value in range(len(attributes[attribute])):
                matches = filter(
                    lambda obj: obj.attributes[attribute] == attribute_value, self.objects)
                nodes.append(Node(list(matches), self.max_depth))

            splits.append(nodes)

        available_attributes_for_children = available_attributes.copy()

        if (len(available_attributes) == 1):
            self.attribute = available_attributes[0]
            self.children = splits[0]
            del available_attributes_for_children[0]
        else:
            max_gain_ratio_index = 0
            max_gain_ratio = 0
            for index in range(len(available_attributes)):
                gain_ratio = self.gain_ratio(splits[index])
                if (gain_ratio > max_gain_ratio):
                    gain_ratio
                    max_gain_ratio = gain_ratio
                    max_gain_ratio_index = index

            self.attribute = available_attributes[max_gain_ratio_index]
            self.children = splits[max_gain_ratio_index]

            del available_attributes_for_children[max_gain_ratio_index]

        for child in self.children:
            child.split(available_attributes_for_children)

    def gain_ratio(self, nodes: list) -> float:
        gain_info = self.entropy(self)
        for node in nodes:
            if (len(node.objects) == 0):
                continue

            entropy = self.entropy(node)
            ratio = len(node.objects) / len(self.objects)
            gain_info -= ratio*entropy

        split_info = 0
        for node in nodes:
            if (len(node.objects) == 0):
                continue

            ratio = len(node.objects) / len(self.objects)
            split_info -= ratio * log(ratio, 2)

        if (split_info == 0):
            return 0

        return gain_info / split_info

    def get_label(self, obj: Object) -> int:
        if (self.label != -1):
            return self.label

        attribute_value = obj.attributes[self.attribute]
        return self.children[attribute_value].get_label(obj)

    def get_training_error_counts(self) -> int:
        if (self.label != -1):
            matches = filter(
                lambda obj: obj.attributes[-1] != self.label, self.objects)
            return len(list(matches))

        count = 0

        for child in self.children:
            count += child.get_training_error_counts()

        return count

    @staticmethod
    def majority(objects: list[Object]) -> int:
        counts = [0] * len(attributes[-1])

        for obj in objects:
            counts[obj.attributes[-1]] += 1

        max_value, max_index = -1, -1
        for index in range(len(counts)):
            if (counts[index] > max_value):
                max_value = counts[index]
                max_index = index

        return max_index

    @staticmethod
    def entropy(node) -> float:
        entropy = 0.0

        for label in range(len(attributes[-1])):
            matches = filter(
                lambda obj: obj.attributes[-1] == label, node.objects)

            ratio = len(list(matches)) / len(node.objects)
            if (ratio != 0):
                entropy -= ratio * log(ratio, 2)

        return entropy
