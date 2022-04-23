from objects import Object
from attributes import attributes
from node import Node

file = open("./static/car.data", "r")
lines = file.readlines()

objects = []

for line in lines:
    splits = line.removesuffix('\n').split(',')
    objects.append(Object(splits))


# -------------------------------------------------> using 90% of the objects

max_depth = len(attributes) - 1

division_number = round(len(objects) * (90/100))
trainee_objects = objects[:division_number]
test_objects = objects[division_number:]

root = Node(trainee_objects, max_depth)
available_attributes = list(range(0, max_depth))
root.split(available_attributes)

training_error = root.get_training_error_counts()
print(f'training error: {(training_error/division_number)*100}')

testing_error = 0
for test_object in test_objects:
    label = root.get_label(test_object)
    if (test_object.attributes[-1] != label):
        testing_error += 1

print(f'test error: {(testing_error/(len(objects)-division_number))*100}')

# -------------------------------------------------> k-fold cross validation


def k_fold(k: int, max_depth: int = len(attributes)-1) -> float:
    k_fold_division_count = round(len(objects) / k)
    k_fold_training_error = 0
    k_fold_testing_error = 0

    for index in range(k):
        start = k_fold_division_count * index
        end = k_fold_division_count*(index+1)
        trainee_objects = objects[0:start] + objects[end:]
        test_objects = objects[start:end]

        root = Node(trainee_objects, max_depth)
        available_attributes = list(range(0, len(attributes)-1))
        root.split(available_attributes)

        k_fold_training_error += root.get_training_error_counts()

        for test_object in test_objects:
            label = root.get_label(test_object)
            if (test_object.attributes[-1] != label):
                k_fold_testing_error += 1

    return (k_fold_testing_error/(k_fold_division_count))/k


print()
k_values = [3, 5, 10]
for value in k_values:
    test_error = k_fold(value) * 100
    print(f'test error with {value}-fold cross-validation: {test_error}')

# -------------------------------------------------> find optimal depth

optimal_depth_value = 0
optimal_depth = float("inf")

print()
for depth in range(len(attributes) - 1):
    k = 10
    test_error = k_fold(k, depth+1)
    complexity_error = 0.05 * (depth+1) / max_depth
    sum_error = (test_error + complexity_error) * 100

    if (sum_error < optimal_depth):
        optimal_depth = sum_error
        optimal_depth_value = depth+1

    print(
        f'test error with {k}-fold cross-validation and depth: {depth+1}: {sum_error}')

print()
print(
    f'maximum depth is: {max_depth}, optimal depth is: {optimal_depth_value}')
