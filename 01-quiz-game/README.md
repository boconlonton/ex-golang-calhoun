# Exercise 01: Quiz Game

## Part 1
- Create a program that will read in a quiz provided viz a CSV file and will then give the quiz to a user
keeping track of how many questions they get right and how many they get
incorrect. Regardless of whether the answer is correct or wrong the next question
should be asked immediately afterwards.

- The csv file should default to `problems.csv` but the user should be able
to customize the filename via a flag.

## Part 2
- Add a timer: The default time limit should be 30 seconds, but should also be
customizable via a flag.

- Your quiz should stop as soon as the time limit has exceeded. That is, you
shouldn't wait for the user to answer one final questions but should
ideally stop the quiz entirely even if you are currently waiting on an answers
from the end user.

- User should be asked to press `enter` before the timer starts, and
then the questions should be printed out to the screen one at a time
until the user provides an answer.

- At the end of the quiz, the program should still output the