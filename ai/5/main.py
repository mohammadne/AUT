# --------------------------------------------------------------------------> import libraries

from sklearn.metrics import accuracy_score
from sklearn.ensemble import RandomForestClassifier
from sklearn.neighbors import KNeighborsClassifier
from sklearn.linear_model import LogisticRegression
from sklearn.svm import SVC
import numpy as np
import pandas as pd

import matplotlib.pyplot as plt
import matplotlib.cm as cm
%matplotlib inline


# --------------------------------------------------------------------------> import data

training_data = pd.read_csv('train.csv')
testing_data = pd.read_csv('test.csv')
training_data.head()

# --------------------------------------------------------------------------> labels

# Get X and y for training data
X_train = training_data.drop(columns=['Activity', 'subject'])
y_train = training_data["Activity"]

# Get X and y for testing data
y_test = testing_data['Activity']
X_test = testing_data.drop(columns=['Activity', 'subject'])

# --------------------------------------------------------------------------> visualization

count_of_each_activity = np.array(y_train.value_counts())
activities = sorted(y_train.unique())
colors = cm.rainbow(np.linspace(0, 1, 4))
plt.figure(figsize=(10, 6))
plt.bar(activities, count_of_each_activity, width=0.3, color=colors)
plt.xticks(rotation=45, fontsize=12)
plt.yticks(rotation=45, fontsize=12)

Acc = 0
Gyro = 0
other = 0

for value in X_train.columns:
    if "Acc" in str(value):
        Acc += 1
    elif "Gyro" in str(value):
        Gyro += 1
    else:
        other += 1

plt.figure(figsize=(12, 8))
plt.bar['Accelerometer', 'Gyroscope', 'Others'], [
    Acc, Gyro, other], color = ('r', 'g', 'b')

# --------------------------------------------------------------------------> modeling

accuracy_scores = np.zeros(4)

# Support Vector Classifier
clf = SVC().fit(X_train, y_train)
prediction = clf.predict(X_test)
accuracy_scores[0] = accuracy_score(y_test, prediction)*100
print('Support Vector Classifier accuracy: {}%'.format(accuracy_scores[0]))

# Logistic Regression
clf = LogisticRegression().fit(X_train, y_train)
prediction = clf.predict(X_test)
accuracy_scores[1] = accuracy_score(y_test, prediction)*100
print('Logistic Regression accuracy: {}%'.format(accuracy_scores[1]))

# K Nearest Neighbors
clf = KNeighborsClassifier().fit(X_train, y_train)
prediction = clf.predict(X_test)
accuracy_scores[2] = accuracy_score(y_test, prediction)*100
print('K Nearest Neighbors Classifier accuracy: {}%'.format(
    accuracy_scores[2]))

# Random Forest
clf = RandomForestClassifier().fit(X_train, y_train)
prediction = clf.predict(X_test)
accuracy_scores[3] = accuracy_score(y_test, prediction)*100
print('Random Forest Classifier accuracy: {}%'.format(accuracy_scores[3]))

# --------------------------------------------------------------------------> accuracy visualization

plt.figure(figsize=(12, 8))
colors = cm.rainbow(np.linspace(0, 1, 4))
labels = ['Support Vector Classifier', 'Logsitic Regression',
          'K Nearest Neighbors', 'Random Forest']
plt.bar(labels,
        accuracy_scores,
        color=colors)
plt.xlabel('Classifiers', fontsize=18)
plt.ylabel('Accuracy', fontsize=18)
plt.title('Accuracy of various algorithms', fontsize=20)
plt.xticks(rotation=45, fontsize=12)
plt.yticks(fontsize=12)
